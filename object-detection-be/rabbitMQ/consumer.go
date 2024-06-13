package rabbitmq

/*
AMQP_URL_CLOUD=amqps://uxyfztvb:h1OhPqSTa18sjMH7Rbb7QWSPWzlf9-Ee@armadillo.rmq.cloudamqp.com/uxyfztvb
AMQP_URL_CLOUD_LAB=
AMQP_MESSAGE_QUEUE=q1
*/
import (
	"WebDetection/database"
	"WebDetection/model"
	"WebDetection/repository"
	"encoding/json"
	"fmt"
	"log"

	"github.com/labstack/echo"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumeRepo struct {
	Repo repository.Repo
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (u ConsumeRepo) Receiver(ctx echo.Context, sql *database.Sql) {
	conn, err := amqp.Dial("amqps://uxyfztvb:h1OhPqSTa18sjMH7Rbb7QWSPWzlf9-Ee@armadillo.rmq.cloudamqp.com/uxyfztvb")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for msg := range msgs {

			message := Message{}
			if err := json.Unmarshal([]byte(msg.Body), &message); err != nil {
				fmt.Printf("%s", err)
			}
			//fmt.Println(message)

			camResult := []model.Camera{}
			modelResult := []model.Model{}
			locResult := []model.Location{}

			sql.Db.Select(&camResult, `SELECT * FROM camera WHERE camera_id = $1`, message.Camera.ID)
			sql.Db.Select(&modelResult, `SELECT * FROM model WHERE model_id = $1`, message.Model.ID)
			sql.Db.Select(&locResult, `SELECT * FROM location WHERE location_id = $1`, message.Location.ID)

			//fmt.Println(camResult, modelResult, locResult)

			if len(camResult) == 0 {
				if _, err := sql.Db.Exec(`INSERT INTO camera (camera_id, type, description_cam) VALUES ($1, $2, $3)`,
					message.Camera.ID,
					message.Camera.Type,
					message.Camera.Description); err != nil {
					fmt.Println("loi", err)
				}
			}

			if len(modelResult) == 0 {
				sql.Db.Exec(`INSERT INTO model (model_id, description_mdl) VALUES ($1, $2)`,
					message.Model.ID,
					message.Model.Description)
			}

			if len(locResult) == 0 {
				sql.Db.Exec(`INSERT INTO location (location_id, lat, lon, alt) VALUES ($1,$2,$3,$4)`,
					message.Location.ID,
					message.Location.Lat,
					message.Location.Lon,
					message.Location.Alt)
			}

			//chen mess database
			MessElement := model.Message{
				Message_id:      message.MessageID,
				Timestamp:       message.Timestamp,
				Location_id:     message.Location.ID,
				Camera_id:       message.Camera.ID,
				Model_id:        message.Model.ID,
				NumberOfObjects: message.NumberOfObjects,
				NumberOfEvents:  message.NumberOfEvents,
				Image_URL:       message.ImageURL,
				Video_URL:       message.VideoURL,
			}
			sql.Db.Exec(`INSERT INTO message (message_id, timestamp, location_id, model_id, camera_id, number_of_objects, number_of_events, image_URL, video_URL)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			ON CONFLICT (message_id) DO UPDATE
			SET timestamp = EXCLUDED.timestamp,
				location_id = EXCLUDED.location_id,
				model_id = EXCLUDED.model_id,
				camera_id = EXCLUDED.camera_id,
				number_of_objects = EXCLUDED.number_of_objects,
				number_of_events = EXCLUDED.number_of_events, 
				image_URL = EXCLUDED.image_URL,
				video_URL = EXCLUDED.video_URL`,
				MessElement.Message_id, MessElement.Timestamp, MessElement.Location_id, MessElement.Model_id, MessElement.Camera_id, MessElement.NumberOfObjects, MessElement.NumberOfEvents, MessElement.Image_URL, MessElement.Video_URL)

			// chen object vao 1 arr roi chen tung phan tu vao bang
			ObjectSlice := []model.Object{}

			for _, object := range message.ObjectList {
				switch obj := object.(type) {
				case map[string]interface{}:
					if humanDetails, ok := obj["Human"]; ok {
						// Process human details (function can be added here)
						humanInfor := humanDetails.(map[string]interface{})
						//fmt.Println(humanInfor["age"], humanInfor["gender"], humanInfor["id"])

						humanLoc := humanInfor["bbox"].(map[string]interface{})
						//fmt.Println(humanLoc["bottomrightx"], humanLoc["bottomrighty"], humanLoc["topleftx"], humanLoc["toplefty"])

						objectElement := model.Object{
							Message_id:        message.MessageID,
							Object_type:       "Human",
							Bbox_topleftx:     int(humanLoc["topleftx"].(float64)),
							Bbox_toplefty:     int(humanLoc["toplefty"].(float64)),
							Bbox_bottomrightx: int(humanLoc["bottomrightx"].(float64)),
							Bbox_bottomrighty: int(humanLoc["bottomrighty"].(float64)),
						}
						if id, ok := humanInfor["id"].(string); ok {
							objectElement.Object_id = id
						}
						if gender, ok := humanInfor["gender"].(string); ok {
							objectElement.Gender = gender
						}
						if age, ok := humanInfor["age"].(string); ok {
							objectElement.Age = age
						}
						ObjectSlice = append(ObjectSlice, objectElement)

					}

					if vehicleDetails, ok := obj["Vehicle"]; ok {
						// Process vehicle details
						vehicleInfor := vehicleDetails.(map[string]interface{})
						//fmt.Println(vehicleInfor["Licence"], vehicleInfor["brand"], vehicleInfor["color"], vehicleInfor["id"], vehicleInfor["type"])
						vehicleLoc := vehicleInfor["bbox"].(map[string]interface{})
						//fmt.Println(reflect.TypeOf(vehicleLoc["bottomrightx"]), vehicleLoc["bottomrighty"], vehicleLoc["topleftx"], vehicleLoc["toplefty"])
						objectElement := model.Object{
							Object_type:       "Vehicle",
							Message_id:        message.MessageID,
							Bbox_topleftx:     int(vehicleLoc["topleftx"].(float64)),
							Bbox_toplefty:     int(vehicleLoc["toplefty"].(float64)),
							Bbox_bottomrightx: int(vehicleLoc["bottomrightx"].(float64)),
							Bbox_bottomrighty: int(vehicleLoc["bottomrighty"].(float64)),
						}
						if id, ok := vehicleInfor["id"].(string); ok {
							objectElement.Object_id = id
						}
						if vehicleType, ok := vehicleInfor["type"].(string); ok {
							objectElement.Vehicle_type = vehicleType
						}
						if brand, ok := vehicleInfor["brand"].(string); ok {
							objectElement.Vehicle_brand = brand
						}
						if color, ok := vehicleInfor["color"].(string); ok {
							objectElement.Vehicle_color = color
						}
						if licence, ok := vehicleInfor["licence"].(string); ok {
							objectElement.Vehicle_licence = licence
						}
						ObjectSlice = append(ObjectSlice, objectElement)
					}

				}
			}
			for _, objectElement := range ObjectSlice {
				sql.Db.Exec(`
				INSERT INTO object (message_id, object_id, object_type, gender, age, vehicle_type, vehicle_brand, vehicle_color, vehicle_licence, bbox_topleftx, bbox_toplefty, bbox_bottomrightx, bbox_bottomrighty)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
				ON CONFLICT (message_id, object_id) DO UPDATE
				SET object_type = EXCLUDED.object_type,
					gender = EXCLUDED.gender,
					age = EXCLUDED.age,
					vehicle_type = EXCLUDED.vehicle_type,
					vehicle_brand = EXCLUDED.vehicle_brand,
					vehicle_color = EXCLUDED.vehicle_color,
					vehicle_licence = EXCLUDED.vehicle_licence,
					bbox_topleftx = EXCLUDED.bbox_topleftx,
					bbox_toplefty = EXCLUDED.bbox_toplefty,
					bbox_bottomrightx = EXCLUDED.bbox_bottomrightx,
					bbox_bottomrighty = EXCLUDED.bbox_bottomrighty;
				`, objectElement.Message_id, objectElement.Object_id, objectElement.Object_type, objectElement.Gender, objectElement.Age, objectElement.Vehicle_type, objectElement.Vehicle_brand, objectElement.Vehicle_color, objectElement.Vehicle_licence, objectElement.Bbox_topleftx, objectElement.Bbox_bottomrighty, objectElement.Bbox_bottomrightx, objectElement.Bbox_bottomrighty)
			}
			// ket thuc insert object

			// insert event
			eventArr := []model.Event{}
			eventLists := message.EventList
			for _, eventList := range eventLists {
				switch event := eventList.(type) {
				case map[string]interface{}:
					if eventDetail, ok := event["human_event"].(map[string]interface{}); ok {
						eventElement := model.Event{
							Message_id: message.MessageID,
							Object_id:  eventDetail["object_id"].(string),
							Event_type: "human_event",
							Action:     eventDetail["action"].(string),
						}
						eventArr = append(eventArr, eventElement)
					}
					if eventDetail, ok := event["vehicle_event"].(map[string]interface{}); ok {
						eventElement := model.Event{
							Message_id: message.MessageID,
							Object_id:  eventDetail["object_id"].(string),
							Event_type: "vehicle_event",
							Action:     eventDetail["action"].(string),
						}
						eventArr = append(eventArr, eventElement)
					}
				}
			}
			for _, event := range eventArr {
				if _, err := sql.Db.Exec(`
				INSERT INTO event (message_id, object_id, event_type, action)
				VALUES ($1, $2, $3, $4)
				ON CONFLICT (object_id, message_id) DO UPDATE
				SET event_type = EXCLUDED.event_type,
					action = EXCLUDED.action;
			`, message.MessageID, event.Object_id, event.Event_type, event.Action); err != nil {
					fmt.Println("loi o chen event", err)
				}
			}

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
