package library

import "errors"

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (musics *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if 0 == len(m.musics) {
		return nil
	}
	for _, v := range m.musics {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removedMusic := &m.musics[index]
	if 0 == index {
		m.musics = m.musics[1:]
	} else if index == len(m.musics)-1 {
		m.musics = m.musics[:len(m.musics)-1]
	} else {
		m.musics = append(m.musics[0:index-1], m.musics[index+1:]...)
	}
	return removedMusic
}

func (m *MusicManager) RemoveByName(name string) (removed []MusicEntry) {
	var left []MusicEntry
	for i, v := range m.musics {
		if v.Name == name {
			removed = append(removed, m.musics[i])
		} else {
			left = append(left, m.musics[i])
		}
	}
	m.musics = left
	return
}
