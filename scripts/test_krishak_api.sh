#!/bin/bash
set -euo pipefail

echo "🔐 Logging in..."

LOGIN_RESPONSE=$(curl -s -w "\n%{http_code}" -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin"}')

HTTP_BODY=$(echo "$LOGIN_RESPONSE" | sed '$d')
HTTP_STATUS=$(echo "$LOGIN_RESPONSE" | tail -n1)

echo "Raw login response:"
echo "$HTTP_BODY"
echo "HTTP Status: $HTTP_STATUS"

if [ "$HTTP_STATUS" -ne 200 ]; then
  echo "❌ Login failed"
  exit 1
fi

TOKEN=$(echo "$HTTP_BODY" | jq -r '.token')
echo "✅ Token acquired"
echo "🔑 JWT Token:"
echo "$TOKEN"
echo

echo "📤 Creating Contractor..."
CONTRACTOR_RESPONSE=$(curl -s -X POST http://localhost:8080/contractors \
  -H "Authorization: $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"name": "Ravi Contractor", "village": "Rampur", "phone": "1234567890"}')
echo "$CONTRACTOR_RESPONSE" | jq .
CONTRACTOR_ID=$(echo "$CONTRACTOR_RESPONSE" | jq -r '.id')

echo
echo "📤 Creating Land..."
LAND_RESPONSE=$(curl -s -X POST http://localhost:8080/lands \
  -H "Authorization: $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"village": "Rampur", "khasra": "K123", "acre": 2.5}')
echo "$LAND_RESPONSE" | jq .
LAND_ID=$(echo "$LAND_RESPONSE" | jq -r '.id')

echo
echo "📤 Creating Contract..."
CONTRACT_RESPONSE=$(curl -s -X POST http://localhost:8080/contracts \
  -H "Authorization: $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{
    \"contractor_id\": \"$CONTRACTOR_ID\",
    \"land_id\": \"$LAND_ID\",
    \"start_date\": \"2025-01-01\",
    \"end_date\": \"2025-12-31\",
    \"expected_yield\": 150.5
  }")
echo "$CONTRACT_RESPONSE" | jq .

echo
echo "✅ All requests completed successfully."
