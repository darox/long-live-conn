# long-live-conn

A client and server written in Golang that can maintain a long-lived HTTP2 connection. The TCP keep-alive intervals are configurable and can be disabled. The server and client can be deployed using Docker or Kubernetes.

## Deployment

A few notes on the deployment of the server and client.

### Docker

This repo includes a docker-compose file that can be used to deploy the server and client. Use the `make local-up` command to launch the server and client. This will auto-create TLS certificates that are required for the server. 

The images are available on Docker Hub:

```
dariomader/long-live-connection-server:v0.0.2
dariomader/long-live-connection-client:v0.0.2
```

### Kubernetes

This repo includes k8s manifests under the `install/kubernetes` directory. Use the `make k8s-up` command to deploy the server and client to a k8s cluster. This will auto-create TLS certificates that are required for the server.


### Client

- CLIENT_KEEP_ALIVE_INTERVAL_SECONDS - The interval in seconds between keep alives. Default is `15` by the net/dialer package.
- CLIENT_KEEP_ALIVE_ENABLED - Enables keep alives in the client. If the server has keep alives enabled, the client will not send keep alives. Default is `true`
- CLIENT_TLS_DISABLE_VERIFICATION - Disables TLS verification. Default is `true`
- CLIENT_REQUEST_INTERVAL_SECONDS - The interval in seconds between requests. Default is `30` seconds
- CLIENT_SERVER_URL - The URL of the server to connect to. Default is `http://localhost:8080`


### Server

- SERVER_KEEP_ALIVE_ENABLED - Enables keep alives in the server. Default is `true`
- SERVER_KEEP_ALIVE_INTERVAL_SECONDS - The interval in seconds between keep alives. Default is `15` by the net/TCPConn package.
- SERVER_TLS_CERT_PATH - The path to the TLS certificate. Default is `cert.pem`
- SERVER_TLS_KEY_PATH - The path to the TLS key. Default is `server.key`


## Example Wireshark captures

This section lists some examples and their corresponding Wireshark captures.

### Keep alive by the server

This requires no change in the config, it's the default.

![Keep alive active and done by the server](./assets/keep-alive-by-server.png)

Note: that the Wireshark traces were taken on MacOS. The default TCP keep-alive in Go is 15 seconds. 
If you disable the TCP keep alive in the client via `CLIENT_KEEP_ALIVE_ENABLED`, the client will send a RST
packet directly after sending the HTTP request. In this case the client will not receive a TCP keep alive packet from the server. See below image for an example.

![Keep alive disabled in the client](./assets/keep-alive-disabled-in-client.png)


### Keep alive by the client

To disable keep alives in the server, set ENV `SERVER_KEEP_ALIVE_ENABLED` to `false`. In this case the client will send keep alives.

![Keep alive active and done by the client](./assets/keep-alive-by-client.png)

