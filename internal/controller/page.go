package controller;

import (
	pageModel "iform/internal/models/page"

	"net/http"
	"sync"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PageOrderRequest struct {
	PageIds []int64 `json:"page_ids" binding:"required"`
}

/**
 * Creates page record
 * Updates page if Id is passed
*/
func CreatePage (c *gin.Context) {
	request := pageModel.Page{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }
	if request.Id == 0 {
		request.IsActive = true;
		request.RelativeOrder = pageModel.GetNewRelativeOrder(request.FormId);
		err := pageModel.CreatePage(&request);
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				*ErrorResponse("Unable to create page!"),
			);
			return;
		}
	} else {
		err := pageModel.UpdatePage(&request);
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				*ErrorResponse("Unable to update page!"),
			);
			return;
		}
	}
	response := SuccessResponse(request);
	c.JSON(http.StatusOK, *response);
}


func UpdatePagesOrder (c *gin.Context) {
	request := PageOrderRequest{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }

    wg := sync.WaitGroup{};   
    for idx, pageId := range request.PageIds {
    	wg.Add(1);
    	go func (id int64, order int) {
    		pageModel.UpdateRelativeOrder(id, order);
    		wg.Done();	
    	}(pageId, idx+1);
    }
    wg.Wait();
	response := SuccessResponse(request);
	c.JSON(http.StatusOK, *response);
}

/**
 * Deactivates page
*/
func DeletePage (c *gin.Context) {
	request := pageModel.Page{};
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
			*ErrorResponse("Page not found!"),
		);
		return
	} else {
		err := pageModel.DeletePage(request.Id);
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				*ErrorResponse("Unable to delete page!"),
			);
			return;
		}
	}
	response := SuccessResponse(request);
	c.JSON(http.StatusOK, *response);
}

/**
 * Get all active pages related to form
*/
func GetFormPages (c *gin.Context) {
	request := pageModel.Page{};
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(err);
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }
	forms, err := pageModel.ListFormPages(request.FormId);
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Unable to load form pages!"),
		);
		return;
	}
	response := SuccessResponse(forms);
	c.JSON(http.StatusOK, *response);
}