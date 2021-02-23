#!/bin/bash

# Simple helper to run the command, communicate between us what was run for which operations, and to serve as a memory.

export TF_CLI_CONFIG_FILE=./examples/environment/01_create/.terraformrc
cd ./examples/environment/

terraform init

terraform apply

terraform state show 