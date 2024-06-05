/**
 * Routes related to question types
*/
package routes;

import (
	"iform/internal/controller"
	"iform/internal/middleware"

	"github.com/gin-gonic/gin"
)

func initSubmissionRoutes(r *gin.Engine) {
	submissionRouter := r.Group("/submission");
	submissionRouter.Use(middleware.IfUserFromSessionID());
	submissionRouter.POST("/init", controller.InitSubmission);
	submissionRouter.POST("/submit_answers", controller.SubmitAnswers);
	submissionRouter.POST("/list", controller.GetFormSubmissions);
	submissionRouter.POST("/submission_data", controller.GetSubmissionData);
}