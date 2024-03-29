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
        "/api/basket": {
            "post": {
                "description": "List the recipes you want",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "basket"
                ],
                "summary": "Create a new basket based on many recipes",
                "parameters": [
                    {
                        "description": "create a new basket based on some payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/basket.basketPayload"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/api/category": {
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
                "summary": "Create a new category",
                "parameters": [
                    {
                        "description": "Create a new category",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/category.CreateCategoryPayload"
                        }
                    }
                ],
                "responses": {}
            }
        },
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
        },
        "/api/recipes": {
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
                "summary": "Create a new recipe",
                "parameters": [
                    {
                        "description": "Create a new recipe",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/recipe.createRecipePayload"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "basket.basketPayload": {
            "type": "object",
            "properties": {
                "recipes": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "category.CreateCategoryPayload": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "category.SetCategoryPayload": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "chef.createChefPayload": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "food.FindFoodPayload": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "food.createFoodPayload": {
            "type": "object",
            "properties": {
                "animal_type": {
                    "type": "string"
                },
                "average_amount": {
                    "$ref": "#/definitions/measurements.UnitType"
                },
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
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "ingredient.IngredientPayload": {
            "type": "object",
            "properties": {
                "amount": {
                    "$ref": "#/definitions/measurements.UnitType"
                },
                "food": {
                    "$ref": "#/definitions/food.FindFoodPayload"
                }
            }
        },
        "measurements.UnitType": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        },
        "recipe.createRecipePayload": {
            "type": "object",
            "properties": {
                "average_amount": {
                    "$ref": "#/definitions/measurements.UnitType"
                },
                "category": {
                    "$ref": "#/definitions/category.SetCategoryPayload"
                },
                "cook_duration": {
                    "type": "integer"
                },
                "directions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/step.StepPayload"
                    }
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ingredient.IngredientPayload"
                    }
                },
                "introduction": {
                    "type": "string"
                },
                "preparation_time": {
                    "$ref": "#/definitions/time.Duration"
                },
                "serving": {
                    "type": "integer"
                },
                "summary": {
                    "type": "string"
                },
                "yield": {
                    "type": "integer"
                }
            }
        },
        "step.StepPayload": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "order": {
                    "type": "integer"
                }
            }
        },
        "time.Duration": {
            "type": "integer",
            "enum": [
                1,
                1000,
                1000000,
                1000000000,
                60000000000,
                3600000000000
            ],
            "x-enum-varnames": [
                "Nanosecond",
                "Microsecond",
                "Millisecond",
                "Second",
                "Minute",
                "Hour"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3010",
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
