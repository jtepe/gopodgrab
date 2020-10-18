package pod

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// addPod adds the podcast to the configuration file.
// If  creating/writing of the file fails, an error is returned.
func addPod(pod *Podcast) error {
	pods, err := readPods()
	if err != nil {
		return err
	}

	pods = append(pods, pod)

	buf, err := json.Marshal(&pods)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(confFile(), buf, 0644); err != nil {
		return err
	}

	return nil
}

// podExists checks whether a podcast by that name is
// already present in the configuration file.
func podExists(name string) bool {
	pods, err := readPods()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read configuration file at %s: %v", confFile(), err)
		os.Exit(1)
	}

	for _, pod := range pods {
		if pod.Name == name {
			return true
		}
	}

	return false
}

// confFile returns the storage location of gopodcrabs configuration
// file. It uses the user's home directory as base if the HOME
// environment variable is set. Otherwise, the current working
// directory is used. '/.gopodgrab' is appended to the base in any case.
func confFile() string {
	h := os.Getenv("HOME")
	if h == "" {
		h = "."
	}

	return h
}

// readPods retrieves the list of podcasts from the configuration
// file, if it exists. If it doesn't an empty list is returned.
// Errors reading the file are passed back to the caller.
func readPods() ([]*Podcast, error) {
	var pods []*Podcast

	f, err := os.Open(confFile())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return pods, nil
		}

		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(buf, &pods); err != nil {
		return nil, err
	}

	return pods, nil
}
