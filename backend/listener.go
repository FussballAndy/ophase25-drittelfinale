package main

import (
	"crypto/tls"
	"net"
	"net/http"

	"github.com/FussballAndy/webtransport-go"
	"github.com/quic-go/quic-go/http3"
)

func CreateServer(addr, certFile, keyFile string, handler http.Handler) (*webtransport.Server, error) {
	var err error
	certs := make([]tls.Certificate, 1)
	certs[0], err = tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: certs,
	}

	if addr == "" {
		addr = ":https"
	}

	if handler == nil {
		handler = http.DefaultServeMux
	}

	quicServer := webtransport.Server{
		H3: http3.Server{
			TLSConfig: config,
			Handler:   handler,
		},
	}
	return &quicServer, nil
}

func ListenAndServeTLS(server *webtransport.Server, certFile, keyFile string) error {
	addr := server.H3.Addr
	handler := server.H3.Handler
	// Open the listeners
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return err
	}
	defer udpConn.Close()

	hErr := make(chan error, 1)
	qErr := make(chan error, 1)
	go func() {
		hErr <- http.ListenAndServeTLS(addr, certFile, keyFile, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			server.H3.SetQUICHeaders(w.Header())
			handler.ServeHTTP(w, r)
		}))
	}()
	go func() {
		qErr <- server.Serve(udpConn)
	}()

	select {
	case err := <-hErr:
		server.Close()
		return err
	case err := <-qErr:
		// Cannot close the HTTP server or wait for requests to complete properly :/
		return err
	}
}
