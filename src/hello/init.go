package hello

import "github.com/GoogleCloudPlatform/functions-framework-go/functions"

func init() {
	functions.HTTP("HelloWorld", HelloWorld)
	functions.HTTP("HelloDynamic", HelloDynamic)
}
