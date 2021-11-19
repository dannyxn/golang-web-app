package main

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"log"
	"net/http"
	"os"
)

func main() {
	session, err := initializeBackend()
	if err != nil {
		log.Fatalf("Error occured during backend initializaiton %v\nExiting...", err)
		return
	}
	defer session.Close()
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initializeBackend() (neo4j.Session, error) {
	username := "neo4j"
	passwordSecretResourceID := "projects/golang-web-app-331620/secrets/neo4j-golang-web-project-password/versions/latest"
	uri := "neo4j+s://a864ff6f.databases.neo4j.io"
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "sa-private-key.json")
	password := getNeo4jPassword(passwordSecretResourceID)

	if len(password) == 0 {
		return nil, errors.New("can't fetch Neo4j password from GCP Secrets Manager")
	}

	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	defer driver.Close()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	greeting, err := testDatabaseConnection(session)
	if err != nil || len(greeting) == 0 {
		log.Fatalf("Can't fetch greeting message or errors occurred: %v\nCheck database connection", err)
	} else {
		log.Print(greeting)
	}

	return session, nil
}

func getNeo4jPassword(passwordSecretResourceID string) string {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	defer client.Close()

	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: passwordSecretResourceID,
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
		return ""
	} else {
		log.Printf("Successfully fetched '%v'", passwordSecretResourceID)
	}
	return string(result.Payload.Data[:])
}

func testDatabaseConnection(session neo4j.Session) (string, error) {
	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]interface{}{"message": "hello, world"})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}
	return greeting.(string), nil
}
