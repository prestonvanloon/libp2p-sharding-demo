package main

import (
	"bufio"
	"flag"
	"log"
	"sync"

	"github.com/golang/protobuf/proto"
	net "github.com/libp2p/go-libp2p-net"
	host "github.com/prestonvanloon/libp2p-sharding-demo/host"
	protos "github.com/prestonvanloon/libp2p-sharding-demo/proto"
)

var (
	port     = flag.Int("p", 10000, "The port to listen on")
	randSeed = flag.Int64("s", 0, "")
	peers    = []*protos.Peer{}
	mutex    = &sync.Mutex{}
)

func main() {
	flag.Parse()
	n, err := host.MakeBasicHost(*port, true /*secio*/, *randSeed)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("I am %s\n", host.FullAddr(n))

	n.SetStreamHandler("/p2p/1.0.0", handleStream)
	select {} // hang forever
}

func handleStream(s net.Stream) {
	//defer log.Println("Stream disconnected")
	log.Println("New stream connected")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go readData(rw)
	go writeData(rw)
}

func readData(rw *bufio.ReadWriter) {
	// register new peer request
	for {
		data := make([]byte, 4096)
		n, err := rw.Read(data)
		if err != nil {
			log.Printf("failed to read %v", err)
			return
		}
		request := &protos.RegisterPeerRequest{}
		err = proto.Unmarshal(data[:n], request)
		if err != nil {
			log.Printf("failed to unmarshall: %v", err)
		}

		log.Printf("Received request to register peer: %+v", request.Peer)
		mutex.Lock()
		peers = append(peers, request.Peer)
		mutex.Unlock()
	}
}

func writeData(rw *bufio.ReadWriter) {
	// send current peers

	response := &protos.GetPeersResponse{Peers: peers}
	data, err := proto.Marshal(response)
	if err != nil {
		log.Printf("failed to marshall data: %v", err)
		return
	}

	rw.Write(data)
	rw.Flush()
}
