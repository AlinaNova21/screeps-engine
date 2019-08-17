package rpc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

type Frame struct {
	ID      int              `json:"id,omitempty"`
	Method  string           `json:"method,omitempty"`
	Channel string           `json:"channel,omitempty"`
	Error   string           `json:"error,omitempty"`
	Args    *json.RawMessage `json:"args,omitempty"`
	Result  interface{}      `json:"result,omitempty"`
}

func (f *Frame) GetArgs(v interface{}) error {
	err := json.Unmarshal(*f.Args, v)
	return err
}

type Method func(*Frame) (interface{}, error)

func NewServer () *Server {
	return &Server{
		Methods: make(map[string]Method),
	}
}

type Server struct {
	Methods map[string]Method
}

func (s *Server) Listen(addr string) error {
	if addr == "" {
		addr = ":21027"
	}
	taddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return fmt.Errorf("Error parsing address %s: %v", addr, err)
	}
	ln, err := net.ListenTCP("tcp", taddr)
	if err != nil {
		return fmt.Errorf("Error listening on address %s: %v", addr, err)
	}
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			return fmt.Errorf("Error listening for client: %v", err)
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn *net.TCPConn) {
	var payloadLen int32
	var err error
	for {
		err = binary.Read(conn, binary.BigEndian, &payloadLen)
		if err != nil {
			log.Printf("Error reading len: %v", err)
			break
		}
		buf := make([]byte, payloadLen)
		off := 0
		for {
			read, err := conn.Read(buf[off:])
			if err != nil {
				log.Printf("Error reading: %v", err)
				break
			}
			off += read
			if int32(off) >= payloadLen {
				break
			}
		}
		frame := Frame{}
		err := json.Unmarshal(buf, &frame)
		if err != nil {
			log.Printf("Error parsing json '%s': %v", string(buf), err)
			break
		}
		if frame.Method == "subscribe" {
			// TODO: Handle subscribe
		} else if fn, ok := s.Methods[frame.Method]; ok {
			ret, err := fn(&frame)
			retFrame := Frame{
				ID: frame.ID,
			}
			if err != nil {
				retFrame.Error = err.Error()
			} else {
				retFrame.Result = ret
			}
			payload, err := json.Marshal(retFrame)
			if err != nil {
				log.Printf("Cannot marshal resp: %v", err)
				break
			}
			payloadLen = int32(len(payload))
			binary.Write(conn, binary.BigEndian, payloadLen)
			_, err = io.Copy(conn, bytes.NewReader(payload))
			if err != nil {
				log.Printf("Error writing payload: %v", err)
				break
			}
		}
	}
}
