#!/bin/bash

UTILITY_DIR="utility-scripts"

./$UTILITY_DIR/create-certs.sh
./$UTILITY_DIR/sign-scripts.sh
./$UTILITY_DIR/build-server-binary.sh

# Start the server in the background
./x509-validation-server &
# Capture the PID
server_pid=$!
sleep 5

./$UTILITY_DIR/run-tests.sh

kill $server_pid