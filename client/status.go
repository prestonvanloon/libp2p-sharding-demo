package main

import (
	"bufio"
	"log"

	"github.com/golang/protobuf/proto"
	net "github.com/libp2p/go-libp2p-net"
	protos "github.com/prestonvanloon/libp2p-sharding-demo/proto"
)

func handleStatusStream(s net.Stream) {
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	// Send client status
	status := &protos.Status{
		// TODO: Populate with real data
		ProtocolVersion: 1,
		ShardId:         1,
		HeadHeight:      1,
		HeadHash:        1,
	}
	data, err := proto.Marshal(status)
	if err != nil {
		log.Printf("failed to marshal data: %v", err)
		return
	}

	rw.Write(data)
	rw.Flush()
}
