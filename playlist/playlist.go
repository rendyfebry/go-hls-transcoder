package playlist

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	cfg "github.com/rendyfebry/go-hls-transcoder/config"
)

type Variant struct {
	URL        string
	Bandwidth  string
	Resolution string
	Codecs     string
}

// GenerateHLSVariant will generate variants info from the given resolutions
func GenerateHLSVariant(resOptions []string, locPrefix string) (variants []*Variant, err error) {
	if len(resOptions) == 0 {
		return nil, errors.New("Please give at least 1 resolutions.")
	}

	for _, r := range resOptions {
		c, err := cfg.GetConfig(r)
		if err != nil {
			continue
		}

		url := fmt.Sprintf("%s.m3u8", c.Name)
		if locPrefix != "" {
			url = locPrefix + "/" + url
		}

		v := &Variant{
			URL:        url,
			Bandwidth:  c.Bandwidth,
			Resolution: c.Resolution,
		}

		variants = append(variants, v)
	}

	if len(variants) == 0 {
		return nil, errors.New("No valid resolutions found.")
	}

	return variants, nil
}

// GeneratePlaylist will generate playlist file from the given variants
func GeneratePlaylist(variants []*Variant, targetPath, filename string) {
	// Set default filename
	if filename == "" {
		filename = "playlist.m3u8"
	}

	// M3U Header
	data := "#EXTM3U\n"
	data += "#EXT-X-VERSION:3\n"

	// Add M3U Info for each variant
	for _, v := range variants {
		// URL & bandwidth is required,
		// if not found we will excluded them from the playlist
		if v.URL == "" || v.Bandwidth == "" {
			continue
		}

		data += "#EXT-X-STREAM-INF:"
		data += fmt.Sprintf("BANDWIDTH=%s", v.Bandwidth)
		if v.Resolution != "" {
			data += fmt.Sprintf(",RESOLUTION=%s", v.Resolution)
		}
		if v.Codecs != "" {
			data += fmt.Sprintf(",CODECS=%s", v.Codecs)
		}

		data += fmt.Sprintf("\n%s\n", v.URL)
	}

	// Write everything to the file
	f, _ := os.Create(filepath.Join(targetPath, filename))
	defer f.Close()

	f.Write([]byte(data))
}
