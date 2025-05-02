# test/test_picChan.sh
#!/bin/bash
set -e

source "$(dirname "$0")/../.env"

curl -X POST http://localhost:8080/picchan \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "text=:an_amiya_keepworking:&token=$PICCHAN_VERIFY_TOKEN&user_name=tester&channel_id=$TEST_CHANNEL_ID"
