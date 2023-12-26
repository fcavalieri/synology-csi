#!/bin/bash
set -e
SOURCE_PATH=$(realpath "$(dirname "${BASH_SOURCE}")")

kubectl apply -f "${SOURCE_PATH}/test.yaml"
