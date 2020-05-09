package hls

import (
	"os"
	"path"
	"strings"
	"testing"
)

func TestCmdExecuteFfmpeg(t *testing.T) {
	base, _ := os.Getwd()
	base = strings.Replace(base, "/hls", "", 1)

	targetPath := path.Join(base, "assets", "hls")
	srcPath := path.Join(base, "assets", "raw", "sample.mov")
	ffmpegPath := "/usr/local/bin/ffmpeg"

	GenerateHLS(ffmpegPath, srcPath, targetPath, "480p")
}
