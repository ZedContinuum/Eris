package Eris

import (
	"fmt"
	"time"
)

type eTime struct {
	EngineComponent
	frames    int
	total     int
	second    float32
	current   int64
	lastFrame int64
}

func (t *eTime) Init(m *Eris) error {
	t.master = m
	t.lastFrame = time.Now().UnixNano()
	return nil
}

func (t *eTime) Update() {
	t.current = time.Now().UnixNano()
	t.master.DeltaTime = float32((t.current - t.lastFrame)) / float32(time.Second)
	t.second += t.master.DeltaTime
	if t.second >= 1 {
		fmt.Println("Seconds passed: ", t.total, " | Frames: ", t.frames, " | Actual Second: ", t.second)
		t.second = 0
		t.frames = 0
		t.total++
	}
	t.frames++
	t.lastFrame = t.current
	if t.total >= 10 {
		t.master.Quit()
	}
}

func (t *eTime) OnVisible() {
	t.lastFrame = time.Now().UnixNano()
	t.current = 0
}
