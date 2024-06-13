package handler

import (
	"WebDetection/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type SiteHandler struct {
	Repo repository.Repo
}

// get /
func (u *SiteHandler) Index(ctx echo.Context) error {

	result, err := u.Repo.IndexRepo(ctx.Request().Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, echo.Map{
			"Message": "loi o index",
			"Error":   err,
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"Message": "Succesfully",
		"data":    result,
	})
}

//

func (u *SiteHandler) FilterSite(ctx echo.Context) error {

	eventType := ctx.Param("eventType")
	timeFrom := ctx.Param("timeFrom")
	timeTo := ctx.Param("timeTo")
	cameraID := ctx.Param("cameraID")
	status := ctx.Param("status")

	result, err := u.Repo.FilterRepo(ctx.Request().Context(), eventType, timeFrom, timeTo, cameraID, status)
	if err != nil {
		fmt.Println("loi o filter handle")
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "error",
			"error":   err,
		})
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "succesfully",
		"data":    result,
	})
}

func (u *SiteHandler) Accept(ctx echo.Context) error {
	messageID := ctx.Param("message_id")
	if err := u.Repo.AcceptRepo(ctx.Request().Context(), messageID); err != nil {
		fmt.Println("loi o accept handle")
		ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "accepted",
	})
}

func (u *SiteHandler) Reject(ctx echo.Context) error {
	messageID := ctx.Param("message_id")
	if err := u.Repo.RejectRepo(ctx.Request().Context(), messageID); err != nil {
		fmt.Println("loi o reject handle")
		ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "rejected",
	})
}

func (u *SiteHandler) DiscardAck(ctx echo.Context) error {
	messageID := ctx.Param("message_id")

	if err := u.Repo.DiscardAckRepo(ctx.Request().Context(), messageID); err != nil {
		fmt.Println("loi o discard ack handle")
		ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "done",
	})
}
