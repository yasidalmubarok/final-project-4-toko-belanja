// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "GLNG-KS07 - Group 5",
            "url": "https://github.com/yusrililhm/final-project-4-toko-belanja"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories": {
            "get": {
                "description": "Get categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Get categories",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Bearer",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Add category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Bearer",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body request for add category",
                        "name": "dto.CategoriesRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CategoriesRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryResponse"
                        }
                    }
                }
            }
        },
        "/categories/{categoryId}": {
            "delete": {
                "description": "Delete category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Delete category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Bearer",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Category Id",
                        "name": "CategoryId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Categories"
                ],
                "summary": "Update category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Bearer",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Category Id path",
                        "name": "CategoryId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body request for update category",
                        "name": "dto.CategoriesRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CategoriesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CategoryResponse"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Get products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get products",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Bearer",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProductResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Add product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Bearer",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body request for add product",
                        "name": "dto.ProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ProductRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.ProductResponse"
                        }
                    }
                }
            }
        },
        "/products/{productId}": {
            "put": {
                "description": "Update product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Bearer",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Product Id",
                        "name": "ProductId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body request for update product",
                        "name": "dto.ProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProductResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Bearer",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Product Id",
                        "name": "ProductId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.ProductResponse"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "post": {
                "description": "Add transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Add transaction",
                "parameters": [
                    {
                        "description": "body request for add transaction",
                        "name": "dto.TransactionRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TransactionRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Bearer",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.TransactionHistoryResponse"
                        }
                    }
                }
            }
        },
        "/transactions/my-transactions": {
            "get": {
                "description": "Get my transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get my transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Bearer",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.TransactionHistoryResponse"
                        }
                    }
                }
            }
        },
        "/transactions/user-transactions": {
            "get": {
                "description": "Get user transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get user transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Bearer",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.TransactionHistoryResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "User login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "body request for user login",
                        "name": "dto.UsersLoginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UsersLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "User register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User register",
                "parameters": [
                    {
                        "description": "body request for user register",
                        "name": "dto.CreateNewUsersRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateNewUsersRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    }
                }
            }
        },
        "/users/topup": {
            "patch": {
                "description": "User topup",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User topup",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "body request for user topup",
                        "name": "dto.UsersTopUpRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UsersTopUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CategoriesRequest": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string",
                    "example": "jersey"
                }
            }
        },
        "dto.CategoryResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.CreateNewUsersRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "maguire.harry@mufc.com"
                },
                "full_name": {
                    "type": "string",
                    "example": "Harry Maguire"
                },
                "password": {
                    "type": "string",
                    "example": "secret"
                }
            }
        },
        "dto.ProductRequest": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer",
                    "example": 1
                },
                "price": {
                    "type": "integer",
                    "example": 120000
                },
                "stock": {
                    "type": "integer",
                    "example": 10
                },
                "title": {
                    "type": "string",
                    "example": "Jersey King MU 2023/2024"
                }
            }
        },
        "dto.ProductResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.TransactionHistoryResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.TransactionRequest": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "integer",
                    "example": 1
                },
                "quantity": {
                    "type": "integer",
                    "example": 3
                }
            }
        },
        "dto.UserResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.UsersLoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "maguire.harry@mufc.com"
                },
                "password": {
                    "type": "string",
                    "example": "secret"
                }
            }
        },
        "dto.UsersTopUpRequest": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer",
                    "example": 150000
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Toko Belanja",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
