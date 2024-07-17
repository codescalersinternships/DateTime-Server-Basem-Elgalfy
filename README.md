# DateTime-Server-Basem-Elgalfy

This project implements a simple DateTime server in GO, in both http nad gin-based servers retrieving the current date and time.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine.

## Prerequisites

- Go (1.22 or later)
- Docker (optional but preferable for easy running)
- Makefile

## Important Notes

- output is in JSON format 
- method is GET 
- for /datetime
- will be hosted on localhost
- for Docker gin will be on port 5000 and http will be on port 5001
- for single builds i.e. gin/http they will be on port 8080
  
### Installation

First clone the repository to your local machine:

```sh
git clone github.com/codescalersinternships/DateTime-Server-Basem-Elgalfy
cd DateTime-Server-Basem-Elgalfy
```


### Build

If you have docker run:

```sh
make docker-build-images
```

or if you dont for http 

```sh
make build-http 
```

and for gin 

```sh
make build-gin
```

### Run

For Docker:

```sh
make docker-run-images
```

For http

```sh
./http
```

For gin

```sh
./gin
```

## Usage

Open a web browser to [http://localhost:{portnumber}/datetime]

Or

```sh
curl http://localhost:{portnumber}/datetime
```

## Test

To run tests: 

```sh
make test
```






