curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"AndrewLondon\"}}){id}}"}' | jq --raw-output '.data.user[0].id'
