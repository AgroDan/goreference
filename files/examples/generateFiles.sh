#!/bin/bash

# So I don't fill the repo with tons of nonsense, I have this
# bash script that will generate 100 1K files filled with random
# data that can be used for hashing or generally working with.
# You can delete when you'd like.

for i in $(seq 1 100); do
    dd if=/dev/random of=./rfile$i bs=1K count=1
done