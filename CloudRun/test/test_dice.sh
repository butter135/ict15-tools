# test/test_dice.sh
#!/bin/bash
set -e

source "$(dirname "$0")/../.env"

curl -X POST http://localhost:8080/dice \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "text=5&token=$DICE_VERIFY_TOKEN"
