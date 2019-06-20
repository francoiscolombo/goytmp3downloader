package play

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	"github.com/gdamore/tcell"
)

type audioPanel struct {
	sampleRate beep.SampleRate
	streamer   beep.StreamSeeker
	ctrl       *beep.Ctrl
	resampler  *beep.Resampler
	volume     *effects.Volume
}

func newAudioPanel(sampleRate beep.SampleRate, streamer beep.StreamSeeker) *audioPanel {
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer)}
	resampler := beep.ResampleRatio(4, 1, ctrl)
	volume := &effects.Volume{Streamer: resampler, Base: 2}
	return &audioPanel{sampleRate, streamer, ctrl, resampler, volume}
}

func (ap *audioPanel) play() {
	speaker.Play(ap.volume)
}

func drawTextLine(screen tcell.Screen, x, y int, s string, style tcell.Style) {
	for _, r := range s {
		screen.SetContent(x, y, r, nil, style)
		x++
	}
}

func (ap *audioPanel) draw(screen tcell.Screen, title string) {

	mainStyle := tcell.StyleDefault.
		Background(tcell.NewRGBColor(51, 102, 0)).
		Foreground(tcell.NewRGBColor(204, 255, 204))

	titleStyle := mainStyle.
		Foreground(tcell.NewRGBColor(204, 255, 51)).
		Bold(true)

	statusStyle := mainStyle.
		Foreground(tcell.NewRGBColor(0, 204, 102)).
		Bold(true)

	for i := 0; i <= 8; i++ {
		drawTextLine(screen, 0, i, "                                                                      ", mainStyle)
	}

	drawTextLine(screen, 2, 1, "You are now playing", mainStyle)
	drawTextLine(screen, 22, 1, title, titleStyle)
	drawTextLine(screen, 2, 3, "Press [ESC] to quit.", mainStyle)
	drawTextLine(screen, 2, 4, "Press [SPACE] to pause/resume.", mainStyle)

	speaker.Lock()
	position := ap.sampleRate.D(ap.streamer.Position())
	length := ap.sampleRate.D(ap.streamer.Len())
	volume := ap.volume.Volume
	speaker.Unlock()

	positionStatus := fmt.Sprintf("%v / %v", position.Round(time.Second), length.Round(time.Second))
	volumeStatus := fmt.Sprintf("%.1f", volume)

	drawTextLine(screen, 2, 6, "Position [<-] / [->]", mainStyle)
	drawTextLine(screen, 24, 6, positionStatus, statusStyle)

	drawTextLine(screen, 2, 7, "Volume    [-] / [+]", mainStyle)
	drawTextLine(screen, 24, 7, volumeStatus, statusStyle)

}

func (ap *audioPanel) handle(event tcell.Event) (changed, quit bool) {

	switch event := event.(type) {
	case *tcell.EventKey:

		if event.Key() == tcell.KeyESC {
			return false, true
		}

		if (event.Key() == tcell.KeyRight) || (event.Key() == tcell.KeyLeft) {
			speaker.Lock()
			newPos := ap.streamer.Position()
			if event.Key() == tcell.KeyRight {
				newPos += ap.sampleRate.N(time.Second)
			} else if event.Key() == tcell.KeyLeft {
				newPos -= ap.sampleRate.N(time.Second)
			}
			if newPos < 0 {
				newPos = 0
			} else if newPos >= ap.streamer.Len() {
				newPos = ap.streamer.Len() - 1
			}
			if err := ap.streamer.Seek(newPos); err != nil {
				return false, true
			}
			speaker.Unlock()
			return true, false
		}

		if event.Key() != tcell.KeyRune {
			return false, false
		}

		if event.Rune() == ' ' {
			speaker.Lock()
			ap.ctrl.Paused = !ap.ctrl.Paused
			speaker.Unlock()
			return false, false
		} else if event.Rune() == '-' {
			speaker.Lock()
			ap.volume.Volume -= 0.1
			speaker.Unlock()
			return true, false
		} else if event.Rune() == '+' {
			speaker.Lock()
			ap.volume.Volume += 0.1
			speaker.Unlock()
			return true, false
		}

	}
	return false, false
}

// Mp3 allow you to play mp3 directly from the command line
func Mp3(mp3Path string) error {

	f, err := os.Open(mp3Path)
	if err != nil {
		return err
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/30))

	screen, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	err = screen.Init()
	if err != nil {
		return err
	}
	defer screen.Fini()

	ap := newAudioPanel(format.SampleRate, streamer)

	title := filepath.Base(mp3Path)

	screen.Clear()
	ap.draw(screen, title)
	screen.Show()

	ap.play()

	seconds := time.Tick(time.Second)
	events := make(chan tcell.Event)
	go func() {
		for {
			events <- screen.PollEvent()
		}
	}()

loop:
	for {
		select {
		case event := <-events:
			changed, quit := ap.handle(event)
			if quit {
				break loop
			}
			if changed {
				ap.draw(screen, title)
				screen.Show()
			}
		case <-seconds:
			ap.draw(screen, title)
			screen.Show()
		}
	}

	return nil
}
