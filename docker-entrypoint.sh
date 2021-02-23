#!/bin/sh
set -ex
cmd='/app/bin/ketch'

if [ "$1" = 'validate' ]; then
  exec ${cmd} "$1" "$2" "$3"
elif [ "$1" = 'publish' ]; then
  exec ${cmd} "$@"
else
  exec ${cmd}
fi
