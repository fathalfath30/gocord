/*
//
//  ______    _   _           _  __      _   _     ____   ___
// |  ____|  | | | |         | |/ _|    | | | |   |___ \ / _ \
// | |__ __ _| |_| |__   __ _| | |_ __ _| |_| |__   __) | | | |
// |  __/ _` | __| '_ \ / _` | |  _/ _` | __| '_ \ |__ <| | | |
// | | | (_| | |_| | | | (_| | | || (_| | |_| | | |___) | |_| |
// |_|  \__,_|\__|_| |_|\__,_|_|_| \__,_|\__|_| |_|____/ \___/
//
// Written by Fathalfath30.
// Email : fathalfath30@gmail.com
// Follow me on:
//  Github : https://github.com/fathalfath30
//  Gitlab : https://gitlab.com/Fathalfath30
//
*/

package gocord

import (
	"net/http"

	"github.com/fathalfath30/gocord/gocord/channel"
	"github.com/fathalfath30/gocord/gocord/guild"
)

//go:generate mockery --name IDiscordClient --filename discord_client.mock.go --structname DiscordClientMock
//go:generate mockery --name IGoCord --filename gocord.mock.go --structname GoCordMock
type (
	IDiscordClient interface {
		Do(req *http.Request) (*http.Response, error)
	}

	IGoCord interface {
		Guild() guild.IGuild
		Channel() channel.IChannel
	}

	Config struct {
		Proxy string
	}

	GoCord struct {
		guild   guild.IGuild
		channel channel.IChannel
	}

	Constructor struct {
		Guild   guild.IGuild
		Channel channel.IChannel
	}
)

func New(constructor *Constructor) (IGoCord, error) {
	var gc *GoCord
	if constructor != nil {
		gc = &GoCord{
			guild:   constructor.Guild,
			channel: constructor.Channel,
		}
	}

	return gc, nil
}

func (gc *GoCord) Guild() guild.IGuild {
	return nil
}

func (gc *GoCord) Channel() channel.IChannel {
	return nil
}
