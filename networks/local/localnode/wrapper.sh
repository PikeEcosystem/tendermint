#!/usr/bin/env sh

##
## Input parameters
##
BINARY=/tendermint/${BINARY:-tendermint}
ID=${ID:-0}
LOG=${LOG:-tendermint.log}

##
## Assert linux binary
##
if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'tendermint' E.g.: -e BINARY=tendermint_my_test_version"
	exit 1
elif ! [ -x "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") is not executable."
	exit 1
fi

##
## Run binary with all parameters
##
export OCHOME="/tendermint/node${ID}"

if [ -d "`dirname ${OCHOME}/${LOG}`" ]; then
  "$BINARY" "$@" | tee "${OCHOME}/${LOG}"
else
  "$BINARY" "$@"
fi

chmod 777 -R /tendermint

