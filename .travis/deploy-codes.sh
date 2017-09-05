#!/bin/bash

# Pull requests and commits to other branches shouldn't try to deploy
if [ "$TRAVIS_PULL_REQUEST" != "false" -o "$TRAVIS_BRANCH" != "master" ]; then
    echo "Skipping deploy codes"
    exit 0
fi

# Pull current codes
mkdir -p ./sdk/ovirtsdk4-git/
cd ./sdk/ovirtsdk4-git/
git init

git config --global user.email "travis@travis-ci.org"
git config --global user.name "GooVirtRobot@TravisCI"

git remote add origin https://${GH_TOKEN}@github.com/imjoey/go-ovirt.git

git pull origin master

# Use newly generated codes to override the pulled ones
rm -fr *.go README.md
cp -r ../ovirtsdk4/* ./

# Copy examples/ and push into go-ovirt repository
rm -fr ./examples
cp -r ../examples ./

# Push back to github
git add -A 

git commit --message "Generator commit ID: ${TRAVIS_COMMIT:0:7} with message: $TRAVIS_COMMIT_MESSAGE. Travis build: $TRAVIS_BUILD_NUMBER."

git push origin master
