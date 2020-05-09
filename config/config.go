package hls

import "errors"

type config struct {
	Name         string
	VideoBitrate string
	Maxrate      string
	BufSize      string
	AudioBitrate string
}

var preset = map[string]*config{
	"360p": {
		Name:         "360p",
		VideoBitrate: "800k",
		Maxrate:      "856k",
		BufSize:      "1200k",
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

// GetConfig return config from the available preset
func GetConfig(res string) (*config, error) {
	cfg, ok := preset[res]
	if !ok {
		return nil, errors.New("Preset not found")
	}

	return cfg, nil
}
