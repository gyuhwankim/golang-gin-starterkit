// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-07-09 11:03:10.49197541 +0000 UTC m=+0.068352204

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is a sample gin starter server.",
        "title": "Go Gin Starter API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "gyuhwan.a.kim@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/gghcode/go-gin-starterkit/blob/master/LICENSE"
        },
        "version": "1.0"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/healthy": {
            "get": {
                "description": "Get server healthy",
                "tags": [
                    "App API"
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/api/todos": {
            "get": {
                "description": "Get all todos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo API"
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/todo.TodoResponse"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create new todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo API"
                ],
                "parameters": [
                    {
                        "description": "todo payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.CreateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.TodoResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid todo payload",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/todos/{id}": {
            "get": {
                "description": "Get todo by todo id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo API"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.TodoResponse"
                        }
                    },
                    "404": {
                        "description": "Not found entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update todo by todo id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo API"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "todo payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.CreateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.TodoResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid todo payload",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not found entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove todo by todo id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Todo API"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/todo.TodoResponse"
                        }
                    },
                    "404": {
                        "description": "Not found entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.APIError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "common.ErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.APIError"
                    }
                }
            }
        },
        "todo.CreateTodoRequest": {
            "type": "object",
            "required": [
                "contents",
                "title"
            ],
            "properties": {
                "contents": {
                    "type": "string",
                    "example": "\u003cnew contents\u003e"
                },
                "title": {
                    "type": "string",
                    "example": "\u003cnew title\u003e"
                }
            }
        },
        "todo.TodoResponse": {
            "type": "object",
            "properties": {
                "contents": {
                    "type": "string"
                },
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
