package actions

import "fmt"

// Trigger represents trigger in OO
type Trigger struct {
	FlowUUID string `json:"flowUuid"`
}

// TriggerResponse represents trigger reponse in OO
type TriggerResponse int

// HandleResponse - print content packs
func (triggerResp TriggerResponse) HandleResponse() {
	fmt.Printf("Flow Triggered.\nExecution ID: %v\n", triggerResp)
}
