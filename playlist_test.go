package hls

import (
	"errors"
	"fmt"
	"os"
	"path"
	"testing"
)

func TestGenerateHLSVariant(t *testing.T) {
	cases := []struct {
		Name          string
		ResOptions    []string
		Prefix        string
		IsError       bool
		ExpectedError string
	}{
		{
			Name:       "test 1",
			ResOptions: []string{"360p", "480p"},
		},
		{
			Name:          "test 2",
			ResOptions:    []string{},
			IsError:       true,
			ExpectedError: "Please give at least 1 resolutions.",
		},
		{
			Name:          "test 3",
			ResOptions:    []string{"invalid"},
			IsError:       true,
			ExpectedError: "No valid resolutions found.",
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(tester *testing.T) {
			variants, err := GenerateHLSVariant(test.ResOptions, test.Prefix)

			if test.IsError {
				if err.Error() != test.ExpectedError {
					fmt.Println(err)
					panic(errors.New("Not expected error"))
				}
			} else {
				if err != nil {
					panic(err)
				}

				if len(variants) != len(test.ResOptions) {
					panic(errors.New("Result should be the same"))
				}
			}
		})
	}
}

func TestGeneratePlaylist(t *testing.T) {
	variants, err := GenerateHLSVariant([]string{"480p"}, "")
	if err != nil {
		panic(err)
	}

	base, _ := os.Getwd()
	targetPath := path.Join(base, "static")

	GeneratePlaylist(variants, targetPath, "")
}
