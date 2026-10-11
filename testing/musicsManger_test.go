package testing

import (
	"testing"

	"example.com/hello/music"
)

func TestMusicsCrud(t *testing.T) {
	entry0 := &music.MusicEntry{Name: "you and me", Artist: "Huan", Source: "Beijing", Type: "Pop"}
	manager := music.NewMusicManager()
	manager.Add(entry0)

	entry1 := &music.MusicEntry{Name: "Glory Box", Artist: "Portishead", Source: "UK", Type: "Jazz"}
	t.Log(manager.Remove(0))

	manager.Add(entry0, entry1)
	t.Log(manager.Find("Glory Box"))

	music.Play("UK", "mp3")
}
