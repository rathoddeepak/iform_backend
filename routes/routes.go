/**
 * Routes related to question types
*/
package routes;

import (
	"github.com/gin-gonic/gin"
    "iform/internal/middleware"
)

func InitRoutes() *gin.Engine {
	r := gin.Default();	
	r.Use(middleware.CORSMiddleware());

    // Init Routes
	initQuestionTypeRoutes(r);
    initSubmissionRoutes(r);
    initQuestionRoutes(r);
    initUserRoutes(r);
    initFormRoutes(r);
    initPageRoutes(r);    
	return r;
}