certs:
	openssl req -x509 -newkey rsa:4096 -nodes -keyout server/key.pem -out server/cert.pem -days 365 \
    -subj "/C=CH/ST=Berne/L=Berne/O=long-live-connection/OU=developers/CN=server"


local-up: certs
	docker-compose rm -f
	docker-compose -f compose.yaml up --build

build:
	cd client && docker build -t dariomader/long-live-connection-client:v0.0.1 .
	cd server && docker build -t dariomader/long-live-connection-server:v0.0.1 .
  