# klaka
Go lambda job for tweeting from a user

## Setup

### Export Twitter credentials
 `export TWITTER_CONSUMER_KEY=<twitter_consumer_key>`
 
 `export TWITTER_CONSUMER_SECRET=<twitter_consumer_secret>`
 
  `export TWITTER_ACCESS_TOKEN=<twitter_access_key>`
  
  `export TWITTER_ACCESS_TOKEN_SECRET=<twitter_access_token_secret>`
  
### Start the db
    docker-compose up -d
 
## Run Tests
  `go test ./...`
