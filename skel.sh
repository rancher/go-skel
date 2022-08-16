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
./.golangci.json
./.drone.yml
./.gitignore
./LICENSE
./main.go
./Makefile
./package/Dockerfile
./README.md.in
./scripts/boilerplate.go.txt
./scripts/build
./scripts/ci
./scripts/default
./scripts/entry
./scripts/package
./scripts/release
./scripts/test
./scripts/validate
./scripts/validate-ci
./scripts/version
./pkg/apis/some.api.group/v1/types.go
./pkg/codegen/cleanup/main.go
./pkg/codegen/main.go
./pkg/foo/controller.go
./pkg/version/version.go
./go.mod.in
./.github/workflows/stale.yml
./CODEOWNERS
./CODE_OF_CONDUCT
"

rm -rf $APP
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
mv -f go.mod.in go.mod
go mod download
go mod tidy
go mod verify
go generate
make .dapper
./.dapper -m bind goimports -w .
./.dapper -m bind rm -rf .cache dist bin

git init
git add -A
git commit -m "Initial Commit"
while ! git gc; do
    sleep 2
done

make ci

echo Created $APP in ./$APP
