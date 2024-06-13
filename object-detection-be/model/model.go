package model

import "time"

type TotalInfo struct {
	MessageID        string    `json:"message_id" db:"message_id"`
	Timestamp        time.Time `json:"timestamp"  db:"timestamp"`
	LocationID       string    `json:"location_id" db:"location_id"`
	Lon              float32   `json:"lon" db:"lon"`
	Lat              float32   `json:"lat" db:"lat"`
	Alt              float32   `json:"alt" db:"alt"`
	ModelID          string    `json:"model_id" db:"model_id"`
	DescriptionModel string    `json:"description_mdl" db:"description_mdl"`
	CameraID         string    `json:"camera_id" db:"camera_id"`
	Type             string    `json:"type" db:"type"`
	DescriptionCam   string    `json:"description_cam" db:"description_cam"`
	NumberOfObjects  int       `json:"number_of_objects" db:"number_of_objects"`
	NumberOfEvents   int       `json:"number_of_events" db:"number_of_events"`
	Status           bool      `json:"status" db:"status"`
	ImageURL         string    `json:"image_url" db:"image_url"`
	VideoURL         string    `json:"video_url" db:"video_url"`
	ObjectID         string    `json:"object_id"`
	ObjectType       string    `json:"object_type"`
	Gender           string    `json:"gender"`
	Age              string    `json:"age"`
	VehicleType      string    `json:"vehicle_type"`
	VehicleBrand     string    `json:"vehicle_brand"`
	VehicleColor     string    `json:"vehicle_color"`
	VehicleLicense   string    `json:"vehicle_license"`
	BboxTopleftX     int       `json:"bbox_topleftx"`
	BboxTopleftY     int       `json:"bbox_toplefty"`
	BboxBottomrightX int       `json:"bbox_bottomrightx"`
	BboxBottomrightY int       `json:"bbox_bottomrighty"`
	EventType        string    `json:"event_type"`
	Action           string    `json:"action"`
}

type Index struct {
	Message_id      string    `json:"message_id" db:"message_id"`
	Timestamp       time.Time `json:"timestamp"  db:"timestamp"`
	Location_id     string    `json:"location_id" db:"location_id"`
	Camera_id       string    `json:"camera_id" db:"camera_id"`
	Type            string    `json:"type" db:"type"`
	Description_cam string    `json:"description_cam" db:"description_cam"`
	Model_id        string    `json:"model_id" db:"model_id"`
	Description_mdl string    `json:"description_mdl" db:"description_mdl"`
	NumberOfObjects int       `json:"number_of_objects" db:"number_of_objects"`
	NumberOfEvents  int       `json:"number_of_events" db:"number_of_events"`
	Object_list     []Object  `json:"object_list"`
	Event_list      []Event   `json:"event_list"`
	Status          *bool     `json:"status" db:"status"`
	Image_URL       string    `json:"image_url" db:"image_url"`
	Video_URL       string    `json:"video_url" db:"video_url"`
}

type Filter struct {
	Message_id      string    `json:"message_id" db:"message_id"`
	Timestamp       time.Time `json:"timestamp"  db:"timestamp"`
	Location_id     string    `json:"location_id" db:"location_id"`
	Camera_id       string    `json:"camera_id" db:"camera_id"`
	Type            string    `json:"type" db:"type"`
	Description_cam string    `json:"description_cam" db:"description_cam"`
	Model_id        string    `json:"model_id" db:"model_id"`
	Description_mdl string    `json:"description_mdl" db:"description_mdl"`
	NumberOfObjects int       `json:"number_of_objects" db:"number_of_objects"`
	NumberOfEvents  int       `json:"number_of_events" db:"number_of_events"`
	Status          *bool     `json:"status" db:"status"`
	Image_URL       string    `json:"image_url" db:"image_url"`
	Video_URL       string    `json:"video_url" db:"video_url"`
}
type GetEvent struct {
	Message_id        string `json:"message_id"`
	Object_id         string `json:"object_id"`
	Event_type        string `json:"event_type"`
	Action            string `json:"action"`
	Object_type       string `json:"object_type"`
	Gender            string `json:"gender"`
	Age               string `json:"age"`
	Vehicle_type      string `json:"vehicle_type"`
	Vehicle_brand     string ` json:"vehicle_brand"`
	Vehicle_color     string `json:"vehicle_color"`
	Vehicle_license   string `json:"vehicle_license"`
	Bbox_topleftx     int    `json:"bbox_topleftx"`
	Bbox_toplefty     int    `json:"bbox_toplefty"`
	Bbox_bottomrightx int    `json:"bbox_bottomrightx"`
	Bbox_bottomrighty int    `json:"bbox_bottomrighty"`
}
