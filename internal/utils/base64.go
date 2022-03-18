package utils

import (
	"encoding/base64"
	"errors"
	"net/http"
)

// SPECIAL THANKS:
// https://freshman.tech/snippets/go/image-to-base64/

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func GetURLEncodedBase64StringFromImageBytes(bytes []byte) (string, error) {
	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	case "image/x-icon":
		base64Encoding += "data:image/x-icon;base64,"
	default:
		return "", errors.New("Unsupported MIME type!" + mimeType)
	}

	// Append the base64 encoded output
	base64Encoding += toBase64(bytes)

	return base64Encoding, nil
}
