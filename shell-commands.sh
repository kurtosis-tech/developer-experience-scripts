#!/bin/sh

docker pull cassandra

docker network create cassandra-network

### Before starting cassandra networks, change the Memory available per container in the Docker Container Engine resource tab to at least 8GB.

### Docker commands to set up sample cluster accessible locally
### NOTE: IF NODES 2 OR 3 FAIL TO STARTUP, YOU MIGHT NEED TO WAIT MORE TIME BETWEEN STARTING THEM UP
### If there is no database initialized when the container starts, then a default database will be created. 
### While this is the expected behavior, this means that it will not accept incoming connections until such initialization completes.
### This may cause issues when using automation tools, such as docker-compose, which start several containers simultaneously.
docker run --name cassandra-1 -p 0.0.0.0:9042:9042 --network cassandra-network -d -e CASSANDRA_LISTEN_ADDRESS=localhost cassandra:latest
docker run --name cassandra-2 -p 0.0.0.0:9043:9042 --network cassandra-network -e CASSANDRA_SEEDS=cassandra-1 -e CASSANDRA_LISTEN_ADDRESS=localhost -d cassandra:latest
docker run --name cassandra-3 -p 0.0.0.0:9044:9042 --network cassandra-network -e CASSANDRA_SEEDS=cassandra-1 -e CASSANDRA_LISTEN_ADDRESS=localhost -d cassandra:latest

### Install gocql for cassandra
go get github.com/gocql/gocql
### Run go file
go run go-cassandra-queries.go
