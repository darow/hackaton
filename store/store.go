package store

import (
	"hackaton/internal/models"
	"sync"
)

type Store struct {
	SessionRepo *SessionRepository `json:"sessionRepo"`
	sync.RWMutex
}

func New() *Store {
	m := make(map[int]*models.Session)
	m[1] = &models.Session{}

	return &Store{
		SessionRepo: &SessionRepository{
			Sessions: m,
		},
	}
}

func (s *Store) Session() *SessionRepository {
	return s.SessionRepo
}
