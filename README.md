# dchat
distributed chat system  
this repo is private  

## How to boot up  
  
sudo docker compose up  
  
## Architecture
there are 5 services:
1. **dchatctl**: cli of project (does not implemented) ("Sorry, I don’t have the time to do that")  
2. **presence service**: control user session snd room session using a distributed in-memory kv store and distributed locks (i choose ETCD)  
3. **room service**: traditional APIes using grpc, postgres & redis for cache  
4. **chat service**: handles the user live connection using gRPC bidirectional streaming  
5. **message service**: messages after chat should save in message service for later (if anyone wants to load messages)
i want to use a kv store with high capabilities for clustering and IO (my choice is scylla db) (does not implemented) ("Sorry, I don’t have the time to do that") 

![architecture](https://raw.githubusercontent.com/rezamokaram/dchat/refs/heads/main/docs/arch.png)

## How to use
  
import the proto files in postman and just enjoy :)
