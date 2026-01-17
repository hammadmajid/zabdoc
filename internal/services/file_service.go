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
// Order is preserved to match the original upload order.
func (fs *FileService) ProcessUploadedImages(fileHeaders []*multipart.FileHeader) ([]dto.Image, error) {
	type result struct {
		index int
		img   dto.Image
		err   error
	}
	results := make(chan result, len(fileHeaders)) // Buffered channel to collect results from goroutines

	for i, fileHeader := range fileHeaders {
		// Launch a goroutine for each file to process them concurrently
		go func(idx int, fh *multipart.FileHeader) {
			file, err := fh.Open()
			if err != nil {
				fs.Logger.Printf("Error opening file: %v", err)
				results <- result{index: idx, err: err}
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
				results <- result{index: idx, err: err}
				return
			}

			if len(fileData) == 0 {
				fs.Logger.Printf("Empty image file: %s", fh.Filename)
				results <- result{index: idx, err: err}
				return
			}

			base64Data := base64.StdEncoding.EncodeToString(fileData)

			image := dto.Image{
				Data:     base64Data,
				MimeType: mimeType,
			}
			fs.Logger.Printf("Image processed: %s (%d bytes → %d base64 chars)", mimeType, len(fileData), len(base64Data))
			results <- result{index: idx, img: image} // Send result back to main goroutine
		}(i, fileHeader)
	}

	// Pre-allocate slice to maintain order
	images := make([]dto.Image, len(fileHeaders))

	// Collect results from all goroutines
	for range fileHeaders {
		res := <-results // Wait for each goroutine to send its result
		if res.err != nil {
			return nil, res.err
		}
		images[res.index] = res.img // Place at original index
	}
	return images, nil
}
