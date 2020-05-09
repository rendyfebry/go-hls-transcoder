package hls

import (
	"fmt"
	"os"
	"path/filepath"

	cfg "github.com/rendyfebry/go-hls-transcoder/config"
)

// GenerateManifest will generate manifest file based on given resolutions
func GenerateManifest(targetPath string, resOptions []string) {
	f, _ := os.Create(filepath.Join(targetPath, "manifest.m3u8"))
	defer f.Close()

	data := "#EXTM3U\n"
	data += "#EXT-X-VERSION:3\n"

	for _, r := range resOptions {
		c, err := cfg.GetConfig(r)
		if err != nil {
			continue
		}

		data += fmt.Sprintf("#EXT-X-STREAM-INF:BANDWIDTH=%s,RESOLUTION=%s\n", c.Bandwidth, c.Resolution)
		data += fmt.Sprintf("%s.m3u8\n", c.Name)
	}

	f.Write([]byte(data))
}
