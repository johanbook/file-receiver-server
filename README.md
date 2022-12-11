# file-receiver-server

![version](https://img.shields.io/github/v/tag/johanbook/file-receiver-server)
![docker pulls](https://img.shields.io/docker/pulls/johanbook/file-receiver-server)
![vulnerabilities](https://img.shields.io/snyk/vulnerabilities/github/johanbook/file-receiver-server)

**file-receiver-server** is a simplistic Go HTTP server for receiving, unzipping and storing archives.

The project aims to be minimal production dependency.

## Get started

The server can be run either using a Golang compilter or through Docker.

### Using Golang

Install and run the server using

```sh
ROOT_FILE_PATH=/path/to/folder ./scripts/run
```

### Using Docker

The server can be run through Docker by mounting in the served directory as a
volume. To the serve the `build` folder in the current directory, run

```sh
docker run \
	--detach \
	--publish 8080:8080 \
	--restart unless-stopped \
	--volume ${PWD}/build/:/app/build/ \
	johanbook/file-receiver-server:latest
```

The Docker images are available for multiple architectures, including armv7.

## Configuration

The following can be configured as environment variables:

- **ACCESS_LOGS** if access logs are showed (default `false`).
- **LOG_LEVEL** used log level. Available values are `debug`, `info`, `warning`
  and `error` (default `info`).
- **PORT** the port the server should listen on (default 8080).
- **ROOT_FILE_PATH** root directory of served file tree (default `build`).

## Documentation

For documentation on technologies, security aspect and similar, see the
[contributions file](./CONTRIBUTING.md).

## Contributing

For contributions and development procedures, see the
[contributions file](./CONTRIBUTING.md).
