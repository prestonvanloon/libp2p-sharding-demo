# libp2p sharding demo

This is a work in progress exploration of libp2p in a sharded system.

Requirements: [bazel](bazel.io)

**WORK IN PROGRESS**

## Bootnode

Run the bootnode with a seed flag `-s` for consistant address

```
bazel build //bootnode && bazel-bin/bootnode/linux_amd64_stripped/bootnode -s 1
```

## Run a client

```
bazel build //client && bazel-bin/client/linux_amd64_stripped/client -b /ip4/127.0.0.1/tcp/10000/ipfs/QmexAnfpHrhMmAC5UNQVS8iBuUUgDrMbMY17Cck2gKrqeX
```

# Communication on a sharded network.

For this experimentation, I'll be using objects that closely represent the
objects outlined in the spec ~1 for a sharded Ethereum. Specifically, 
"collations" in place of "blocks". This design is meant to explore a Torus
shaped network where nodes would connect to peers interested in their shard and
neighboring shards. Nodes would propigate messages to their peers in the
direction of the message destination.

For example, if a node was interested in shard 15 and they have 4 peers looking
for shards [11, 14, 17, 19] and the node receives a message destined for shard
55 then this node would propagate the message to the latter 2 peers (listening
on 17, 19). The benefit here is that we can optimize the distance traveled for
a given message when nodes only accept peers listening on shards within a
narrow range. 


## Communication Scenarios

### Handshake / status 

Whenever a node connects to another, they must exchange a status message which
consists of the shard they are interested in, protocol version, and some
information about their sync status.

### New collation hashes

When a node comes online for the first time or after being offline for some
time, they will be behind on collations. As such, a node should send all
collation hashes that the connecting node may not have. 

### Request for collation headers 

A node may request collation headers from since a particular collation. This
data is also generally available from the SMC, but the node may not be
connected to the main chain. 

### Response for collation headers 

In response to a request for collation headers, the node may return zero or
headers.

### Request for collation bodies

A node may request another node for collation bodies for given collations.

### Response for collation bodies

In response to a request for collation bodies, the node would return zero or
more collation bodies for the collations requested.

### A new collation

New collations should be propagated to relevant peers. Either by receiving the 
message from the SMC or from another node in the network (especially if the 
running node is not also connected to the main chain / SMC).

