/**
 * Routes related to form
*/
package routes;

import (
	"iform/internal/controller"
	"iform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initFormRoutes(r *gin.Engine) {
	formRouter := r.Group("/form");	
	formRouter.Use(middleware.UserFromSessionID());
	formRouter.POST("/create", controller.CreateForm);
	formRouter.POST("/delete", controller.DeleteForm);
	formRouter.POST("/list", controller.GetUserForms);	
	formRouter.POST("/id", controller.FormData);
}