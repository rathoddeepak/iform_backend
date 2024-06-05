/**
 * Routes related to page
*/
package routes;

import (
	"iform/internal/controller"
	"iform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initPageRoutes(r *gin.Engine) {
	pageRouter := r.Group("/page");
	pageRouter.Use(middleware.UserFromSessionID());
	pageRouter.POST("/create", controller.CreatePage);
	pageRouter.POST("/delete", controller.DeletePage);
	pageRouter.POST("/list", controller.GetFormPages);
	pageRouter.POST("/order", controller.UpdatePagesOrder);	
}


