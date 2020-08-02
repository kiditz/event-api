// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/digital-staff": {
            "post": {
                "description": "Create a new digital staff category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "digitalStaff"
                ],
                "summary": "AddDigitalStaff used to categories help digital staff",
                "parameters": [
                    {
                        "description": "New DigitalStaff",
                        "name": "digitallStaff",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.DigitalStaff"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/translate.ResultSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entity.DigitalStaff"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/translate.ResultErrors"
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "AddUser api used to signup",
                "parameters": [
                    {
                        "description": "New User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/translate.ResultSuccess"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entity.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/translate.ResultErrors"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.DigitalStaff": {
            "type": "object",
            "required": [
                "image",
                "subtitle",
                "title"
            ],
            "properties": {
                "followers": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "subtitle": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "type"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "translate.ResultErrors": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Dynamic message"
                },
                "status": {
                    "type": "string",
                    "example": "Bad Request"
                },
                "status_code": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "translate.ResultSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "status": {
                    "type": "string",
                    "example": "OK"
                },
                "status_code": {
                    "type": "integer",
                    "example": 200
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8000",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Spgku Application",
	Description: "This is event staffing application management",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
