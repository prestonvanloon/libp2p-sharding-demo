# libp2p sharding demo

This is a work in progress exploration of libp2p in a sharded system.

Requirements: [bazel](bazel.io)


## Bootnode

Run the bootnode with a seed flag `-s` for consistant address

```
bazel build //bootnode && bazel-bin/bootnode/linux_amd64_stripped/bootnode -s 1
```

## Run a client

```
bazel build //client && bazel-bin/client/linux_amd64_stripped/client -b /ip4/127.0.0.1/tcp/10000/ipfs/QmexAnfpHrhMmAC5UNQVS8iBuUUgDrMbMY17Cck2gKrqeX
```
