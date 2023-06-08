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

package guild

import (
	"github.com/fathalfath30/gocord/gocord/client/discord"
)

//go:generate mockery --name IGuild --filename guild.mock.go --structname GuildMock
type (
	IGuild interface {
		Create(guild *Guild) (*Guild, error)
		Get(id string) (*Guild, error)
		GetPreview(id string) (*Guild, error)
	}

	Guilds struct {
		client discord.IClient
	}

	Constructor struct {
		Client discord.IClient
	}
)

func New(constructor *Constructor) (IGuild, error) {
	g := new(Guild)
	return g, nil
}
