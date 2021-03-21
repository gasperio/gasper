package gasper

import (
	"archive/tar"
	"bytes"
	"fmt"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/empty"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/remote/transport"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Push(img, src string, opt ...Option) error {
	o := makeOptions(opt...)

	ref, err := name.ParseReference(img)
	if err != nil {
		return err
	}

	image, err := mutate.CreatedAt(empty.Image, v1.Time{Time: time.Now()})
	if err != nil {
		return err
	}

	layer, err := layerFromDir(src)
	if err != nil {
		return err
	}

	ri, err := remote.Image(ref, o.remote...)
	if err == nil {
		man, err := ri.Manifest()
		if err != nil {
			return err
		}

		h, err := layer.Digest()
		if err != nil {
			return err
		}

		for _, l := range man.Layers {
			if l.Digest.Algorithm == h.Algorithm && l.Digest.Hex == h.Hex {
				return nil
			}
		}
	} else if terr, ok := err.(*transport.Error); ok && terr.StatusCode == http.StatusNotFound {
		// That's ok
	} else {
		return err
	}

	newImage, err := mutate.AppendLayers(image, layer)
	if err != nil {
		return err
	}

	if err := remote.Write(ref, newImage, o.remote...); err != nil {
		return err
	}

	ih, err := newImage.Digest()
	if err != nil {
		return err
	}

	fmt.Println(ih.Hex)

	return nil
}

func layerFromDir(root string) (v1.Layer, error) {
	var b bytes.Buffer

	tw := tar.NewWriter(&b)

	err := filepath.Walk(root, func(fp string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		rel, err := filepath.Rel(root, fp)
		if err != nil {
			return err
		}

		hdr := &tar.Header{
			//Name: path.Join(root, filepath.ToSlash(rel)),
			Name: filepath.ToSlash(rel),
			Mode: int64(info.Mode()),
		}

		if !info.IsDir() {
			hdr.Size = info.Size()
		}

		if info.Mode().IsDir() {
			hdr.Typeflag = tar.TypeDir
		} else if info.Mode().IsRegular() {
			hdr.Typeflag = tar.TypeReg
		} else {
			return err
		}

		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		if !info.IsDir() {
			f, err := os.Open(fp)
			if err != nil {
				return err
			}

			if _, err := io.Copy(tw, f); err != nil {
				return err
			}

			f.Close()
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if err := tw.Close(); err != nil {
		return nil, err
	}

	return tarball.LayerFromReader(&b)
}
