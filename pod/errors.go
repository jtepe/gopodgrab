package pod

import "errors"

var (
	ErrPodExists = errors.New("podcast by that name already exists")
	ErrNoEntry   = errors.New("no podcast is managed by that name")
)
