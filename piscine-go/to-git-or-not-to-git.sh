#!/bin/sh

curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"jaali\"}}){id}}"}' | cut -b "24-27"

