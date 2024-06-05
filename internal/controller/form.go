package controller;

import (
	formModel "iform/internal/models/form"
	"iform/pkg/helpers/input"

	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * Creates form records
 * Updates form if Id is passed
*/
func CreateForm (c *gin.Context) {
	request := formModel.Form{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }
    request.UserId = c.MustGet("user_id").(int64);
	if request.Id == 0 {
		request.IsActive = true;
		err := formModel.CreateForm(&request);
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				*ErrorResponse("Unable to create form!"),
			);
			return;
		}
		code := input.RandomStringWithInteger(5, int(request.Id));
		// Updates LinkCode
		formModel.UpdateLinkCode(
			request.Id,
			code,
		);
		request.LinkCode = code;
	} else {
		err := formModel.UpdateForm(&request);
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				*ErrorResponse("Unable to update form!"),
			);
			return;
		}
	}
	response := SuccessResponse(request);
	c.JSON(http.StatusOK, *response);
}

/**
 * Deactivates form
*/
func DeleteForm (c *gin.Context) {
	request := formModel.Form{};
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
			*ErrorResponse("Form not found!"),
		);
		return
	} else {
		err := formModel.DeleteForm(request.Id);
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				*ErrorResponse("Unable to delete form!"),
			);
			return;
		}
	}
	response := SuccessResponse(request);
	c.JSON(http.StatusOK, *response);
}

/**
 * Get all active forms created by user
*/
func GetUserForms (c *gin.Context) {
	request := formModel.Form{};
	request.UserId = c.MustGet("user_id").(int64);
	forms, err := formModel.ListUserForms(request.UserId);
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Unable to load user forms!"),
		);
		return;
	}
	response := SuccessResponse(forms);
	c.JSON(http.StatusOK, *response);
}


/**
 * Get all active forms created by user
*/
func FormData (c *gin.Context) {
	request := formModel.Form{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }
	form, err := formModel.GetFormById(request.Id);
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Unable to load form!"),
		);
		return;
	}
	response := SuccessResponse(form);
	c.JSON(http.StatusOK, *response);
}
