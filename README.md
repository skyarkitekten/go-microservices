# Building Microservices in Go

## Introduction

We building microservices in Go

### Technology Used

1. Go (a.k.a. "golang")
1. gRPC
1. Protobuf - including compiler
1. Docker
1. Kubernetes
1. Istio
1. GCP

## Getting Started

Install the following:

1. gRPC
1. protobuf (move to usr/bin)

## Running the services

### Everything

`docker-compose up -s`

### User Service

1. `docker run -d -p 5432:5432 postgres`
1. `make build`
1. `make run`

## Docker Commands

1. List docker images in table: `docker images --format "table {{.ID}}\t{{.Repository}}\t{{.Tag}}"`
1. Purge images: `docker images -f dangling=true`


## UI

https://auth0.com/blog/our-engineering-experience-with-react-and-storybook/
