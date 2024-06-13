package rabbitmq
type Location struct {
    ID  string  `json:"id"`
    Lat float64 `json:"lat"`
    Lon float64 `json:"lon"`
    Alt float64 `json:"alt"`
}

type Model struct {
    ID          string `json:"id"`
    Description string `json:"description"`
}

type Camera struct {
    ID          string `json:"id"`
    Type        string `json:"type"`
    Description string `json:"description"`
}

type BoundingBox struct {
    TopLeftX     int `json:"topleftx"`
    TopLeftY     int `json:"toplefty"`
    BottomRightX int `json:"bottomrightx"`
    BottomRightY int `json:"bottomrighty"`
}

type Object struct {
    ID     string      `json:"id"`
    Gender string      `json:"gender,omitempty"`
    Age    string      `json:"age,omitempty"`
    Type   string      `json:"type,omitempty"`
    Brand  string      `json:"brand,omitempty"`
    Color  string      `json:"color,omitempty"`
    Licence string     `json:"Licence,omitempty"`
    BBox   BoundingBox `json:"bbox"`
}


type HumanEvent struct {
    ObjectID string `json:"object_id"`
    Action   string `json:"action"`
}

type VehicleEvent struct {
    ObjectID string `json:"object_id"`
    Action   string `json:"action"`
}

type Message struct {
    MessageID       string        `json:"message_id"`
    Timestamp       string        `json:"timestamp"`
    Location        Location      `json:"location"`
    Model           Model         `json:"model"`
    Camera          Camera        `json:"camera"`
    NumberOfObjects int           `json:"number_of_objects"`
    ObjectList      []interface{}      `json:"object_list"`
    NumberOfEvents  int           `json:"number_of_events"`
    EventList       []interface{} `json:"event_list"`
    ImageURL        string        `json:"image_URL"`
    VideoURL        string        `json:"video_URL"`
}
