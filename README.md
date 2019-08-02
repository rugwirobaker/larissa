# Larissa
Larissa is an elegent image server accessible via a S3 like HTTP API.

## Project Status
As of now larissa is still a prototype with the design principles not yet set.

## Try it (HTTP routes)
**steps to run it on your machine(linux)**
1. `git clone github.com/rugwirobaker/larissa`
2. `make build`
3. `./bin/larissa`

**availabe routes(not implemented)**
* `/put`: is meant to upload new images              
* `/get`: is meant to download an image
* `/del`: is meant to delete an image
* `/exists`: verifies whether an image has been saved
* `/build`: gives larissa build information(version, buildDate)