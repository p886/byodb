# Build Your Own Database

A simple key value store I built as a learning-project. **Not** intended for production, use a proper database for that ;).

## Installation

With go installed, compile the code:

```
go build
```
## Usage

### REPL

#### Store Data

Inside the REPL type `PUT` followed by a key and a value:

```
PUT hello world
```
the database will respond with a success or error message.

#### Retrieve data

In the REPL type `GET` followed by a key:

```
GET hello
```
the database will respond with a success message containing the most recent value stored for that key. In case of error an error message will be printed.

### CLI

#### Store Data

```
./byod put [key] [value]
```

#### Retrieve Data

```
./byod get [key]
```

## Configuration

Configuration is stored in a `.env` file. See `.env.sample` for config option names and example values.

## TODOS

- Proper console / REPL using a readline library
- ~~CLI interface~~
- ACIDity
- Indexing
- Transactions
- HTTP interface
- Dockerization
- Distributed cluster mode

… and much, much more.
