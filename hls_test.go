package hls

import (
	"os"
	"path"
	"testing"
)

func TestCmdExecuteFfmpeg(t *testing.T) {
	base, _ := os.Getwd()

	targetPath := path.Join(base, "static")
	srcPath := path.Join(base, "static", "sample.mov")
	ffmpegPath := "/usr/local/bin/ffmpeg"

	err := GenerateHLS(ffmpegPath, srcPath, targetPath, "480p")
	if err != nil {
		panic(err)
	}
}
