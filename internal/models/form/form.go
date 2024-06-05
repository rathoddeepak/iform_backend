package users;

import (
	dbm "iform/pkg/managers/db"
	"errors"
	"time"	
)

type Form struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	UserId int64 `json:"user_id"`
	Title string `json:"title"`
	LinkCode string `json:"link_code"`
	Timeout int `json:"timeout"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func (Form) TableName() string {
	return "forms"
}

/**
 * Creates new form record
*/
func CreateForm (form *Form) error {
	db := dbm.GetConnection();

	result := db.Create(form)
	if result.Error != nil {
		return result.Error;
	}
	if result.RowsAffected == 0 {
		return errors.New("Unable to create form");
	}

	return nil;
}

/**
 * Updates form title and timeout
*/
func UpdateForm (form *Form) error {	
	result := dbm.GetConnection().
		Model(&Form{}).
		Where("id = ?", form.Id).
		Updates(map[string]interface{}{
	        "title":  form.Title,
	        "timeout": form.Timeout,
	    });

	if err := result.Error; err != nil {
		return err;
	}

	return nil;
}

/**
 * Updates form link code
*/
func UpdateLinkCode (formId int64, code string) error {	
	result := dbm.GetConnection().
		Model(&Form{}).
		Where("id = ?", formId).
		Update("link_code", code);

	if err := result.Error; err != nil {
		return err;
	}

	return nil;
}

/**
 * Returns for by formId
*/
func GetFormById (formId int64) (*Form, error) {	
	form := &Form{};
	result := dbm.GetConnection().Model(&Form{}).First(form, formId);
	if err := result.Error; err != nil {
		return nil, err;
	}
	return form, nil;
}

/**
 * Delete form record
*/
func DeleteForm (formId int64) error {	
	result := dbm.GetConnection().
		Model(&Form{}).
		Where("id = ?", formId).
		Update("is_active", false);

	if err := result.Error; err != nil {
		return err;
	}

	return nil;
}

/**
 * Returns forms created by user
*/
func ListUserForms (userId int64) (*[]Form, error) {	
	forms := []Form{};

	result := dbm.GetConnection().
	Model(&Form{}).
	Order("created_at desc").
	Where("user_id = ?", userId).
	Where("is_active", true).Find(&forms);

	if err := result.Error; err != nil {
		return nil, err;
	}

	return &forms, nil;
}

