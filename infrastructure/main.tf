resource "google_project_service" "run_api" {
  service = "run.googleapis.com"

  disable_on_destroy = true
}

resource "google_cloud_run_service" "run_service" {
  name = var.name
  location = "europe-central2"

  template {
    spec {
      containers {
        image = var.image_gcr_path
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  depends_on = [google_project_service.run_api]
}

output "service_url" {
  value = google_cloud_run_service.run_service.status[0].url
}