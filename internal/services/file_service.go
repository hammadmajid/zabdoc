package services

import (
	"encoding/base64"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"path/filepath"
	"zabdoc/internal/api/dto"
)

// FileService handles file/image processing.
type FileService struct {
	Logger *log.Logger
}

func NewFileService(logger *log.Logger) *FileService {
	return &FileService{Logger: logger}
}

// ProcessUploadedImages processes uploaded image files and returns Image slices.
func (fs *FileService) ProcessUploadedImages(fileHeaders []*multipart.FileHeader) ([]dto.Image, error) {
	type result struct {
		img dto.Image
		err error
	}
	results := make(chan result, len(fileHeaders)) // Buffered channel to collect results from goroutines

	for _, fileHeader := range fileHeaders {
		// Launch a goroutine for each file to process them concurrently
		go func(fh *multipart.FileHeader) {
			file, err := fh.Open()
			if err != nil {
				fs.Logger.Printf("Error opening file: %v", err)
				results <- result{err: err}
				return
			}
			defer file.Close()

			ext := filepath.Ext(fh.Filename)
			mimeType := mime.TypeByExtension(ext)
			if mimeType == "" {
				mimeType = "image/jpeg"
			}

			fileData, err := io.ReadAll(file)
			if err != nil {
				fs.Logger.Printf("Error reading file: %v", err)
				results <- result{err: err}
				return
			}

			if len(fileData) == 0 {
				fs.Logger.Printf("Empty image file: %s", fh.Filename)
				results <- result{err: err}
				return
			}

			base64Data := base64.StdEncoding.EncodeToString(fileData)

			image := dto.Image{
				Data:     base64Data,
				MimeType: mimeType,
			}
			fs.Logger.Printf("Image processed: %s (%d bytes → %d base64 chars)", mimeType, len(fileData), len(base64Data))
			results <- result{img: image} // Send result back to main goroutine
		}(fileHeader)
	}

	var images []dto.Image
	// Collect results from all goroutines
	for range fileHeaders {
		res := <-results // Wait for each goroutine to send its result
		if res.err != nil {
			return nil, res.err
		}
		images = append(images, res.img)
	}
	return images, nil
}
