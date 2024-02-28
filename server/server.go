package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/net/http2"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Welcome to the HTTP/2 server!"))

	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (ln tcpKeepAliveListener) Accept() (net.Conn, error) {
	conn, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	// Set keep-alive options on the accepted connection
	err = conn.SetKeepAlive(ln.keepAliveEnabled)
	if err != nil {
		return nil, err
	}

	if ln.keepAliveIntervalSeconds != -1 {
		err = conn.SetKeepAlivePeriod(time.Duration(ln.keepAliveIntervalSeconds) * time.Second)
	}

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func runServer(c serverConfig) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)

	server := &http.Server{
		Addr:    c.listenAddress,
		Handler: mux,
	}

	// Load server certificate and key
	cert, err := tls.LoadX509KeyPair(c.tlsCertPath, c.tlsKeyPath)
	if err != nil {
		return err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{http2.NextProtoTLS}, // Enable HTTP/2
	}

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return err
	}

	// Wrap the listener to configure keep-alive for new connections
	tlsLn := tls.NewListener(tcpKeepAliveListener{
		ln.(*net.TCPListener),
		c.keepAliveIntervalSeconds,
		c.keepAliveEnabled,
	}, tlsConfig)

	return server.Serve(tlsLn)
}

func setupServerConfig() serverConfig {
	viper.SetDefault("SERVER_KEEP_ALIVE_INTERVAL_SECONDS", -1)
	viper.SetDefault("SERVER_KEEP_ALIVE_ENABLED", true)
	viper.SetDefault("SERVER_LISTEN_ADDRESS", ":8080")
	viper.SetDefault("SERVER_TLS_CERT_PATH", "cert.pem")
	viper.SetDefault("SERVER_TLS_KEY_PATH", "key.pem")

	viper.AutomaticEnv()

	return serverConfig{
		keepAliveIntervalSeconds: viper.GetFloat64("SERVER_KEEP_ALIVE_INTERVAL_SECONDS"),
		keepAliveEnabled:         viper.GetBool("SERVER_KEEP_ALIVE_ENABLED"),
		listenAddress:            viper.GetString("SERVER_LISTEN_ADDRESS"),
		tlsCertPath:              viper.GetString("SERVER_TLS_CERT_PATH"),
		tlsKeyPath:               viper.GetString("SERVER_TLS_KEY_PATH"),
	}
}

type serverConfig struct {
	keepAliveIntervalSeconds float64
	keepAliveEnabled         bool
	tlsCertPath              string
	tlsKeyPath               string
	listenAddress            string
}

// tcpKeepAliveListener sets TCP keep-alive on accepted connections.
type tcpKeepAliveListener struct {
	*net.TCPListener
	keepAliveIntervalSeconds float64
	keepAliveEnabled         bool
}
