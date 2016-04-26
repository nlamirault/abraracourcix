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

The API supports `application/json` , encoded in UTF-8. XML is not on the roadmap.

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


## Errors

| Code | Name                  | Retry Behavior                                                                                                                                    |
|------|-----------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| 400  | Bad Request           | The body of the request invalid. The request either must be changed before being retried or depends on another request being processed before it. |
| 404  | Not Found             | The requested resource could not be found. The request must be changed before being retried.                                                      |
| 500  | Internal Server Error | The server encountered an error while processing the request. This request should be retried without change.                                      |

Example:

```json
HTTP/1.1 500 Internal Server Error
Content-Length : 28
Content-Type : application/json; charset=utf-8
Date : Tue, 26 Apr 2016 00:42:10 GMT


{
  "error": "Can't decode url"
}
```


## URLs

### GET /urls/`:url`

#### Description

Retrieve an URL using the specified short URL.

#### Example request

```
GET http://localhost:8080/api/v1/urls/8WAAfWwiID
```

#### Example response

```json
HTTP/1.1 200 OK
Content-Length : 103
Content-Type : application/json; charset=utf-8
Date : Tue, 26 Apr 2016 00:46:04 GMT

{
  "key": "8WAAfWwiID",
  "url": "http://www.google.com",
  "creation_date": "2016-04-25T23:45:49.94186238+02:00"
}
```

### POST /urls

#### Description

creates a new short URL.

#### Example request

```json
POST http://localhost:8080/api/v1/urls

{
  "url": "http://www.google.com"
}
```

#### Example response

```json
HTTP/1.1 200 OK
Content-Type : application/json; charset=utf-8
Date : Tue, 26 Apr 2016 00:02:49 GMT
Content-Length : 103

{
  "key": "8WAAfWwiID",
  "url": "http://www.google.com",
  "creation_date": "2016-04-26T00:02:49.94186238+02:00"
}
```


### GET /stats/`:url`

#### Description

Get analytics for an URL

#### Example request

```
GET http://localhost:8080/api/v1/stats/8WAAfWwiID
```

#### Example response

```json
HTTP/1.1 200 OK
Content-Type : application/json; charset=utf-8
Date : Tue, 26 Apr 2016 11:56:18 GMT
Content-Length : 72

{
  "longUrlClicks": "1",
  "shortUrlClicks": "2",
  "user_agents": {
    "bat/0.1.0": 2
  }
}
```
