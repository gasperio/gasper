package gasper

import (
	"errors"
	"fmt"
	"github.com/docker/cli/cli/config"
	"github.com/docker/cli/cli/config/types"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"io/ioutil"
	"os"
	"strings"
)

func Login(serverAddress, username, password string, passwordStdin bool) error {
	if passwordStdin {
		contents, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		password = strings.TrimSuffix(string(contents), "\n")
		password = strings.TrimSuffix(password, "\r")
	}

	if username == "" || password == "" {
		return errors.New("Username and password required")
	}

	cf, err := config.Load(os.Getenv("DOCKER_CONFIG"))
	if err != nil {
		return err
	}

	creds := cf.GetCredentialsStore(serverAddress)
	if serverAddress == name.DefaultRegistry {
		serverAddress = authn.DefaultAuthKey
	}

	if err := creds.Store(types.AuthConfig{
		ServerAddress: serverAddress,
		Username:      username,
		Password:      password,
	}); err != nil {
		return err
	}

	fmt.Println("Login succeeded")

	return nil
}
