# API

Welcome to the Abraracourcix API.

## Overview

### Verbs

The verbs used in the API are:

* `GET` : Used to retrieve resources.
* `POST` : Used to create resources, or perform custom actions
* `PUT` : Used to replace the resources or collections.
* `DELETE` : Used for deleting resources.

### Media types

The API supports `application/json` , encoded in UTF -8.
XML is not on the roadmap.

### Version

We use a management system of the versions of the URI.
The first element of the URI contains the identifier of the target version,
for example :

    http://abraracourix:8080/api/num_version

The current version is `v1`


## Authentication

The API supports Basic Authentication as defined in
[RFC2617](http://www.ietf.org/rfc/rfc2617.txt).

To use Basic Authentication, simply send the username and password associated with the account (cURL will prompt you to enter the password) :

    $ curl -u <username> http://abraracourix:8080/user


## URLs

### Get an URL

Request :

    GET /urls/:url

Response :

    {
        "url": "http://www.google.com",
        "key": "0ZooGs0wiB"
    }


### Creates a shorten URL

Request :

    POST /urls

Parameters :

    url :  a long URL to be shortened

Response :

    {
        "url": "http://www.google.com",
        "key": "0ZooGs0wiB"
    }
