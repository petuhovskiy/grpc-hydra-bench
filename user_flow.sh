#!/bin/bash
source .env && docker-compose exec hydra hydra \
  token user \
  --endpoint http://127.0.0.1:4444/ \
  --client-id $CLIENT_ID \
  --client-secret $CLIENT_SECRET \
  --scope openid,offline \
  --port 5555