#!/bin/bash
RUN_NAME=word_statistic
mkdir -p output/
export GO111MODULE=on && go build -o output/${RUN_NAME} -tags=jsoniter



