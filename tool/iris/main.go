package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	mvc.New(app).Handle(new(ExampleController))
	app.Run(iris.Addr(":8080"))
}

type ExampleController struct{}

func (c *ExampleController) Get() hero.Result {
	return hero.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}
func (c *ExampleController) GetPing() string {
	return "pong"
}
func (c *ExampleController) GetHello() any {
	return map[string]string{"message": "Hello Iris!"}
}
func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from the custom handler without following the naming guide"
}

func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	middleware := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.Next()
	}
	b.Handle(
		"GET",
		"/custom_path",
		"CustomHandlerWithoutFollowingTheNamingGuide",
		middleware,
	)
}
