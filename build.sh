#!/bin/sh

go build

if test ! -f regression-test.sh
then
    wget https://raw.githubusercontent.com/sverrehu/ccalc/refs/heads/main/regression-test.sh
    chmod 755 regression-test.sh
fi
./regression-test.sh
