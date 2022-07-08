// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Matheus Carmo",
            "url": "http://www.swagger.io/support",
            "email": "mematheuslc@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/chefs": {
            "post": {
                "description": "You just need your name and your e-mail",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chef"
                ],
                "summary": "Create a new chef",
                "parameters": [
                    {
                        "description": "Create a new chef",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/chef.createChefPayload"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/foods": {
            "post": {
                "description": "Creates a new food which can be used within recipes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "food"
                ],
                "summary": "Create a new food",
                "parameters": [
                    {
                        "description": "Create a new food",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/food.createFoodPayload"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "chef.createChefPayload": {
            "type": "object",
            "properties": {
                "Name": {
                    "type": "string"
                },
                "Role": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "food.createFoodPayload": {
            "type": "object",
            "properties": {
                "family": {
                    "type": "string"
                },
                "genus": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "order": {
                    "type": "string"
                },
                "scientific_name": {
                    "type": "string"
                },
                "specie": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7777",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Guiomar API",
	Description:      "Guiomar private and public API docs.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
