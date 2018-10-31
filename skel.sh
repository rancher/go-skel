#!/bin/bash
set -e

if [ "$#" != 1 ] || [ "$1" = "-h" ] || [ "$1" = "--help" ]; then
    echo Usage: $0 NEW-APP-PACKAGE
    exit 1
fi

if ! echo $1 | grep -q /; then
    echo "use full package name like github.com/rancher/widget"
    exit 1
fi

BASE=$(dirname $0)

PKG=$1
APP=$(basename $1)
REPO=$(basename $(dirname $1))
IMAGE=$REPO/$APP
FILES="
./Dockerfile.dapper
./.dockerignore
./.drone.yml
./.gitignore
./LICENSE
./main.go
./Makefile
./package/Dockerfile
./pkg/controllers/foo/controller.go
./pkg/server/server.go
./README.md.in
./scripts/build
./scripts/ci
./scripts/entry
./scripts/package
./scripts/release
./scripts/test
./scripts/validate
./scripts/version
./types/apis/some.api.group/v1/types.go
./types/codegen/cleanup/main.go
./types/codegen/main.go
./vendor.conf
"


mkdir -p $APP

for i in $FILES; do
    mkdir -p $APP/$(dirname $i)
    echo Creating $APP/$i
    sed \
        -e "s!%REPO%!$REPO!g" \
        -e "s!%PKG%!$PKG!g" \
        -e "s!%APP%!$APP!g" \
        -e "s!%IMAGE%!$IMAGE!g" \
        $BASE/$i > $APP/$i
    if echo $i | grep -q scripts; then
        echo chmod +x $APP/$i
        chmod +x $APP/$i
    fi
done

cd ./$APP
mv README.md.in README.md
make deps
./.dapper -m bind env go generate
./.dapper -m bind env chown -R $(id -u) types
make deps

rm vendor.conf

git init
git add -A
git commit -m "Initial Commit"
while ! git gc; do
    sleep 2
done

make

echo Created $APP in ./$APP
