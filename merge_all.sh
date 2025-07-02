#!/bin/bash

function check_tools() {
    if ! command -v git /dev/null; then
        echo "Coult not find Git in your operating system. Aborting."
        exit 1
    fi 
}

function merge_all() {
    git switch master
    git merge dev master
    git push origin master

    git switch stable
    git merge master stable
    git push origin stable

    git switch dev
}

function main() {
    check_tools
    merge_all
}

main