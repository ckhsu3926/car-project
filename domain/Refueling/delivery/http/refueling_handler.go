package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"car-record/entities"
	"car-record/tools"
)

type refuelingHttpHandler struct {
	Usecase entities.RefuelingUsecase
}

func NewRefuelingHttpHandler(r *gin.RouterGroup, u entities.RefuelingUsecase) {
	handler := &refuelingHttpHandler{
		Usecase: u,
	}

	r.POST("/add", handler.Add)
	r.GET("/list", handler.GetList)
	r.PUT("/update", handler.Update)
	r.DELETE("/delete", handler.Delete)
}

type refueling struct {
	ID                uint            `json:"id"`
	VehicleID         uint            `json:"vehicleID" binding:"required"`
	Date              string          `json:"date" binding:"required"`
	Station           string          `json:"station" binding:"required"`
	OctaneNumber      int             `json:"octaneNumber" binding:"required"`
	UnitPrice         decimal.Decimal `json:"unitPrice"`
	Count             decimal.Decimal `json:"count"`
	Value             int             `json:"value" binding:"required"`
	Mileage           decimal.Decimal `json:"mileage" binding:"required"`
	MointorFuelRecord decimal.Decimal `json:"mointorFuelRecord"`
}

type refuelingResponse struct {
	tools.GinResponse
	Data []entities.RefuelingRecord `json:"data"`
}

// Refueling godoc
// @Summary      Add
// @Description  Add vehicle's refueling record
// @Tags         refueling
// @Param        body  body      refueling  false  "post body"
// @Success      200                {object}  refuelingResponse
// @Router       /api/refueling/add [post]
func (h *refuelingHttpHandler) Add(c *gin.Context) {
	response := refuelingResponse{}

	var body refueling
	getBodyErr := c.ShouldBindJSON(&body)
	if getBodyErr != nil {
		response.Error = getBodyErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}
	record := entities.RefuelingRecord{
		ID:                body.ID,
		VehicleID:         body.VehicleID,
		Date:              body.Date,
		Station:           body.Station,
		OctaneNumber:      body.OctaneNumber,
		UnitPrice:         body.UnitPrice,
		Count:             body.Count,
		Value:             body.Value,
		Mileage:           body.Mileage,
		MointorFuelRecord: body.MointorFuelRecord,
	}

	addErr := h.Usecase.Add(c.Request.Context(), record)
	if addErr != nil {
		response.ErrorResponse(c, addErr)
		return
	}

	list, getListerr := h.Usecase.GetList(c.Request.Context(), body.VehicleID)
	if getListerr != nil {
		response.ErrorResponse(c, getListerr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Refueling godoc
// @Summary      GetList
// @Description  Get vehicle's refueling record list
// @Tags         refueling
// @Param        vehicleID  query     uint  true  "Vehicle ID"
// @Success      200   {object}  refuelingResponse
// @Router       /api/refueling/list [get]
func (h *refuelingHttpHandler) GetList(c *gin.Context) {
	response := refuelingResponse{}

	vehicleIDString := c.Query("vehicleID")
	vehicleID, atoiErr := strconv.Atoi(vehicleIDString)
	if vehicleIDString == "" || atoiErr != nil {
		response.Error = "invalid vehicle id"
		c.AbortWithStatusJSON(400, response)
		return
	}

	list, getListErr := h.Usecase.GetList(c.Request.Context(), uint(vehicleID))
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Refueling godoc
// @Summary      Update
// @Description  Update vehicle's refueling record
// @Tags         refueling
// @Param        body  body      refueling  false  "post body"
// @Success      200   {object}  refuelingResponse
// @Router       /api/refueling/update [put]
func (h *refuelingHttpHandler) Update(c *gin.Context) {
	response := refuelingResponse{}

	var body refueling
	getBodyErr := c.ShouldBindJSON(&body)
	if getBodyErr != nil {
		response.Error = getBodyErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}
	record := entities.RefuelingRecord{
		ID:                body.ID,
		VehicleID:         body.VehicleID,
		Date:              body.Date,
		Station:           body.Station,
		OctaneNumber:      body.OctaneNumber,
		UnitPrice:         body.UnitPrice,
		Count:             body.Count,
		Value:             body.Value,
		Mileage:           body.Mileage,
		MointorFuelRecord: body.MointorFuelRecord,
	}

	updateErr := h.Usecase.Update(c.Request.Context(), record)
	if updateErr != nil {
		response.ErrorResponse(c, updateErr)
		return
	}

	list, getListerr := h.Usecase.GetList(c.Request.Context(), body.VehicleID)
	if getListerr != nil {
		response.ErrorResponse(c, getListerr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Refueling godoc
// @Summary      Delete
// @Description  Delete vehicle's refueling record
// @Tags         refueling
// @Param        vehicleID          query     uint  true  "Vehicle ID"
// @Param        refuelingRecordID  query     uint  true  "Refueling Record ID"
// @Success      200        {object}  refuelingResponse
// @Router       /api/refueling/delete [delete]
func (h *refuelingHttpHandler) Delete(c *gin.Context) {
	response := refuelingResponse{}

	vehicleIDString := c.Query("vehicleID")
	vehicleID, atoiErr := strconv.Atoi(vehicleIDString)
	if vehicleIDString == "" || atoiErr != nil {
		response.Error = "invalid vehicle id"
		c.AbortWithStatusJSON(400, response)
		return
	}

	refuelingRecordIDString := c.Query("refuelingRecordID")
	refuelingRecordID, atoiErr := strconv.Atoi(refuelingRecordIDString)
	if refuelingRecordIDString == "" || atoiErr != nil {
		response.Error = "invalid refueling record id"
		c.AbortWithStatusJSON(400, response)
		return
	}

	deleteErr := h.Usecase.Delete(c.Request.Context(), uint(refuelingRecordID))
	if deleteErr != nil {
		response.ErrorResponse(c, deleteErr)
		return
	}

	list, getListerr := h.Usecase.GetList(c.Request.Context(), uint(vehicleID))
	if getListerr != nil {
		response.ErrorResponse(c, getListerr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}
