package handler

import (
	"net/http"
	db "sensors/DB"
	"sensors/helpers"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReadingData struct {
	Sensor_type  string
	Sensor_value string
	Id1          string
	Id2          string
	Time_stamp   string
}

// @Summary      Retrieve paginated sensor data
// @Description  Returns sensor readings with pagination support
// @Tags         sensors
// @Produce      json
// @Param        page   query     int  false  "Page number (default is 1)"
// @Param        limit  query     int  false  "Number of records per page (default is 10)"
// @Success      200    {array}   ReadingData
// @Failure      500    {object}  map[string]string
// @Router       /sensors [get]
func GetData(c echo.Context) error {
	debug := new(helpers.HelperStruct)
	debug.Init()
	debug.SetUid(c.Request())
	debug.Log(helpers.Statement, "GetData (+)")
	lRecords := []ReadingData{}
	pageStr := c.QueryParam("page")
	limitStr := c.QueryParam("limit")

	// Default values if not provided
	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	offset := (page - 1) * limit

	lSqlString := `SELECT IFNULL(sensor_type,''), IFNULL(sensor_value,''), IFNULL(id1,''), IFNULL(id2,''), IFNULL(time_stamp,'') FROM readings_table  LIMIT ? OFFSET ? ;`

	lRow, lerr := db.Gdb.Query(lSqlString, limit, offset)
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
