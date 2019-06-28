#!/bin/sh
rm /usr/bin/simulator
go build -o /usr/bin/simulator
exec /usr/bin/simulator
