#!/bin/bash
source .env && docker-compose exec hydra hydra \
  clients create \
  --endpoint http://127.0.0.1:4445/ \
  --id $CLIENT_ID \
  --secret $CLIENT_SECRET \
  --grant-types authorization_code,refresh_token \
  --response-types code,id_token \
  --scope openid,offline \
  --callbacks http://127.0.0.1:5555/callback