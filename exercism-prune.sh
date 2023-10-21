#!/usr/bin/env bash

find ./exercism -type d -name '.exercism' -exec rm -rf {} +
find ./exercism -type f -name '*.md' -delete
