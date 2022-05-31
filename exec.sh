#!/bin/bash

if [[ ! -d data ]]; then
    mkdir -p data
    curl -o data/data.deb -L "https://github.com/tiechui1994/actions/releases/download/nginx_1.20.2/nginx_1.20.2_ubuntu_16.04_amd64.deb"
    curl -o data/nginx.deb -L "https://github.com/tiechui1994/actions/releases/download/nginx_1.20.2/nginx_1.20.2_ubuntu_18.04_amd64.deb"
fi


for (( i=0 ; i<100; i++)); do
    key="$(date +%s)_$i"
    data="@data/data.deb"
    if [[ $(expr $i % 2) -eq 0 ]]; then
        data="@data/nginx.deb"
    fi

    echo "$key $data"
    curl -X POST http://127.0.0.1:8080/api/add?key=$key --data $data

    sleep 0.5

    curl -X GET http://127.0.0.1:8080/api/get?key=$key -o /dev/null -s -w "http: %{http_code}\nsize_download: %{size_download} bytes\n\n"
done

