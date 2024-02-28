# golang:1.22.0
FROM golang@sha256:03082deb6ae090a0caa4e4a8f666bc59715bc6fa67f5fd109f823a0c4e1efc2a as build

WORKDIR /go/src/server
COPY server.go main.go ../go.mod ../go.sum ./ 

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/server

FROM gcr.io/distroless/static-debian11@sha256:9235ad98ee7b70ffee7805069ba0121b787eb1afbd104f714c733a8da18f9792
COPY --from=build /go/bin/server /server/server
CMD ["/server/server"]