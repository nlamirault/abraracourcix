package pb 

const (
api = `{"swagger":"2.0","basePath":"","info":{"title":"Abraracourcix REST API","version":"1.0.0","description":"\nFor more information about the usage of the Abraracourcix REST API, see\n[https://github.com/nlamirault/abraracourcix](https://github.com/nlamirault/abraracourcix).\n"},"schemes":null,"consumes":["application/json"],"produces":["application/json"],"paths":{"/v2beta/urls":{"get":{"operationId":"List","responses":{"200":{"description":"","schema":{"$ref":"#/definitions/v2betaGetUrlsResponse"}}},"summary":"List returns all available Urls","tags":["UrlService"]},"post":{"operationId":"Create","parameters":[{"in":"body","name":"body","required":true,"schema":{"$ref":"#/definitions/v2betaCreateUrlRequest"}}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/v2betaCreateUrlResponse"}}},"summary":"Create creates a new Url","tags":["UrlService"]}},"/v2beta/urls/{key}":{"get":{"operationId":"Get","parameters":[{"format":"string","in":"path","name":"key","required":true,"type":"string"}],"responses":{"200":{"description":"","schema":{"$ref":"#/definitions/v2betaGetUrlResponse"}}},"summary":"Get return a Url","tags":["UrlService"]}}},"definitions":{"v2betaCreateUrlRequest":{"properties":{"link":{"format":"string","type":"string"}},"type":"object"},"v2betaCreateUrlResponse":{"properties":{"url":{"$ref":"#/definitions/v2betaUrl"}},"type":"object"},"v2betaGetUrlRequest":{"properties":{"key":{"format":"string","type":"string"}},"type":"object"},"v2betaGetUrlResponse":{"properties":{"Url":{"$ref":"#/definitions/v2betaUrl"}},"type":"object"},"v2betaGetUrlsRequest":{"type":"object"},"v2betaGetUrlsResponse":{"properties":{"Urls":{"items":{"$ref":"#/definitions/v2betaUrl"},"type":"array"}},"type":"object"},"v2betaUrl":{"properties":{"creation":{"format":"string","type":"string"},"key":{"format":"string","type":"string"},"link":{"format":"string","type":"string"}},"type":"object"}}}
`
urls = `{
  "swagger": "2.0",
  "info": {
    "title": "urls.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v2beta/urls": {
      "get": {
        "summary": "List returns all available Urls",
        "operationId": "List",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v2betaGetUrlsResponse"
            }
          }
        },
        "tags": [
          "UrlService"
        ]
      },
      "post": {
        "summary": "Create creates a new Url",
        "operationId": "Create",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v2betaCreateUrlResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v2betaCreateUrlRequest"
            }
          }
        ],
        "tags": [
          "UrlService"
        ]
      }
    },
    "/v2beta/urls/{key}": {
      "get": {
        "summary": "Get return a Url",
        "operationId": "Get",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/v2betaGetUrlResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "key",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "string"
          }
        ],
        "tags": [
          "UrlService"
        ]
      }
    }
  },
  "definitions": {
    "v2betaCreateUrlRequest": {
      "type": "object",
      "properties": {
        "link": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "v2betaCreateUrlResponse": {
      "type": "object",
      "properties": {
        "url": {
          "$ref": "#/definitions/v2betaUrl"
        }
      }
    },
    "v2betaGetUrlRequest": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string",
          "format": "string"
        }
      }
    },
    "v2betaGetUrlResponse": {
      "type": "object",
      "properties": {
        "Url": {
          "$ref": "#/definitions/v2betaUrl"
        }
      }
    },
    "v2betaGetUrlsRequest": {
      "type": "object"
    },
    "v2betaGetUrlsResponse": {
      "type": "object",
      "properties": {
        "Urls": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v2betaUrl"
          }
        }
      }
    },
    "v2betaUrl": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string",
          "format": "string"
        },
        "link": {
          "type": "string",
          "format": "string"
        },
        "creation": {
          "type": "string",
          "format": "string"
        }
      }
    }
  }
}
`
)
