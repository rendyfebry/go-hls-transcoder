// Package hls provides a few functionalities to generate HLS
// files using ffmpeg.
package hls

import (
	"os"
	"os/exec"
)

// GenerateHLS will generate HLS file based on resolution presets.
// The available resolutions are: 360p, 480p, 720p and 1080p.
func GenerateHLS(srcPath, targetPath, resolution string) error {
	options, err := getOptions(srcPath, targetPath, resolution)
	if err != nil {
		return err
	}

	return GenerateHLSCustom(options)
}

// GenerateHLSCustom will generate HLS using the flexible options params.s
// options is array of string that accepted by ffmpeg command
func GenerateHLSCustom(options []string) error {
	cmd := exec.Command("ffmpeg", options...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	return err
}
