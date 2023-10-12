package storage

import (
	"time"

	"name_message/framework/db/model"
)

func (n *NameStorage) SaveName(name string) (int, error) {
	nid := len(n.store.db.Data)
	n.store.db.Data = append(n.store.db.Data, &model.NameMessageDB{
		ID:        nid,
		Name:      name,
		CreatedAt: time.Now().Unix(),
	})
	return nid, nil
}
