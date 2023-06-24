package store

import (
	"sync"
)

type Store struct {
	sessionRepository  *SessionRepository
	vineyardRepository *VineyardRepository
	sync.RWMutex
}

var store Store

func New() *Store {
	return &Store{
		sessionRepository: &SessionRepository{},
	}
}

func (s *Store) Store() Store {
	return store
}

func (s *Store) Session() *SessionRepository {
	return s.sessionRepository
}

func (s *Store) Vineyard() *VineyardRepository {
	return s.vineyardRepository
}
