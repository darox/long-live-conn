# long-live-conn

Is a repository that contains a client and server written in Golang that can maintain a long-lived HTTP2 connection. The following variables are configurable in the server and client through ENV variables:

## Client

- CLIENT_KEEP_ALIVE_SECONDS - The interval in seconds between keep alives. Default is `15` seconds by the net dialer default
- CLIENT_KEEP_ALIVE_ENABLED - Enables keep alives in the client. If the server has keep alives enabled, the client will not send keep alives. Default is `true`
- CLIENT_TLS_DISABLE_VERIFICATION - Disables TLS verification. Default is `true`
- CLIENT_REQUEST_INTERVAL_SECONDS - The interval in seconds between requests. Default is `30` seconds
- CLIENT_SERVER_URL - The URL of the server to connect to. Default is `http://localhost:8080`


## Server

- SERVER_KEEP_ALIVE_ENABLED - Enables keep alives in the server. Default is `true`
- SERVER_KEEP_ALIVE_INTERVAL_SECONDS - The interval in seconds between keep alives. This usually comes from the OS. Default is `-1` meaning it takes the OS default
- SERVER_TLS_CERT_PATH - The path to the TLS certificate. Default is `cert.pem`
- SERVER_TLS_KEY_PATH - The path to the TLS key. Default is `server.key`


## Examples

This section lists some examples and their corresponding Wireshark captures.

### Keep alive active and executed by the server

This requires no change in the config, it's the default.

![Keep alive active and done by the server](./assets/keep-alive-by-server.png)

### Keep alive active and executed by the client

To disable keep alives in the server, set ENV `SERVER_KEEP_ALIVE_ENABLED` to `false`. In this case the client will send keep alives.

![Keep alive active and done by the client](./assets/keep-alive-by-client.png)

## Docker images

Only arch64 images are available at the moment

```
dariomader/long-live-connection-server:v0.0.1
dariomader/long-live-connection-client:v0.0.1
```
