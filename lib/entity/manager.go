package entity

type Manager struct {
	Entities []*Entity
}

func NewManager() *Manager {
	return &Manager{
		Entities: make([]*Entity, 0),
	}
}

func (m *Manager) Update() error {
	for _, e := range m.Entities {
		e.Update(m.Entities)
	}
	return nil
}

func (m *Manager) AddEntity(e *Entity) {
	m.Entities = append(m.Entities, e)
}
