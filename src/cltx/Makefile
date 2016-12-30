#!/bin/bash -xe

function getversion() {
    branch=`git rev-parse --abbrev-ref HEAD`
    commitid=`git rev-parse --short HEAD`
    builddate=`date +%Y%m%d-%H%M%S`
    echo $branch-$commitid-$builddate
}

cd `dirname $0`

# use go vendor
export GO15VENDOREXPERIMENT=1

version=`getversion`
pkgpath="version.go"

echo "build log-linux"
GOOS=linux GOARCH=amd64 go build -ldflags "-X main.VersionStr=${version}"  -o bin/log-linux ${pkgpath}

#echo "build log-osx"
#GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.VersionStr=${version}"  -o bin/log-osx ${pkgpath}

#echo "build log-win
#GOOS=windows GOARCH=386  go build -ldflags "-X main.versionStr=${version}"  -o bin/log-win ${pkgpath}

echo "done"
