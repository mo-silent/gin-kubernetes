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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/base/captcha": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "生成验证码,返回包括随机数id,base64,验证码长度",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/base/initdata": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "生成验证码,返回包括随机数id,base64,验证码长度",
                "parameters": [
                    {
                        "description": "初始化数据库参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/conf.Mysql"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/base/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名，密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/deployment/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment"
                ],
                "summary": "创建 Deployment",
                "parameters": [
                    {
                        "description": "Deployment simple configuration",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.DeploymentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/deployment/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment"
                ],
                "summary": "删除单个 Deployment",
                "parameters": [
                    {
                        "type": "string",
                        "default": "default",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "deployment 名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/deployment/get": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment"
                ],
                "summary": "获取单个 Deployment",
                "parameters": [
                    {
                        "type": "string",
                        "default": "default",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "deployment 名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/deployment/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment"
                ],
                "summary": "获取命名空间下的所有 Deployment",
                "parameters": [
                    {
                        "type": "string",
                        "default": "default",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/deployment/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deployment"
                ],
                "summary": "更新 Deployment 的镜像版本和副本集",
                "parameters": [
                    {
                        "description": "Deployment configuration information that needs to be changed",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.DeployUpdateMessage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/namespace/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Namespace"
                ],
                "summary": "创建 Namespace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Namespace name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/namespace/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Namespace"
                ],
                "summary": "删除单个 Namespace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace 名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/namespace/get": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Namespace"
                ],
                "summary": "获取单个 Namespace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "namespace 名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/namespace/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Namespace"
                ],
                "summary": "获取所有的 Namespace",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/pod/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod"
                ],
                "summary": "创建 Pod",
                "parameters": [
                    {
                        "description": "Pod simple configuration",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PodRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/pod/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod"
                ],
                "summary": "删除单个 Pod",
                "parameters": [
                    {
                        "type": "string",
                        "default": "default",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "pod名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/pod/get": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod"
                ],
                "summary": "获取单个 Pod 信息",
                "parameters": [
                    {
                        "type": "string",
                        "default": "default",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "pod名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/pod/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod"
                ],
                "summary": "获取命名空间中所有的  Pod 信息",
                "parameters": [
                    {
                        "type": "string",
                        "default": "default",
                        "description": "命名空间",
                        "name": "namespace",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/pod/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pod"
                ],
                "summary": "更新 Pod",
                "parameters": [
                    {
                        "description": "Pod configuration information that needs to be changed",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PodUpdateMessage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "注册用户",
                "parameters": [
                    {
                        "description": "用户名，密码，昵称，手机和邮箱",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/system.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "description": "用户名，密码，昵称，手机和邮箱",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.GetName"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/user/get": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/user/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "列出所有用户",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "用户更新密码",
                "parameters": [
                    {
                        "description": "用户名，原密码，新密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.NewPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CommonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "conf.Mysql": {
            "type": "object",
            "properties": {
                "config": {
                    "description": "数据库连接参数",
                    "type": "string"
                },
                "db-name": {
                    "description": "数据库名称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "path": {
                    "description": "数据库路径",
                    "type": "string"
                },
                "port": {
                    "description": "数据库端口",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "request.DeployUpdateMessage": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "newImage": {
                    "type": "string"
                },
                "replicasNumber": {
                    "type": "integer"
                }
            }
        },
        "request.DeploymentRequest": {
            "type": "object",
            "properties": {
                "namespace": {
                    "type": "string"
                },
                "object": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "request.GetName": {
            "type": "object",
            "properties": {
                "forced": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.Login": {
            "type": "object",
            "properties": {
                "captchaId": {
                    "type": "string"
                },
                "captchaValue": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.NewPassword": {
            "type": "object",
            "properties": {
                "newPassword": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.PodRequest": {
            "type": "object",
            "properties": {
                "containerName": {
                    "type": "string"
                },
                "containerPort": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "namespace": {
                    "type": "string"
                },
                "podName": {
                    "type": "string"
                }
            }
        },
        "request.PodUpdateMessage": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "namespace": {
                    "type": "string"
                },
                "newImage": {
                    "type": "string"
                }
            }
        },
        "response.CommonResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "system.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "description": "主键ID",
                    "type": "integer"
                },
                "nickName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "K8S Client API",
	Description:      "This is a k8s cluster management service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
