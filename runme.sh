#!/bin/bash

docker build -t collatz/server .
docker run --rm -d -p 8080:8080 collatz/server