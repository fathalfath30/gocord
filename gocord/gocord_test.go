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
package gocord_test

import (
	"context"
	"github.com/fathalfath30/gocord/gocord"
	channelMock "github.com/fathalfath30/gocord/gocord/channel/mocks"
	guildMock "github.com/fathalfath30/gocord/gocord/guild/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

// GoCordTestSuite is used to simplify generating test case
type (
	GoCordTestSuite struct {
		suite.Suite
		t   *testing.T
		ctx context.Context
	}
)

// Test_RunGoCordTestSuite Running the test suite
func Test_RunGoCordTestSuite(t *testing.T) {
	suite.Run(t, &GoCordTestSuite{t: t})
}

// SetupSuite this function executes before the test suite begins
// execution
func (ts *GoCordTestSuite) SetupSuite() {
	// set StartingNumber to one
}

// TearDownSuite which will run after all the tests in the suite
// have been run.
func (ts *GoCordTestSuite) TearDownSuite() {
	// todo: TearDownSuite
}

// SetupTest SetupTestSuite has a SetupTest method, which will run
// before each test in the suite.
func (ts *GoCordTestSuite) SetupTest() {
	ts.ctx = context.Background()
}

// TearDownTest TearDownTestSuite has a TearDownTest method, which
// will run after each test in the suite.
func (ts *GoCordTestSuite) TearDownTest() {
	// todo: TearDownTest
}

func (ts *GoCordTestSuite) Test_GoCord_CanCreateWithMockedConstructor() {
	gc, err := gocord.New(&gocord.Constructor{
		Guild:   guildMock.NewGuildMock(ts.t),
		Channel: channelMock.NewChannelMock(ts.t),
	})

	ts.Require().NotNil(gc)
	ts.Require().Nil(err)

	ts.Require().IsType(&guildMock.GuildMock{}, gc.Guild())
	ts.Require().IsType(&channelMock.ChannelMock{}, gc.Channel())
}
