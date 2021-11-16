variable "project_id" {
  type        = string
  description = "GCP project id"
}

variable "image_gcr_path" {
  type        = string
  description = "Path to Docker image on GCR"
}