#!/bin/bash

remove_executable() {
    rm -rf golab
}
trap remove_executable EXIT

go build . && ./golab $@