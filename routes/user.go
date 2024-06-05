/**
 * Routes related to user
*/
package routes;

import (
	"iform/internal/controller"

	"github.com/gin-gonic/gin"
)

func initUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/user");
	userRoutes.POST("/login_user", controller.LoginUser);
	userRoutes.POST("/validate_otp", controller.ValidateOTP);
}