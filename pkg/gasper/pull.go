package gasper

import (
	"archive/tar"
	"fmt"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/remote/transport"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Pull(img, dst string, opt ...Option) error {
	o := makeOptions(opt...)

	ref, err := name.ParseReference(img)
	if err != nil {
		return err
	}

	image, err := remote.Image(ref, o.remote...)
	if err != nil {
		if terr, ok := err.(*transport.Error); ok && terr.StatusCode == http.StatusNotFound {
			fmt.Printf("Image not found: %s", img)
			return nil
		} else {
			return err
		}
	}

	layers, err := image.Layers()
	if err != nil {
		return err
	}

	for _, l := range layers {
		r, err := l.Uncompressed()
		if err != nil {
			return err
		}

		if err := untar(dst, r); err != nil {
			return err
		}
	}

	return nil
}

func untar(dst string, r io.Reader) error {
	tr := tar.NewReader(r)

	for {
		header, err := tr.Next()

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}

		target := filepath.Join(dst, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			f.Close()
		}
	}
}
