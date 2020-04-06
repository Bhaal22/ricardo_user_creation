package events

import (
    "encoding/json"
)
type Event struct {
    Name string `json: name`
    Status string `json: status`
    Reason string `json: reason`
}

func (e *Event) ToJson() string {
    eventJson, _ := json.Marshal(e)
    return string(eventJson)
}