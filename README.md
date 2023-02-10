# JSON Microservice
### Currently a WIP, this project is intended to be something small at the beginning and serve as a foundation for more to come

## Getting Started

These instructions will help you set up the project and run it on your local machine.
Prerequisites

- Go version 1.14 or later
- Docker and docker-compose (if you want to run the service using docker-compose)

## Installing

Clone the repository to your local machine:

`$ git clone https://github.com/suedoh/json-microservice.git`

## Project Structure

This project contains a basic implementation of a JSON microservice in Go that uses the net package to handle concurrent connections over TCP. The service provides a simple protocol for communication between clients and the server.
Project Structure

The project contains the following three Go files:

- `protocol.go`: This file contains the structs and functions used to define the communication protocol. It includes the Request struct that is used to encode and decode messages between the client and server, as well as the processMessage function that handles incoming requests.

- `server.go`: This file contains the code for the server component of the microservice. It listens for incoming connections, creates a new goroutine for each connection to handle incoming messages, and uses the processMessage function to handle the messages.

- `client.go`: This file contains the code for the client component of the microservice. It connects to the server, sends a request, and prints the response.

## Running the Microservice with Make

To run the service using the make command, simply run the following in the root of the project:

`$ make all`

The included Makefile provides the following commands:

`make server`: This command will compile the server.go file into an executable named server.

`make client`: This command will compile the client.go file into an executable named client.

`make all`: This command will compile both the server and client components.

To run the microservice, first start the server by executing the server executable. Then, in another terminal window, start the client by executing the client executable. The client will connect to the server, send a request, and print the response.

## Running the Service using Docker Compose

To run the service using docker-compose, navigate to the root of the project and run the following command:

`$ docker-compose up`

This will build a Docker image for the service using the Dockerfile included in the project, and run the image as a container.

## Built With

- [Go](https://golang.org/) - The programming language used
- [Docker](https://www.docker.com/) - Containerization technology used in the project

## Author

[suedoh] - Initial work - [Github](github.com/suedoh)

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](./LICENSE) file for details.
