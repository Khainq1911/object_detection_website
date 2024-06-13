package repoimpl

import (
	"WebDetection/database"
	"WebDetection/model"
	"WebDetection/repository"
	"context"
	"fmt"
	"strconv"
	"time"
)

type RepoImpl struct {
	sql *database.Sql
}

func NewDataRepo(sql *database.Sql) repository.Repo {
	return (&RepoImpl{
		sql: sql,
	})
}

func (db *RepoImpl) IndexRepo(ctx context.Context) ([]model.Index, error) {
	index := []model.Index{}

	query :=
		`SELECT DISTINCT * FROM message                                
	INNER JOIN model ON message.model_id = model.model_id
	INNER JOIN camera ON message.camera_id = camera.camera_id
	`

	if err := db.sql.Db.Select(&index, query); err != nil {
		fmt.Println("loi o indexRepo", err)
		return nil, err
	}

	for i := range index {
		object := []model.Object{}
		event := []model.Event{}
		query_object := `select * from object where message_id = $1`
		if err := db.sql.Db.Select(&object, query_object, index[i].Message_id); err != nil {
			fmt.Println("loi o indexRepo", err)
			return nil, err
		}
		query_event := `select * from event where message_id = $1`
		if err := db.sql.Db.Select(&event, query_event, index[i].Message_id); err != nil {
			fmt.Println("loi o indexRepo", err)
			return nil, err
		}

		index[i].Object_list = object
		index[i].Event_list = event
	}
	return index, nil
}

func (db *RepoImpl) FilterRepo(ctx context.Context, objectType string, timeFrom string, timeTo string, cameraID string, status string) ([]model.Index, error) {
	data := []model.Index{}

	/* object := []model.Object{} */
	queryParams := []interface{}{}

	// Build base query with JOINs
	query := `SELECT DISTINCT * FROM message
		INNER JOIN model ON message.model_id = model.model_id
		INNER JOIN camera ON message.camera_id = camera.camera_id`

	// Dynamic filtering conditions with placeholders
	if objectType != "" && objectType != "all" {
		queryParams = append(queryParams, objectType)
		query += ` WHERE EXISTS ( SELECT * FROM event WHERE event.message_id = message.message_id AND event_type = $` + strconv.Itoa(len(queryParams)) + `)`
	}

	if timeFrom != "" && timeTo != "" && timeFrom != "all" && timeTo != "all" {
		// Enhanced error handling for time parsing
		tFrom, err := time.Parse(time.RFC3339, timeFrom)
		if err != nil {
			fmt.Println("error parsing timeFrom:", err)
			return nil, err
		}
		queryParams = append(queryParams, tFrom)
		query += ` AND message.timestamp >= $` + strconv.Itoa(len(queryParams))
		tTo, err := time.Parse(time.RFC3339, timeTo)
		if err != nil {
			fmt.Println("error parsing timeTo:", err)
			return nil, err
		}
		queryParams = append(queryParams, tTo)
		query += ` AND message.timestamp <= $` + strconv.Itoa(len(queryParams))
	}

	if cameraID != "" && cameraID != "all" {
		queryParams = append(queryParams, cameraID)
		query += ` AND message.camera_id = $` + strconv.Itoa(len(queryParams))
	}

	if status != "" && status != "all" {
		if status == "null" {
			query += ` AND status IS NULL`
		} else {
			queryParams = append(queryParams, status)
			query += ` AND message.status = $` + strconv.Itoa(len(queryParams))
		}
	}

	// Execute query and handle potential errors
	if err := db.sql.Db.Select(&data, query, queryParams...); err != nil {
		fmt.Println("Error filtering repo:", err)
		return nil, err
	}

	for i := range data {
		event := []model.Event{}
		if objectType != "all" {
			queryEvent := `SELECT * FROM event WHERE message_id = $1 AND event_type = $2`
			if err := db.sql.Db.Select(&event, queryEvent, data[i].Message_id, objectType); err != nil {
				return nil, err
			}
		} else {
			queryEvent := `SELECT * FROM event WHERE message_id = $1`
			if err := db.sql.Db.Select(&event, queryEvent, data[i].Message_id); err != nil {
				return nil, err
			}
		}

		data[i].Event_list = event
	}

	for i := range data {
		object := []model.Object{}
		if objectType == "all" {
			queryObject := `select * from object WHERE message_id = $1`
			if err := db.sql.Db.Select(&object, queryObject, data[i].Message_id); err != nil {
				fmt.Println("error in select * object", err)
				return nil, err
			}
		} else if objectType == "vehicle_event" {
			obj := "Vehicle"
			queryObject := `select * from object WHERE message_id = $1 and object_type = $2`
			if err := db.sql.Db.Select(&object, queryObject, data[i].Message_id, obj); err != nil {
				fmt.Println("error in select * object", err)
				return nil, err
			}
		} else {
			obj := "Human"
			queryObject := `select * from object WHERE message_id = $1 and object_type = $2`
			if err := db.sql.Db.Select(&object, queryObject, data[i].Message_id, obj); err != nil {
				fmt.Println("error in select * object", err)
				return nil, err
			}
		}
		data[i].Object_list = object
	}
	return data, nil
}

func (db *RepoImpl) AcceptRepo(ctx context.Context, id string) error {

	query := `UPDATE message SET status = true WHERE message_id = $1`

	_, err := db.sql.Db.Exec(query, id)
	if err != nil {
		fmt.Println("loi thuc hien query accept repo")
		return err
	}

	return nil
}

func (db *RepoImpl) RejectRepo(ctx context.Context, id string) error {

	query := `UPDATE message SET status = false WHERE message_id = $1`

	_, err := db.sql.Db.Exec(query, id)
	if err != nil {
		fmt.Println("loi thuc hien query reject repo")
		return err
	}

	return nil
}

func (db *RepoImpl) DiscardAckRepo(ctx context.Context, id string) error {

	query := `UPDATE message SET status = NULL WHERE message_id = $1`

	_, err := db.sql.Db.Exec(query, id)
	if err != nil {
		fmt.Println("loi thuc hien query discard ack repo")
		return err
	}

	return nil
}
