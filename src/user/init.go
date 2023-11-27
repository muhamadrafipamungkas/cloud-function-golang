package user

import "github.com/GoogleCloudPlatform/functions-framework-go/functions"

func init() {
	functions.HTTP("HelloUser", HelloUser)
}
