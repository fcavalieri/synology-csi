#!/bin/bash
set -ex
SOURCE_PATH=$(realpath "$(dirname "${BASH_SOURCE}")")
ROOT_PATH=${SOURCE_PATH}/..

helm upgrade --install synology-csi \
             --create-namespace \
             --namespace synology-csi \
             -f ${SOURCE_PATH}/values.yaml \
             ${ROOT_PATH}/deploy/helm
kubectl apply -n synology-csi -f ${SOURCE_PATH}/cifs-csi-secret.yaml
