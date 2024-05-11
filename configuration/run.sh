#!/bin/bash
set -e

ENVIRONMENT=dev
ENVIRONMENT_DIRECTORY=environments/$ENVIRONMENT

ansible-playbook \
  -i $ENVIRONMENT_DIRECTORY \
  site.yaml