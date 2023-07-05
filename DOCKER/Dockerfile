# stage 1 Generate tendermint Binary
FROM golang:1.18-alpine as builder
RUN apk update && \
    apk upgrade && \
    apk add --no-cache git make gcc libc-dev build-base curl jq bash file gmp-dev clang libtool autoconf automake
COPY / /tendermint
WORKDIR /tendermint
RUN make build-linux

# stage 2
FROM golang:1.18-alpine
LABEL maintainer="hello@blockchain.line.me"

# tendermint will be looking for the genesis file in /tendermint/config/genesis.json
# (unless you change `genesis_file` in config.toml). You can put your config.toml and
# private validator file into /tendermint/config.
#
# The /tendermint/data dir is used by tendermint to store state.
ENV OCHOME /tendermint

# OS environment setup
# Set user right away for determinism, create directory for persistence and give our user ownership
# jq and curl used for extracting `pub_key` from private validator while
# deploying tendermint with Kubernetes. It is nice to have bash so the users
# could execute bash commands.
RUN apk update && \
    apk upgrade && \
    apk add --no-cache git make gcc libc-dev build-base curl jq bash file gmp-dev clang libtool autoconf automake && \
    addgroup tendermint && \
    adduser -S -G tendermint tendermint -h "$OCHOME"

# Run the container with tendermint by default. (UID=100, GID=1000)
USER tendermint

WORKDIR $OCHOME

# p2p, rpc and prometheus port
EXPOSE 26656 26657 26660

STOPSIGNAL SIGTERM

COPY --from=builder /tendermint/build/tendermint /usr/bin/tendermint

# You can overwrite these before the first run to influence
# config.json and genesis.json. Additionally, you can override
# CMD to add parameters to `tendermint node`.
ENV PROXY_APP=kvstore MONIKER=dockernode CHAIN_ID=dockerchain

COPY ./DOCKER/docker-entrypoint.sh /usr/local/bin/

ENTRYPOINT ["docker-entrypoint.sh"]
CMD ["node"]

# Expose the data directory as a volume since there's mutable state in there
VOLUME [ "$OCHOME" ]

