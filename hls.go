package hls

import (
	"os"
	"os/exec"
	"path/filepath"
)

type Config struct {
	Name         string
	VideoBitrate string
	Maxrate      string
	BufSize      string
	AudioBitrate string
}

var resConfig = map[string]*Config{
	"360p": {
		Name:         "360p",
		VideoBitrate: "800k",
		Maxrate:      "856k",
		BufSize:      "1200sk",
		AudioBitrate: "96k",
	},
	"480p": {
		Name:         "480p",
		VideoBitrate: "1400k",
		Maxrate:      "1498k",
		BufSize:      "2100k",
		AudioBitrate: "128k",
	},
	"720p": {
		Name:         "720p",
		VideoBitrate: "2800k",
		Maxrate:      "2996k",
		BufSize:      "4200k",
		AudioBitrate: "128k",
	},
	"1080p": {
		Name:         "1080p",
		VideoBitrate: "5000k",
		Maxrate:      "5350k",
		BufSize:      "7500k",
		AudioBitrate: "192k",
	},
}

func getOptions(srcPath, targetPath, res string) []string {
	config := resConfig[res]

	filenameTS := filepath.Join(targetPath, res+"_%03d.ts")
	filenameM3U8 := filepath.Join(targetPath, res+".m3u8")

	options := []string{
		"-hide_banner",
		"-y",
		"-i", srcPath,
		"-vf", "scale=trunc(oh*a/2)*2:1080",
		"-c:a", "aac",
		"-ar", "48000",
		"-c:v", "h264",
		"-profile:v", "main",
		"-crf", "20",
		"-sc_threshold", "0",
		"-g", "48",
		"-keyint_min", "48",
		"-hls_time", "10",
		"-hls_playlist_type", "vod",
		"-b:v", config.VideoBitrate,
		"-maxrate", config.Maxrate,
		"-bufsize", config.BufSize,
		"-b:a", config.AudioBitrate,
		"-preset", "ultrafast",
		"-hls_segment_filename", filenameTS,
		filenameM3U8,
	}

	return options
}

func GenerateHLS(ffmpegPath, srcPath, targetPath, res string) error {
	options := getOptions(srcPath, targetPath, res)

	cmd := exec.Command(ffmpegPath, options...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	return err
}
