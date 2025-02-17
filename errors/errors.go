package errors

import "errors"

var (
	ErrApiUrlEmpty         = errors.New("apiURL cannot be empty")
	ErrApiKeyEmpty         = errors.New("apiKey cannot be empty")
	ErrPlayerIDEmpty       = errors.New("player ID cannot be empty")
	ErrPlayerNicknameEmpty = errors.New("nickname cannot be empty")
	ErrMatchIDEmpty        = errors.New("match ID cannot be empty")
	ErrRequestFailed       = errors.New("request failed")
	ErrBadRequest          = errors.New("request failed")
)
