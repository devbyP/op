package storage

import (
	"name_message/core/domain"
	"time"
)

func (n *NameStorage) GetName() *domain.NameMessage {
	data := n.store.db.Data[0]
	return &domain.NameMessage{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: time.Unix(data.CreatedAt, 0),
	}
}
