package tests

import (
	"testing"

	"github.com/golang/mock/gomock"
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
		// These are the expected return types for google drive specific mimeTypes
		{"application/vnd.google-apps.audio", ".mp3", true},
		{"application/vnd.google-apps.document", ".docx", true},
		{"application/vnd.google-apps.drive-sdk", ".unknown", true},
		{"application/vnd.google-apps.drawing", ".png", true},
		{"application/vnd.google-apps.file", ".unknown", true},
		{"application/vnd.google-apps.folder", ".folder", true},
		{"application/vnd.google-apps.form", ".form", true},
		{"application/vnd.google-apps.fusiontable", ".table", true},
		{"application/vnd.google-apps.jam", ".jam", true},
		{"application/vnd.google-apps.mail-layout", ".email", true},
		{"application/vnd.google-apps.map", ".map", true},
		{"application/vnd.google-apps.photo", ".jpg", true},
		{"application/vnd.google-apps.presentation", ".pptx", true},
		{"application/vnd.google-apps.script", ".js", true},
		{"application/vnd.google-apps.shortcut", ".shortcut", true},
		{"application/vnd.google-apps.site", ".html", true},
		{"application/vnd.google-apps.spreadsheet", ".xlsx", true},
		{"application/vnd.google-apps.unknown", ".unknown", true},
		{"application/vnd.google-apps.vid", ".mp4", true},
		{"application/vnd.google-apps.video", ".mp4", true},
	}
}
