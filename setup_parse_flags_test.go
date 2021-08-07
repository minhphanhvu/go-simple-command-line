package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestSetupParseFlags(t *testing.T) {
	testCases := []struct {
		args           []string
		err            error
		expectedConf   projectConfig
		expectedOutput string
	}{
		{
			args: []string{"-n=Project1", "-d=./repo/Project1", "-r=https://github.com/finn/Project1"},
			err:  nil,
			expectedConf: projectConfig{
				Name:         "Project1",
				localPath:    "./repo/Project1",
				RepoURL:      "https://github.com/finn/Project1",
				StaticAssets: false,
			},
		},
		{
			args:         []string{"wComd"},
			err:          errors.New("Wrong command or no positional parameters expected."),
			expectedConf: projectConfig{},
		},
		{
			args:           []string{"-h"},
			err:            errors.New("flag: help requested"),
			expectedConf:   projectConfig{},
			expectedOutput: "Usage of scaffold-gen:",
		},
	}

	byteBuf := new(bytes.Buffer)
	count := 1
	for _, tc := range testCases {
		fmt.Fprintf(os.Stdout, "Test number: %d\n", count)
		conf, err := setupParseFlags(byteBuf, tc.args)
		if tc.err == nil && err != nil {
			t.Errorf("Expected non-nil error, Got: %v", err)
		}

		if tc.err != nil {
			if err == nil || err.Error() != tc.err.Error() {
				t.Errorf("Expected error: %v, got: %v", tc.err, err)
			}
		}

		if conf != tc.expectedConf {
			t.Errorf("Expected: %+v, Got: %+v", tc.expectedConf, conf)
		}

		if len(tc.expectedOutput) != 0 {
			actualOutput := byteBuf.String()
			if strings.Index(actualOutput, tc.expectedOutput) == -1 {
				t.Errorf("Expected output: %s, Got: %s", tc.expectedOutput, actualOutput)
			}
		}
		byteBuf.Reset()
		count += 1
	}
}
