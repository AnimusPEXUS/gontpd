package gontpd

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"
)

type Server struct {
	port int

	working bool

	OnStarted     func()
	OnStopped     func()
	OnAcceptError func(error)
	OnHandleError func(error)
}

func NewServer(port int) (*Server, error) {
	self := new(Server)
	self.port = port
	return self, nil
}

func (self *Server) IsWorking() bool {
	return self.working
}

func (self *Server) sigStarted() {
	if self.OnStarted != nil {
		self.OnStarted()
	}
}

func (self *Server) sigStopped() {
	if self.OnStopped != nil {
		self.OnStopped()
	}
}

func (self *Server) sigAcceptError(err error) {
	if self.OnAcceptError != nil {
		self.OnAcceptError(err)
	}
}

func (self *Server) sigHandleError(err error) {
	if self.OnHandleError != nil {
		self.OnHandleError(err)
	}
}

func (self *Server) Run() error {
	self.working = true
	log.Println("info", "Time server is starting")

	defer func() {
		self.working = false
		log.Println("info", "Time server is stopped")
		go self.sigStopped()
	}()

	srv, err := net.Listen("tcp", fmt.Sprintf(":%d", self.port))
	if err != nil {
		return err
	}

	go self.sigStarted()
	log.Println("info", "Time server is started", srv.Addr().String())

	for {
		conn, err := srv.Accept()
		if err != nil {
			go self.sigAcceptError(err)
			return err
		}
		go self.handleConn(conn)
	}
}

func (self *Server) Start() {
	go self.Run()
}

func (self *Server) handleConn(conn net.Conn) {

	log.Println("info", "new connection from client "+conn.RemoteAddr().String())

	now := int32(time.Now().UTC().Unix())

	err := binary.Write(conn, binary.BigEndian, int32(UnixToRfc(int64(now))))
	if err != nil {
		go self.sigHandleError(err)
		log.Println(
			"error",
			"client "+conn.RemoteAddr().String(),
			"can't write response data",
			err,
		)
	}

	err = conn.Close()
	if err != nil {
		go self.sigHandleError(err)
		log.Println(
			"error",
			"client "+conn.RemoteAddr().String(),
			"can't close connection",
			err,
		)
	}

	log.Println("info", "responce written. connection closed")
}
