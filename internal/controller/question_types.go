package controller;

import (
	questionTypeModel "iform/internal/models/question_types"

	"net/http"

	"github.com/gin-gonic/gin"	
)

/**
 * Returns list of all question types
*/
func GetDefaultTypes(c *gin.Context) {	
	result := []questionTypeModel.ElementBase{};
	
	// Text Element
	result = append(result, *questionTypeModel.CreateDefaultText());

	// Email Element
	result = append(result, *questionTypeModel.CreateDefaultEmail());

	// Phone Element
	result = append(result, *questionTypeModel.CreateDefaultPhone());

	// Multiple Choice Element
	result = append(result, *questionTypeModel.CreateDefaultMultipleChoice());

	// Checkbox Element
	result = append(result, *questionTypeModel.CreateDefaultCheckbox());

	response := SuccessResponse(result);

	c.JSON(http.StatusOK, *response);
}