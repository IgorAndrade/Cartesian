# Cartesian API

The API receives a point and a distance and returns a JSON list of points that are within the distance of x, y, using the Manhattan distance method.

## Makefile commands

### Run all tests

    $ make test
### Run app
Run the app from source code,
but before running you need to add the environment variable `FILE_PATH` which indicates where the initial data is
 
    $ make run

### build the docker image
Generated image name is `cartesian:1.0` 
 
    $ make build
### Run app from docker
Generated container name is `app` and port `8080` 
 
    $ make docker_run

## How to user
Execute:

    $ make docker_run

CURL:

    $ curl 'localhost:8080/api/points?x=65&y=-70&distance=50'    