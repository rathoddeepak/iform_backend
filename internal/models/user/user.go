package users;

import (
	dbm "iform/pkg/managers/db"
	"gorm.io/gorm"
	"errors"
	"time"	
)

type User struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	PhoneNumber string `json:"phone_number"`
	OTP string `json:"otp"`
	CreatedAt time.Time `json:"created_at"`
}

func IsUserByPhone(phone string) *User {
	var user User;
	db := dbm.GetConnection();	
	if err := db.Where("phone_number = ?", phone).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil;
		}
		return nil;
	}
	return &user;
}

func CreateUser (user *User) error {
	db := dbm.GetConnection();

	insertedUser := IsUserByPhone(user.PhoneNumber);

	if insertedUser == nil {
		result := db.Create(user)
		if result.Error != nil {
			return result.Error;
		}
		if result.RowsAffected == 0 {
			return errors.New("Unable to create user");
		}
	} else {
		user.Id = insertedUser.Id;
		user.CreatedAt = insertedUser.CreatedAt;
	}

	/**
	 * This function is not implemented
	 * as we are not actually sending OTP
	 */
	// err := smsHelper.SendSMSOTP();

	return nil;
}

func VerifyOTP (phoneNumber string, otp string) (bool, int64) {
	var user User;
	db := dbm.GetConnection();	
	result := db.Where("phone_number = ?", phoneNumber).Where("otp = ?", otp).First(&user);
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, user.Id;
		}
		return false, user.Id;
	}
	return true, user.Id;
}