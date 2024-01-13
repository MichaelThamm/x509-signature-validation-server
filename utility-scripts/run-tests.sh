#!/bin/bash

TEST_DIR="test-scripts"
for file in $TEST_DIR/*.sh; do
    echo "$file"
    curl -X POST --data-binary @"$file" http://localhost:8080/execute
done