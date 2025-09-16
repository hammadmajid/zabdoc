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

// ProcessUploadedImages processes uploaded image files and returns SectionImage slices.
func (fs *FileService) ProcessUploadedImages(fileHeaders []*multipart.FileHeader, filesKey string) ([]dto.SectionImage, error) {
	var images []dto.SectionImage
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			fs.Logger.Printf("Error opening file: %v", err)
			return nil, err
		}
		defer file.Close()

		ext := filepath.Ext(fileHeader.Filename)
		mimeType := mime.TypeByExtension(ext)
		if mimeType == "" {
			mimeType = "image/jpeg"
		}

		fileData, err := io.ReadAll(file)
		if err != nil {
			fs.Logger.Printf("Error reading file: %v", err)
			return nil, err
		}

		if len(fileData) == 0 {
			fs.Logger.Printf("Empty image file: %s", fileHeader.Filename)
			return nil, err
		}

		base64Data := base64.StdEncoding.EncodeToString(fileData)

		sectionImage := dto.SectionImage{
			Data:     base64Data,
			MimeType: mimeType,
		}
		images = append(images, sectionImage)
		fs.Logger.Printf("Image processed: %s (%d bytes → %d base64 chars)", mimeType, len(fileData), len(base64Data))
	}
	return images, nil
}
