.PHONY: build clean install go-format go-lint go-doc tf-doc

###
## Provider Variables
##
PROVIDER_PLUGINS_DIR = ~/.terraform.d/plugins
PROVIDER_HOSTNAME = truemark.io
PROVIDER_NAMESPACE = terraform
PROVIDER_NAME = truemark-confluent-cloud
PROVIDER_PROD_NAME = terraform-provider-$(PROVIDER_NAME)
PROVIDER_VERSION = 1.0.0
PROVIDER_OS_ARCH = darwin_amd64
PROVIDER_BUILT_FILENAME = $(PROVIDER_PROD_NAME)_$(PROVIDER_VERSION)_$(PROVIDER_ARCHOS)

###
## Build Variables
##
BUILD_COMPILER = go
BUILD_COMMAND = build

BUILD_OPTS_OUTDIR = build
BUILD_OPTS_VERBOSE = -v -work -x 
BUILD_OPTS_INSTALL_DEPS = -i 
BUILD_OPTS_REBUILD = -a 
BUILD_OPTS_MODULE = -mod=mod

BUILD_MKDIR = mkdir
BUILD_OPTS = $(BUILD_OPTS_VERBOSE)
BUILD_OUT = ./$(BUILD_OPTS_OUTDIR)/$(PROVIDER_PROD_NAME)

###
## Clean Commands
CLEAN_CMD = rm
CLEAN_OPTS = -rf 

###
## Install Commands
INSTALL_MK_CMD = mkdir
INSTALL_MK_OPTS = -p
INSTALL_MV_CMD = mv
INSTALL_MV_OPTS =
INSTALL_LOCATION = $(PROVIDER_PLUGINS_DIR)/$(PROVIDER_HOSTNAME)/$(PROVIDER_NAMESPACE)/$(PROVIDER_NAME)/$(PROVIDER_VERSION)/$(PROVIDER_OS_ARCH)


build:
	$(BUILD_MKDIR) ./$(BUILD_OPTS_OUTDIR)
	$(BUILD_COMPILER) $(BUILD_COMMAND) $(BUILD_OPTS) -o $(BUILD_OUT)


clean:
	$(CLEAN_CMD) $(CLEAN_OPTS) $(BUILD_OPTS_OUTDIR)


install:
	## We produce this in order to run to run:
	##    PLUGINS_DIR/HOSTNAME/NAMESPACE/NAME/VERSION/OS_ARCH/PRODUCT_NAME
	## That will look something like:
	##    ~/.terraform.d/plugins
	##    └── truemark.io
	##        └── terraform
	##            └── truemark-confluent-cloud
	##                └── 1.0.0
	##                    └── darwin_amd64
	##                        └── terraform-provider-truemark-confluent-cloud
	$(INSTALL_MK_CMD) $(INSTALL_MK_OPTS) $(INSTALL_LOCATION)
	$(INSTALL_MV_CMD) $(INSTALL_MV_OPTS) $(BUILD_OUT) $(INSTALL_LOCATION)


go-format:
	go fmt ./...


go-lint:


go-doc:


tf-doc:
