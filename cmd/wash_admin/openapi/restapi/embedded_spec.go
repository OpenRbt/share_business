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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "microservice for the bonus system of self-service car washes",
    "title": "wash-admin",
    "version": "1.0.0"
  },
  "paths": {
    "/healthCheck": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "Standard"
        ],
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "ok": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    },
    "/wash-server": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash_servers"
        ],
        "operationId": "get",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashServerGet"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/WashServer"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "WashServer not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash_servers"
        ],
        "operationId": "add",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashServerAdd"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Success creation"
          },
          "404": {
            "description": "WashServer not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash_servers"
        ],
        "operationId": "delete",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashServerDelete"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "404": {
            "description": "WashServer not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash_servers"
        ],
        "operationId": "update",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashServerUpdate"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Success update"
          },
          "404": {
            "description": "WashServer not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "WashServer": {
      "type": "object",
      "properties": {
        "api_key": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "WashServerAdd": {
      "required": [
        "name"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "WashServerDelete": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "WashServerGet": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "WashServerUpdate": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "authKey": {
      "description": "Session token inside Authorization header.",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "authKey": []
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "microservice for the bonus system of self-service car washes",
    "title": "wash-admin",
    "version": "1.0.0"
  },
  "paths": {
    "/healthCheck": {
      "get": {
        "security": [
          {}
        ],
        "tags": [
          "Standard"
        ],
        "operationId": "healthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "ok": {
                  "type": "boolean"
                }
              }
            }
          }
        }
      }
    },
    "/wash-server": {
      "get": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash_servers"
        ],
        "operationId": "get",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashServerGet"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/WashServer"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "404": {
            "description": "WashServer not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash_servers"
        ],
        "operationId": "add",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashServerAdd"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Success creation"
          },
          "404": {
            "description": "WashServer not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash_servers"
        ],
        "operationId": "delete",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashServerDelete"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "OK"
          },
          "404": {
            "description": "WashServer not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "security": [
          {
            "authKey": []
          }
        ],
        "tags": [
          "wash_servers"
        ],
        "operationId": "update",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/WashServerUpdate"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Success update"
          },
          "404": {
            "description": "WashServer not exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "500": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "WashServer": {
      "type": "object",
      "properties": {
        "api_key": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "WashServerAdd": {
      "required": [
        "name"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "WashServerDelete": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "WashServerGet": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "WashServerUpdate": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "Either same as HTTP Status Code OR \u003e= 600.",
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "authKey": {
      "description": "Session token inside Authorization header.",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "security": [
    {
      "authKey": []
    }
  ]
}`))
}
