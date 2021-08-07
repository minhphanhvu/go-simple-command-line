package main

import (
	"errors"
	"testing"
)

func TestValidateConf(t *testing.T) {
	testCases := []struct{
		expectedErrors []error
		conf projectConfig
	}{
		{
			conf: projectConfig{
				Name: "Project1",
				localPath: "./repo/Project1",
				RepoURL: "https://github.com/minhphanh.vu/Project1",
				StaticAssets: false,
			},
			expectedErrors: []error{},
		},
		{
			conf: projectConfig{
				RepoURL: "https://github.com/minhphanh.vu/Project1",
				StaticAssets: false,
			},
			expectedErrors: []error{
				errors.New("Project name cannot be empty."),
				errors.New("Project location cannot be empty."),
			},
		},
	}
	for _, tc := range testCases {
		errs := validateConf(tc.conf)
		if len(tc.expectedErrors) == 0 && len(errs) != 0 {
			t.Errorf("Expected no errors, got: %v", errs)
		}

		if len(tc.expectedErrors) != 0{
			for i, e := range tc.expectedErrors {
				if errs[i] == nil || errs[i].Error() != e.Error() {
					t.Errorf("Expected error: %v, got: %v", e, errs[i])
				}
			}
		}
	}
}