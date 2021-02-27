#!/usr/bin/env bash

PROBLEMS_URL=https://raw.githubusercontent.com/davidcorbin/euler-offline/master/project_euler_problems.txt
OUTPUT_FILE=project_euler_problems.txt

curl "$PROBLEMS_URL" --output "$OUTPUT_FILE"