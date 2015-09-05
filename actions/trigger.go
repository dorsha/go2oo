package actions

import "fmt"

// Trigger represents trigger in OO
type Trigger struct {
	FlowUUID string `json:"flowUuid"`
}

// TriggerResponse represents trigger reponse in OO
type TriggerResponse int

// HandleResponse - prints trigger response (execution ID)
func (triggerResp TriggerResponse) HandleResponse() {
	fmt.Printf("Flow Triggered.\nExecution ID: %v\n", triggerResp)
}
