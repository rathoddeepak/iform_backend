/**
 * Routes related to page
*/
package routes;

import (
	"iform/internal/controller"
	"iform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initQuestionRoutes(r *gin.Engine) {
	questionRouter := r.Group("/question");
	questionRouter.Use(middleware.UserFromSessionID());
	questionRouter.POST("/create", controller.CreateQuestion);
	questionRouter.POST("/delete", controller.DeleteQuestion);
	questionRouter.POST("/list", controller.GetPageQuestions);
	questionRouter.POST("/order", controller.UpdateQuestionsOrder);	
}
