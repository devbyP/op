package storage

import (
	"time"

	"op/internal/model"
)

type (
	IMessage interface {
		GetMessageByID(id int) (*model.MessageData, error)
		GetMessages() ([]*model.MessageData, error)
		SaveMessage(message string, t uint8) (int, error)
	}

	MessageStorage struct {
		store *Storage[*model.MessageData]
	}
)

func NewMessageStorage(path string) *MessageStorage {
	return &MessageStorage{
		store: Must[*model.MessageData](New[*model.MessageData](path)),
	}
}

func (ms *MessageStorage) GetMessages() ([]*model.MessageData, error) {
	return ms.store.db.Data, nil
}

func (ms *MessageStorage) GetMessageByID(id int) (*model.MessageData, error) {
	return ms.store.db.Data[id], nil
}

func (ms *MessageStorage) SaveMessage(mess string, t model.MessageType) (id int, err error) {
	id = GenerateIncrementalID(ms.store.db.Data)
	payload := &model.MessageData{
		ID:        id,
		Message:   mess,
		Type:      t,
		CreatedAt: time.Now().Unix(),
	}
	ms.store.db.Data = append(ms.store.db.Data, payload)
	return
}
