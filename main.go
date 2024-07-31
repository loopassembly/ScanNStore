package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// getModuleRootDir returns the root directory of the module.
func getModuleRootDir() string {
	// Assuming the module root directory is the directory where the main.go resides
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	return dir
}

var testDataDir = filepath.Join(getModuleRootDir(), "genai", "testdata")

// uploadFile uploads the given file to the service and returns a [genai.File]
// representing it. mimeType optionally specifies the MIME type of the data in
// the file; if set to "", the service will try to automatically determine the
// type from the data contents.
func uploadFile(ctx context.Context, client *genai.Client, path, mimeType string) (*genai.File, error) {
	osf, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer osf.Close()

	opts := &genai.UploadFileOptions{MIMEType: mimeType}
	file, err := client.UploadFile(ctx, "", osf, opts)
	if err != nil {
		return nil, err
	}

	for file.State == genai.FileStateProcessing {
		log.Printf("processing %s", file.Name)
		time.Sleep(5 * time.Second)
		var err error
		file, err = client.GetFile(ctx, file.Name)
		if err != nil {
			return nil, err
		}
	}
	if file.State != genai.FileStateActive {
		return nil, fmt.Errorf("uploaded file has state %s, not active", file.State)
	}
	return file, nil
}

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyBhJ709JA3k4OIscgN1yWW5HmrKJrpPpes"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	imagePath := filepath.Join(testDataDir, "/home/loopassembly/Documents/ScanNStore/iu.jpg")
	file, err := uploadFile(ctx, client, imagePath, "")
	if err != nil {
		log.Fatal(err)
	}
	defer client.DeleteFile(ctx, file.Name)

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx,
		genai.FileData{URI: file.URI},
		genai.Text("Can you tell me about the instruments in this photo?"))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
