# simple-pixel-server
A simple pixel server.

## Build

    go get
    go build
  
## Run the server

    ./simple-pixel [port]
  
The default port is 8080. The server accepts requests to the path `/tracking.gif` with query parameters `uid` and `domain`. The service responds with a small image. The service logs request to stdout in the following format:

    timestamp, uid, domain, latitude, longitude

This service makes use of https://github.com/oschwald/geoip2-golang to resolve geolocation data.
