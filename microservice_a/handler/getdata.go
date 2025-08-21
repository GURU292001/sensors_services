package handler

import (
	"net/http"
	db "sensors/DB"
	"sensors/helpers"

	"github.com/labstack/echo/v4"
)

type ReadingData struct {
	Sensor_type  string
	Sensor_value string
	Id1          string
	Id2          string
	Time_stamp   string
}

// @Summary      Retrieve all sensor data
// @Description  Returns all sensor readings from the database
// @Tags         sensors
// @Produce      json
// @Success      200  {array}   ReadingData
// @Failure      500  {object}  map[string]string
// @Router       /sensors [get]
func GetData(c echo.Context) error {
	debug := new(helpers.HelperStruct)
	debug.Init()
	debug.SetUid(c.Request())
	debug.Log(helpers.Statement, "GetData (+)")
	lRecords := []ReadingData{}

	lRow, lerr := db.Gdb.Query("SELECT nvl(sensor_type,''), nvl(sensor_value,''), nvl(id1,''), nvl(id2,''), nvl(time_stamp,'') FROM readings_table;")
	if lerr != nil {
		debug.Log(helpers.Elog, "GD01", lerr)
		return lerr
	}
	defer lRow.Close()
	for lRow.Next() {
		var lrecord ReadingData
		lerr = lRow.Scan(&lrecord.Sensor_type, &lrecord.Sensor_value, &lrecord.Id1, &lrecord.Id2, &lrecord.Time_stamp)
		if lerr != nil {
			debug.Log(helpers.Elog, "GD02", lerr)
			return lerr
		}
		lRecords = append(lRecords, lrecord)
	}
	debug.Log(helpers.Statement, "GetData (-)")
	return c.JSON(http.StatusOK, lRecords)
	// return c.String(http.StatusOK, "Hello, World! from getData")

}
