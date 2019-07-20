#!/bin/sh
cname=`cat ./cname`
docker build ./ -t $cname
docker run -t --init --name $cname -v `pwd`:/work -p 2345:2345 --restart=always $cname
