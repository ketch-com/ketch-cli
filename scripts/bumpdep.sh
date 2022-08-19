#!/usr/bin/env sh

# Code generated by shipbuilder init 1.16.9. DO NOT EDIT.

if [ ! -f "./scripts/check.sh" ]; then
  cd $(command dirname -- "$(command readlink -f "$(command -v -- "$0")")")/..
fi

if [ -f "./features/shipbuilder/.env" ]; then
  . ./features/shipbuilder/.env
fi

if [ -f "go.mod" ]; then
  . ./scripts/check.sh go jq
fi

if [ -f "package.json" ]; then
  . ./scripts/check.sh ncu npm
fi

if [ "$1" != "--local" ]; then
  . ./scripts/check.sh git github

  if [ -n $($git status --porcelain --untracked-files=no) ]; then
    echo "ERROR: you have unstaged changes"
    exit 1
  fi

  $git checkout main
  $git pull

  if [ -n "$($git branch --list "$USER/chore/bump-deps")" ]; then
    $git branch -D $USER/chore/bump-deps
  fi
  $git checkout -b $USER/chore/bump-deps
fi

set -e

if [ -f "go.mod" ]; then
  export module=$($go mod edit -json | $jq -r .Module.Path)

  echo "module $module" > go.mod
  echo "" >> go.mod
  echo "go ${shipbuilder_go_version:-1.18}" >> go.mod
  echo "" >> go.mod
  rm go.sum

  $go get -u -t ./...
  $go mod tidy
  $git add go.mod go.sum
fi

if [ -f "package.json" ]; then
  $ncu -u --target latest
  $npm install
  $git add package.json package-lock.json
fi

if [ "$1" != "--local" ]; then
  ./scripts/build.sh

  $git commit -m "chore: bump dependencies"
  $git push -f
  $git checkout main
  $github pr create -H $USER/chore/bump-deps -t "chore: bump dependencies" -F ./features/shipbuilder/DEPENDENCY_PULL_REQUEST.md
fi
