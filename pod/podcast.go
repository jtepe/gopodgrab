package pod

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
)

// Podcast represents a podcast. It has a feed URL, name
// and additional metadata.
type Podcast struct {
	FeedURL    string `json:"feed_url"`    // URL to retrieve the podcast feed from
	Name       string `json:"name"`        // The name under which this podcast is managed
	LocalStore string `json:"local_store"` // Directory path of the local store for this podcast
}

// NewPodcast creates a new podcast and intializes the
// local storage for it. If creation of the local storage
// fails, or a podcast by that name is already managed by
// gopodgrab, an error is returned.
// If the refresh of the feed, or adding the configuration
// of the podcast fails, an error is returned, as well.
func NewPodcast(name, feedURL, storageDir string) (*Podcast, error) {
	if podExists(name) {
		return nil, ErrPodExists
	}

	pod := &Podcast{
		Name:       name,
		FeedURL:    feedURL,
		LocalStore: storageDir,
	}

	if err := pod.RefreshFeed(); err != nil {
		return nil, err
	}

	if err := addPod(pod); err != nil {
		return nil, err
	}

	return pod, nil
}

// List returns the list of managed podcasts from
// the configuration file.
// Failure to read the configuration file results in a error.
func List() ([]*Podcast, error) {
	pods, err := readPods()
	if err != nil {
		return nil, err
	}

	res := make([]*Podcast, 0, len(pods))
	for _, p := range pods {
		res = append(res, p)
	}

	return res, nil
}

// RefreshFeed updates the locally stored feed from remote.
func (pod *Podcast) RefreshFeed() error {
	resp, err := http.Get(pod.FeedURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = pod.storeExists()
	if err != nil {
		return err
	}

	f, err := os.Create(pod.feedFile())
	if err != nil {
		return err
	}
	defer f.Close()

	zipper := zip.NewWriter(f)

	file, err := zipper.Create(pod.Name)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	err = zipper.Close()
	if err != nil {
		return err
	}

	return nil
}

// storeExists ensures that the podcast storage directory is present.
func (pod *Podcast) storeExists() error {
	if err := os.MkdirAll(pod.LocalStore, os.ModeDir|0755); err != nil {
		return err
	}

	return nil
}

// feedFile returns the full file path of the locally stored, zipped feed.
func (pod *Podcast) feedFile() string {
	return pod.LocalStore + "/feed.zip"
}

// dirExists checks whether the directory specified by path exists.
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	return info.IsDir()
}
