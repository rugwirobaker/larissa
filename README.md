[![Build Status](https://cloud.drone.io/api/badges/rugwirobaker/larissa/status.svg)](https://cloud.drone.io/rugwirobaker/larissa)
[![codebeat badge](https://codebeat.co/badges/b8191688-451a-4562-b1c6-9222dc589dca)](https://codebeat.co/projects/github-com-rugwirobaker-larissa-master)

# Larissa
Larissa is an elegent image server accessible via a S3 like HTTP API.

## Project Status
As of now larissa is still a prototype with the design principles not yet set.

## Goals
* support more backends(mongodb, minio, CloudStore, S3, etc...)
* filter/limit file mimetype (png, jpeg, gif initialy)
* multifile upload
* parameter configured processing
* access sdk for major languages(Go, Js, Python, Php,...)

## Run your own

**Natively**

If you're inside GOPATH, make sure GO111MODULE=on, if you're outside GOPATH, then Go Modules are on by default. 
The main package is inside cmd and is run like any go project as follows:

1. clone the repository:

    ```$ git clone https://github.com/rugwirobaker/larissa.git```

3. cd into the larissa directory:
    
    ```$ cd larissa```
2. build the binary:

    ```$ make build```
3. the outputed binary is saved at the bin directory, and tou can run it with

    ```$ ./bin/larissa```
4. You can also pass it an optional configuraion file:

    ```$ ./bin/larissa --config_file "path/config_file"```

**using docker**

using docker compose this options does the following for you:
1. build the larissa binary from scratch
2. add larissa configuration
3. run larissa server

For this to work you need `docker` and `docker-compose` setup on your computer

```$ make dev```

To rebuild the docker image

```$ make dev-build```

By default in all cases larissa is listening at:

```localhost:3000```
## usage

| **function**                | **Endpoint**            | **Options**                   | **Method**  | 
|:---------------------       |:----------------------- |:------------------------------|:------------|
| larissa buiild info         |`"/build"`               |                               | **GET**     |
| server health status        |`"/health"`              |                               | **GET**     | 
| upload image                |`"/put/:bucket"`         |Form Field:`image:"image_name"`| **PUT**     |
| list alll images a in bucket|`"/list/:bucket"`        |                               | **GET**     |
| dowload image               |`"/get/:bucket/:image"`  |                               | **GET**     |           
| delete image                |`"/del/:bucket/:image"`  |                               | **DELETE**  |           
| verify image existance      |`"/exists:bucket/:image"`|                               | **GET**     |


## Development
Tips:

* the errors package follows the example of [upspin project's errors](https://commandcenter.blogspot.com/2017/12/error-handling-in-upspin.html)
## Contributing
As of now PRs are not yet accepted. The project's direaction is still shaping up
