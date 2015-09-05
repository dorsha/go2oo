package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dorsha/go2oo/actions"
	"github.com/dorsha/go2oo/restUtil"
)

const (
	getConfigItems  = "get-config-items"
	getContentPacks = "show-content-packs"
	trigger         = "trigger"
)

const (
	restURI = "/rest/latest/"
)

var (
	url      = flag.String("url", "http://localhost:8080/oo", "The URL of Central (i.e. http://localhost:8080/oo)")
	user     = flag.String("user", "", "User name for Central")
	password = flag.String("password", "", "Password for Central")
	action   = flag.String("action", "", "What do you want to do? Avialble actions: "+getConfigItems+","+getContentPacks+","+trigger)
	uuid     = flag.String("uuid", "", "Flow uuid")
)

func main() {
	flag.Parse()

	if len(*action) == 0 {
		fmt.Println("Action must be specified")
	}

	client := restUtil.CreateHTTPClient()
	restURL := *url + restURI

	switch {
	case *action == getConfigItems:
		fmt.Println("Getting configuration items from: " + *url)
		ci := &actions.ConfigItem{}
		restUtil.Get(*client, restURL+"/config", &ci.Props, *user, *password)
		ci.HandleResponse()
	case *action == getContentPacks:
		fmt.Println("Getting content packs items from: " + *url)
		cps := &actions.ContentPacks{}
		restUtil.Get(*client, restURL+"/content-packs", cps, *user, *password)
		cps.HandleResponse()
	case *action == trigger:
		if len(*uuid) == 0 {
			fmt.Println("UUID must be specidied (--uuid)")
			os.Exit(1)
		}
		fmt.Println("Triggering flow: " + *uuid)
		trigger := &actions.Trigger{FlowUUID: *uuid}
		triggerResp := new(actions.TriggerResponse)
		restUtil.Post(*client, restURL+"executions", trigger, triggerResp, *user, *password)
		triggerResp.HandleResponse()
	}
}
