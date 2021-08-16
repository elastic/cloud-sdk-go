#!/bin/bash

SDK_PATH=$(pwd)
BRANCH=ms-${1}
SWAGGER_DEF_PATH=scala-services/adminconsole/src/main/resources
GIT_CLOUD_REPO=https://github.com/elastic/cloud
REPO_PATH=/tmp/cloud

if [[ -d ${REPO_PATH} ]]; then
    cd ${REPO_PATH}
    git fetch
    git checkout ${BRANCH}
    git pull
else
    git clone --single-branch -b ${BRANCH} ${GIT_CLOUD_REPO} ${REPO_PATH}
fi

cp ${REPO_PATH}/${SWAGGER_DEF_PATH}/apidocs.json ${SDK_PATH}/api/
cp ${REPO_PATH}/${SWAGGER_DEF_PATH}/apidocs-user.json ${SDK_PATH}/api/
