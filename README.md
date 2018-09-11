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

## Docker Commands

1. List docker images in table: `docker images --format "table {{.ID}}\t{{.Repository}}\t{{.Tag}}"`
1. Purge images: `docker images -f dangling=true`