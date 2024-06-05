package controller;

import (
	userModel "iform/internal/models/user"
	sessionModel "iform/internal/models/session"
	"iform/pkg/helpers/input"

	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type ValidateRequest struct {
	OTP string `json:"otp" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type ValidateResponse struct {
	SesssionId string `json:"session_id"`
}

const (
	defaultOTP = "123456"
)

/**
 * Creates user if not created and sends OTP via SMS
*/
func LoginUser (c *gin.Context) {
	request := LoginRequest{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }
	
	if isValidPhone := input.IsValidPhone(request.PhoneNumber); !isValidPhone {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Invalid phone number!"),
		);
		return;
	}

	user := userModel.User{
		OTP: defaultOTP,
		PhoneNumber: request.PhoneNumber,
	};

	err := userModel.CreateUser(&user);
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Unable to create user!"),
		);
		return;
	}

	response := SuccessResponse("OTP sent successfully");

	c.JSON(http.StatusOK, *response);
}

/**
 * Verify otp for pertiuclar user
*/
func ValidateOTP (c *gin.Context) {
	request := ValidateRequest{};
	if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Incomplete request!"),
		);
        return
    }

    // Verify OTP
	isVerified, userId := userModel.VerifyOTP(request.PhoneNumber, request.OTP);
	if !isVerified {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Wrong OTP entered!"),
		);
		return;
	}

	// Create New Session If Verified
	session := sessionModel.Session {
		UserId: userId,
	}
	err := sessionModel.CreateSession(&session);
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			*ErrorResponse("Unable to create session!"),
		);
		return;
	}



	response := SuccessResponse(ValidateResponse {
		SesssionId: session.Id,
	});

	c.JSON(http.StatusOK, *response);
}