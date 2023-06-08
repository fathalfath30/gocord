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

func New(constructor *Constructor) (IGoCord, error) {
	var gc *GoCord
	if constructor != nil {
		gc = &GoCord{
			guild:   constructor.Guild,
			channel: constructor.Channel,
		}

		return gc, nil
	}
	
	gc = new(GoCord)
	var err error
	gc.guild, err = guild.New(nil)
	if err != nil {
		return nil, err
	}

	gc.channel, err = channel.New()
	if err != nil {
		return nil, err
	}

	return gc, nil
}

func (gc *GoCord) Guild() guild.IGuild {
	return gc.guild
}

func (gc *GoCord) Channel() channel.IChannel {
	return gc.channel
}
