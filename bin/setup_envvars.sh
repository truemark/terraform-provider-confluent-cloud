#!/bin/bash


###
## TerraForm Variables
##
export TF_PLUGINS_DIR="/Users/brian/.terraform.d/plugins"

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
export PROVIDER_HOSTNAME="truemark.io"
export PROVIDER_NAMESPACE="terraform"
export PROVIDER_NAME="truemark-confluent-cloud"
export PROVIDER_PROD_NAME="terraform-provider-$PROVIDER_NAME"
export PROVIDER_VERSION="1.0.0"
export PROVIDER_OS_ARCH="darwin_amd64"
export PROVIDER_BUILT_FILENAME="$PROVIDER_PROD_NAME"_"$PROVIDER_VERSION"_"$PROVIDER_ARCHOS"

###
## Build Variables
##
# export BUILD_OUTPUT="./$BUILD_DIR/$PROVIDER_BUILT_FILENAME"
# export TF_INSTALL_DIR="$TF_PLUGINS_DIR/$PROVIDER_ORG/$PROVIDER_NAME/$PROVIDER_VERSION/$PROVIDER_ARCHOS"


export CLEAN_OPTS="-i -x -r -cache -modcache -testcache"