package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/net/http2"
)

func doRequest(client *http.Client, c clientConfig) (httpStatus, error) {
	resp, err := client.Get(c.serverURL)
	if err != nil {
		return "", err

	}
	defer resp.Body.Close()

	s := httpStatus(resp.Status)

	return s, nil
}

func setupClient(c clientConfig) (http.Client, error) {
	dialer := &net.Dialer{
		KeepAlive: time.Duration(c.keepAliveIntervalSeconds) * time.Second,
	}

	// Create a custom http.Transport with the custom dialer
	tr := &http.Transport{
		DisableKeepAlives: !c.keepAliveEnabled,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: c.tlsDisableVerification},
		DialContext:       dialer.DialContext, // Use the custom dialer for the keep alive settings
	}

	// Ensure that HTTP/2 is used.
	err := http2.ConfigureTransport(tr)

	if err != nil {
		return http.Client{}, err
	}

	// Create an HTTP client with the custom Transport
	client := http.Client{
		Transport: tr,
	}
	return client, nil
}

func setupClientConfig() clientConfig {
	viper.SetDefault("CLIENT_KEEP_ALIVE_INTERVAL_SECONDS", 15) // 15 is the net dialer default
	viper.SetDefault("CLIENT_KEEP_ALIVE_ENABLED", true)
	viper.SetDefault("CLIENT_TLS_DISABLE_VERIFICATION", true)
	viper.SetDefault("CLIENT_REQUEST_INTERVAL_SECONDS", 30)
	viper.SetDefault("CLIENT_SERVER_URL", "https://localhost:8080")

	viper.AutomaticEnv()

	return clientConfig{
		keepAliveIntervalSeconds: viper.GetFloat64("CLIENT_KEEP_ALIVE_INTERVAL_SECONDS"),
		keepAliveEnabled:         viper.GetBool("CLIENT_KEEP_ALIVE_ENABLED"),
		tlsDisableVerification:   viper.GetBool("CLIENT_TLS_DISABLE_VERIFICATION"),
		RequestIntervalSeconds:   viper.GetFloat64("CLIENT_REQUEST_INTERVAL_SECONDS"),
		serverURL:                viper.GetString("CLIENT_SERVER_URL"),
	}
}

type clientConfig struct {
	keepAliveIntervalSeconds float64
	keepAliveEnabled         bool
	tlsDisableVerification   bool
	RequestIntervalSeconds   float64
	serverURL                string
}

type httpStatus string
