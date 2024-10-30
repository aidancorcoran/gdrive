# gdrive - Google Drive CLI Tool

`gdrive` is a Command Line Interface (CLI) tool, written in Go, designed to help you interact with your Google Drive seamlessly from the terminal. The CLI simplifies tasks like uploading, downloading, and listing files, enabling a Git-like workflow for your Google Drive content.

## Features

- **List** files in your Google Drive directory
- **Upload** files directly to your Google Drive
- **Download** files from Google Drive to your local machine
- **OAuth2 Authentication** is pre-configured, so users don't need to set up credentials.
- **Extensible Commands** for future expansion (similar to `git`)

## Getting Started

### Prerequisites

- **Go version 1.23 or later** (to build from source)

### Installation

Currently, `gdrive` can be compiled from source:

```bash
# Clone the repository
git clone https://github.com/aidancorcoran/gdrive.git
cd gdrive

# Build the executable
go build

# Build the binary and add it to your PATH
go build -o gdrive
```

### Unit Testing

`gdrive` includes a suite of unit tests to ensure that the application functions as expected. These tests cover various scenarios, including handling of different MIME types, error cases, and validation of core functionality.

To run the unit tests, follow these steps:

1. **Navigate to the Project Directory**: Make sure you are in the root directory of the project (where `go.mod` is located).

2. **Run All Tests**: Use the following command to run all unit tests in the project:

   ```bash
   go test ./tests
   ```