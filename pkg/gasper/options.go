package gasper

import (
	"github.com/google/go-containerregistry/pkg/authn"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"net/http"
)

// Option is a functional option for gasper
type Option func(*options)

type options struct {
	remote   []remote.Option
	platform *v1.Platform
}

func makeOptions(opts ...Option) options {
	opt := options{
		remote: []remote.Option{
			remote.WithAuthFromKeychain(authn.DefaultKeychain),
		},
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

func WithTransport(t http.RoundTripper) Option {
	return func(o *options) {
		o.remote = append(o.remote, remote.WithTransport(t))
	}
}

func WithPlatform(platform *v1.Platform) Option {
	return func(o *options) {
		if platform != nil {
			o.remote = append(o.remote, remote.WithPlatform(*platform))
		}
		o.platform = platform
	}
}

func WithAuthFromKeychain(keys authn.Keychain) Option {
	return func(o *options) {
		o.remote[0] = remote.WithAuthFromKeychain(keys)
	}
}

func WithAuth(auth authn.Authenticator) Option {
	return func(o *options) {
		o.remote[0] = remote.WithAuth(auth)
	}
}

func WithUserAgent(ua string) Option {
	return func(o *options) {
		o.remote = append(o.remote, remote.WithUserAgent(ua))
	}
}
