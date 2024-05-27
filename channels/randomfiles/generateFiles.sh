#!/bin/bash

# So I don't fill the repo with tons of nonsense, I have this
# bash script that will generate 300 5M files filled with random
# data that can be used for hashing or generally working with.
# You can delete when you'd like.

for i in $(seq 1 300); do
    dd if=/dev/random of=./rfile$i bs=1M count=5
done
