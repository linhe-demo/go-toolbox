#!/bin/bash

git pull https://github.com/linhe-demo/go-toolbox.git

# shellcheck disable=SC2164
cd toolbox

go build

supervisorctl restart toolbox-service

