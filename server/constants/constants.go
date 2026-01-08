package constants

import (
	common "server/models"
)

var Env common.DotEnv

const (
	JWTContextKey = "user"
)

type LocalKeyType struct{}

var OrganizationEmailKey LocalKeyType
var OrganizationIDKey LocalKeyType
