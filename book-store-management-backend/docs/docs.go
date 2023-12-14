// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Bui Vi Quoc",
            "url": "https://www.facebook.com/bviquoc/",
            "email": "21520095@gm.uit.edu.vn"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/authors": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Get all authors",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/authormodel.ResListAuthor"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Create author with name",
                "parameters": [
                    {
                        "description": "Create author",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authormodel.ReqCreateAuthor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "author id",
                        "schema": {
                            "$ref": "#/definitions/authormodel.ResCreateAuthor"
                        }
                    }
                }
            }
        },
        "/booktitles": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "booktitles"
                ],
                "summary": "Get all booktitles",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 10,
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 11,
                        "name": "total",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1709500431,
                        "name": "createdAtFrom",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1709500431,
                        "name": "createdAtTo",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "example": true,
                        "name": "isActive",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/booktitlemodel.ResListBookTitle"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "booktitles"
                ],
                "summary": "Create booktitle name, desc, authors, categories.",
                "parameters": [
                    {
                        "description": "Create booktitle",
                        "name": "booktitle",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/booktitlemodel.ReqCreateBookTitle"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "booktitle id",
                        "schema": {
                            "$ref": "#/definitions/booktitlemodel.ResCreateBookTitle"
                        }
                    }
                }
            }
        },
        "/booktitles/:id": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "booktitles"
                ],
                "summary": "Delete booktitle by id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ResSuccess"
                        }
                    }
                }
            }
        },
        "/categories": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Get all categories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/categorymodel.ResListCategory"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "categories"
                ],
                "summary": "Create category with name",
                "parameters": [
                    {
                        "description": "Create category",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/categorymodel.ReqCreateCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "category id",
                        "schema": {
                            "$ref": "#/definitions/categorymodel.ResCreateCategory"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authormodel.Author": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "authormodel.Filter": {
            "type": "object",
            "properties": {
                "searchKey": {
                    "type": "string"
                }
            }
        },
        "authormodel.ReqCreateAuthor": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Nguyễn Nhật Ánh"
                }
            }
        },
        "authormodel.ResCreateAuthor": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "authormodel.ResListAuthor": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/authormodel.Author"
                    }
                },
                "filter": {
                    "$ref": "#/definitions/authormodel.Filter"
                },
                "paging": {
                    "$ref": "#/definitions/common.Paging"
                }
            }
        },
        "booktitlemodel.BookTitle": {
            "type": "object",
            "properties": {
                "authorIds": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "categoryIds": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isActive": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "booktitlemodel.Filter": {
            "type": "object",
            "properties": {
                "createdAtFrom": {
                    "type": "integer",
                    "example": 1709500431
                },
                "createdAtTo": {
                    "type": "integer",
                    "example": 1709500431
                },
                "isActive": {
                    "type": "boolean",
                    "example": true
                },
                "searchKey": {
                    "type": "string",
                    "example": ""
                }
            }
        },
        "booktitlemodel.ReqCreateBookTitle": {
            "type": "object",
            "properties": {
                "authorIds": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "tgnna"
                    ]
                },
                "categoryIds": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "dmtt",
                        "dmtruyen"
                    ]
                },
                "desc": {
                    "type": "string",
                    "example": "Tôi Là Bêtô là tác phẩm của nhà văn chuyên viết cho thanh thiếu niên Nguyễn Nhật Ánh."
                },
                "id": {
                    "type": "string",
                    "example": "bookId"
                },
                "name": {
                    "type": "string",
                    "example": "Tôi là Bêtô"
                }
            }
        },
        "booktitlemodel.ResBookTitleDetail": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "booktitlemodel.ResCreateBookTitle": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "booktitlemodel.ResListBookTitle": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/booktitlemodel.BookTitle"
                    }
                },
                "filter": {
                    "$ref": "#/definitions/booktitlemodel.Filter"
                },
                "paging": {
                    "$ref": "#/definitions/common.Paging"
                }
            }
        },
        "categorymodel.Category": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "categorymodel.Filter": {
            "type": "object",
            "properties": {
                "searchKey": {
                    "type": "string"
                }
            }
        },
        "categorymodel.ReqCreateCategory": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Trinh thám"
                }
            }
        },
        "categorymodel.ResCreateCategory": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "categorymodel.ResListCategory": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/categorymodel.Category"
                    }
                },
                "filter": {
                    "$ref": "#/definitions/categorymodel.Filter"
                },
                "paging": {
                    "$ref": "#/definitions/common.Paging"
                }
            }
        },
        "common.Paging": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer",
                    "example": 10
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "total": {
                    "type": "integer",
                    "example": 11
                }
            }
        },
        "common.ResSuccess": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "boolean",
                    "example": true
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Book Store Management API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
