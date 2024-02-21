package entity

import "github.com/hajimehoshi/ebiten/v2"

type Manager struct {
	Entities []*Entity
}

func NewManager() *Manager {
	return &Manager{
		Entities: make([]*Entity, 0),
	}
}

func (m *Manager) Update() error {
	player := m.Entities[0]
	for _, e := range m.Entities {
		if e != player {
			e.MoveTowards(player)
		}
	}
	return nil
}

func (m *Manager) AddEntity(e *Entity) {
	m.Entities = append(m.Entities, e)
}

func (m *Manager) Draw(screen *ebiten.Image) {
	for _, e := range m.Entities {
		e.Draw(screen)
	}
}

func (m *Manager) RemoveEntity(e *Entity) {
	for i, entity := range m.Entities {
		if entity == e {
			m.Entities = append(m.Entities[:i], m.Entities[i+1:]...)
			break
		}
	}
}

func (m *Manager) MovePlayer(x, y int) {
	m.Entities[0].x += float64(x) * m.Entities[0].speed
	m.Entities[0].y += float64(y) * m.Entities[0].speed
}
