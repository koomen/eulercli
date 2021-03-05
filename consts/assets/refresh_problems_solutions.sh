#!/usr/bin/env bash

PROBLEMS_URL=https://raw.githubusercontent.com/davidcorbin/euler-offline/master/project_euler_problems.txt
PROBLEMS_OUTPUT_FILE=project_euler_problems.txt

curl "$PROBLEMS_URL" --output "$PROBLEMS_OUTPUT_FILE"

SOLUTIONS_URL=https://raw.githubusercontent.com/luckytoilet/projecteuler-solutions/master/Solutions.md
SOLUTIONS_OUTPUT_FILE=Solutions.md

curl "$SOLUTIONS_URL" --output "$SOLUTIONS_OUTPUT_FILE"