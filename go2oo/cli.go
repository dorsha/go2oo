package main

import (
	"flag"
	"fmt"

	"github.com/dorsha/go2oo/actions"
	"github.com/dorsha/go2oo/restUtil"
)

const (
	getConfigItems  = "get-config-items"
	getContentPacks = "show-content-packs"
)

var (
	url      = flag.String("url", "http://localhost:8080/oo", "The URL of Central (i.e. http://localhost:8080/oo)")
	user     = flag.String("user", "", "User name for Central")
	password = flag.String("password", "", "Password for Central")
	action   = flag.String("action", "", "What do you want to do? Avialble actions: "+getConfigItems+","+getContentPacks)
)

func main() {
	flag.Parse()

	if len(*action) == 0 {
		fmt.Println("Action must be specified")
	}

	client := restUtil.CreateHTTPClient()
	restURL := *url + "/rest/"

	switch {
	case *action == getConfigItems:
		fmt.Println("Getting configuration items from: " + *url)
		ci := &actions.ConfigItem{}
		restUtil.Get(*client, restURL+"/config", &ci.Props)
		ci.HandleResponse()
	case *action == getContentPacks:
		fmt.Println("Getting content packs items from: " + *url)
		cps := &actions.ContentPacks{}
		restUtil.Get(*client, restURL+"/content-packs", cps)
		cps.HandleResponse()
	}
}
