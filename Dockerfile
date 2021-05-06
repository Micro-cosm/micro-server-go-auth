

FROM	golang:1.16-buster as builder

ARG		EXECUTABLE
ENV		EXECUTABLE=${EXECUTABLE}

WORKDIR	/app

COPY	go.* ./
RUN		go mod download
COPY	./main.go ./main.go

RUN		go build -v -o $EXECUTABLE


FROM	debian:buster-slim

RUN		set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates &&  rm -rf /var/lib/apt/lists/*

COPY	--from=builder /app/$EXECUTABLE /app/$EXECUTABLE
COPY	docker-entrypoint.sh /

ENTRYPOINT ["/docker-entrypoint.sh"]