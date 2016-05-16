#!/bin/bash

#
# Documentation:
# This script generates the HTML and place them at the correct location.
#
bundle exec middleman build --clean
cp -r build/* .
rm -rf build
