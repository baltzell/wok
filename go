#!/usr/bin/bash

export PATH=$PWD/coatjava/coatjava/bin:$PATH

# get the electron efficiency:
trutheff ./out_electronproton.hipo | \
    grep ^{ | \
    jq '.effs["11"][0]'

