package pod

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

// DownloadEpisode downloads the podcast episode e.
// The download size of the episode in bytes is recorded
// in Episode.Bytes.
func DownloadEpisode(e *Episode) error {
	u, err := url.Parse(e.File.URL)
	if err != nil {
		return err
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(path.Base(u.Path))
	if err != nil {
		return err
	}
	defer f.Close()

	n, err := io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	e.Bytes = n

	return nil
}
