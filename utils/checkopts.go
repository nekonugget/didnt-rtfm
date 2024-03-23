package utils

import (
	"errors"
	"net/url"
)

func CheckUrl(urlInput string) (*url.URL, error) {
	parsed, err := url.Parse(urlInput)
	if err != nil {
		return nil, errors.New("url parse error")
	}
	return parsed, err
}
