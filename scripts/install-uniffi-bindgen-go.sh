#!/usr/bin/env bash
set -euo pipefail

UNIFFI_BINDGEN_GO_TAG="${UNIFFI_BINDGEN_GO_TAG:-v0.6.0+v0.30.0}"
UNIFFI_BINDGEN_GO_REPO="${UNIFFI_BINDGEN_GO_REPO:-https://github.com/NordSecurity/uniffi-bindgen-go}"

if command -v uniffi-bindgen-go >/dev/null 2>&1; then
    INSTALLED_VERSION="$(uniffi-bindgen-go --version 2>/dev/null | awk '{print $2}')"
    REQUESTED_VERSION="${UNIFFI_BINDGEN_GO_TAG#v}"

    if [ "${INSTALLED_VERSION}" = "${REQUESTED_VERSION}" ]; then
        echo "uniffi-bindgen-go already available: $(command -v uniffi-bindgen-go) (${INSTALLED_VERSION})"
        exit 0
    fi

    echo "Updating uniffi-bindgen-go from ${INSTALLED_VERSION:-unknown} to ${UNIFFI_BINDGEN_GO_TAG}"
fi

cargo install uniffi-bindgen-go \
    --git "${UNIFFI_BINDGEN_GO_REPO}" \
    --tag "${UNIFFI_BINDGEN_GO_TAG}" \
    --locked \
    --force

echo "Installed uniffi-bindgen-go ${UNIFFI_BINDGEN_GO_TAG}"
