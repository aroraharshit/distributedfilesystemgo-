Distributed Decentralized File Storage System

This project implements a Distributed Decentralized File Storage System using P2P (Peer-to-Peer) networking. The system allows multiple nodes (peers) to store and retrieve files in a distributed manner. Each node can act as both a file server and a client, storing files locally while also being capable of fetching files from other nodes in the network.
Features
  File Storage: Each node can store files locally and make them available to other nodes.
  File Retrieval: Nodes can request files from other peers in the network if they do not have the file stored locally.
  P2P Networking: Peer-to-peer (P2P) communication between nodes for file transfer.
  File Encryption: Files are encrypted before storage and decrypted when retrieved, ensuring secure file transfers.
  Network Bootstrapping: Nodes can connect to the network using a list of bootstrap nodes for initial peer discovery.
Architecture
  P2P Transport Layer: Handles the peer-to-peer communication and file transfer over the network.
  File Server: Each node acts as a file server, storing files locally and serving them to other nodes.
  File Storage: Files are stored locally and encrypted for security. Files can be retrieved by hash key.
  Message Protocol: A custom protocol for communication between nodes, using message types like MessageStoreFile and MessageGetFile.
