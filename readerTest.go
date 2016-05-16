package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	ss := "abc"
	log.Println(bytes.NewBufferString(ss))

	// check read
	checkStringRead()
	time.Sleep(time.Duration(1) * time.Second)

	checkByteRead()
}

func checkStringRead() {
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	b := make([]byte, 20)

	fmt.Printf("len %v size %v brBuffed->%v\n", s.Len(), s.Size(), br.Buffered())

	n, err := br.Read(b)
	fmt.Printf("%-20s %-2v %v %v\n", b[:n], n, err, s.Len())
	// ABCDEFGHIJKLMNOPQRST 20 <nil>

	fmt.Printf("len %v size %v brBuffed->%v\n", s.Len(), s.Size(), br.Buffered())

	n, err = br.Read(b)
	fmt.Printf("%-20s %-2v %v %v\n", b[:n], n, err, s.Len())
	// UVWXYZ1234567890 16 <nil>
	fmt.Printf("len %v size %v brBuffed->%v\n", s.Len(), s.Size(), br.Buffered())

	n, err = br.Read(b)
	fmt.Printf("%-20s %-2v %v %v\n", b[:n], n, err, s.Size())
	// 0 EOF
}

func checkByteRead() {
	//str := strings.NewReader("1234567890")
	//str := strings.NewReader("abcdef")
	buf := []byte{0x05, 0x00, 0x00, 0x02, 0xea, 0x00, 0x00, 0x00, 0x03, 0x97, 0x98, 0x99,
		0x05, 0x00, 0x00, 0x03, 0xea, 0x00, 0x00, 0x00, 0x04, 0x97, 0x98, 0x99, 0x10,
	}
	b := bytes.NewReader(buf) // complete pack
	re := bufio.NewReader(b)

	//buf1 := []byte{0x05,0x00,0x00,0x03,0xea,0x00,0x00,0x00,0x03,0x97,0x98,0x99,}
	//buf2 := []byte{0x05,0x00,0x00,0x03,0xea,0x00,0x00,0x00,0x03,0x97,0x98,0x99,}
	//buf3 := []byte{0x05,0x00,0x00,0x03,0xea,0x00,0x00,0x00,0x03,0x97,0x98,0x99,}
	//b1 := bytes.NewReader(buf1)
	//b2 := bytes.NewReader(buf2)
	//b3 := bytes.NewReader(buf3)
	//re1  := bufio.NewReader(b1)
	//re2  := bufio.NewReader(b2)
	//re3  := bufio.NewReader(b3)

	// complete pack
	head, err := re.Peek(9)
	if err != nil {
		log.Print("err ", string(head), " e->", err, " buffed->", re.Buffered())
	}
	log.Println("head ", head, " buffed->", re.Buffered())

	header := head[0]
	log.Println("header ", header, " buffed->", re.Buffered())

	cmd := binary.BigEndian.Uint32(head[1:5])
	log.Println("cmd ", cmd, " buffed->", re.Buffered())

	length := binary.BigEndian.Uint32(head[5:9])
	log.Println("length ", length, " buffed->", re.Buffered())

	data := make([]byte, 9+length)
	// func A read
	//len := 0
	////for {
	////	n,err := re.Read(data[len:])
	//	if err != nil {
	//		//break
	//	}
	//
	//	log.Println("data", data, " n->", n,  " err", err, " buffed->", re.Buffered())
	//	len += n
	//if len == (int)(9 + length) {
	//	log.Println("one packet", data)
	//}

	// func B peek + discard
	data, err = re.Peek((int)(9 + length))
	log.Println("data", data, " err", err, " buffed->", re.Buffered())

	re.Discard((int)(9 + length))
	log.Println("data", data, " err", err, " buffed->", re.Buffered())

	//}
}
