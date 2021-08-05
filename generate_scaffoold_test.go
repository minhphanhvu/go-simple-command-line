package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestGenerateScaffold(t *testing.T) {
	testCases := []struct {
		conf projectConfig
		expectOutput string
	}{
		{
			conf: projectConfig{
				Name: "Project1",
				localPath: "/mydir/project1",
				RepoURL: "https://github.com/minhphanhvu/project1",
				StaticAssets: false,
			},
			expectOutput: "Generating scaffold for project Project1 in /mydir/project1\n",
		},
		{
			conf: projectConfig{
				Name: "Project2",
				localPath: "/mydir/project2",
				RepoURL: "https://github.com/minhphanhvu/project2",
				StaticAssets: false,
			},
			expectOutput: "Generating scaffold for project Project2 in /mydir/project2\n",
		},
	}

	byteBuf := new(bytes.Buffer)
	for _, tc := range testCases {
		generateScaffold(byteBuf, tc.conf)
		returnedOutput := byteBuf.String()
		if strings.Index(returnedOutput, tc.expectOutput) == -1 {
			t.Errorf("Expected output: %s, Got: %s", tc.expectOutput, returnedOutput)
		}
		byteBuf.Reset()
	}
}