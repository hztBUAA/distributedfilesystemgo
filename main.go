package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/anthdm/foreverstore/p2p"
)

func makeServer(listenAddr string, nodes ...string) *FileServer {
	// Opts is the set of rules to initialize an object
	// tcp needs the address of listening, needs the so-called handshakeFunc, and needs the so-called decoder,which is correspongding to the Enc of the following FileServer 
	tcptransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tcpTransport := p2p.NewTCPTransport(tcptransportOpts)


	// fileserver needs a tcpTransport
	fileServerOpts := FileServerOpts{
		EncKey:            newEncryptionKey(), // provide safe encoder
		StorageRoot:       listenAddr + "_network", // the root position (folder)
		PathTransformFunc: CASPathTransformFunc, // a function to transform the keys in files into the full path of file system
		Transport:         tcpTransport, // provide the transimission of network, dealing with the network communication
		BootstrapNodes:    nodes, // definite the initial nodes, which is a slice of strings, including the initial nodes int the network
	}

	s := NewFileServer(fileServerOpts)

	tcpTransport.OnPeer = s.OnPeer

	return s
}

func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":7000", "")
	s3 := makeServer(":5000", ":3000", ":7000") // s3 会使用到 bootstrap 能够快进入节点 知道这个网络的所有节点

	go func() { log.Fatal(s1.Start()) }()
	time.Sleep(500 * time.Millisecond) // 0.5 s
	go func() { log.Fatal(s2.Start()) }()

	time.Sleep(2 * time.Second) // 2 s , then start the s3

	go s3.Start()
	time.Sleep(2 * time.Second)

	for i := 0; i < 20; i++ {
		key := fmt.Sprintf("picture_%d.png", i) // i: 0->19  20 pictures stored into key variable
		data := bytes.NewReader([]byte("my big data file here!")) // data is corresponding to the key ?  bytes' Reader will produce the data bytes ?
		s3.Store(key, data) // a key value to be stored in s3=====here i need to have a great understanding of the store 

		if err := s3.store.Delete(s3.ID, key); err != nil {
			log.Fatal(err)
		}

		r, err := s3.Get(key)
		if err != nil {
			log.Fatal(err)
		}

		b, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
	}
}
