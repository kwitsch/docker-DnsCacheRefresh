# DNS Cache Refresh

## Description
Service for periodically requesting dns resolution against an defined resolver. </br>
If a service name instead of an IP address is used as an resolver, the resolver name resolution will be processed through 127.0.0.11.

## Features
- Golang static build
- self contained init service
- small size through scratch image with only necessarry executable
- all configurations done throug environment variables

## Environment
- DCR.RESOLVER: service name or ip of the targeted dns server
- DCR.START: duration from start until first request
- DCR.REFRESH: duration between request blocks
- DCR.VERBOSE: detailed log messages(true or false)
- DCR.DOMAINS[*n*]: list of domains to be requested(n represents a numeric value)

## Example-Stack
```YAML
version: "3.8"

services:
  unbound:
    image: kwitsch/unbound:monthly
    networks: 
      - dnscom
  dnscacherefresh:
    image: kwitsch/dnscacherefresh
    environment:
      - DCR.RESOLVER=unbound
      - DCR.START=10s
      - DCR.REFRESH=2h
      - DCR.VERBOSE=True
      - DCR.DOMAINS[0]=github.com
      - DCR.DOMAINS[1]=hub.docker.com
    networks: 
      - dnscom

networks:
  dnscom:
```