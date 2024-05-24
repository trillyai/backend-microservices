// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "API Support",
            "email": "fiber@swagger.io"
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
        "/comments": {
            "get": {
                "description": "Get a list of comments by username or post ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Get comments by username or post ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Post ID",
                        "name": "postId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
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
                            "$ref": "#/definitions/shared.Comments"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update an existing comment with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Update an existing comment",
                "parameters": [
                    {
                        "description": "Comment to update",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shared.UpdateCommentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/shared.UpdateCommentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new comment with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Create a new comment",
                "parameters": [
                    {
                        "description": "Comment to create",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shared.CreateCommentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/shared.CreateCommentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete an existing comment with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Delete an existing comment",
                "parameters": [
                    {
                        "description": "Comment to delete",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shared.DeleteCommentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/shared.DeleteCommentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/comments/{commentId}": {
            "get": {
                "description": "Get a comment by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "comments"
                ],
                "summary": "Get a comment by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comment ID",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/shared.Comment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/likes": {
            "get": {
                "description": "Get a list of likes by post ID or comment ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "likes"
                ],
                "summary": "Get likes by post ID or comment ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Post ID",
                        "name": "postId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Comment ID",
                        "name": "commentId",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
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
                            "$ref": "#/definitions/shared.Likes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new like for a post or a comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "likes"
                ],
                "summary": "Create a new like",
                "parameters": [
                    {
                        "description": "Like to create",
                        "name": "like",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shared.CreateLikeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/shared.CreateLikeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a like by providing the details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "likes"
                ],
                "summary": "Delete a like",
                "parameters": [
                    {
                        "description": "Like to delete",
                        "name": "like",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shared.DeleteLikeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/shared.DeleteLikeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Check if the API is running",
                "tags": [
                    "health"
                ],
                "summary": "Check API status",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "description": "Get a list of posts by user ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get posts by user ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
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
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/shared.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update an existing post with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Update an existing post",
                "parameters": [
                    {
                        "description": "Post to update",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shared.UpdatePostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/shared.UpdatePostResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new post with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Create a new post",
                "parameters": [
                    {
                        "description": "Post to create",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shared.CreatePostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/shared.CreatePostResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete an existing post with the given details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Delete an existing post",
                "parameters": [
                    {
                        "description": "Post to delete",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/shared.DeletePostRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/shared.DeletePostReesponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/posts/{postId}": {
            "get": {
                "description": "Get a post by its ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get a post by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Post ID",
                        "name": "postId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/shared.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "shared.Comment": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "postId": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "shared.Comments": {
            "type": "object",
            "properties": {
                "commentCount": {
                    "type": "integer"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/shared.Comment"
                    }
                }
            }
        },
        "shared.CreateCommentRequest": {
            "type": "object",
            "required": [
                "comment",
                "postId"
            ],
            "properties": {
                "comment": {
                    "type": "string"
                },
                "postId": {
                    "type": "string"
                },
                "username": {
                    "description": "extract from token",
                    "type": "string"
                }
            }
        },
        "shared.CreateCommentResponse": {
            "type": "object",
            "properties": {
                "createdDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "shared.CreateLikeRequest": {
            "type": "object",
            "properties": {
                "commentId": {
                    "type": "string"
                },
                "postId": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "shared.CreateLikeResponse": {
            "type": "object",
            "properties": {
                "createdDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "shared.CreatePostRequest": {
            "type": "object",
            "required": [
                "description",
                "tripId"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "tripId": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "shared.CreatePostResponse": {
            "type": "object",
            "properties": {
                "createdDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "shared.DeleteCommentRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "shared.DeleteCommentResponse": {
            "type": "object",
            "properties": {
                "deletedDate": {
                    "type": "string"
                }
            }
        },
        "shared.DeleteLikeRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "shared.DeleteLikeResponse": {
            "type": "object",
            "properties": {
                "deletedDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "shared.DeletePostReesponse": {
            "type": "object",
            "properties": {
                "deletedDate": {
                    "type": "string"
                }
            }
        },
        "shared.DeletePostRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "shared.Like": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "shared.Likes": {
            "type": "object",
            "properties": {
                "likeCount": {
                    "type": "integer"
                },
                "likes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/shared.Like"
                    }
                }
            }
        },
        "shared.Post": {
            "type": "object",
            "properties": {
                "commentCount": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "likeCount": {
                    "type": "integer"
                },
                "stations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/shared.Station"
                    }
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "tripId": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "shared.Station": {
            "type": "object",
            "properties": {
                "details": {},
                "lat": {
                    "type": "string"
                },
                "long": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "shared.UpdateCommentRequest": {
            "type": "object",
            "required": [
                "comment",
                "id"
            ],
            "properties": {
                "comment": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "shared.UpdateCommentResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "updateDate": {
                    "type": "string"
                }
            }
        },
        "shared.UpdatePostRequest": {
            "type": "object",
            "required": [
                "description",
                "id",
                "tripId"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "tripId": {
                    "type": "string"
                }
            }
        },
        "shared.UpdatePostResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "tripId": {
                    "type": "string"
                },
                "updateDate": {
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
	Version:          "1.0",
	Host:             "localhost:8082",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Interaction Server API",
	Description:      "This is an authentication server API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
