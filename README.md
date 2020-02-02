# Kingdom

Your favourite princesses API.

## Install

```sh
$ go get -u github.com/kingdom-tower/kingdom
```

## Usage

```sh
$ kingdom --help
Usage:
  kingdom [command]

Available Commands:
  help        Help about any command
  server      Runs the kingdom API server

Flags:
  -h, --help   help for kingdom

Use "kingdom [command] --help" for more information about a command.
```

### Run server

To run the server, you need an existing folder that would contain the images of your princesses (by default `./data`):

```sh
$ kingdom server --help
Runs the kingdom API server

Usage:
  kingdom server [flags]

Flags:
  -d, --dir string   The directory where the server reads and stores the files (default "./data")
  -h, --help         help for server
  -p, --port int     The port to run the server at (default 8080)
```

Server flags can be set as well through the usage of environment variables. The variable names are the same as the flag long names, but in uppercase and with the prefix `KINGDOM_`, so `--port` would be `KINGDOM_PORT` and `--dir` would be `KINGDOM_DIR`.

## API

| Method | Endpoint         | Description               | Body                                          |
|--------|------------------|---------------------------|-----------------------------------------------|
| GET    | /princess/{id}   | Returns a princess image  |                                               |
| POST   | /princess        | Uploads a princess        | Type form, file = the file you want to upload |
| GET    | /princess/random | Returns a random princess |                                               |

To upload a file:

```sh
$ curl -X POST http://localhost:8080 -F 'file=@my/file/path'
```
