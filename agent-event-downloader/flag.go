package main

import (
	"agent-event-downloader/util"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"time"
)

const (
	ApplicationIdFlag = "application"
	UsernameFlag      = "user"
	OtpFlag           = "otp"
	CompanyFlag       = "company"
	AgentFlag         = "agent"
	AfterFlag         = "after"
	BeforeFlag        = "before"
	OutputFlag        = "output"
	SilentFlag        = "silent"
)

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:  UsernameFlag,
		Usage: "user to use when logging into the IXON Cloud",
	},
	cli.StringFlag{
		Name:  OtpFlag,
		Usage: "otp token to use when logging into the IXON Cloud",
	},
	cli.StringFlag{
		Name:  ApplicationIdFlag,
		Usage: "application ID used when communicating with the IXON API",
	},
	cli.StringFlag{
		Name:  CompanyFlag,
		Usage: "company ID of company to get event data from",
	},
	cli.StringFlag{
		Name:  AgentFlag,
		Usage: "agent ID of agent to get event data from",
	},
	cli.StringFlag{
		Name:  AfterFlag,
		Usage: "starting date for event data (UNIX timestamp)",
	},
	cli.StringFlag{
		Name:  BeforeFlag,
		Usage: "ending date for event data (UNIX timestamp)",
	},
	cli.BoolFlag{
		Name:  SilentFlag,
		Usage: "do not print any log output",
	},
	cli.StringFlag{
		Name:  OutputFlag,
		Usage: "export agent event data to `FILE`",
	},
}

func verifyFlags(c *cli.Context) error {
	if !c.IsSet(ApplicationIdFlag) {
		return errors.New("please enter an application ID")
	}

	if !c.IsSet(UsernameFlag) {
		return errors.New("please enter a username")
	}

	if !c.IsSet(CompanyFlag) {
		return errors.New("please enter a company id")
	}

	if !c.IsSet(AgentFlag) {
		return errors.New("please enter an agent id")
	}

	if !c.IsSet(AfterFlag) {
		return errors.New("please enter a start date for event data")
	}

	err := validateDateFlag(c, AfterFlag)

	if err != nil {
		return err
	}

	if !c.IsSet(BeforeFlag) {
		return errors.New("please enter an end date for event data")
	}

	err = validateDateFlag(c, BeforeFlag)

	if err != nil {
		return err
	}

	if c.Bool(SilentFlag) {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	return nil
}

func validateDateFlag(c *cli.Context, flagName string) error {
	_, err := time.Parse(util.ISO8601, c.String(flagName))

	if err != nil {
		return errors.New(fmt.Sprintf("please enter a valid ISO 8601 timestamp for flag %s", flagName))
	}

	return nil
}
