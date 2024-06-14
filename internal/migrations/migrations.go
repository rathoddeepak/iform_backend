/**
 * Migrations
 * 
 * Models may look redundant
 * they declared again to remove complexity,
 * and they will be merged with logical models later
 * if required
*/
package migrations;

import (
	"time"

	dbm "iform/pkg/managers/db"

	"github.com/google/uuid"
)

type User struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	PhoneNumber string    `gorm:"unique;not null"`
	OTP         string    `gorm:"default:null"`
	CreatedAt   time.Time `gorm:"default:now()"`
	Forms       []Form    `gorm:"foreignKey:UserID"`
	Sessions    []Session `gorm:"foreignKey:UserID"`
	Submissions []Submission `gorm:"foreignKey:UserID"`
}

type Form struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	UserID    int64     `gorm:"not null"`
	Title     string    `gorm:"not null"`
	LinkCode  string    `gorm:"default:null"`
	Timeout   int       `gorm:"default:null"`
	CreatedAt time.Time `gorm:"default:now()"`
	IsActive  bool      `gorm:"default:true"`
	Pages     []Page    `gorm:"foreignKey:FormID"`
	Submissions []Submission `gorm:"foreignKey:FormID"`
}

type Page struct {
	ID            int64     `gorm:"primaryKey;autoIncrement"`
	FormID        int64     `gorm:"not null"`
	Title         string    `gorm:"not null"`
	RelativeOrder int16     `gorm:"default:0"`
	IsActive      bool      `gorm:"default:true"`
	Questions     []Question `gorm:"foreignKey:PageID"`
}

type Question struct {
	ID            int64       `gorm:"primaryKey;autoIncrement"`
	PageID        int64       `gorm:"not null"`
	RelativeOrder int16       `gorm:"default:0"`
	QuestionText  string      `gorm:"not null"`
	QuestionType  string      `gorm:"not null"`
	Validations   interface{} `gorm:"type:jsonb;default:'{}'"`
	Config        interface{} `gorm:"type:jsonb;default:'{}'"`
	IsActive      bool        `gorm:"default:true"`
	SubmissionAnswers []SubmissionAnswer `gorm:"foreignKey:QuestionID"`
}

type Submission struct {
	ID         int64     `gorm:"primaryKey;autoIncrement"`
	FormID     int64     `gorm:"not null"`
	UserID     *int64    `gorm:"default:null"`  // Optional, so use a pointer
	CreatedAt  time.Time `gorm:"default:now()"`
	UpdatedAt  time.Time `gorm:"default:now()"`
	Submitted  bool      `gorm:"default:false"`
	SubmissionAnswers []SubmissionAnswer `gorm:"foreignKey:SubmissionID"`
}

type SubmissionAnswer struct {
	ID           int64     `gorm:"primaryKey;autoIncrement"`
	SubmissionID int64     `gorm:"not null"`
	QuestionID   int64     `gorm:"not null"`
	Answer       interface{} `gorm:"type:jsonb;not null"`
}

type Session struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    int64     `gorm:"not null"`
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:"default:now()"`
}

/**
 * Create Tables and Updates tables for database
*/ 
func StartAutoMigrate() error {
	db := dbm.GetConnection();
	// Migrate the schema
	err := db.AutoMigrate(&User{}, &Form{}, &Page{}, &Question{}, &Submission{}, &SubmissionAnswer{}, &Session{})
	return err;
}
