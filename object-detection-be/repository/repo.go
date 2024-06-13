package repository

import (
	"WebDetection/model"
	"context"
)

type Repo interface {
	IndexRepo(ctx context.Context) ([]model.Index, error)
	FilterRepo(ctx context.Context, objectType string, timeFrom string, timeTo string, cameraID string, status string) ([]model.Index, error)
	AcceptRepo(ctx context.Context, id string) error
	RejectRepo(ctx context.Context, id string) error
	DiscardAckRepo(ctx context.Context, id string) error
}
