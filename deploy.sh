#!/bin/bash

_IMAGE_NAME="dev_webapp_image"
_GCP_PROJECT_NAME="golang-web-app-331620"
_image_gcr_path=gcr.io/"${_GCP_PROJECT_NAME}"/"${_IMAGE_NAME}"
_TF_VERSION=1.0.11
_USER="vagrant"

function prepare_image() {
  docker build -t "${_image_gcr_path}" .
  docker push "${_image_gcr_path}"
}

function deploy() {
  cd infrastructure || exit
  tfenv use "${_TF_VERSION}"
  terraform init
  terraform apply -var=project_id=${_GCP_PROJECT_NAME} -var=image_gcr_path="${_image_gcr_path}" -lock=false
}

export PATH=/home/"${_USER}"/.tfenv/bin:$PATH
prepare_image
deploy
