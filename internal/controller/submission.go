package controller;

import (
	submissionAnswerModel "iform/internal/models/submission_answer"
	submissionModel "iform/internal/models/submission"
	formModel "iform/internal/models/form"
	pageModel "iform/internal/models/page"
	"net/http"
	"github.com/gin-gonic/gin"
	"sync"
);

type InitSubmissionRequest struct {
	SubmissionId int64 `json:"submission_id"`
	UserId int64 `json:"user_id"`
	FormId int64 `json:"form_id" binding:"required"`
}

type InitSubmissionResponse struct {
	Timeleft int `json:"timeleft"`
	HasTimeout bool `json:"timeout"`
	SubmissionId int64 `json:"submission_id"`
	Pages *[]pageModel.Page `json:"pages"`
	FormData *formModel.Form `json:"form"`
	Submitted bool `json:"submitted"`
}

type SubmitAnswerRequest struct	 {
	SubmissionId int64 `json:"submission_id" binding:"required"`
	PageId int64 `json:"page_id" binding:"required"`
	FormId int64 `json:"form_id" binding:"required"`
	Answers []submissionAnswerModel.SubmissionAnswer `json:"answers" binding:"required"`
}

type FormSubmissionDataRequest struct	 {
	PageId int64 `json:"page_id" binding:"required"`
	SubmissionId int64 `json:"submission_id" binding:"required"`
	IncludeConfig bool `json:"include_config"`
}

type FormSubmissionDataResponse struct {
	Fields []submissionModel.DetailedSubmissionRow `json:"fields"`
	FormData *formModel.Form `json:"form"`
}

/** --- If Submission Id is passed ---
* Create function to check either time has expired
* If expired then return expired true else return expired false
*** --- If Submission Id is not passed ---
* Check if user_id is passed
* If it is passed then, get current submission Id and return
* If submission is not created then created one and return

**/
func InitSubmission (c *gin.Context) {
	request := InitSubmissionRequest{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    };    
    submission := submissionModel.Submission{
		FormId: request.FormId,
	};	

	formData, err := formModel.GetFormById(request.FormId);
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Form not found!"),
		);
        return
	}

	pages, err := pageModel.ListFormPages(request.FormId)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Form pages not found!"),
		);
        return
	}

	hasTimeout := formData.Timeout > 0;
	timeleft := formData.Timeout;
	request.UserId = c.MustGet("user_id").(int64);

	response := InitSubmissionResponse{
		HasTimeout: hasTimeout,
		Pages: pages,
		FormData: formData,
	};

    if request.SubmissionId != 0 {
    	response.SubmissionId = request.SubmissionId;
    	submissionExpiryData, err := submissionModel.GetSubmissionExpiry(
    		request.SubmissionId,
    	);

    	if err != nil {
			c.JSON(
				http.StatusBadRequest,
				*ErrorResponse("Unable to get submission!"),
			);
			return;
		}
    	if hasTimeout && !submissionExpiryData.Submitted {
    		if submissionExpiryData.Timeleft <= 0 {
	    		submissionAnswerModel.DeleteAnswerBySubmissionId(request.SubmissionId);
	    		submissionModel.UpdateAtNow(request.SubmissionId);
			} else {
				timeleft = submissionExpiryData.Timeleft;
			}
    	}
    	response.Submitted = submissionExpiryData.Submitted;
    } else {
    	if request.UserId != 0 {
	    	submission.UserId = request.UserId;

	    	submissionId, _ := submissionModel.GetSubmissionIdByUserAndForm(
	    		request.UserId,
	    		request.FormId,
	    	);
	    	submission.Id = submissionId;
    	}
    	if submission.Id == 0 {
    		err := submissionModel.CreateSubmission(&submission);
    		if err != nil {
				c.JSON(
					http.StatusBadRequest,
					*ErrorResponse("Unable to get submission!"),
				);
				return;
			}			
    	}
    	response.SubmissionId = submission.Id; 	
    }
    response.Timeleft = timeleft;    
    c.JSON(
		http.StatusOK,
		*SuccessResponse(response),
	);
}

/**
 * Records answers related to submission_id and question_id
*/
func SubmitAnswers (c *gin.Context) {
	request := SubmitAnswerRequest{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }

    isLastPage := pageModel.IsLastFormPage(request.FormId, request.PageId);

    wg := sync.WaitGroup{};   
    // Use of transactions would have improved it a lot
    for _, answer := range request.Answers {
    	wg.Add(1);
    	go func (id int64, answer submissionAnswerModel.SubmissionAnswer) {
    		answer.SubmissionId = id;
    		submissionAnswerModel.CreateSubmissionAnswer(&answer);
    		wg.Done();	
    	}(request.SubmissionId, answer);
    }
    wg.Wait();

    if isLastPage {
    	submissionModel.MarkAsSubmitted(request.SubmissionId);
    }

	response := SuccessResponse("Submitted Successfully!");
	c.JSON(http.StatusOK, *response);
}

/**
 * Get Response List
 * Returns response list, with form title
*/
func GetFormSubmissions (c *gin.Context) {
	request := InitSubmissionRequest{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    };

	submissions, err := submissionModel.GetSubmissions(request.FormId);
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Unable to load submissions!"),
		);
		return;
	}

	response := SuccessResponse(submissions);
	c.JSON(http.StatusOK, *response);
}


/**
 * Get Response
 * Join with form question and submission answer
 * Sort by relative order
 * Group as pages
 * Add option for adding validation and config for submitter
*/
func GetSubmissionData (c *gin.Context) {
	request := FormSubmissionDataRequest{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    };

    fields, err := submissionModel.GetFormSubmissionData(
    	request.SubmissionId,
    	request.PageId,
    	request.IncludeConfig,
    );
	
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Unable to load fields!"),
		);
		return;
	}

	response := SuccessResponse(fields);
	c.JSON(http.StatusOK, *response);
}
