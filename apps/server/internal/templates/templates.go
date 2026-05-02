// internal/templates/templates.go
package templates

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed base-single.docx
var baseSingle []byte

//go:embed base-multi.docx
var baseMulti []byte

// Templates holds the embedded docx files and exposes typed access.
type Templates struct {
	single []byte
	multi  []byte
}

// New constructs a Templates, validating both embedded files are readable zips.
func New() (*Templates, error) {
	if err := validateZip(baseSingle, "base-single.docx"); err != nil {
		return nil, err
	}
	if err := validateZip(baseMulti, "base-multi.docx"); err != nil {
		return nil, err
	}
	return &Templates{single: baseSingle, multi: baseMulti}, nil
}

// BaseSingleReader returns a zip.Reader over base-single.docx.
func (t *Templates) BaseSingleReader() (*zip.Reader, error) {
	return zip.NewReader(bytes.NewReader(t.single), int64(len(t.single)))
}

// BaseMultiReader returns a zip.Reader over base-multi.docx.
func (t *Templates) BaseMultiReader() (*zip.Reader, error) {
	return zip.NewReader(bytes.NewReader(t.multi), int64(len(t.multi)))
}

// BaseSingleBytes returns the raw bytes of base-single.docx.
func (t *Templates) BaseSingleBytes() []byte { return t.single }

// BaseMultiBytes returns the raw bytes of base-multi.docx.
func (t *Templates) BaseMultiBytes() []byte { return t.multi }

func validateZip(data []byte, name string) error {
	_, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return fmt.Errorf("templates: %s is not a valid zip/docx: %w", name, err)
	}
	return nil
}
