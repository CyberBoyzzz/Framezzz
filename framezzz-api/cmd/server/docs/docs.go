// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "swagger": "2.0",
    "info": {
        "description": "Api Endpoints for Go Server",
        "title": "Go Rest Api",
        "contact": {},
        "version": "1.0.0"
    },
    "host": "localhost:7312",
    "basePath": "/",
    "schemes": ["http"],
    "paths": {
        "/api/comic/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comics"
                ],
                "summary": "Add a specific comic",
                "parameters": [
                    {
                        "description": "Comic object that needs to be added",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddComicRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.IDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/comic/delete/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comics"
                ],
                "summary": "Delete a specific comic",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comic ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetComicResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/comic/update": {
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comics"
                ],
                "summary": "Update a specific comic",
                "parameters": [
                    {
                        "description": "Comic object that needs to be updated",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateComicRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.IDResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/comic/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comics"
                ],
                "summary": "Get a specific comic",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Comic ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.GetComicResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/comics": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comics"
                ],
                "summary": "Get all comics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.GetComicResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.GetComicResponse": {
            "type": "object",
            "required": ["id", "title", "coverUrl", "postUrl"],
            "properties": {
                "coverUrl": {
                    "type": "string",
                    "description": "URL of the comic's cover image"
                },
                "id": {
                    "type": "integer",
                    "description": "Unique identifier of the comic"
                },
                "postUrl": {
                    "type": "string",
                    "description": "URL where the comic can be accessed"
                },
                "title": {
                    "type": "string",
                    "description": "Title of the comic"
                }
            }
        },
        "model.AddComicRequest": {
            "type": "object",
            "required": ["title", "coverUrl", "postUrl"],
            "properties": {
                "title": {
                    "type": "string",
                    "description": "Title of the comic"
                },
                "coverUrl": {
                    "type": "string",
                    "description": "URL of the comic's cover image"
                },
                "postUrl": {
                    "type": "string",
                    "description": "URL where the comic can be accessed"
                }
            }
        },
        "model.UpdateComicRequest": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string",
                    "description": "Title of the comic"
                },
                "coverUrl": {
                    "type": "string",
                    "description": "URL of the comic's cover image"
                },
                "postUrl": {
                    "type": "string",
                    "description": "URL where the comic can be accessed"
                }
            }
        },
        "model.IDResponse": {
            "type": "object",
            "required": ["id"],
            "properties": {
                "id": {
                    "type": "integer",
                    "description": "Unique identifier of the comic"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Go Rest Api",
	Description:      "Api Endpoints for Go Server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
