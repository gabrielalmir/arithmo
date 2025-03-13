# Arithmo

Arithmo is a simple in-memory key-value store written in Go. It supports basic operations such as setting, getting, incrementing, and decrementing values.

## Features

- Set and get values
- Increment and decrement integer values
- Check existence of keys
- Delete keys
- Determine the type of stored values

## Installation

To build the project, run:

```sh
make build
```

To build the project for Windows, run:

```sh
make buildwin
```

## Usage

To start the server, run:

```sh
make run
```

The server will start on port `6379`.

## Commands

- `SET key value` - Set the value of a key
- `GET key` - Get the value of a key
- `INCR key` - Increment the integer value of a key
- `DECR key` - Decrement the integer value of a key
- `TYPE key` - Get the type of the value stored at key
- `DEL key` - Delete a key
- `EXISTS key` - Check if a key exists
- `QUIT` - Close the connection

## License

This project is licensed under the MIT License.
