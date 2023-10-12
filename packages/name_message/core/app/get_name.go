package app

import "name_message/core/domain"

type GetNameResponse struct {
	data *domain.NameMessage
}

func (n *NameMessageAdapter) GetName() (*GetNameResponse, error) {
	res := &GetNameResponse{data: n.store.GetName()}
	return res, nil
}
