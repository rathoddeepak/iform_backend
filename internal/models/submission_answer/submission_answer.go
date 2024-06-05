package submission_answer;

import (
	dbm "iform/pkg/managers/db"
	"iform/pkg/helpers/scanners"
	"gorm.io/gorm/clause"
	"errors"
)

type SubmissionAnswer struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	SubmissionId int64 `json:"submission_id"`
	QuestionId int64 `json:"question_id"`
	Answer scanners.JSONB `json:"answer" gorm:"type:jsonb"`
}

func (SubmissionAnswer) TableName() string {
	return "submission_answers";
}

/**
 * Delete Submission Answer
*/
func DeleteAnswerBySubmissionId (submissionId int64) error {
	result := dbm.GetConnection().
	Where("submission_id = ?", submissionId).
	Delete(&SubmissionAnswer{});
	if result.Error != nil {
		return result.Error;
	}
	return nil;
}

/**
 * Creates sumbission record
*/
func CreateSubmissionAnswer (submisisonAnswer *SubmissionAnswer) error {
	db := dbm.GetConnection();

	result := db.Clauses(clause.OnConflict{
	  Columns:   []clause.Column{
	  	{Name: "submission_id"},
	  	{Name: "question_id"},
	  },
	  DoUpdates: clause.Assignments(map[string]interface{}{
	  	"answer": submisisonAnswer.Answer,
	 }),
	}).Create(submisisonAnswer)

	if result.Error != nil {
		return result.Error;
	}
	if result.RowsAffected == 0 {
		return errors.New("Unable to create submission");
	}

	return nil;
}