#!/bin/bash

# Pull requests shouldn't try to deploy
if [ "$TRAVIS_PULL_REQUEST" != "false" ]; then
    echo "Skipping deploy codes for pull request"
    exit 0
fi

# Commits to other branches except master shouldn't try to deploy
if [ "$TRAVIS_BRANCH" != "master" ]; then
    if [ -z "$TRAVIS_TAG" ];then
        echo "Skipping deploy code for non master branch commits"
        exit 0
    fi
fi

# Pull current codes
mkdir -p ./sdk/ovirtsdk-git/
cd ./sdk/ovirtsdk-git/
git init

git config --global user.email "travis@travis-ci.org"
git config --global user.name "GooVirtRobot@TravisCI"

git remote add origin https://${GH_TOKEN}@github.com/imjoey/go-ovirt.git

git pull origin master

# Use newly generated codes to override the pulled ones
rm -fr *.go README.md
cp -r ../ovirtsdk/* ./

# Copy examples/ and push into go-ovirt repository
rm -fr ./examples
cp -r ../examples ./

# Push back to github
git add -A 

git commit --message "Generator commit ID: ${TRAVIS_COMMIT:0:7} with message: $TRAVIS_COMMIT_MESSAGE. Travis build: $TRAVIS_BUILD_NUMBER."

# For builds triggered by a tag, TRAVIS_BRANCH is the same as 
# the name of the tag (TRAVIS_TAG).
if [ "$TRAVIS_BRANCH" != "master" ];then
    git tag -a ${TRAVIS_TAG} -m "New version release: ${TRAVIS_TAG}"
fi

git push origin ${TRAVIS_BRANCH}
