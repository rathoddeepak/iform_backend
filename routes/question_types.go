/**
 * Routes related to question types
*/
package routes;

import (
	"iform/internal/controller"

	"github.com/gin-gonic/gin"
)

func initQuestionTypeRoutes(r *gin.Engine) {
	questionTypeRouter := r.Group("/question_types");
	questionTypeRouter.POST("/get_default_types", controller.GetDefaultTypes);
}