package hls

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	cfg "github.com/rendyfebry/go-hls-transcoder/config"
)

type Variant struct {
	Filename   string
	Prefix     string
	Bandwidth  string
	Resolution string
	Codecs     string
}

// GenerateHLSVariant will generate variants info from the given resolutions
func GenerateHLSVariant(resOptions []string) (variants []*Variant, err error) {
	if len(resOptions) == 0 {
		return nil, errors.New("Please give at least 1 resolutions.")
	}

	for _, r := range resOptions {
		c, err := cfg.GetConfig(r)
		if err != nil {
			continue
		}

		v := &Variant{
			Filename:   c.Name,
			Bandwidth:  c.Bandwidth,
			Resolution: c.Resolution,
		}

		variants = append(variants, v)
	}

	if len(resOptions) == 0 {
		return nil, errors.New("No valid resolutions found.")
	}

	return variants, nil
}

// GeneratePlaylist will generate playlist file from the given variants
func GeneratePlaylist(targetPath, filename, varLocPrefix string, variants []*Variant) {
	// Set default filename
	if filename == "" {
		filename = "playlist.m3u8"
	}

	// M3U Header
	data := "#EXTM3U\n"
	data += "#EXT-X-VERSION:3\n"

	// Add M3U Info for each variant
	for _, v := range variants {
		// Filename & bandwidth is required,
		// if not found we will excluded them from the playlist
		if v.Filename == "" || v.Bandwidth == "" {
			continue
		}

		data += "#EXT-X-STREAM-INF:"
		data += fmt.Sprintf("BANDWIDTH:%s", v.Bandwidth)
		if v.Resolution != "" {
			data += fmt.Sprintf("RESOLUTION:%s", v.Resolution)
		}
		if v.Codecs != "" {
			data += fmt.Sprintf("CODECS:%s", v.Codecs)
		}

		varLoc := v.Filename
		if v.Prefix != "" {
			varLoc = fmt.Sprintf("%s/%s", v.Prefix, v.Filename)
		}

		data += fmt.Sprintf("\n%s", varLoc)
	}

	// Write everything to the file
	f, _ := os.Create(filepath.Join(targetPath, filename))
	defer f.Close()

	f.Write([]byte(data))
}
