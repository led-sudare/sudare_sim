#!/bin/sh
docker build ./ -t led-sudare-simulator
docker run -t --init --name led-sudare-simulator -v `pwd`:/go/src/simulator/ -p 2345:2345 led-sudare-simulator