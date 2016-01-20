# simple-pixel-server
A simple pixel server.

## Build

    go get
    go build
  
## Run the server

    ./simple-pixel [port]
  
The default port is 8080. The server accepts requests to the path `/tracking.gif` with query parameters `uid`, `domain`, `lat`, and `lon`. The service responds with a small image. The service logs request to stdout in the following format:

    {"Timestamp":1453289550,"Uid":"123112352455","Domain":"test1.com","Geo":{"Latitude":37.7576171,"Longitude":-122.5776844}}

