#!/bin/sh
docker run -it --rm -v `pwd`:/go-system-programming/17 --platform linux/amd64 --name debian gcc:11-bullseye
