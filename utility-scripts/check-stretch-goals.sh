#!/bin/bash

KEYS_DIR="keys"

echo "Stretch Goal # 1 - Check the certificate extension for code signing:"
curl -X POST --data-binary @./test-scripts/1.sh http://localhost:8080/execute?code-sign=true

echo "Stretch Goal # 2 - Accept concurrent requests:"
curl -X POST --data-binary @./test-scripts/1.sh http://localhost:8080/execute?id=1 & \
curl -X POST --data-binary @./test-scripts/2.sh http://localhost:8080/execute?id=2

echo "Stretch Goal # 3 - Verify signature from a set of certificates:"
curl -X POST --data-binary @./test-scripts/1.sh http://localhost:8080/execute?key-dir=true