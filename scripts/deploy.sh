#!/usr/bin/env bash
SCRIPTPATH="$( cd "$(dirname "$0")" ; pwd -P )"

echo "Provisioning..."
cd $SCRIPTPATH
cd ../terraform 
terraform taint google_compute_instance.stranger-services
terraform apply