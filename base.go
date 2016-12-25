package sdetools

type SDE struct {
	BaseDir string

	loadedNames     bool
	systemNamesById map[int]string
}
