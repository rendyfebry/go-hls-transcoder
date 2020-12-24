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

	err := GenerateHLS(srcPath, targetPath, "480p")
	if err != nil {
		panic(err)
	}
}
