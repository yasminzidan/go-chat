# Go RPC Chat Server

Simple RPC chat server and client in Go.  
Server is Dockerized.

## Run Server with Docker
cd server
docker build -t rose0505/chat-server .
docker run -p 1234:1234 rose0505/chat-server

## Run Client
cd client
go run main.go

## Docker Hub Image
https://hub.docker.com/r/rose0505/chat-server
https://hub.docker.com/repositories/rose0505

