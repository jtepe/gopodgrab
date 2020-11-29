package pod

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// addPod adds the podcast to the configuration file.
// If  creating/writing of the file fails, an error is returned.
func addPod(pod *Podcast) error {
	pods, err := readPods()
	if err != nil {
		return err
	}

	pods[pod.Name] = pod

	buf, err := json.MarshalIndent(&pods, "", "  ")
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
		fmt.Fprintf(os.Stderr, "failed to read configuration file at %s: %v\n", confFile(), err)
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
// file. It uses the user's default config directory as base the exact
// location of which is OS dependent. Otherwise, the current working
// directory is used. '/.gopodgrab' is appended to the base in any case.
func confFile() string {
	cf, err := os.UserConfigDir()
	if err != nil {
		cf = "."
	}

	cf = cf + "/gopodgrab/gopodgrab.json"
	return cf
}

// readPods retrieves all podcasts from the configuration file
// indexed by the name of the podcast. If the file doesn't, it
// is created and an empty map of podcasts is returned.
// Errors reading the file are passed back to the caller.
func readPods() (map[string]*Podcast, error) {
	pods := make(map[string]*Podcast)

	cf := confFile()

	f, err := os.Open(cf)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = createConfFile(cf)
			return pods, err
		}

		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// The config file is still empty.
	if len(buf) == 0 {
		return pods, nil
	}

	if err := json.Unmarshal(buf, &pods); err != nil {
		return nil, err
	}

	return pods, nil
}

// createConfFile creates an empty config file at location cf.
// It also attempts to create all missing directories in the path.
func createConfFile(cf string) error {
	p := path.Dir(cf)
	if err := os.MkdirAll(p, 0755); err != nil {
		return err
	}

	f, err := os.Create(cf)
	if err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
