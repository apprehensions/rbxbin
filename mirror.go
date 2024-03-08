package rbxbin

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"
)

// Mirror represents an available Roblox deployment mirror.
type Mirror string

// PackageManifestSuffix can be added at the end of a mirror URL
// to retrieve the package manifest URL path.
const PackageManifestSuffix = "-rbxPkgManifest.txt"

var (
	ErrNoMirrorFound = errors.New("no accessible deploy mirror found")

	// As of 2024-02-03:
	//   setup-cfly.rbxcdn.com = roblox-setup.cachefly.net
	//   setup.rbxcdn.com = setup-ns1.rbxcdn.com = setup-ak.rbxcdn.com
	//   setup-hw.rbxcdn.com = setup-ll.rbxcdn.com = does not exist
	KnownMirrors = []Mirror{
		// Sorted by speed
		Mirror("https://setup.rbxcdn.com"),
		Mirror("https://setup-cfly.rbxcdn.com"),
		Mirror("https://s3.amazonaws.com/setup.roblox.com"),
	}
)

// URL returns the mirror's URL with the given optional channel.
func (m Mirror) URL(channel string) string {
	if channel == "" {
		return string(m)
	}

	// Ensure that the channel is lowercased, since internally in
	// ClientSettings it will be lowercased, but not on the deploy mirror.
	channel = strings.ToLower(channel)

	return string(m) + "/channel/" + channel
}

// Mirror returns an available Mirror from [Mirrors].
func GetMirror() (Mirror, error) {
	slog.Info("Finding an accessible deploy mirror")

	for _, m := range KnownMirrors {
		resp, err := http.Head(m.URL("") + "/" + "version")
		if err != nil {
			slog.Error("Bad deploy mirror", "mirror", m, "error", err)

			continue
		}
		resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			slog.Info("Found deploy mirror", "mirror", m)

			return m, nil
		}
	}

	return "", ErrNoMirrorFound
}
