package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Define our own Context - db, cache, global var, etc..
type appContext struct {
	db string
}

//Wraper struct for func that need own Context
type appHandler struct {
	//our own context
	*appContext
	//inject our appContext
	H func(a *appContext, c *gin.Context)
}

//HandlerFunc implement gin.HandlerFunc interface
func (ah appHandler) HandlerFunc(c *gin.Context) {
	ah.H(ah.appContext, c)
}

//showDB work between our context and gin.Context
func showDB(a *appContext, c *gin.Context) {
	c.String(200, a.db)
}

func main() {
	r := gin.Default()
	//initial appContext, normally is something like `NewAppContext()`
	context := &appContext{db: "Initial mock db"}

	//Create our appHandler and pass HandlerFunc to gin router
	r.GET("/", appHandler{context, showDB}.HandlerFunc)
	r.Run(":8080")
}
