#!/bin/bash
set -ex
cd $(dirname "${BASH_SOURCE[0]}")

if [ ! -d "java-langserver" ]; then
    git clone git@github.com:sourcegraph/java-langserver.git java-langserver
else
    cd java-langserver && git fetch origin && git checkout origin/master && cd ..
fi

cd java-langserver
mvn clean compile assembly:single

cd ..
mv java-langserver/target/java-language-server.jar docker
cp java-langserver/add-android-support-libs.sh docker/

cd docker
if [ -d artifacts ]; then
    cd ./artifacts && git fetch origin && git checkout origin/master && cd -
else
    git clone --depth 1 https://github.com/sourcegraph/java-artifacts artifacts
fi

if [ -d android-sdk-jars ]; then
    cd ./android-sdk-jars && git fetch origin && git checkout origin/master && cd -
else
    git clone --depth 1 https://github.com/sourcegraph/android-sdk-jars
fi

docker build -t ${IMAGE-"xlang-java"} .
