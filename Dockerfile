FROM golang AS build-env

RUN apt update
RUN apt install gcc
ADD src /src
WORKDIR /src
RUN go get github.com/ramr/go-reaper
RUN go build -ldflags "-linkmode external -extldflags -static" -o dnscacherefresh

FROM scratch
COPY --from=build-env /src/dnscacherefresh /dnscacherefresh

ENTRYPOINT ["/dnscacherefresh"]