// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Key management done with Go",
    "title": "KEYSERVER",
    "version": "1.0.0"
  },
  "paths": {
    "/create": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "single key operation"
        ],
        "summary": "Return a new key",
        "operationId": "createKey",
        "parameters": [
          {
            "x-exportParamName": "Body",
            "description": "Entity object that needs to be stored",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateKey"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/KeyData"
              }
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/key/{strKey}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "single key operation"
        ],
        "summary": "Returns a single key Data by given Key as a string",
        "operationId": "keyGet",
        "parameters": [
          {
            "type": "string",
            "x-exportParamName": "strKey",
            "name": "strKey",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/KeyData"
              }
            }
          },
          "404": {
            "description": "No key was found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "CreateKey": {
      "type": "object",
      "properties": {
        "CreatedBy": {
          "type": "string"
        },
        "Expire": {
          "type": "string",
          "pattern": "/(?:[0-9]{4})-(?:[0-9]{2})-(?:[0-9]{2})/",
          "example": "2019-05-17"
        }
      }
    },
    "KeyData": {
      "type": "object",
      "properties": {
        "CreatedBy": {
          "type": "string"
        },
        "Data": {
          "type": "string"
        },
        "Expire": {
          "type": "string",
          "pattern": "/(?:[0-9]{4})-(?:[0-9]{2})-(?:[0-9]{2})/",
          "example": "2019-05-17"
        },
        "ID": {
          "type": "integer",
          "format": "int64"
        },
        "Valid": {
          "type": "boolean"
        }
      },
      "xml": {
        "name": "Entity"
      },
      "example": {
        "Description": "Description",
        "ID": 0,
        "Name": "Name",
        "Status": "Status"
      }
    }
  },
  "tags": [
    {
      "name": "single key operation"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Key management done with Go",
    "title": "KEYSERVER",
    "version": "1.0.0"
  },
  "paths": {
    "/create": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "single key operation"
        ],
        "summary": "Return a new key",
        "operationId": "createKey",
        "parameters": [
          {
            "x-exportParamName": "Body",
            "description": "Entity object that needs to be stored",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/CreateKey"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/KeyData"
              }
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/key/{strKey}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "single key operation"
        ],
        "summary": "Returns a single key Data by given Key as a string",
        "operationId": "keyGet",
        "parameters": [
          {
            "type": "string",
            "x-exportParamName": "strKey",
            "name": "strKey",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/KeyData"
              }
            }
          },
          "404": {
            "description": "No key was found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "CreateKey": {
      "type": "object",
      "properties": {
        "CreatedBy": {
          "type": "string"
        },
        "Expire": {
          "type": "string",
          "pattern": "/(?:[0-9]{4})-(?:[0-9]{2})-(?:[0-9]{2})/",
          "example": "2019-05-17"
        }
      }
    },
    "KeyData": {
      "type": "object",
      "properties": {
        "CreatedBy": {
          "type": "string"
        },
        "Data": {
          "type": "string"
        },
        "Expire": {
          "type": "string",
          "pattern": "/(?:[0-9]{4})-(?:[0-9]{2})-(?:[0-9]{2})/",
          "example": "2019-05-17"
        },
        "ID": {
          "type": "integer",
          "format": "int64"
        },
        "Valid": {
          "type": "boolean"
        }
      },
      "xml": {
        "name": "Entity"
      },
      "example": {
        "Description": "Description",
        "ID": 0,
        "Name": "Name",
        "Status": "Status"
      }
    }
  },
  "tags": [
    {
      "name": "single key operation"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
}