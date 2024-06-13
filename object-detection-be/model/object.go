package model

type Object struct{
	Message_id string `json:"message_id"`
	Object_id string `json:"object_id"`
	Object_type string `json:"object_type"`
	Gender string `json:"gender"`
	Age string `json:"age"`
	Vehicle_type string `json:"vehicle_type"`
	Vehicle_brand string ` json:"vehicle_brand"`
	Vehicle_color string `json:"vehicle_color"`
	Vehicle_licence string `json:"vehicle_license"`
	Bbox_topleftx int `json:"bbox_topleftx"`
	Bbox_toplefty int `json:"bbox_toplefty"`
	Bbox_bottomrightx int `json:"bbox_bottomrightx"`
	Bbox_bottomrighty int `json:"bbox_bottomrighty"`
}
