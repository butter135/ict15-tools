# test/tokumei.sh
#!/bin/bash
set -e

source "$(dirname "$0")/../.env"

curl -X POST http://localhost:8080/tokumei \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "text=:an_amiya_keepworking:&token=$TOKUMEI_VERIFY_TOKEN1"
