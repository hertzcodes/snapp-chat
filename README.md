# Snapp Chat

Snapp Chat is a real-time chatroom application built using Go (version 1.22.0). It leverages NATS JetStream for messaging and PostgreSQL for data storage. The project consists of a client and server, allowing users to communicate seamlessly in a chatroom environment.

## Features

- Real-time messaging using NATS JetStream
- Persistent message storage with Jetstream storage
- Chat history
- Client-server architecture
- Easy to set up and run using Docker Compose

## Technologies Used

- Go 1.22.0
- NATS JetStream
- PostgreSQL
- Docker & Docker Compose

## Usage & Installation

- Fill the `config.json` files for both client and server. You can check the samples in each directory.
- Download the client from [releases page](https://github.com/hertzcodes/snapp-chat/releases/)
- Run the server using `docker-compose up -d` in the project directory
- You can also download a NATs server and run `nats-server-c nats-server.conf` & then run the binary version of the server in the [releases page](https://github.com/hertzcodes/snapp-chat/releases/). Make sure you have your Postgresql database set. 