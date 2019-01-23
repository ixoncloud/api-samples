package main

import (
	"agent-event-downloader/ixon"
	"agent-event-downloader/util"
	"encoding/csv"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"time"
)

const (
	PasswordEnv = "IXON_IXPLATFORM_PASSWORD"
)

func setupApp() *cli.App {
	app := cli.NewApp()
	app.Version = "0.1.0"
	app.Name = "agent-event-dl"
	app.Compiled = time.Now()
	app.Author = "IXON"
	app.Usage = "Export agent event data from the IXON Cloud"

	app.Flags = Flags

	app.Action = onAction

	return app
}

func onAction(c *cli.Context) error {
	err := verifyFlags(c)

	if err != nil {
		return err
	}

	password := os.Getenv(PasswordEnv)

	if password == "" {
		return cli.NewExitError(fmt.Sprintf("Please specify your password in the %s environment variable", PasswordEnv), 1)
	}

	applicationId := c.String(ApplicationIdFlag)
	username := c.String(UsernameFlag)
	otp := c.String(OtpFlag)
	companyId := c.String(CompanyFlag)
	agentId := c.String(AgentFlag)

	afterDateString := c.String(AfterFlag)
	afterDate, _ := time.Parse(util.ISO8601, afterDateString)
	beforeDateString := c.String(BeforeFlag)
	beforeDate, _ := time.Parse(util.ISO8601, beforeDateString)

	outputFile := os.Stdout
	if c.IsSet(OutputFlag) {
		outputFile, err = openOrCreateFile(c.String(OutputFlag))

		if err != nil {
			return cli.NewExitError("Could not open or create output file", 1)
		}
	}

	client := ixon.NewClient(applicationId, companyId)

	log.Info("Retrieving IXapi discovery")

	err = client.Discovery.DiscoverApiEndpoints()

	if err != nil {
		return err
	}

	log.Info("Logging in")

	err = client.Auth.Login(username, otp, password)

	if err != nil {
		return cli.NewExitError("Could not login to IXON Cloud", 1)
	}

	log.Info("Retrieving agent event data")

	// Get event data
	eventData, err := client.Agent.GetEventList(agentId, afterDate, beforeDate)

	if err != nil {
		return cli.NewExitError("Could not fetch agent event data", 1)

	}

	if len(eventData) == 0 {
		return cli.NewExitError("No agent events found", 1)
	}

	csvWriter := csv.NewWriter(outputFile)

	headers := make([]string, 0)

	for k := range eventData[0] {
		headers = append(headers, k)
	}

	sortedHeaders, err := writeHeaders(csvWriter, headers)

	if err != nil {
		return cli.NewExitError("Could not write headers to file", 1)

	}

	err = writeData(csvWriter, sortedHeaders, eventData)

	if err != nil {
		return cli.NewExitError("Could not write data to file", 1)

	}

	log.Info("Done!")

	return nil
}

func main() {
	app := setupApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
