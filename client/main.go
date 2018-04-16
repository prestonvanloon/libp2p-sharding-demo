package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	"github.com/prestonvanloon/libp2p-sharding-demo/host"
	protos "github.com/prestonvanloon/libp2p-sharding-demo/proto"

	libp2pHost "github.com/libp2p/go-libp2p-host"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	ma "github.com/multiformats/go-multiaddr"
)

var (
	bootnodeAddr = flag.String("b", "", "Bootnode address")
	port         = flag.Int("p", 10100, "Port to listen on")
)

func main() {
	flag.Parse()

	h, err := host.MakeBasicHost(*port, true, 0)
	if err != nil {
		log.Fatal(err)
	}

	var peers *[]*protos.Peer

	if *bootnodeAddr != "" {
		log.Println("getting peers")
		peers, err = getPeersFromBootnode(*bootnodeAddr, h)
		if err != nil {
			log.Fatal(err)
		}
	}

	_ = peers

	log.Printf("peers: %v", peers)
}

func getPeersFromBootnode(target string, h libp2pHost.Host) (*[]*protos.Peer, error) {
	ipfsaddr, err := ma.NewMultiaddr(target)
	if err != nil {
		return nil, err
	}

	// TODO: What is this??
	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	if err != nil {
		return nil, err
	}

	peerid, err := peer.IDB58Decode(pid)
	if err != nil {
		return nil, err
	}

	// Decapsulate the /ipfs/<peerID> part from the target
	// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
	targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

	// We have a peer ID and a targetAddr so we add it to the peerstore
	// so LibP2P knows how to contact it
	h.Peerstore().AddAddr(peerid, targetAddr, pstore.PermanentAddrTTL)

	log.Println("opening stream")
	// make a new stream from host B to host A
	// it should be handled on host A by the handler we set above because
	// we use the same /p2p/1.0.0 protocol
	s, err := h.NewStream(context.Background(), peerid, "/p2p/1.0.0")
	if err != nil {
		return nil, err
	}
	// Create a buffered stream so that read and writes are non blocking.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	// Send our address
	request := &protos.RegisterPeerRequest{
		Peer: &protos.Peer{Address: host.FullAddr(h)},
	}

	data, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}

	rw.Write(data)
	rw.Flush()

	data = make([]byte, 4096)
	n, err := rw.Read(data)
	if err != nil {
		log.Fatalf("failed to read data %v", err)
	}
	response := &protos.GetPeersResponse{}
	err = proto.Unmarshal(data[:n], response)

	return &response.Peers, nil
}
