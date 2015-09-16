# Abraracourcix

[![Travis](https://img.shields.io/travis/nlamirault/abraracourcix.svg)]()

This tool is a simple URL Shortener.
Storage backends supported are :

* [x] BoltDB
* [x] LevelDB
* [x] Redis
* [x] MongoDB

Some analytics are available for URLs

## Usage

We will use [bat](https://github.com/astaxie/bat) to make HTTP request

Launch web service :

    $ bin/abraracourcix
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
