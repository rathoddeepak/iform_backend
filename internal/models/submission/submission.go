package submission;

import (
	dbm "iform/pkg/managers/db"
	"iform/pkg/helpers/scanners"
	"errors"
	"time"
	"fmt"
)

type Submission struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	FormId int64 `json:"form_id"`
	UserId int64 `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SubmissionExpiry struct {
	Id int64 `json:"id" gorm:"primaryKey"`
	Timeleft int `json:"timeleft"`
	Submitted bool `json:"submitted"`
}

type SubmissionForm struct {
	Id int64 `json:"id"`
	Title string `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type DetailedSubmissionRow struct {
	SubmissionAnswerId int64 `json:"submission_answer_id"`
	SubmissionAnswer scanners.JSONB `json:"submission_answer" gorm:"type:jsonb"`	
	
	QuestionId int64 `json:"question_id"`
	QuestionText string `json:"question_text"`
	QuestionType string `json:"question_type"`
	QuestionConfig scanners.JSONB `json:"question_config,omitempty" gorm:"type:jsonb"`
	QuestionValidations scanners.JSONBArray `json:"question_validations,omitempty" gorm:"type:jsonb"`

	PageId int64 `json:"page_id"`
	PageTitle string `json:"page_title"`
}

func (Submission) TableName() string {
	return "submissions";
}

/**
 * Creates sumbission record
*/
func CreateSubmission (submisison *Submission) error {
	db := dbm.GetConnection();

	result := db.Create(submisison)
	if result.Error != nil {
		return result.Error;
	}
	if result.RowsAffected == 0 {
		return errors.New("Unable to create submission");
	}

	return nil;
}

/**
 * Get SubmissionId by User and Form Id
*/
func GetSubmissionIdByUserAndForm(userId int64, formId int64) (int64, error) {
	var submission Submission;

	result := dbm.GetConnection().
	Select("id").
	Where("user_id = ?", userId).
	Where("form_id = ?", formId).
	First(&submission);

	if err := result.Error; err != nil {
		return 0, err;
	}

	return submission.Id, nil;
}

/**
 * Returns timeleft fro submission 
*/ 
func GetSubmissionExpiry(submissionId int64) (*SubmissionExpiry, error) {
	var submission SubmissionExpiry;
	result := dbm.GetConnection().
	Raw(
		`select 
		sm.id,
		sm.submitted,
		date_part('minutes', (fm.timeout || ' minutes')::interval - AGE(NOW(), sm.updated_at))::int timeleft
		from submissions sm
		join forms fm on fm.id = sm.form_id
		where sm.id = ?`, 
		submissionId,
	).First(&submission);
	if err := result.Error; err != nil {
		return nil, err;
	}
	return &submission, nil;
}

/**
 * Get Submission by submission id
*/
func GetSubmissionById(submissionId int64) (*Submission, error) {
	var submission Submission;
	result := dbm.GetConnection().First(&submission, submissionId);
	if err := result.Error; err != nil {
		return nil, err;
	}
	return &submission, nil;
}

/**
 * Get Submissions with title
*/
func GetSubmissions(formId int64) ([]SubmissionForm, error) {
	
	var submissions []SubmissionForm;
	result := dbm.GetConnection().Raw(`
		select sm.id, fm.title, sm.created_at from submissions sm
		join forms fm on fm.id = sm.form_id
		WHERE sm.form_id = ?`, formId,
	).Find(&submissions);

	if err := result.Error; err != nil {
		return nil, err;
	}

	return submissions, nil;
}

/**
 * Updates submission update_at
*/
func UpdateAtNow (submissionId int64) error {	
	result := dbm.GetConnection().
		Model(&Submission{}).
		Where("id = ?", submissionId).
		Updates(map[string]interface{}{
	        "updated_at":  time.Now(),
	    });

	if err := result.Error; err != nil {
		return err;
	}

	return nil;
}

/**
 * Updates submission update_at
*/
func MarkAsSubmitted (submissionId int64) error {	
	result := dbm.GetConnection().
		Model(&Submission{}).
		Where("id = ?", submissionId).
		Update("submitted", true);
	if err := result.Error; err != nil {
		return err;
	}
	return nil;
}


/**
 * Get Form Submission Detailed Data
*/
func GetFormSubmissionData(submissionId int64, pageId int64, includeConfig bool) ([]DetailedSubmissionRow, error) {	
	var fields []DetailedSubmissionRow;
	
	extraCols := "";
	if includeConfig {
		extraCols = `qs.validations question_validations,
		qs.config question_config,`;
	}

	query := fmt.Sprintf(`
		select
		  pg.id page_id,
		  pg.title page_title,
		  qs.id question_id,
		  qs.question_text,
		  qs.question_type,
		  %s
		  sa.id submission_answer_id,
		  sa.answer submission_answer
		from submissions sb
		join forms fm on fm.id = sb.form_id
		join pages pg on pg.form_id = fm.id
		join questions qs on qs.page_id = pg.id and qs.is_active
		left join submission_answers sa on sa.submission_id = sb.id and sa.question_id = qs.id
		WHERE sb.id = ? and pg.id = ?
		order by qs.relative_order`, extraCols);

	result := dbm.GetConnection().Raw(query, submissionId, pageId).Find(&fields);

	if err := result.Error; err != nil {
		return nil, err;
	}

	return fields, nil;
}

