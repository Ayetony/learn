// Package music - manager the data
package music

import (
	"errors"
	"fmt"
	"time"
)

type MusicEntry struct {
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

// Player 定义播放接口
type Player interface {
	Play(source string) (progress int)
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index > len(m.musics) {
		return nil, errors.New("index out of range")
	}
	return &m.musics[index], err
}

func (m *MusicManager) Find(name string) (entry *MusicEntry) {
	if len(m.musics) == 0 {
		return nil
	}
	for _, entry := range m.musics {
		if entry.Name == name {
			return &entry
		}
	}
	return nil
}

func (musicEntry *MusicEntry) ModifyMusicEntry(entry *MusicEntry) *MusicEntry {
	if entry == nil {
		return nil
	}
	if entry.Artist != "" {
		musicEntry.Artist = entry.Artist
	}
	if entry.Name != "" {
		musicEntry.Name = entry.Name
	}
	if entry.Source != "" {
		musicEntry.Source = entry.Source
	}
	if entry.Type != "" {
		musicEntry.Type = entry.Type
	}
	return musicEntry
}

func (m *MusicManager) Add(entrys ...*MusicEntry) {
	if len(entrys) > 0 {
		for _, entry := range entrys {
			m.musics = append(m.musics, *entry)
		}
	}
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index > len(m.musics) {
		return nil
	}
	removedMusic := &m.musics[index]
	//删除后，调整数组
	// index 处在中间索引
	if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 {
		// 否则只剩下一个
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = m.musics[:index-1]
	}
	return removedMusic
}

type Mp3Player struct {
	stat     int
	progress int
}

func (p *Mp3Player) Play(source string) (progress int) {

	fmt.Printf("Playing the music %s", source)
	p.progress = 0

	for p.progress < 100 {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("*")
		p.progress += 10
	}
	return p.progress
}

func Play(source, mtype string) {
	var p Player
	switch mtype {
	case "mp3":
		p = &Mp3Player{}
		fmt.Println(p.Play(source))
	default:
		fmt.Println("Music End")
	}

}
