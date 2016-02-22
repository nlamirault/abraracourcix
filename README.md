# Abraracourcix

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fabraracourcix.svg)](https://badge.fury.io/gh/nlamirault%2Fabraracourcix)

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/abraracourcix/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/abraracourcix/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/abraracourcix/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/abraracourcix/tree/develop)

This tool is a simple URL Shortener.
Storage backends supported are :

* [BoltDB][]
* [LevelDB][]
* [Redis][]
* [MongoDB][]

Some analytics are available for URLs.

Rest API : [API.md](API.md)

## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_linux_386) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_openbsd_386) / [windows](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/abraracourcix-0.8.0_netbsd_arm) ]


## Usage

We will use [bat](https://github.com/astaxie/bat) to make HTTP request

Launch web service :

    $ abraracourcix
    INFO[0000] Launch Abraracourcix on 8080 using boltdb backend

Store an URL :

    $ bat POST http://localhost:8080/api/v1/urls url:="http://www.google.com"
    POST /api/v1/urls HTTP/1.1
    Host: localhost:8080
    Accept: application/json
    Accept-Encoding: gzip, deflate
    Content-Type: application/json
    User-Agent: bat/0.0.2

    {"url":"http://www.google.com"}


    HTTP/1.1 200 OK
    Content-Type : application/json; charset=utf-8
    Date : Tue, 08 Sep 2015 00:10:48 GMT
    Content-Length : 51


    {
        "url": "http://www.google.com",
        "key": "0ZooGs0wiB"
    }

Retrieve it using the key :

    $ bat http://localhost:8080/api/v1/urls/0ZooGs0wiB
    GET /api/v1/urls/0ZooGs0wiB HTTP/1.1
    Host: localhost:8080
    Accept: application/json
    Accept-Encoding: gzip, deflate
    User-Agent: bat/0.0.2


    HTTP/1.1 200 OK
    Content-Length : 51
    Content-Type : application/json; charset=utf-8
    Date : Tue, 08 Sep 2015 00:11:19 GMT


    {
        "url": "http://www.google.com",
        "key": "0ZooGs0wiB"
    }


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

[BoltDB]: https://github.com/boltdb/bolt
[LevelDB]: http://leveldb.org/
[Redis]: http://redis.io/
[MongoDB]: https://www.mongodb.org/
