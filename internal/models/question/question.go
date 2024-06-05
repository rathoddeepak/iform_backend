package question;

import (
	dbm "iform/pkg/managers/db"
	"iform/pkg/helpers/scanners"
	"errors"
)

type Question struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	PageId int64 `json:"page_id" binding:"required"`
	QuestionText string `json:"question_text"`
	QuestionType string `json:"question_type"`	
	RelativeOrder int `json:"relative_order"`
	Validations scanners.JSONBArray `json:"validations" gorm:"type:jsonb"`
	Config scanners.JSONB `json:"config" gorm:"type:jsonb"`
	IsActive bool `json:"is_active"`
}

func (Question) TableName() string {
	return "questions"
}

/**
 * Creates new question record
*/
func CreateQuestion (question *Question) error {
	db := dbm.GetConnection();

	result := db.Create(question)
	if result.Error != nil {
		return result.Error;
	}
	if result.RowsAffected == 0 {
		return errors.New("Unable to create question");
	}

	return nil;
}

/**
 * Updates question data
*/
func UpdateQuestion (question *Question) error {	
	result := dbm.GetConnection().
		Model(&Question{}).
		Where("id = ?", question.Id).
		Updates(map[string]interface {} {
			"question_text": question.QuestionText,
			"question_type": question.QuestionType,
			"validations": question.Validations,
			"config": question.Config,
		});
	if err := result.Error; err != nil {
		return err;
	}

	return nil;
}

/**
 * Updates relative order of a question
*/
func UpdateRelativeOrder (questionId int64, relativeOrder int) error {
	result := dbm.GetConnection().
	Model(&Question{}).
	Where("id = ?", questionId).
	Update("relative_order", relativeOrder);
	if err := result.Error; err != nil {
		return err;
	}
	return nil;
}

/**
 * Get question record
*/
func DeleteQuestion (question int64) error {	
	result := dbm.GetConnection().
		Model(&Question{}).
		Where("id = ?", question).
		Update("is_active", false);

	if err := result.Error; err != nil {
		return err;
	}

	return nil;
}

/**
 * Returns latest relative order + 1
*/
func GetNewRelativeOrder(pageId int64) int {
	var question Question;
	var relativeOrder int;

	result := dbm.GetConnection().
	Select("relative_order").
	Order("relative_order desc").
	Where("page_id = ?", pageId).
	First(&question);

	if err := result.Error; err == nil {
		relativeOrder = question.RelativeOrder;
	}

	return relativeOrder + 1;
}

/**
 * Returns questions of page
*/
func ListPageQuestions (pageId int64) (*[]Question, error) {	
	questions := []Question{};

	result := dbm.GetConnection().
	Model(&Question{}).
	Order("relative_order asc").
	Where("page_id = ?", pageId).
	Where("is_active", true).Find(&questions);

	if err := result.Error; err != nil {
		return nil, err;
	}

	return &questions, nil;
}