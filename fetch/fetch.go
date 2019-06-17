package fetch

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/francoiscolombo/goytmp3downloader/utube"
	"github.com/schollz/progressbar"
	"gopkg.in/gookit/color.v1"
)

func download(video *utube.VideoMetaData, videoPath string, dontExtractMp3 bool) (string, error) {

	if dontExtractMp3 == false {
		_, err := exec.LookPath("ffmpeg")
		if err != nil {
			return "", fmt.Errorf("ffmpeg not found, no way to extract audio, sorry")
		}
	}

	wspace := regexp.MustCompile(`\W+`)
	title := wspace.ReplaceAllString(video.Title, "-")
	if len(title) > 64 {
		title = title[:64]
	}
	title = strings.TrimRight(strings.ToLower(title), "-")
	filename := fmt.Sprintf("%s/%s.mp4", videoPath, title)

	req, _ := http.NewRequest("GET", video.Format.URL, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var out io.Writer
	f, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	out = f
	defer f.Close()

	bar := progressbar.NewOptions(
		int(resp.ContentLength),
		progressbar.OptionSetBytes(int(resp.ContentLength)),
	)
	out = io.MultiWriter(out, bar)
	io.Copy(out, resp.Body)

	fmt.Println()

	// Extract audio from downloaded video using ffmpeg
	if dontExtractMp3 == false {
		ffmpeg, _ := exec.LookPath("ffmpeg")
		color.C256(25).Println("Extracting audio ...")
		fname := filename
		mp3 := strings.TrimRight(fname, filepath.Ext(fname)) + ".mp3"
		cmd := exec.Command(ffmpeg, "-y", "-i", fname, "-vn", mp3)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return "", fmt.Errorf("Failed to extract audio: %s", err)
		}
		fmt.Println()
		color.C256(62).Println("Extracted audio:", mp3)
	}

	return filename, nil
}

/*
VideoFromYoutube fetch a video from youtube and extract the mp3 if you don't set the flag 'full' to true
*/
func VideoFromYoutube(id, path string, full bool) {
	video, err := utube.GetVideoMetaData(id)
	if err == nil {
		filename, err := download(&video, path, full)
		if err != nil {
			log.Fatal(err)
		}
		if full == false {
			e := os.Remove(filename)
			if e != nil {
				log.Fatal(err)
			}
		}
	}
}
