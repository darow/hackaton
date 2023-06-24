package store

import (
	"hackaton/internal/models"
)

type VineyardRepository struct {
	store      *Store
	vineyards  map[int]*models.Vineyard
	vineyardID int
}

func (r *VineyardRepository) Post(v *models.Vineyard) {
	if r.vineyardID == 0 {
		r.vineyardID++
	}

	_, ok := r.vineyards[v.ID]
	if !ok {
		v.ID = r.vineyardID
	}
	r.vineyards[v.ID] = v

	return
}
