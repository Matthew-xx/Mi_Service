# GetSmsCd Service

This is the GetSmsCd service

Generated with

```
micro new Mi_house/GetSmsCd --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.GetSmsCd
- Type: srv
- Alias: GetSmsCd

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./GetSmsCd-srv
```

Build a docker image
```
make docker
```