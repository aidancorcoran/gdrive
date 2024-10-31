/*
Copyright © 2024 Aidan Corcoran <aidancorcoran.dev@gmail.com>
*/

package tests

import (
	"testing"

	"github.com/aidancorcoran/gdrive/cmd"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/drive/v3"
)

// Mocking the Google Drive Service
type MockDriveService struct {
	mockCtrl *gomock.Controller
	service  *drive.Service
}

func TestGetFileExtension(t *testing.T) {
	tests := []struct {
		mimeType    string
		expectedExt string
		expectedErr bool
	}{
		// Valid cases
		// These are the expected return types for google drive specific mimeTypes
		{"application/vnd.google-apps.audio", ".mp3", false},
		{"application/vnd.google-apps.document", ".docx", false},
		{"application/vnd.google-apps.drive-sdk", ".unknown", false},
		{"application/vnd.google-apps.drawing", ".png", false},
		{"application/vnd.google-apps.file", ".unknown", false},
		{"application/vnd.google-apps.folder", ".folder", false},
		{"application/vnd.google-apps.form", ".form", false},
		{"application/vnd.google-apps.fusiontable", ".table", false},
		{"application/vnd.google-apps.jam", ".jam", false},
		{"application/vnd.google-apps.mail-layout", ".email", false},
		{"application/vnd.google-apps.map", ".map", false},
		{"application/vnd.google-apps.photo", ".jpg", false},
		{"application/vnd.google-apps.presentation", ".pptx", false},
		{"application/vnd.google-apps.script", ".js", false},
		{"application/vnd.google-apps.shortcut", ".shortcut", false},
		{"application/vnd.google-apps.site", ".html", false},
		{"application/vnd.google-apps.spreadsheet", ".xlsx", false},
		{"application/vnd.google-apps.unknown", ".unknown", false},
		{"application/vnd.google-apps.vid", ".mp4", false},
		{"application/vnd.google-apps.video", ".mp4", false},
		// These are the expected return types for other mimeTypes
		{"application/vnd.oasis.opendocument.text", ".odt", false},
		{"application/vnd.openxmlformats-officedocument.wordprocessingml.document", ".docx", false},
		{"application/rtf", ".rtf", false},
		{"application/pdf", ".pdf", false},
		{"text/plain", ".txt", false},
		{"application/zip", ".zip", false},
		{"application/epub+zip", ".epub", false},
		{"text/markdown", ".md", false},
		{"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", ".xlsx", false},
		{"application/x-vnd.oasis.opendocument.spreadsheet", ".ods", false},
		{"text/csv", ".csv", false},
		{"text/tab-separated-values", ".tsv", false},
		{"application/vnd.openxmlformats-officedocument.presentationml.presentation", ".pptx", false},
		{"application/vnd.oasis.opendocument.presentation", ".odp", false},
		{"image/jpeg", ".jpg", false},
		{"image/png", ".png", false},
		{"image/svg+xml", ".svg", false},
		{"application/vnd.google-apps.script+json", ".json", false},

		// Error cases: unknown MIME type
		{"random/mime-type", "", true},
	}

	for _, test := range tests {
		ext, err := cmd.GetFileExtension(test.mimeType)
		if test.expectedErr {
			assert.Error(t, err) // Ensure an error was returned by GetFileExtension
		} else {
			assert.NoError(t, err)
			assert.Equal(t, test.expectedExt, ext) // Ensure nil was returned by GetFileExtension
		}
	}
}

func TestGetFileIdAndMimeType(t *testing.T) {

}
