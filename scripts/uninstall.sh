#!/bin/bash
set -e
SOURCE_PATH=$(realpath "$(dirname "${BASH_SOURCE}")")
ROOT_PATH=${SOURCE_PATH}/..

helm uninstall synology-csi