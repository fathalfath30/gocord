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
	"github.com/fathalfath30/gocord/gocord/channel"
	"github.com/fathalfath30/gocord/gocord/guild"
)

//go:generate mockery --name IGoCord --filename gocord.mock.go --structname GoCordMock
type (
	IGoCord interface {
		// Guild will return guild.IGuild to manage discord server
		Guild() guild.IGuild

		// Channel will return channel.IChannel to manage channel
		Channel() channel.IChannel
	}
)
