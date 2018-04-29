package main

import (
	"dino/communicationlayer/dinoproto3"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"strings"

	"github.com/gogo/protobuf/proto"
)

/*
	1- We will serialize some data via proto
	2- We will send this data via TCP to a different service
	3- We will deserialize the data via proto3, and print ou tht extracted values

	A- A TCP client needs to be written and sent the data
	B- A TCP server to receive the data
*/
func main() {
	op := flag.String("op", "s", "s for server, c for client") // proto3test -op s => will run a as a server, proto3test -op c will run as a client
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		RunProto3Server()
	case "c":
		RunProto3Client()
	}
}

func RunProto3Server() {
	l, err := net.Listen("tcp", ":8550")
	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := l.Accept() // listen for incoming tcp client requests
		if err != nil {
			log.Fatal(err)
		}

		defer l.Close() // close listener
		go func(c net.Conn) {
			defer c.Close()
			data, err := ioutil.ReadAll(c)
			if err != nil {
				return
			}

			a := &dinoproto3.Animal{}
			err = proto.Unmarshal(data, a)	// unserialize binary protobuf data
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(a)
		}(c)
	}
}

func RunProto3Client() {
	a := &dinoproto2.Animal{
		Id:         1,
		AnimalType: "Raptor",
		Nickname: 	"rapto",
		Zone: 		3,
		Age: 		20
	}
	data, err := proto.Marshal(a)	// serialize data into protobuf binary
	if err != nil {
		log.Fatal(err)
	}

	SendData(data)
}

func SendData(data []byte) {
	c, err := net.Dial("tcp", "127.0.0.1:8550")	// establish connection with tcp server
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Write(data)	// will write data to tcp channel
}
