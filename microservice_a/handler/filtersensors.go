package handler

import (
	"log"
	"net/http"
	db "sensors/DB"
	"sensors/helpers"

	"github.com/labstack/echo/v4"
)

// Filter by ID1 + ID2
// GetByID godoc
// @Summary      Retrieve sensor data by ID1 and ID2
// @Description  Returns all sensor readings that match the given ID1 and ID2
// @Tags         sensors
// @Produce      json
// @Param        id1   query     string  true  "ID1 value (e.g., 'A')"
// @Param        id2   query     string  true  "ID2 value (e.g., '100')"
// @Success      200   {array}   ReadingData
// @Failure      500   {object}  map[string]string
// @Router       /filter-byid [get]
func GetByID(c echo.Context) error {
	debug := new(helpers.HelperStruct)
	debug.Init()
	debug.SetUid(c.Request())
	debug.Log(helpers.Statement, "GetByID (+)")
	id1 := c.QueryParam("id1")
	id2 := c.QueryParam("id2")

	rows, err := db.Gdb.Query(`
	SELECT nvl(sensor_type,''), nvl(sensor_value,''), nvl(id1,''), nvl(id2,''), nvl(time_stamp,'')
	FROM readings_table
	WHERE id1 = ? AND id2 = ?`, id1, id2)
	if err != nil {
		debug.Log(helpers.Elog, "GBI01 ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	defer rows.Close()

	var results []ReadingData
	for rows.Next() {
		var r ReadingData
		if err := rows.Scan(&r.Sensor_type, &r.Sensor_value, &r.Id1, &r.Id2, &r.Time_stamp); err != nil {
			debug.Log(helpers.Elog, "GBI02 ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		results = append(results, r)
	}

	debug.Log(helpers.Statement, "GetByID (-)")
	return c.JSON(http.StatusOK, results)
}

// GetByTime godoc
// @Summary      Retrieve sensor data by time range
// @Description  Returns all sensor readings between a start and end timestamp
// @Tags         sensors
// @Produce      json
// @Param        start   query     string  true  "Start timestamp (format: YYYY-MM-DD HH:MM:SS)"
// @Param        end     query     string  true  "End timestamp (format: YYYY-MM-DD HH:MM:SS)"
// @Success      200     {array}   ReadingData
// @Failure      500     {object}  map[string]string
// @Router       /sensors/filter-bytime [get]
func GetByTime(c echo.Context) error {
	debug := new(helpers.HelperStruct)
	debug.Init()
	debug.Log(helpers.Statement, "GetByTime (+)")
	debug.SetUid(c.Request())
	start := c.QueryParam("start") // e.g., "2020-05-04 11:00:00"
	end := c.QueryParam("end")     // e.g., "2020-06-04 14:00:00"
	log.Printf("start : %v ,end : %v", start, end)
	rows, err := db.Gdb.Query(`
	SELECT nvl(sensor_type,''), nvl(sensor_value,''), nvl(id1,''), nvl(id2,''), nvl(time_stamp,'')
	FROM readings_table
	WHERE time_stamp BETWEEN ? AND ?`, start, end)
	if err != nil {
		debug.Log(helpers.Elog, "GBT01 ", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "GBT01 " + err.Error()})
	}
	defer rows.Close()

	var results []ReadingData
	for rows.Next() {
		var r ReadingData
		if err := rows.Scan(&r.Sensor_type, &r.Sensor_value, &r.Id1, &r.Id2, &r.Time_stamp); err != nil {
			debug.Log(helpers.Elog, "GBT01 ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "GBT02 " + err.Error()})
		}
		results = append(results, r)
	}

	debug.Log(helpers.Statement, "GetByTime (-)")
	return c.JSON(http.StatusOK, results)
}
