# Natsbeat

Natsbeat is an elastic [Beat](https://www.elastic.co/products/beats) that reads
metrics from [NATS](https://nats.io/) monitoring endpoints and indexes them into Elasticsearch database.

## Description

> [NATS](https://nats.io/) is an open source messaging system for cloud native applications, IoT messaging, and microservices architectures

When the monitoring port is enabled, the NATS server runs a lightweight web server on port 8222 which exposes several endpoints that provide metrics in JSON format.
Natsbeat collects these metrics.

## Getting Started with Natsbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Natsbeat and also install the
dependencies, run the following command:

```sh
go get github.com/nfvsap/natsbeat
cd $GOPATH/src/github.com/nfvsap/natsbeat
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Natsbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/nfvsap/natsbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Natsbeat run the command below. This will generate a binary
in the same directory with the name natsbeat.

```sh
cd $GOPATH/src/github.com/nfvsap/natsbeat
make
```

### Run

To run Natsbeat with debugging output enabled, run:

```sh
./natsbeat -c natsbeat.yml -e -d "*"
```

### Setup Kibana

To automatically load the basic kibana dashboard for Natsbeat, run:

```sh
./natsbeat -v -e -d "*" setup -E setup.dashboards.enabled=true
```


### Test

To test Natsbeat, run the following command:

```sh
make testsuite
```

alternatively:
```sh
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```sh
make update
```


### Cleanup

To clean  Natsbeat source code, run the following command:

```sh
make fmt
```

To clean up the build directory and generated artifacts, run:

```sh
make clean
```


### Clone

To clone Natsbeat from the git repository, run the following commands:

```sh
mkdir -p ${GOPATH}/src/github.com/nfvsap/natsbeat
git clone https://github.com/nfvsap/natsbeat ${GOPATH}/src/github.com/nfvsap/natsbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```sh
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
