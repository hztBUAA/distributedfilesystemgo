package p2p

import "net"

// Peer is an interface that represents the remote node.
// send is for send messages to it's Peer from itself(its'Peers should implements the peer)
type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// Transport is anything that handles the communication
// between the nodes in the network. This can be of the
// form (TCP, UDP, websockets, ...)
type Transport interface {
	Addr() string // returns an address of "ip" + "port"
	Dial(string) error // establish the connection to the specified address
	ListenAndAccept() error // to accept a connection from other nodes
	Consume() <-chan RPC // it returns a channel that only accept the RPC, which is used for receiving the RPC from other nodes 
	Close() error
}
