package routes

import (
	"lock/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(r *gin.RouterGroup, db *gorm.DB) {

	r.POST("/signup", handlers.SignUp)
	r.POST("/login",handlers.Login)
	r.POST("/select-app",handlers.SelectApp)

}
