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
            "name": "GO feature flag relay proxy",
            "url": "https://gofeatureflag.org",
            "email": "contact@gofeatureflag.org"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/thomaspoignant/go-feature-flag/blob/main/LICENSE"
        },
        "version": "{{.Version}}",
        "x-logo": {
            "url": "https://raw.githubusercontent.com/thomaspoignant/go-feature-flag/main/logo_128.png"
        }
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "Making a **GET** request to the URL path ` + "`" + `/health` + "`" + ` will tell you if the relay proxy is ready to serve\ntraffic.\n\nThis is useful especially for loadbalancer to know that they can send traffic to the service.",
                "produces": [
                    "application/json"
                ],
                "summary": "Health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.HealthResponse"
                        }
                    }
                }
            }
        },
        "/info": {
            "get": {
                "description": "Making a **GET** request to the URL path ` + "`" + `/info` + "`" + ` will give you information about the actual state\nof the relay proxy.\n\nAs of Today the level of information is small be we can improve this endpoint to returns more\ninformation.",
                "produces": [
                    "application/json"
                ],
                "summary": "Info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.InfoResponse"
                        }
                    }
                }
            }
        },
        "/metrics": {
            "get": {
                "description": "This endpoint is providing metrics about the relay proxy in the prometheus format.",
                "produces": [
                    "text/plain"
                ],
                "summary": "Prometheus endpoint",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ofrep/v1/evaluate/flags": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Making a **POST** request to the URL ` + "`" + `/ofrep/v1/evaluate/flags` + "`" + ` will give you the value of the list\nof feature flags for this evaluation context.\n\nIf no flags are provided, the API will evaluate all available flags in the configuration.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Open-Feature Remote Evaluation Protocol bulk evaluation API.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The request will be processed only if ETag doesn't match.",
                        "name": "If-None-Match",
                        "in": "header"
                    },
                    {
                        "description": "Evaluation Context and list of flag for this API call",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.OFREPEvalFlagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OFREP successful evaluation response",
                        "schema": {
                            "$ref": "#/definitions/model.OFREPBulkEvaluateSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad evaluation request",
                        "schema": {
                            "$ref": "#/definitions/model.OFREPCommonErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized - You need credentials to access the API",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    },
                    "403": {
                        "description": "Forbidden - You are not authorized to access the API",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    }
                }
            }
        },
        "/ofrep/v1/evaluate/flags/{flag_key}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Making a **POST** request to the URL ` + "`" + `/ofrep/v1/evaluate/flags/{your_flag_name}` + "`" + ` will give you the\nvalue of the flag for this evaluation context\n",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Evaluate a feature flag using the OpenFeature Remote Evaluation Protocol",
                "parameters": [
                    {
                        "description": "Evaluation Context for this API call",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.OFREPEvalFlagRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Name of your feature flag",
                        "name": "flag_key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.OFREPEvaluateSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.OFREPEvaluateErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    },
                    "404": {
                        "description": "Flag Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.OFREPEvaluateErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    }
                }
            }
        },
        "/v1/allflags": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Making a **POST** request to the URL ` + "`" + `/v1/allflags` + "`" + ` will give you the values of all the flags for\nthis user.\n\nTo get a variation you should provide information about the user.\nFor that you should provide some user information in JSON in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "All flags variations for a user",
                "parameters": [
                    {
                        "description": "Payload of the user we want to challenge against the flag.",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AllFlagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.AllFlags"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    }
                }
            }
        },
        "/v1/data/collector": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "This endpoint is receiving the events of your flags usage to send them in the data collector.\n\nIt is used by the different Open Feature providers to send in bulk all the cached events to avoid\nto lose track of what happen when a cached flag is used.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Endpoint to send usage of your flags to be collected",
                "parameters": [
                    {
                        "description": "List of flag evaluation that be passed to the data exporter",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CollectEvalDataRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/model.CollectEvalDataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    }
                }
            }
        },
        "/v1/feature/{flag_key}/eval": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Making a **POST** request to the URL ` + "`" + `/v1/feature/\u003cyour_flag_name\u003e/eval` + "`" + ` will give you the value of the\nflag for this user.\n\nTo get a variation you should provide information about the user:\n- User information in JSON in the request body.\n- A default value in case there is an error while evaluating the flag.\n\nNote that you will always have a usable value in the response, you can use the field ` + "`" + `failed` + "`" + ` to know if\nan issue has occurred during the validation of the flag, in that case the value returned will be the\ndefault value.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Evaluate a feature flag",
                "parameters": [
                    {
                        "description": "Payload of the user we want to get all the flags from.",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.EvalFlagRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Name of your feature flag",
                        "name": "flag_key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.EvalFlagDoc"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    }
                }
            }
        },
        "/ws/v1/flag/change": {
            "post": {
                "description": "This endpoint is a websocket endpoint to be notified about flag changes, every change\nwill send a request to the client with a model.DiffCache format.\n",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Websocket endpoint to be notified about flag changes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "apiKey use authorize the connection to the relay proxy",
                        "name": "apiKey",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/notifier.DiffCache"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/modeldocs.HTTPErrorDoc"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "exporter.FeatureEvent": {
            "type": "object",
            "properties": {
                "contextKind": {
                    "description": "ContextKind is the kind of context which generated an event. This will only be \"anonymousUser\" for events generated\non behalf of an anonymous user or the reserved word \"user\" for events generated on behalf of a non-anonymous user",
                    "type": "string",
                    "example": "user"
                },
                "creationDate": {
                    "description": "CreationDate When the feature flag was requested at Unix epoch time in milliseconds.",
                    "type": "integer",
                    "example": 1680246000011
                },
                "default": {
                    "description": "Default value is set to true if feature flag evaluation failed, in which case the value returned was the default\nvalue passed to variation. If the default field is omitted, it is assumed to be false.",
                    "type": "boolean",
                    "example": false
                },
                "key": {
                    "description": "Key of the feature flag requested.",
                    "type": "string",
                    "example": "my-feature-flag"
                },
                "kind": {
                    "description": "Kind for a feature event is feature.\nA feature event will only be generated if the trackEvents attribute of the flag is set to true.",
                    "type": "string",
                    "example": "feature"
                },
                "source": {
                    "description": "Source indicates where the event was generated.\nThis is set to SERVER when the event was evaluated in the relay-proxy and PROVIDER_CACHE when it is evaluated from the cache.",
                    "type": "string",
                    "example": "SERVER"
                },
                "userKey": {
                    "description": "UserKey The key of the user object used in a feature flag evaluation. Details for the user object used in a feature\nflag evaluation as reported by the \"feature\" event are transmitted periodically with a separate index event.",
                    "type": "string",
                    "example": "94a25909-20d8-40cc-8500-fee99b569345"
                },
                "value": {
                    "description": "Value of the feature flag returned by feature flag evaluation."
                },
                "variation": {
                    "description": "Variation  of the flag requested. Flag variation values can be \"True\", \"False\", \"Default\" or \"SdkDefault\"\ndepending on which value was taken during flag evaluation. \"SdkDefault\" is used when an error is detected and the\ndefault value passed during the call to your variation is used.",
                    "type": "string",
                    "example": "admin-variation"
                },
                "version": {
                    "description": "Version contains the version of the flag. If the field is omitted for the flag in the configuration file\nthe default version will be 0.",
                    "type": "string",
                    "example": "v1.0.0"
                }
            }
        },
        "flag.ErrorCode": {
            "type": "string",
            "enum": [
                "PROVIDER_NOT_READY",
                "FLAG_NOT_FOUND",
                "PARSE_ERROR",
                "TYPE_MISMATCH",
                "GENERAL",
                "INVALID_CONTEXT",
                "TARGETING_KEY_MISSING",
                "FLAG_CONFIG"
            ],
            "x-enum-varnames": [
                "ErrorCodeProviderNotReady",
                "ErrorCodeFlagNotFound",
                "ErrorCodeParseError",
                "ErrorCodeTypeMismatch",
                "ErrorCodeGeneral",
                "ErrorCodeInvalidContext",
                "ErrorCodeTargetingKeyMissing",
                "ErrorFlagConfiguration"
            ]
        },
        "flag.Flag": {
            "type": "object",
            "properties": {
                "defValue": {
                    "description": "default value (as text); for usage message",
                    "type": "string"
                },
                "name": {
                    "description": "name as it appears on command line",
                    "type": "string"
                },
                "usage": {
                    "description": "help message",
                    "type": "string"
                },
                "value": {
                    "description": "value as set"
                }
            }
        },
        "model.AllFlagRequest": {
            "type": "object",
            "properties": {
                "evaluationContext": {
                    "description": "EvaluationContext The representation of a EvaluationContext for your feature flag system.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.EvaluationContextRequest"
                        }
                    ]
                },
                "user": {
                    "description": "Deprecated: User The representation of a user for your feature flag system.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.UserRequest"
                        }
                    ]
                }
            }
        },
        "model.CollectEvalDataRequest": {
            "type": "object",
            "properties": {
                "events": {
                    "description": "Events is the list of the event we send in the payload",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/exporter.FeatureEvent"
                    }
                },
                "meta": {
                    "description": "Meta are the extra information added during the configuration",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "model.CollectEvalDataResponse": {
            "type": "object",
            "properties": {
                "ingestedContentCount": {
                    "description": "IngestedContentCount number of model.FeatureEvents that have been sent to the data exporter",
                    "type": "integer"
                }
            }
        },
        "model.EvalFlagRequest": {
            "type": "object",
            "properties": {
                "defaultValue": {
                    "description": "The value will we use if we are not able to get the variation of the flag."
                },
                "evaluationContext": {
                    "description": "EvaluationContext The representation of a EvaluationContext for your feature flag system.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.EvaluationContextRequest"
                        }
                    ]
                },
                "user": {
                    "description": "Deprecated: User The representation of a user for your feature flag system.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.UserRequest"
                        }
                    ]
                }
            }
        },
        "model.EvaluationContextRequest": {
            "type": "object",
            "properties": {
                "custom": {
                    "description": "Custom is a map containing all extra information for this user.",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "company": "GO Feature Flag",
                        "email": "contact@gofeatureflag.org",
                        "firstname": "John",
                        "lastname": "Doe"
                    }
                },
                "key": {
                    "description": "Key is the identifier of the UserRequest.",
                    "type": "string",
                    "example": "08b5ffb7-7109-42f4-a6f2-b85560fbd20f"
                }
            }
        },
        "model.HealthResponse": {
            "type": "object",
            "properties": {
                "initialized": {
                    "description": "Set to true if the HTTP server is started",
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "model.InfoResponse": {
            "type": "object",
            "properties": {
                "cacheRefresh": {
                    "description": "This is the last time when your flag file was read and store in the internal cache.",
                    "type": "string",
                    "example": "2022-06-13T11:22:55.941628+02:00"
                }
            }
        },
        "model.OFREPBulkEvaluateSuccessResponse": {
            "type": "object",
            "properties": {
                "flags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OFREPFlagBulkEvaluateSuccessResponse"
                    }
                }
            }
        },
        "model.OFREPCommonErrorResponse": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "$ref": "#/definitions/flag.ErrorCode"
                },
                "errorDetails": {
                    "type": "string"
                }
            }
        },
        "model.OFREPEvalFlagRequest": {
            "type": "object",
            "properties": {
                "context": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "company": "GO Feature Flag",
                        "firstname": "John",
                        "lastname": "Doe",
                        "targetingKey": "4f433951-4c8c-42b3-9f18-8c9a5ed8e9eb"
                    }
                }
            }
        },
        "model.OFREPEvaluateErrorResponse": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "$ref": "#/definitions/flag.ErrorCode"
                },
                "errorDetails": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                }
            }
        },
        "model.OFREPEvaluateSuccessResponse": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "reason": {
                    "type": "string"
                },
                "value": {},
                "variant": {
                    "type": "string"
                }
            }
        },
        "model.OFREPFlagBulkEvaluateSuccessResponse": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "string"
                },
                "errorDetails": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                },
                "metadata": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "reason": {
                    "type": "string"
                },
                "value": {},
                "variant": {
                    "type": "string"
                }
            }
        },
        "model.UserRequest": {
            "type": "object",
            "properties": {
                "anonymous": {
                    "description": "Anonymous set if this is a logged-in user or not.",
                    "type": "boolean",
                    "example": false
                },
                "custom": {
                    "description": "Custom is a map containing all extra information for this user.",
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    },
                    "example": {
                        "company": "GO Feature Flag",
                        "email": "contact@gofeatureflag.org",
                        "firstname": "John",
                        "lastname": "Doe"
                    }
                },
                "key": {
                    "description": "Key is the identifier of the UserRequest.",
                    "type": "string",
                    "example": "08b5ffb7-7109-42f4-a6f2-b85560fbd20f"
                }
            }
        },
        "modeldocs.AllFlags": {
            "description": "AllFlags contains the full list of all the flags available for the user",
            "type": "object",
            "properties": {
                "flags": {
                    "description": "flags is the list of flag for the user.",
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/modeldocs.FlagState"
                    }
                },
                "valid": {
                    "description": "` + "`" + `true` + "`" + ` if something went wrong in the relay proxy (flag does not exists, ...) and we serve the defaultValue.",
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "modeldocs.EvalFlagDoc": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "description": "Code of the error returned by the server.",
                    "type": "string",
                    "example": ""
                },
                "failed": {
                    "description": "` + "`" + `true` + "`" + ` if something went wrong in the relay proxy (flag does not exists, ...) and we serve the defaultValue.",
                    "type": "boolean",
                    "example": false
                },
                "metadata": {
                    "description": "Metadata is a field containing information about your flag such as an issue tracker link, a description, etc ...",
                    "type": "object",
                    "additionalProperties": true
                },
                "reason": {
                    "description": "reason why we have returned this value.",
                    "type": "string",
                    "example": "TARGETING_MATCH"
                },
                "trackEvents": {
                    "description": "` + "`" + `true` + "`" + ` if the event was tracked by the relay proxy.",
                    "type": "boolean",
                    "example": true
                },
                "value": {
                    "description": "The flag value for this user."
                },
                "variationType": {
                    "description": "The variation used to give you this value.",
                    "type": "string",
                    "example": "variation-A"
                },
                "version": {
                    "description": "The version of the flag used.",
                    "type": "string",
                    "example": "1.0"
                }
            }
        },
        "modeldocs.FlagState": {
            "type": "object",
            "properties": {
                "timestamp": {
                    "description": "Timestamp is the time when the flag was evaluated.",
                    "type": "integer",
                    "example": 1652113076
                },
                "trackEvents": {
                    "description": "TrackEvents this flag is trackable.",
                    "type": "boolean",
                    "example": false
                },
                "value": {
                    "description": "Value is the flag value, it can be any JSON types."
                },
                "variationType": {
                    "description": "VariationType is the name of the variation used to have the flag value.",
                    "type": "string",
                    "example": "variation-A"
                }
            }
        },
        "modeldocs.HTTPErrorDoc": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "Message of your error",
                    "type": "string",
                    "example": "An error occurred"
                }
            }
        },
        "notifier.DiffCache": {
            "type": "object",
            "properties": {
                "added": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/flag.Flag"
                    }
                },
                "deleted": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/flag.Flag"
                    }
                },
                "updated": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/notifier.DiffUpdated"
                    }
                }
            }
        },
        "notifier.DiffUpdated": {
            "type": "object",
            "properties": {
                "new_value": {
                    "$ref": "#/definitions/flag.Flag"
                },
                "old_value": {
                    "$ref": "#/definitions/flag.Flag"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Use configured APIKeys in yaml config as authorization keys, disabled when this yaml config is not set.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "GO Feature Flag relay proxy endpoints",
	Description:      "# Introduction\n\nThis API is documented in **OpenAPI format** and describe the REST API of the **GO Feature Flag relay proxy**.\n\nThe relay proxy is a component to evaluate your feature flags remotely when using **GO Feature Flag**.  \nThis API is mostly used by all the OpenFeature providers.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
