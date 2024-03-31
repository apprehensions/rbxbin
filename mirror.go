package rbxbin

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/robloxapi/rbxdhist"
)

// Mirror represents a Roblox deployment mirror.
type Mirror string

// Job represents a deployment build.
type Job rbxdhist.Job

// DefaultMirror is the default deployment mirror that can be
// used in situations where mirror fallbacks are undesired.
const DefaultMirror Mirror = "https://setup.rbxcdn.com"

var (
	ErrNoMirrorFound = errors.New("no accessible deploy mirror found")

	// As of 2024-02-03:
	//   setup-cfly.rbxcdn.com = roblox-setup.cachefly.net
	//   setup.rbxcdn.com = setup-ns1.rbxcdn.com = setup-ak.rbxcdn.com
	//   setup-hw.rbxcdn.com = setup-ll.rbxcdn.com = does not exist
	Mirrors = []Mirror{
		DefaultMirror,
		Mirror("https://roblox-setup.cachefly.net"),
		Mirror("https://s3.amazonaws.com/setup.roblox.com"),
	}
)

// URL returns the mirror's URL with the given optional deployment channel.
func (m Mirror) URL(channel string) string {
	if channel == "" || channel == "LIVE" || channel == "live" {
		return string(m)
	}

	// Ensure that the channel is lowercased, since internally in
	// ClientSettings it will be lowercased, but not on the deploy mirror.
	channel = strings.ToLower(channel)

	return string(m) + "/channel/" + channel
}

// Package returns a URL to a package given a package name
// and a Deployment, relative to the mirror.
func (m Mirror) PackageURL(d Deployment, pkg string) string {
	return m.URL(d.Channel) + "/" + d.GUID + "-" + pkg
}

// Jobs returns the available deployment builds for the mirror.
func (m Mirror) GetJobs() ([]*Job, error) {
	hist, err := http.Get(m.URL("") + "/DeployHistory.txt")
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(hist.Body)
	hist.Body.Close()
	if err != nil {
		return nil, err
	}

	return ParseJobs(body), nil
}

// ParseJobs is a wrapper that returns a list of deployments, parsed
// from the stream of bytes.
//
// See [rbxdhist.Lex] for more information.
func ParseJobs(js []byte) (jobs []*Job) {
	stream := rbxdhist.Lex(js)
	for _, s := range stream {
		j, ok := s.(*rbxdhist.Job)
		if !ok || j == nil {
			continue
		}

		jobs = append(jobs, (*Job)(j))
	}

	return
}

// Mirror returns an available deployment mirror from [Mirrors].
//
// Deployment mirrors may go down, or be blocked by ISPs.
func GetMirror() (Mirror, error) {
	slog.Info("Finding an accessible deploy mirror")

	for _, m := range Mirrors {
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
