package rbxbin

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/robloxapi/rbxdhist"
)

// Mirror represents an available Roblox deployment mirror.
type Mirror string

// Job represents a deployment build.
type Job rbxdhist.Job

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

// Package returns a URL to a package given a package name
// and a Deployment, relative to the mirror.
func (m Mirror) Package(d *Deployment, pkg string) string {
	return m.URL(d.Channel) + "/" + d.GUID + "-" + pkg
}

// PackageManifest returns a URL of the package manifest of
// the given Deployment.
func (m Mirror) PackageManifest(d *Deployment) string {
	return m.URL(d.Channel) + "/" + d.GUID + "-rbxPkgManifest.txt"
}

// Jobs fetches the available deployment builds for the mirror.
func (m Mirror) Jobs() ([]*Job, error) {
	var jobs []*Job

	hist, err := http.Get(m.URL("") + "/DeployHistory.txt")
	if err != nil {
		return jobs, err
	}

	body, err := io.ReadAll(hist.Body)
	hist.Body.Close()
	if err != nil {
		return jobs, err
	}

	stream := rbxdhist.Lex(body)
	for _, s := range stream {
		j, ok := s.(*rbxdhist.Job)
		if !ok || j == nil {
			continue
		}

		jobs = append(jobs, (*Job)(j))
	}

	return jobs, nil
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
