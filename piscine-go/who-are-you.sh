#! bin/bash/

curl -s https://learn.01founders.co/assets/superhero/all.json | jq '.[52] | .name'
