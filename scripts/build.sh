#!/usr/bin/env sh

# Code generated by shipbuilder init 1.21.0. DO NOT EDIT.

if [ ! -f "./scripts/check.sh" ]; then
  cd $(command dirname -- "$(command readlink -f "$(command -v -- "$0")")")/..
fi

if [ -f "./features/shipbuilder/.env" ]; then
  . ./features/shipbuilder/.env
fi

what="${shipbuilder_build:-go}"

for i in $what; do
  if [ -f "./scripts/build$i.sh" ]; then
    . "./scripts/build$i.sh" $@
  fi
done
