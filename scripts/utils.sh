#!/usr/bin/env bash

case "${OSTYPE}" in
linux*)
  HOSTOS=linux
  ;;
darwin*)
  HOSTOS=darwin
  ;;
*)
  echo "unsupported HOSTOS=${OSTYPE}" 1>&2
  exit 1
  ;;
esac

_hostarch=$(uname -m)
case "${_hostarch}" in
*64*)
  HOSTARCH=amd64
  ;;
*386*)
  HOSTARCH=386
  ;;
*686*)
  HOSTARCH=386
  ;;
*)
  echo "unsupported HOSTARCH=${_hostarch}" 1>&2
  exit 1
  ;;
esac

checksum_sha256() {
  if command -v shasum >/dev/null 2>&1; then
    shasum -a 256 -c "${1}"
  elif command -v sha256sum >/dev/null 2>&1; then
    sha256sum -c "${1}"
  else
    echo "missing shasum tool" 1>&2
    return 1
  fi
}

get_shasum() {
  local present_shasum=''
  if command -v shasum >/dev/null 2>&1; then
    present_shasum=$(shasum -a 256 "${1}"| awk -F' ' '{print $1}')
  elif command -v sha256sum >/dev/null 2>&1; then
    present_shasum=$(sha256sum "${1}" | awk -F' ' '{print $1}')
  else
    echo "missing shasum tool" 1>&2
    return 1
  fi
  echo "$present_shasum"
}

ensure_py3_bin() {
  # If given executable is not available, the user Python bin dir is not in path
  # This function assumes the executable to be checked was installed with
  # pip3 install --user ...
  if ! command -v "${1}" >/dev/null 2>&1; then
    echo "User's Python3 binary directory must be in \$PATH" 1>&2
    echo "Location of package is:" 1>&2
    pip3 show ${2:-$1} | grep "Location"
    echo "\$PATH is currently: $PATH" 1>&2
    exit 1
  fi
}

ensure_py3() {
  if ! command -v python3 >/dev/null 2>&1; then
    echo "python3 binary must be in \$PATH" 1>&2
    exit 1
  fi
  if ! command -v pip3 >/dev/null 2>&1; then
    curl -SsL https://bootstrap.pypa.io/get-pip.py -o get-pip.py
    python3 get-pip.py --user
    rm -f get-pip.py
    ensure_py3_bin pip3
  fi
}
