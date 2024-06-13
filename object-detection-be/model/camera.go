package model

type Camera struct {
	Camera_id       string `json:"camera_id" db:"camera_id"`
	Type            string `json:"type" db:"type"`
	Description_cam string `json:"description_cam" db:"description_cam"`
}

type Location struct {
	Location_id string  `json:"location_id" db:"location_id"`
	Lon         float32 `json:"lon" db:"lon"`
	Lat         float32 `json:"lat" db:"lat"`
	Alt         float32 `json:"alt" db:"alt"`
}

type Model struct {
	Model_id        string `json:"model_id" db:"model_id"`
	Description_mdl string `json:"description_mdl" db:"description_mdl"`
}
