package session

import (
	"nextgoBlog/model"
	"nextgoBlog/utils/random"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	sessionCookieName = "r_session"
	sessionMaxAge = 60 * 60 * 24 * 7
)

type session struct {
	token string
	userID string
	createdAt time.Time

	db *gorm.DB
	data map[string]interface{}
	sync.Mutex
}


type Session interface {
	Token() string
	UserID() string
	CreatedAt() time.Time

	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
	Delete(key string) error

	Expired() bool
}

type sessionStore struct {
	db *gorm.DB
	// cache [some cache]
}

type Store interface {
	GetSession(c echo.Context) (Session, error)
	GetSessionByToken(token string) (Session, error)
	GetSessionsByUserID(userID string) ([]Session, error)
	IssueSession(userID string) (Session, error)
}

func newSession(db *gorm.DB, token string, userID string, createdAt time.Time) *session {
	return &session{
		token: token,
		userID: userID,
		createdAt: createdAt,
		db: db,
		data: make(map[string]interface{}),
	}
}

func (s *session) Token() string {
	return s.token;
}

func (s *session) UserID() string {
	return s.userID;
}

func (s *session) CreatedAt() time.Time {
	return s.createdAt;
}

func (s *session) Get(key string) (interface{}, error) {
	s.Lock()
	defer s.Unlock()

	return s.data[key], nil
}

func (s *session) Set(key string, value interface{}) error {
	s.Lock()
	defer s.Unlock()

	s.data[key] = value
	return nil
}

func (s *session) Delete(key string) error {
	s.Lock()
	defer s.Unlock()

	delete(s.data, key)
	return nil
}

func (s *session) Expired() bool {
	
	return time.Since(s.createdAt) > time.Duration(sessionMaxAge)*time.Second
}

func newSessionStore(db *gorm.DB) Store {
	return &sessionStore{db: db}
}

func (ss *sessionStore) GetSession(c echo.Context) (Session, error) {
	cookie, err := c.Cookie(sessionCookieName)
	token := cookie.Value
	if err != nil {
		return nil, err
	}

	return ss.GetSessionByToken(token)
}

func (ss *sessionStore) GetSessionByToken(token string) (Session, error) {
	var r model.SessionRecord
	
	err := ss.db.First(&r, &model.SessionRecord{Token: token}).Error
	if err != nil {
		return nil, err
	}

	return newSession(ss.db, r.Token, r.UserID, r.CreatedAt), nil
}

func (ss *sessionStore) GetSessionsByUserID(userID string) ([]Session, error) {
	var rs []model.SessionRecord

	err := ss.db.Find(&rs, &model.SessionRecord{UserID: userID}).Error
	if err != nil {
		return nil, err
	}

	sessions := make([]Session, len(rs))
	for _, t := range rs {
		sessions = append(sessions, newSession(ss.db, t.Token, t.UserID, t.CreatedAt))
	}

	return sessions, nil
}

func (ss *sessionStore) IssueSession(userID string) (Session, error) {
	r := model.SessionRecord {
		Token: random.SecureAlphaNumeric(64),
		UserID: userID,
		CreatedAt: time.Now(),
	}

	err := ss.db.Create(&r).Error
	if err != nil {
		return nil, err
	}

	return newSession(ss.db, r.Token, r.UserID, r.CreatedAt), nil
}