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
        "/accounts": {
            "get": {
                "description": "Retrieves all bank accounts' details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Get all bank accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/responses.Account"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new bank account with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Create a new bank account",
                "parameters": [
                    {
                        "description": "Account details",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.AccountRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.Account"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/accounts/{id}": {
            "get": {
                "description": "Retrieves a bank account's details by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Get a bank account by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.Account"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/accounts/{id}/transactions": {
            "get": {
                "description": "Retrieves all transactions for the specified bank account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get account transactions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "id",
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
                                "$ref": "#/definitions/responses.Transaction"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new transaction for the specified bank account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Create a new transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Transaction details",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.TransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.Transaction"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        },
        "/transfer": {
            "post": {
                "description": "Transfer funds from one account to another",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Transfer funds between accounts",
                "parameters": [
                    {
                        "description": "Transfer details",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.TransferRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.Message"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "requests.AccountRequest": {
            "type": "object",
            "properties": {
                "inital_balance": {
                    "type": "number"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "requests.TransactionRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "requests.TransferRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "from_acount_id": {
                    "type": "string"
                },
                "to_account_id": {
                    "type": "string"
                }
            }
        },
        "responses.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                }
            }
        },
        "responses.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "responses.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "responses.Transaction": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1/",
	Schemes:          []string{},
	Title:            "Bank Account Manager API",
	Description:      "RESTful API endpoints for Bank Account Management",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
