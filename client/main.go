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
	net "github.com/libp2p/go-libp2p-net"
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

	log.Printf("peers: %v", peers)

	for _, peer := range *peers {
		if err = connectTo(peer, h); err != nil {
			log.Printf("Failed to connect to peer (%v): %v", peer, err)
		}
	}

	h.SetStreamHandler("/p2p/1.0.0", handleStream)
	select {} // hang forever
}

func getPeersFromBootnode(target string, h libp2pHost.Host) (*[]*protos.Peer, error) {
	peerid, targetAddr, err := targetInfo(target)
	if err != nil {
		return nil, err
	}
	// We have a peer ID and a targetAddr so we add it to the peerstore
	// so LibP2P knows how to contact it
	h.Peerstore().AddAddr(*peerid, targetAddr, pstore.PermanentAddrTTL)

	log.Println("opening stream")
	// make a new stream from host B to host A
	// it should be handled on host A by the handler we set above because
	// we use the same /p2p/1.0.0 protocol
	s, err := h.NewStream(context.Background(), *peerid, "/p2p/1.0.0")
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

func targetInfo(target string) (*peer.ID, ma.Multiaddr, error) {
	ipfsaddr, err := ma.NewMultiaddr(target)
	if err != nil {
		return nil, nil, err
	}

	// TODO: What is this??
	pid, err := ipfsaddr.ValueForProtocol(ma.P_IPFS)
	if err != nil {
		return nil, nil, err
	}

	peerid, err := peer.IDB58Decode(pid)
	if err != nil {
		return nil, nil, err
	}

	// Decapsulate the /ipfs/<peerID> part from the target
	// /ip4/<a.b.c.d>/ipfs/<peer> becomes /ip4/<a.b.c.d>
	targetPeerAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", peer.IDB58Encode(peerid)))
	targetAddr := ipfsaddr.Decapsulate(targetPeerAddr)

	return &peerid, targetAddr, nil
}

func connectTo(peer *protos.Peer, h libp2pHost.Host) error {
	peerid, targetAddr, err := targetInfo(peer.Address)
	if err != nil {
		return err
	}

	// We have a peer ID and a targetAddr so we add it to the peerstore
	// so LibP2P knows how to contact it
	h.Peerstore().AddAddr(*peerid, targetAddr, pstore.PermanentAddrTTL)

	log.Println("opening stream")
	// make a new stream from host B to host A
	// it should be handled on host A by the handler we set above because
	// we use the same /p2p/1.0.0 protocol
	s, err := h.NewStream(context.Background(), *peerid, "/p2p/1.0.0")
	if err != nil {
		return err
	}
	// Create a buffered stream so that read and writes are non blocking.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	// TODO
	_ = rw

	return nil
}

func handleStream(s net.Stream) {
	log.Println("New stream connected")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go readData(rw)
	//go writeData(rw)
}

func readData(rw *bufio.ReadWriter) {
	// TODO
}
