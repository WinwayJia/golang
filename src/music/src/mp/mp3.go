package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat     int
	progress int
}

func (m *MP3Player) Play(source string) {
	fmt.Println("playing mp3 music", source)
	m.progress = 0

	for m.progress < 100 {
		time.Sleep(100 * time.Millisecond) // 假装播放
		fmt.Print(".")
		m.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
