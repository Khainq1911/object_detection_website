package model

type Event struct{
	Message_id string `json:"message_id"`
	Object_id string `json:"object_id"`
	Event_type string `json:"event_type"`
	Action string `json:"action"`
}
