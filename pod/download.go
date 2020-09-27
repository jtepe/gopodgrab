package pod

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func Download(e *Episode) error {
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
