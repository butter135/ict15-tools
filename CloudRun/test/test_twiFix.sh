# test/test_twiFix.sh
#!/bin/bash
set -e

source "$(dirname "$0")/../.env"

curl -X POST http://localhost:8080/twifix \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "text=$TEST_TWEET_URL&token=$TWIFIX_VERIFY_TOKEN1&user_name=tester&channel_id=$TEST_CHANNEL_ID"
