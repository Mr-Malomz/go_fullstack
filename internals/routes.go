package internals

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	//views
	app.Router.GET("/", app.indexPageHandler())

	//apis
	app.Router.POST("/", app.createTodoHandler())
	app.Router.DELETE("/:id", app.deleteTodoHandler())
}
