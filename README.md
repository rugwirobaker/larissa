[![Build Status](https://cloud.drone.io/api/badges/rugwirobaker/larissa/status.svg)](https://cloud.drone.io/rugwirobaker/larissa)

# Larissa
Larissa is an elegent image server accessible via a S3 like HTTP API.

## Project Status
As of now larissa is still a prototype with the design principles not yet set.

## Goals
* filter/limit file mimetype (png, jpeg, gif initialy)
* multifile upload
* parameter configured processing

## Run your own
**steps to run it on your machine(linux)**
1. `git clone github.com/rugwirobaker/larissa`
2. `make build`
3. `./bin/larissa`

## usage

| **function**          | **Endpoint**            | **Options** | **Details** | 
|:--------------------- |:----------------------- |:------------|:------------|
| upload image          |`"/put/:bucket"`         |             |             |
| dowload image         |`"/get/:bucket/:image"`  |             |             |           
| delete image          |`"/del/:bucket/:image"`  |             |             |           
| verify image existance|`"/exists:bucket/:image"`|             |             |
| larissa buiild info   |`"/build"`               |             |             |

