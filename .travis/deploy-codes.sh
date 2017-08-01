#!/bin/bash

# Pull requests and commits to other branches shouldn't try to deploy
if [ "$TRAVIS_PULL_REQUEST" != "false" -o "$TRAVIS_BRANCH" != "master" ]; then
    echo "Skipping deploy codes"
    exit 0
fi

cd ./sdk/ovirtsdk4/

git init

# Copy examples/ and push into go-ovirt repository
cp -r ../examples ./

git config --global user.email "travis@travis-ci.org"
git config --global user.name "Travis CI"

git remote add origin https://${GH_TOKEN}@github.com/imjoey/go-ovirt.git

git add -A 

git commit --message "Generator commit ID: ${TRAVIS_COMMIT:0:7} with message: $TRAVIS_COMMIT_MESSAGE. Travis build: $TRAVIS_BUILD_NUMBER."

git push -f --quiet --set-upstream origin master
