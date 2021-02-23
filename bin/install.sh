source ./bin/setup_envvars.sh

echo "TF_INSTALL_DIR: " "$TF_INSTALL_DIR"
echo "BUILD_OUTPUT: " "$BUILD_OUTPUT"
mkdir -p -m +rw "$TF_INSTALL_DIR"
cp ./"$BUILD_OUTPUT" "$TF_INSTALL_DIR"

echo ""
echo ""
echo "Contents of "$TF_INSTALL_DIR""
ls -al $TF_INSTALL_DIR