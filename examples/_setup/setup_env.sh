#!/bin/bash

###
## from: https://www.terraform.io/docs/cli/config/environment-variables.html
## Sets up our operating environment. 
# TF_LOG
# TF_LOG_PATH
# TF_INPUT
# TF_VAR_name
# TF_DATA_DIR
# TF_WORKSPACE
# TF_IN_AUTOMATION
# TF_REGISTRY_DISCOVERY_RETRY
# TF_REGISTRY_CLIENT_TIMEOUT
# TF_CLI_CONFIG_FILE
# TF_IGNORE

export TF_LOG=TRACE 
export TF_LOG_PATH="./tf-logfile.log"

export TRUEMARK_CONFLUENTCLOUD_USERNAME=
export TRUEMARK_CONFLUENTCLOUD_PASSWORD=
            