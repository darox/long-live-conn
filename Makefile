certs:
	openssl req -x509 -newkey rsa:4096 -nodes -keyout server/key.pem -out server/cert.pem -days 365 \
    -subj "/C=CH/ST=Berne/L=Berne/O=long-live-connection/OU=developers/CN=server"

local-up: certs
	docker-compose rm -f
	docker-compose -f compose.yaml up --build


release: 
	@read -p "Are you sure you want to continue releasing a new image? [y/N] " answer; \
	if [ "$$answer" != "y" ]; then \
		exit 1; \
	fi
	cd client && docker buildx build --platform linux/amd64,linux/arm64 -t dariomader/long-live-connection-client:v0.0.2 --push .
	cd server && docker buildx build --platform linux/amd64,linux/arm64 -t dariomader/long-live-connection-server:v0.0.2 --push .

k8s-clean:
	kubectl delete secret long-live-conn-server-certs --ignore-not-found=true
	kubectl delete -f install/kubernetes/server --ignore-not-found=true
	kubectl delete -f install/kubernetes/client --ignore-not-found=true

k8s-up: k8s-clean certs
	kubectl create secret tls long-live-conn-server-certs \
    --cert=server/cert.pem \
    --key=server/key.pem
	kubectl apply -f install/kubernetes/server
	kubectl apply -f install/kubernetes/client