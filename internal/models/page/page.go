package users;

import (
	dbm "iform/pkg/managers/db"
	"errors"
)

type Page struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	FormId int64 `json:"form_id" binding:"required"`
	Title string `json:"title"`
	RelativeOrder int `json:"relative_order"`
	IsActive bool `json:"is_active"`
}

func (Page) TableName() string {
	return "pages"
}

/**
 * Creates new page record
*/
func CreatePage (page *Page) error {
	db := dbm.GetConnection();

	result := db.Create(page)
	if result.Error != nil {
		return result.Error;
	}
	if result.RowsAffected == 0 {
		return errors.New("Unable to create page");
	}

	return nil;
}

/**
 * Updates page title
*/
func UpdatePage (page *Page) error {	
	result := dbm.GetConnection().
		Model(&Page{}).
		Where("id = ?", page.Id).
		Update("title", page.Title);

	if err := result.Error; err != nil {
		return err;
	}

	return nil;
}

/**
 * Updates relative order of a page
*/
func UpdateRelativeOrder (pageId int64, relativeOrder int) error {
	result := dbm.GetConnection().
	Model(&Page{}).
	Where("id = ?", pageId).
	Update("relative_order", relativeOrder);
	if err := result.Error; err != nil {
		return err;
	}
	return nil;
}

/**
 * Get page record
*/
func DeletePage (pageId int64) error {	
	result := dbm.GetConnection().
		Model(&Page{}).
		Where("id = ?", pageId).
		Update("is_active", false);

	if err := result.Error; err != nil {
		return err;
	}

	return nil;
}

/**
 * Returns latest relative order + 1
*/
func GetNewRelativeOrder(formId int64) int {
	var page Page;
	var relativeOrder int;

	result := dbm.GetConnection().
	Select("relative_order").
	Order("relative_order desc").
	Where("form_id = ?", formId).
	First(&page);

	if err := result.Error; err == nil {
		relativeOrder = page.RelativeOrder;
	}

	return relativeOrder + 1;
}

/**
 * Returns pages related to form
*/
func ListFormPages (formId int64) (*[]Page, error) {	
	pages := []Page{};

	result := dbm.GetConnection().
	Model(&Page{}).
	Order("relative_order asc").
	Where("form_id = ?", formId).
	Where("is_active", true).Find(&pages);

	if err := result.Error; err != nil {
		return nil, err;
	}

	return &pages, nil;
}

/**
 * Returns if page_id is last or not
*/
func IsLastFormPage (formId int64, pageId int64) bool {	
	lastPage := Page{};

	result := dbm.GetConnection().
	Model(&Page{}).
	Select("id").
	Order("relative_order desc").
	Where("form_id = ?", formId).
	Where("is_active", true).First(&lastPage);

	if err := result.Error; err != nil {
		return false;
	}

	return lastPage.Id == pageId;
}