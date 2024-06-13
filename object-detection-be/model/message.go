package model


type Message struct {
	Message_id      string `json:"message_id" db:"message_id"`
	Timestamp       string `json:"timestamp"  db:"timestamp"`
	Location_id     string `json:"location_id" db:"location_id"`
	Camera_id       string `json:"camera_id" db:"camera_id"`
	Model_id        string `json:"model_id" db:"model_id"`
	NumberOfObjects int    `json:"number_of_objects" db:"number_of_objects"`
	NumberOfEvents  int    `json:"number_of_events" db:"number_of_events"`
	Status          *bool   `json:"status" db:"status"`
	Image_URL       string `json:"image_URL" db:"image_URL"`
	Video_URL       string `json:"video_URL" db:"video_URL"`
}
