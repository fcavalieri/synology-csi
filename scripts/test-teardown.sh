#!/bin/bash
set -e
SOURCE_PATH=$(realpath "$(dirname "${BASH_SOURCE}")")

kubectl delete deployment -n synology-csi-test test --grace-period=0 --force
kubectl delete -f "${SOURCE_PATH}/test.yaml"
