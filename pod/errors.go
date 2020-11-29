package pod

import "errors"

var (
	ErrPodExists    = errors.New("podcast by that name already exists")
	ErrNoEntry      = errors.New("no podcast is managed by that name")
	ErrReservedName = errors.New("the name " + ReservedPodName + " is reserved by gopodgrab")
	ErrArchiveEmpty = errors.New("feed file zip archive empty")
)
