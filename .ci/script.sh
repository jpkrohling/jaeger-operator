#!/bin/bash

CC_TEST_REPORTER_BIN=$(command -v cc-test-reporter)
if [ -z ${CC_TEST_REPORTER_BIN} ]; then
    curl -sL https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter
    chmod +x ./cc-test-reporter
    CC_TEST_REPORTER_BIN="./cc-test-reporter"
fi

${CC_TEST_REPORTER_BIN} before-build

make vendor install-tools ci
RT=$?
if [ ${RT} != 0 ]; then
    echo "Failed to build the operator."
    exit ${RT}
fi

${CC_TEST_REPORTER_BIN} after-build --exit-code ${RT}
