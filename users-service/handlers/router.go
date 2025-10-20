package handlers

import "github.com/gin-gonic/gin"

type Dependencies struct {
	UsrHandler *UserHandler
}

func RegistrasiRouter(r *gin.Engine, deps *Dependencies) {
	usr := r.Group("/users")
	{
		usr.GET("/:id", deps.UsrHandler.GetUserById)
		usr.POST("/", deps.UsrHandler.CreateUser)
	}
}
