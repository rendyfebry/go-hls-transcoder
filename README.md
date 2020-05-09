# Go HLS Transcoder

[![GoDoc](https://godoc.org/github.com/rendyfebry/go-hls-transcoder?status.svg)](https://godoc.org/github.com/rendyfebry/go-hls-transcoder)

Simple golang HLS transcoder with ffmpeg

## Prerequisite

- Golang 1.14
- ffmpeg

## Examples

Generate HLS files + playlist from a mov file.

```go
package main

import (
	hls "github.com/rendyfebry/go-hls-transcoder"
	hlsPlaylist "github.com/rendyfebry/go-hls-transcoder/playlist"
)

func main() {
	ffmpegPath := "/usr/local/bin/ffmpeg"
	srcPath := "/assets/raw/movie.mov"
	targetPath := "/assets/hls"
	resOptions := []string{"480p", "720p"}

	variants, _ := hlsPlaylist.GenerateHLSVariant(resOptions, "")
	hlsPlaylist.GeneratePlaylist(variants, targetPath, "")

	for _, res := range resOptions {
		hls.GenerateHLS(ffmpeg, srcPath, targetPath, res)
	}
}
```

## License

MIT licensed. See the LICENSE.md file for details.
