#!/bin/bash

set -e

DIR=$(cd $(dirname ${0})/.. && pwd)
cd ${DIR}

VERSION=${VERSION:-0.1.0}

# Compile
./scripts/compile.sh

# Zip all pacakges
mkdir -p ./pkg/dist

for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
    PLATFORM_NAME=$(basename ${PLATFORM})
    ARCHIVE_NAME=gitnpm_${VERSION}_${PLATFORM_NAME}

    if [ $PLATFORM_NAME = "dist" ]; then
        continue
    fi

    pushd ${PLATFORM}
    zip ${DIR}/pkg/dist/${ARCHIVE_NAME}.zip ./*
    popd
done
