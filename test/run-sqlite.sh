#!/bin/sh

set -e

GTS_DB_TYPE=sqlite \
GTS_DB_ADDRESS=':memory:' \
go test ./... ${@}