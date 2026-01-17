package music

import "fmt"

type music struct {
	name       string
	instrument string
	play       bool
}

func NewMember(name, instrument string) *music {
	return &music{
		name:       name,
		instrument: instrument,
		play:       false,
	}
}

func (m *music) GetName() {
	fmt.Println(m.name)
}

func (m *music) PlaySong() {
	m.play = true
	fmt.Println("Song playing...")
}

func (m *music) StopSong() {
	m.play = false
	fmt.Println("Song stopped")
}
