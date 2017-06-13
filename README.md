# Abraracourcix

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fabraracourcix.svg)](https://badge.fury.io/gh/nlamirault%2Fabraracourcix)

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/abraracourcix/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/abraracourcix/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/abraracourcix/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/abraracourcix/tree/develop)

This tool is a simple URL Shortener.
*Abraracourcix* uses [gRPC](http://www.grpc.io) for its message protocol.

The project includes 2 command line utilities:

* *abraracourcixctl*, to communicate with a server *abraracourcixd*
* *abraracourcixadm*, an administration tool to manage the server.

For gRPC-supported languages, *Abraracourcix* provides a JSON gateway, which provides a RESTful proxy that translates HTTP / JSON requests to gRPC messages.

Metrics for [Prometheus](https://prometheus.io/) are exported.

Application traces are available using [OpenTracing](http://opentracing.io/). Supported systems are:

* [X] [Jaeger] (https://github.com/uber/jaeger)
* [X] [Zipkin] (https://github.com/openzipkin)

Storage backends supported are :

* [BoltDB][https://github.com/boltdb/bolt]
* [LevelDB][http://leveldb.org/]
* [Redis][https://redis.io/]
* [MongoDB][https://www.mongodb.org/]


## Installation

You can download the binaries :
* Architecture i386 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_linux_386) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_openbsd_386) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetd-0.2.0_netbsd_arm) ]

### Trinquetctl

* Architecture i386 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_linux_386) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_openbsd_386) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/pilotariak/oss/trinquetctl-0.2.0_netbsd_arm) ]

### Trinquetadm

* Architecture i386 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_linux_386) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_openbsd_386) / [windows](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcixadm-2.0.0_netbsd_arm) ]



## Usage

Launch Zipkin with Docker, and open a browser on 9411:

    $ docker run -d -p 9411:9411 openzipkin/zipkin

or Jaeger with Docker (open a browser on 16686):

    $ docker run -d -p5775:5775/udp -p16686:16686 jaegertracing/all-in-one:latest


Use the *abraracourcixd* CLI to launch a server:

    $ abraracourcixd run -config abraracourcix.toml -v 2 -logtostderr

Configure CLI:

    $ export ABRARACOURCIX_SERVER="localhost:8080"
    $ export ABRARACOURCIX_USERNAME="admin"
    $ export ABRARACOURCIX_APIKEY="nimda"

Use the *abraracourcixctl* CLI to use the URLs informations :

    $ abraracourcixctl url list
    URLs:
    $ abraracourcixctl url add --link https://news.google.fr/
    URL:
    - Key: 5X81VNCDVq
    - Link: https://news.google.fr/
    - Date: 2017-06-13 17:08:48.898220331 +0200 CEST
    $ abraracourcixctl url list
    URLs:
    - 5X81VNCDVq
    $ abraracourcixctl url get --key 5X81VNCDVq
    URL:
    - Key: 5X81VNCDVq
    - Link: https://news.google.fr/
    - Date: 2017-06-13 17:08:48.898220331 +0200 CEST

Use the *abraracourcixadm* CLI to manage the server.

    $ abraracourcixadm health
    +------------+--------+---------+
    |  SERVICE   | STATUS |  TEXT   |
    +------------+--------+---------+
    | UrlService | OK     | SERVING |
    +------------+--------+---------+

    $ abraracourcixadm info
    +----------------+----------------+---------+--------+
    |    SERVICE     |      URI       | VERSION | STATUS |
    +----------------+----------------+---------+--------+
    | abraracourcixd | localhost:8080 | 2.0.0   | OK     |
    +----------------+----------------+---------+--------+


You could explore the API using [Swagger](http://swagger.io/) UI :

    http://localhost:9090/swagger-ui/


## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Start backends :

        $ docker run -d -p 6379:6379 --name redis redis:3
        $ docker run -d -p 27017:27017 --name mongo mongo:3.1

* Launch unit tests :

        $ make test


## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
