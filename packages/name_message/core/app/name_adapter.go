package app

type NameMessageAdapter struct {
	store DbPort
}

func NewNameMessageAdapter(db DbPort) *NameMessageAdapter {
	return &NameMessageAdapter{store: db}
}
