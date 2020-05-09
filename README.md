# Go HLS Transcoder

Simple golang HLS transcoder with ffmpeg

## Prerequisite

- Golang 1.14
- ffmpeg

## Examples

```
package main

import (
	hls "github.com/rendyfebry/go-hls-transcoder"
)

func main() {
	ffmpegPath := "/usr/local/bin/ffmpeg"
	srcPath := "/assets/raw/movie.mov"
	targetPath := "/assets/hls"

	hls.GenerateHLS(ffmpeg, srcPath, targetPath, "720p")
}
```
