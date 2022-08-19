#!/usr/bin/env sh

# Code generated by shipbuilder init 1.21.0. DO NOT EDIT.

system_goos=$(go env GOOS)
system_goarch=$(go env GOARCH)

if [ ! -f "./scripts/check.sh" ]; then
  cd $(command dirname -- "$(command readlink -f "$(command -v -- "$0")")")/..
fi

if [ -f "./features/shipbuilder/.env" ]; then
  . ./features/shipbuilder/.env
fi

. ./scripts/check.sh

export CGO_ENABLED="${shipbuilder_go_cgo_enabled:-0}"
export GOPRIVATE="go.ketch.com"

if [ -d "./cmd" ]; then
  check go

  set -e

  if [ -z "$*" ]; then
    oses="${shipbuilder_go_os:-darwin}"
  else
    oses=$*
  fi

  cmds=$(find cmd -type d -maxdepth 1 -mindepth 1 | cut -f2- -d/)

  for os in $oses; do
    export GOOS="$os"

    case "$GOOS" in
    "linux")
      arches="${shipbuilder_go_linux_arch:-amd64}"
      ;;

    "windows")
      arches="${shipbuilder_go_windows_arch:-amd64}"
      ;;

    "darwin")
      arches="${shipbuilder_go_darwin_arch:-amd64 arm64}"
      ;;

    *)
      echo "$GOOS is not supported"
      ;;
    esac

    for arch in $arches; do
      export GOARCH="$arch"

      echo "Building for $GOOS $GOARCH..."
      for i in $cmds; do
        echo "|- $i"
        $go build -o "./.build/$GOOS-$GOARCH/$i" "./cmd/$i/main.go"

        if [ "$GOOS" = "darwin" -a "$GOARCH" = "$system_goarch" ]; then
          echo "|--> copied to ~/go/bin"
          mkdir -p $HOME/go/bin
          cp "./.build/$GOOS-$GOARCH/$i" "$HOME/go/bin/$i" || true
        fi
      done

      if [ -d "./test/smoketest" ]; then
        # Build the smoketest
        echo "\- smoketest"
        $go test -c -o "./.build/$GOOS-$GOARCH/smoke.test" -json -failfast --tags smoke ./test/smoketest/...
      fi
    done
  done
else
  go build ./...
fi
