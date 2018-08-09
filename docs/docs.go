// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-08-03 15:26:17.189363428 -0400 EDT m=+0.026239986

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is the backend API for the Personal Finance App",
        "title": "Personal Finance App",
        "contact": {
            "name": "Personal Finance Support",
            "url": "http://www.swcraftsman.com",
            "email": "admin@swcraftsman.com"
        },
        "license": {},
        "version": "0.10"
    },
    "paths": {
        "/api/account": {
            "get": {
                "description": "Get all accounts",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts account"
                ],
                "summary": "List Accounts",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Account"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create an account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts account"
                ],
                "summary": "Add Account",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/Account"
                        }
                    }
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}