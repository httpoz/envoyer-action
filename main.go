package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/actions-go/toolkit/core"
)

var (
	exit = os.Exit
	now  = func() time.Time {
		return time.Now()
	}
)

const (
	envoyerURL = "https://envoyer.io/api/projects/%s/deployments"
)

func runMain() {
	sleep, ok := core.GetInput("myInput")
	if !ok {
		core.Error("Unable to find required input myInput")
		exit(1)
	}
	core.Debug(fmt.Sprintf("Waiting %s milliseconds", sleep))
	core.Debug(now().String())
	delay, err := strconv.Atoi(sleep)
	if err != nil {
		core.Error(err.Error())
		exit(1)
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
	core.Debug(now().String())
	core.SetOutput("time", now().String())
	core.Info("All good!")
}

func main() {
	var credentials Credentials
	credentials.collectInputs()

	runMain()
}

type Credentials struct {
	BranchName string
	ProjectID  string
	Token      string
	Timeout    time.Duration
}

func (c *Credentials) collectInputs() {
	projectID, pok := core.GetInput("projectId")
	if !pok {
		core.Error("Unable to find required input projectId")
		exit(1)
	}
	c.ProjectID = projectID

	token, tok := core.GetInput("token")
	if !tok {
		core.Error("Unable to find required input token")
		exit(1)
	}
	c.Token = token

	branchName, bok := core.GetInput("branchName")
	if !bok {
		core.Error("Unable to find required input branchName")
		exit(1)
	}
	c.BranchName = branchName
}

// create a new HTTP client with a user defined timeout
func (c *Credentials) newHTTPClient() *http.Client {
	var timeout time.Duration
	if c.Timeout != 0 {
		timeout = c.Timeout
	} else {
		timeout = 30 * time.Second
	}

	return &http.Client{
		Timeout: timeout,
	}
}
