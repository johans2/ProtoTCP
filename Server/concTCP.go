package main

import (
	"RTSServer/proto/messages"
	"bufio"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/protobuf/proto"
)

const MIN = 1
const MAX = 100

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		// Read the tcp data
		bytes := make([]byte, 4096)
		length, err := bufio.NewReader(c).Read(bytes)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Deserialize the message
		tmsg := &messages.TestMessage{}
		if err := proto.Unmarshal(bytes[:length], tmsg); err != nil {
			fmt.Println(err)
		} else {
			switch tmsg.GetPayload().(type) {
			case *messages.TestMessage_Pl1:
				fmt.Println(tmsg.GetPl1().Msg)
				break
			case *messages.TestMessage_Pl2:
				fmt.Println(tmsg.GetPl2().Number)
				break
			case *messages.TestMessage_Pl3:
				fmt.Println(tmsg.GetPl3().Percent)
				break
			}
		}

		// Serialize and send back
		size := proto.Size(tmsg)
		bytes, err = proto.Marshal(tmsg)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("respons size %s\n", strconv.Itoa(size))
		sizeBytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(sizeBytes, uint32(size))
		c.Write(sizeBytes)
		c.Write(bytes)
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	fmt.Println(messages.Person_WORK)

	PORT := ":" + arguments[1]
	listener, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Started server")
	defer listener.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
