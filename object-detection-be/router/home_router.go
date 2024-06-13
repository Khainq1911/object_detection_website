package router

import (
	handler "WebDetection/handler/home"

	"github.com/labstack/echo"
)

type API struct {
	Echo    *echo.Echo
	Handler handler.SiteHandler
}

func (api *API) SetUpRouter() {
	api.Echo.GET("/", api.Handler.Index)
	api.Echo.GET("/chart", api.Handler.Index)
	api.Echo.GET("/filter/:eventType/:timeFrom/:timeTo/:cameraID/:status", api.Handler.FilterSite)
	api.Echo.GET("chart/filter/:eventType/:timeFrom/:timeTo/:cameraID/:status", api.Handler.FilterSite)
	api.Echo.GET("/accept/:message_id", api.Handler.Accept)
	api.Echo.GET("/reject/:message_id", api.Handler.Reject)
	api.Echo.GET("/discard-ack/:message_id", api.Handler.DiscardAck)
}
