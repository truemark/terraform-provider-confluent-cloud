#!/bin/bash


###
## TerraForm Variables
##
export TF_PLUGINS_DIR="~/.terraform.d/plugins"


###
## Build Variables
##
export BUILD_DIR="build"
## CLI Options to the GoLang Compiler
export VERBOSE=" -v -work -x "
export INSTALL_DEPS=" -i "
export REBUILD=" -a "
export MOD=" -mod=mod "   # (readonly | vendor | mod)


###
## Provider Variables
##
export PROVIDER_ORG="truemark"
export PROVIDER_NAME="truemark-confluent-cloud"
export PROVIDER_TYPE="provider"
export PROVIDER_PROD_NAME="terraform-$PROVIDER_TYPE-$PROVIDER_NAME"
export PROVIDER_VERSION="1.0.0"



###
## Build Variables
##
export BUILD_OUTPUT="./$BUILD_DIR/$PROVIDER_PROD_NAME"
export TF_INSTALL_DIR="$TF_PLUGINS_DIR/$PROVIDER_ORG/$PROVIDER_NAME/$PROVIDER_VERSION"



export CLEAN_OPTS="-i -x -r -cache -modcache -testcache"