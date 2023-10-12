package app

type SaveNameRequest struct {
}
type SaveNameResponse struct {
}

func (n *NameMessageAdapter) SaveName(req *SaveNameRequest) (*SaveNameResponse, error) {
	return nil, nil
}
