package session;

import (
	dbm "iform/pkg/managers/db"
	"errors"
	"time"	
)

type Session struct {
	Id string `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserId int64 `json:"user_id"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func (Session) TableName() string {
  return "sessions"
}

func CreateSession (session *Session) error {
	db := dbm.GetConnection();

	// IsActive false for all previous sessions
	// db
	// .Model(&Session{})
	// .Where("user_id = ?", session.UserId)
	// .Update("is_active", false);

	result := db.Create(session);
	if result.Error != nil {
		return result.Error;
	}

	if result.RowsAffected == 0 {
		return errors.New("Unable to create user");
	}

	return nil;
}

func GetSessionData (sessionId string) (*Session, error) {
	db := dbm.GetConnection();
	session := &Session{};
	result := db.Where("id = ?", sessionId).First(session);
	if result.Error != nil {
		return nil, result.Error;
	}
	return session, nil;
}