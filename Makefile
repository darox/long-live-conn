certs:
	openssl req -x509 -newkey rsa:4096 -nodes -keyout server/key.pem -out server/cert.pem -days 365 \
    -subj "/C=CH/ST=Berne/L=Berne/O=long-live-connection/OU=developers/CN=server"

local-up: certs
	docker-compose rm -f
	docker-compose -f compose.yaml up --build

build-server:
	cd server && docker buildx build -t dariomader/long-live-connection-server --platform linux/amd64,linux/arm64 .

build-client: 
	cd client && docker buildx build -t dariomader/long-live-connection-client --platform linux/amd64,linux/arm64 .

release: release-server release-client

release-server: build-server
	cd server && docker buildx build -t dariomader/long-live-connection-server:v0.0.2 --platform linux/amd64,linux/arm64 --push .

release-client: build-client
	cd server && docker buildx build -t dariomader/long-live-connection-server:v0.0.2 --platform linux/amd64,linux/arm64 --push .
