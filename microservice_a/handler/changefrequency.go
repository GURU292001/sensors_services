package handler

import (
	"log"
	"net/http"
	"sensors/helpers"
	"time"

	"github.com/labstack/echo/v4"
)

// Request represents the input for changing frequency
type Request struct {
	Sensor    string `json:"sensor" example:"TEMPERATURE"` // TEMPERATURE, MOTION, HUMIDITY
	Frequency string `json:"frequency" example:"2s"`       // duration format: "2s", "500ms", "1m"
}

// Response represents the output of ChangeFrequency
type Response struct {
	Status string `json:"status" example:"S"`
	Msg    string `json:"msg" example:"Frequency changed Successfully"`
}

// ChangeFrequency godoc
// @Summary      Change sensor frequency
// @Description  Update the data generation frequency for a specific sensor type (TEMPERATURE, MOTION, HUMIDITY).
// @Tags         sensors
// @Accept       json
// @Produce      json
// @Param        request body Request true "Sensor type and frequency"
// @Success      200 {object} Response
// @Failure      400 {object} map[string]string
// @Router       /sensors/frequency [put]
func ChangeFrequency(c echo.Context) error {
	debug := new(helpers.HelperStruct)
	debug.Init()
	debug.SetUid(c.Request())
	debug.Log(helpers.Statement, "ChangeFrequency(+)")
	lReq := new(Request)
	var lResp Response
	lResp.Status = "S"
	lResp.Msg = "Frequency changed Successfully"
	if err := c.Bind(lReq); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	dur, err := time.ParseDuration(lReq.Frequency)
	if err != nil {
		log.Println("error on frequency time :", err)
	}
	switch lReq.Sensor {
	case "TEMPERATURE":
		TemperatureDuration = dur
	case "MOTION":
		MotionDuration = dur
	case "HUMIDITY":
		HumidityDuration = dur
	}
	log.Printf("request :%+v, Temperature:%v , motion:%v , Humidity:%v ,", lReq, TemperatureDuration, MotionDuration, HumidityDuration)
	debug.Log(helpers.Statement, "ChangeFrequency(-)")
	return c.JSON(http.StatusOK, lResp)

}
