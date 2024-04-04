package rbxbin

import (
	"errors"
	"log/slog"

	"github.com/apprehensions/rbxweb"
	"github.com/apprehensions/rbxweb/clientsettings"
)

// ErrBadChannel indicates if the mentioned deployment channel does not exist
// or out of permission scope for the current authenticated user.
var ErrBadChannel = errors.New("deployment channel is invalid or unauthorized")

// Deployment is a representation of a Binary's deployment or version.
//
// Channel can either be a given channel, or empty - in which Roblox
// will consider the 'default' channel.
//
// In all things related to the Roblox API, the default channel is empty,
// or 'live'/'LIVE' on clientsettings. On the Client/Studio, the default channel
// is (or can be) 'production'. This behavior is undocumented, so it is reccomended
// to simply provide an empty string for the channel.
type Deployment struct {
	Type    clientsettings.BinaryType
	Channel string
	GUID    string
}

// FetchDeployment returns the latest deployment information for the given
// Roblox binary type with the given deployment channel.
func GetDeployment(bt clientsettings.BinaryType, channel string) (Deployment, error) {
	slog.Info("Fetching Binary Deployment", "name", bt, "channel", channel)

	cv, err := clientsettings.GetClientVersion(bt, channel)
	if err != nil {
		var apiError rbxweb.ErrorResponse
		if errors.As(err, &apiError) {
			if apiError.Code == 5 {
				return Deployment{}, ErrBadChannel
			}
		}
		return Deployment{}, err
	}

	slog.Info("Fetched Binary Deployment!",
		"name", bt, "channel", channel, "guid", cv.GUID, "version", cv.Version)

	return Deployment{
		Type:    bt,
		Channel: channel,
		GUID:    cv.GUID,
	}, nil
}
