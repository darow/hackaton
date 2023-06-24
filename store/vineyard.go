package store

import (
	"hackaton/internal/models"
)

type VineyardRepository struct {
	Store      *Store
	Vineyards  map[int]*models.Vineyard
	VineyardID int
}

func (r *VineyardRepository) Post(v *models.Vineyard) {
	if r.VineyardID == 0 {
		r.VineyardID++
	}

	_, ok := r.Vineyards[v.ID]
	if !ok {
		v.ID = r.VineyardID
	}
	r.Vineyards[v.ID] = v

	return
}
