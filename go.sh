#!/bin/bash

curl https://dl.google.com/go/go1.18.linux-amd64.tar.gz -o go.tar.gz

tar xvf go.tar.gz && sudo rm -rf /usr/local/go && \
sudo mv go /usr/local && rm -rf go.tar.gz
