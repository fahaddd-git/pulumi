#!/bin/bash

set -eo pipefail
set -x

# shellcheck disable=SC2043
for i in "pascal v0.0.1"; do
  set -- $i # treat strings in loop as args
  PULUMI_LANG="$1"
  TAG="$2"

  LANG_DIST="goreleaser-lang/${PULUMI_LANG}"
  mkdir -p "$LANG_DIST"
  cd "${LANG_DIST}"

  rm -rf ./*
  gh release download "${TAG}" --repo "pulumi/pulumi-${PULUMI_LANG}" --pattern '*.tar.gz' || true
  gh release download "${TAG}" --repo "pulumi/pulumi-${PULUMI_LANG}" --pattern '*.zip' || true

  for DIST_OS in darwin linux windows; do
    for i in "amd64 x64" "arm64 arm64"; do
      set -- $i # treat strings in loop as args
      DIST_ARCH="$1"
      RENAMED_ARCH="$2" # goreleaser in pulumi/pulumi renames amd64 to x64

      OUTDIR="$DIST_OS-$RENAMED_ARCH"

      mkdir -p $OUTDIR
      find '.' -name "*-$DIST_OS-$DIST_ARCH.tar.gz" -print0 -exec tar -xzvf {} -C $OUTDIR \;
      find '.' -name "*-$DIST_OS-$DIST_ARCH.zip" -print0 -exec unzip {} -d $OUTDIR \;
    done
  done

  cd ../..
done