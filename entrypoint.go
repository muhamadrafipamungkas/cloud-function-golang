package cloudfunctiongolang

import (
	_ "github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	Hello "github.com/muhamadrafipamungkas/cloud-function-golang/src/hello"
	User "github.com/muhamadrafipamungkas/cloud-function-golang/src/user"
)

func init() {
	functions.HTTP("HelloWorld", Hello.HelloWorld)
	functions.HTTP("HelloDynamic", Hello.HelloDynamic)
	functions.HTTP("HelloUser", User.HelloUser)
}
