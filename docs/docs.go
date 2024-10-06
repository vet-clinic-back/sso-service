// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/v1/sign-in": {
            "post": {
                "description": "Sign in user with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign in",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully signed in",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid input body",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    }
                }
            }
        },
        "/auth/v1/sign-up/owner": {
            "post": {
                "description": "Sign up a new owner by providing their details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "owners"
                ],
                "summary": "Sign up a new owner",
                "parameters": [
                    {
                        "description": "Owner input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Owner"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully signed up. Token returned",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid input body",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "409": {
                        "description": "owner with same email already exists",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    }
                }
            }
        },
        "/auth/v1/sign-up/vet": {
            "post": {
                "description": "Sign up a new vet by providing their details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vets"
                ],
                "summary": "Sign up a new vet",
                "parameters": [
                    {
                        "description": "Vet input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Vet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully signed up. Token returned",
                        "schema": {
                            "$ref": "#/definitions/models.SuccessDTO"
                        }
                    },
                    "400": {
                        "description": "Invalid input body",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "409": {
                        "description": "vet with same email already exists",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrorDTO": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Owner": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "description": "password hash",
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "models.SuccessDTO": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "description": "password hash",
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "models.Vet": {
            "type": "object",
            "properties": {
                "clinic_number": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "description": "password hash",
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Vet clinic auth service",
	Description:      "auth service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}