#!/usr/bin/env bash

### THIS FILE WAS PRODUCED USING A STUPID FUCKING AI BOT
# IF YOU FIND ANY MISTAKES, PLEASE REPORT IT

# Install nepdate via `go install`.
#
# Usage:
#   curl -sSL https://raw.githubusercontent.com/UdeshyaDhungana/nepdate/main/install.sh | bash
#
# Override the version installed with:
#   NEPDATE_VERSION=v1.1.0 curl -sSL .../install.sh | bash

set -euo pipefail

MODULE="github.com/UdeshyaDhungana/nepdate"
REF="${NEPDATE_VERSION:-latest}"

if ! command -v go >/dev/null 2>&1; then
  echo "Error: Go is not installed." >&2
  echo "nepdate installs via 'go install', so the Go toolchain is required first." >&2
  echo "Install Go from https://go.dev/doc/install, then re-run this script." >&2
  exit 1
fi

echo "Installing nepdate (${REF}) via go install..."
# GOPROXY=direct bypasses the public Go module proxy (proxy.golang.org) and
# fetches straight from GitHub. This avoids stale @latest resolution if a
# tag was ever deleted/recreated, since the public proxy caches version info
# and doesn't reliably invalidate it.
GOPROXY=direct go install "${MODULE}@${REF}"

GOBIN="$(go env GOBIN)"
if [ -z "$GOBIN" ]; then
  GOBIN="$(go env GOPATH)/bin"
fi

echo ""
echo "nepdate installed to: ${GOBIN}/nepdate"

case ":${PATH}:" in
  *":${GOBIN}:"*)
    ;;
  *)
    echo ""
    echo "Note: ${GOBIN} is not on your PATH."
    echo "Add this to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
    echo "  export PATH=\"\$PATH:${GOBIN}\""
    ;;
esac

echo ""
echo "Run 'nepdate --help' to get started."
