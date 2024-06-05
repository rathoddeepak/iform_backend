package controller;

import (
	questionModel "iform/internal/models/question"
	questionTypesModel "iform/internal/models/question_types"

	"net/http"
	"sync"
	"fmt"

	"github.com/gin-gonic/gin"
)

type QuestionOrderRequest struct {
	QuestionIds []int64 `json:"question_ids" binding:"required"`
}

/**
 * Creates question record
 * Updates question if Id is passed
*/
func CreateQuestion (c *gin.Context) {
	request := questionModel.Question{};
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(err);
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }
	if request.Id == 0 {
		if len(request.QuestionType) == 0 {
			textElement := questionTypesModel.CreateDefaultText();
			request.QuestionType = textElement.Id;
		}
		request.IsActive = true;
		request.RelativeOrder = questionModel.GetNewRelativeOrder(request.PageId);
		err := questionModel.CreateQuestion(&request);
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				*ErrorResponse("Unable to create question!"),
			);
			return;
		}
	} else {
		err := questionModel.UpdateQuestion(&request);
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				*ErrorResponse("Unable to update question!"),
			);
			return;
		}
	}
	response := SuccessResponse(request);
	c.JSON(http.StatusOK, *response);
}

/**
 * Updates relative order of questions
*/
func UpdateQuestionsOrder (c *gin.Context) {
	request := QuestionOrderRequest{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }

    wg := sync.WaitGroup{};   
    for idx, pageId := range request.QuestionIds {
    	wg.Add(1);
    	go func (id int64, order int) {
    		questionModel.UpdateRelativeOrder(id, order);
    		wg.Done();	
    	}(pageId, idx+1);
    }
    wg.Wait();
	response := SuccessResponse(request);
	c.JSON(http.StatusOK, *response);
}

/**
 * Deactivates question
*/
func DeleteQuestion (c *gin.Context) {
	request := questionModel.Question{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }
	if request.Id == 0 {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Question not found!"),
		);
		return
	} else {
		err := questionModel.DeleteQuestion(request.Id);
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				*ErrorResponse("Unable to delete question!"),
			);
			return;
		}
	}
	response := SuccessResponse(request);
	c.JSON(http.StatusOK, *response);
}

/**
 * Get all active questions related to page
*/
func GetPageQuestions (c *gin.Context) {
	request := questionModel.Question{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }
	forms, err := questionModel.ListPageQuestions(request.PageId);
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Unable to load page questions!"),
		);
		return;
	}
	response := SuccessResponse(forms);
	c.JSON(http.StatusOK, *response);
}