#!/bin/bash

set -xe


function _filter_undeploy_conditions() {
    # Pull requests shouldn't try to deploy
    if [ "$TRAVIS_PULL_REQUEST" != "false" ]; then
        echo "Skipping deploy codes for pull request"
        exit 0
    fi

    # Commits to other branches except master shouldn't try to deploy
    if [ "$TRAVIS_BRANCH" != "master" ]; then
        if [ -z "$TRAVIS_TAG" ]; then
            echo "Skipping deploy code for non master branch commits"
            exit 0
        fi
    fi
}


function _load_ssh_key() {
    # Hide keys
    set +x
    openssl aes-256-cbc -K $encrypted_04c5d849a700_key \
        -iv $encrypted_04c5d849a700_iv \
        -in automation/travis_rsa.go_ovirt.enc \
        -out automation/travis_rsa.go_ovirt -d
    # Restore command output
    set -x

    eval "$(ssh-agent)"
    chmod 0600 automation/travis_rsa.go_ovirt
    ssh-add automation/travis_rsa.go_ovirt
}


function _init_git_config() {
    git config --global user.email "travis@travis-ci.org"
    git config --global user.name "GooVirtRobot@TravisCI"
}


function _clone_remote_master() {
    # Clone codes to local
    git clone git@github.com:oVirt/go-ovirt.git go-ovirt
}


function _deploy_to_master() {
    # Remove the original files
    rm -fr go-ovirt/*.go go-ovirt/README.md go-ovirt/examples

    # Copy newly generated codes to override the original
    cp -r sdk/ovirtsdk/* go-ovirt/
    cp -r sdk/examples go-ovirt/

    pushd go-ovirt
    # Push only if there are changes
    if git status --porcelain 2>/dev/null | grep -E "^??|^M"
    then
      git add -A
      git commit --message "Generator commit ID: ${TRAVIS_COMMIT:0:7} with message: $TRAVIS_COMMIT_MESSAGE. Travis build: $TRAVIS_BUILD_NUMBER."
      git push origin HEAD:master
    fi
    popd
}


function _deploy_to_tag() {
    pushd go-ovirt
    git tag -a $TRAVIS_TAG -m "New version release: $TRAVIS_TAG"
    git push origin $TRAVIS_TAG
    popd
}


function _deploy() {
    if [ -z "$TRAVIS_TAG" ]
    then
        _deploy_to_master
    else
        _deploy_to_tag
    fi
}


_filter_undeploy_conditions
_load_ssh_key
_init_git_config
_clone_remote_master
_deploy