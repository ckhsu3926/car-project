package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"car-record/entities"
	"car-record/tools"
)

type maintenanceHttpHandler struct {
	Usecase entities.MaintenanceUsecase
}

func NewMaintenanceHttpHandler(r *gin.RouterGroup, u entities.MaintenanceUsecase) {
	handler := &maintenanceHttpHandler{
		Usecase: u,
	}

	recordRouter := r.Group("record")
	recordRouter.POST("/create", handler.CreateRecord)
	recordRouter.GET("/list", handler.GetRecordList)
	recordRouter.PUT("/update", handler.UpdateRecord)
	recordRouter.DELETE("/delete", handler.DeleteRecord)

	detailRouter := recordRouter.Group("detail")
	detailRouter.GET("/", handler.GetDetailList)
	detailRouter.POST("/", handler.SetDetailList)
}

type maintenanceRecord struct {
	ID        uint            `json:"id"`
	VehicleID uint            `json:"vehicleID" binding:"required"`
	Date      string          `json:"date" binding:"required"`
	Mileage   decimal.Decimal `json:"mileage" binding:"required"`
	Shop      string          `json:"shop" binding:"required"`
	Amount    int             `json:"amount" binding:"required"`
}
type maintenanceRecordResponse struct {
	tools.GinResponse
	Data []entities.MaintenanceRecord `json:"data"`
}

type maintenanceRecordDetail struct {
	Name    string `json:"name" binding:"required"`
	Value   int    `json:"value"`
	Content string `json:"content"`
}
type maintenanceRecordDetailResponse struct {
	tools.GinResponse
	Data []maintenanceRecordDetail `json:"data"`
}

// Record

// Maintenance godoc
// @Summary      CreateRecord
// @Description  Create vehicle's maintenance record
// @Tags         Maintenance
// @Param        body  body      maintenanceRecord  false  "post body"
// @Success      200                  {object}  maintenanceRecordResponse
// @Router       /api/maintenance/record/create [post]
func (h *maintenanceHttpHandler) CreateRecord(c *gin.Context) {
	response := maintenanceRecordResponse{}

	var body maintenanceRecord
	getBodyErr := c.ShouldBindJSON(&body)
	if getBodyErr != nil {
		response.Error = getBodyErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}
	record := entities.MaintenanceRecord{
		VehicleID: body.VehicleID,
		Date:      body.Date,
		Mileage:   body.Mileage,
		Shop:      body.Shop,
		Amount:    body.Amount,
	}

	createErr := h.Usecase.CreateRecord(c.Request.Context(), record)
	if createErr != nil {
		response.ErrorResponse(c, createErr)
		return
	}

	list, getListErr := h.Usecase.GetRecordList(c.Request.Context(), body.VehicleID)
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Maintenance godoc
// @Summary      GetRecordList
// @Description  Get vehicle's maintenance record list
// @Tags         Maintenance
// @Param        vehicleID  query     uint  true  "Vehicle ID"
// @Success      200   {object}  maintenanceRecordResponse
// @Router       /api/maintenance/record/list [get]
func (h *maintenanceHttpHandler) GetRecordList(c *gin.Context) {
	response := maintenanceRecordResponse{}

	vehicleIDString := c.Query("vehicleID")
	vehicleID, atoiErr := strconv.Atoi(vehicleIDString)
	if vehicleIDString == "" || atoiErr != nil {
		response.Error = "invalid vehicle id"
		c.AbortWithStatusJSON(400, response)
		return
	}

	list, getListErr := h.Usecase.GetRecordList(c.Request.Context(), uint(vehicleID))
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Maintenance godoc
// @Summary      UpdateRecord
// @Description  Update vehicle's maintenance record
// @Tags         Maintenance
// @Param        body  body      maintenanceRecord  false  "post body"
// @Success      200   {object}  maintenanceRecordResponse
// @Router       /api/maintenance/record/update [put]
func (h *maintenanceHttpHandler) UpdateRecord(c *gin.Context) {
	response := maintenanceRecordResponse{}

	var body maintenanceRecord
	getBodyErr := c.ShouldBindJSON(&body)
	if getBodyErr != nil {
		response.Error = getBodyErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}
	record := entities.MaintenanceRecord{
		ID:        body.ID,
		VehicleID: body.VehicleID,
		Date:      body.Date,
		Mileage:   body.Mileage,
		Shop:      body.Shop,
		Amount:    body.Amount,
	}

	updateErr := h.Usecase.UpdateRecord(c.Request.Context(), record)
	if updateErr != nil {
		response.ErrorResponse(c, updateErr)
		return
	}

	list, getListerr := h.Usecase.GetRecordList(c.Request.Context(), body.VehicleID)
	if getListerr != nil {
		response.ErrorResponse(c, getListerr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Maintenance godoc
// @Summary      DeleteRecord
// @Description  Delete vehicle's maintenance record
// @Tags         Maintenance
// @Param        vehicleID            query     uint  true  "Vehicle ID"
// @Param        maintenanceRecordID  query     uint  true  "Maintenance Record ID"
// @Success      200        {object}  maintenanceRecordResponse
// @Router       /api/maintenance/record/delete [delete]
func (h *maintenanceHttpHandler) DeleteRecord(c *gin.Context) {
	response := maintenanceRecordResponse{}

	vehicleIDString := c.Query("vehicleID")
	vehicleID, atoiErr := strconv.Atoi(vehicleIDString)
	if vehicleIDString == "" || atoiErr != nil {
		response.Error = "invalid vehicle id"
		c.AbortWithStatusJSON(400, response)
		return
	}

	maintenanceRecordIDString := c.Query("maintenanceRecordID")
	maintenanceRecordID, atoiErr := strconv.Atoi(maintenanceRecordIDString)
	if maintenanceRecordIDString == "" || atoiErr != nil {
		response.Error = "invalid refueling record id"
		c.AbortWithStatusJSON(400, response)
		return
	}

	deleteErr := h.Usecase.DeleteRecord(c.Request.Context(), uint(maintenanceRecordID))
	if deleteErr != nil {
		response.ErrorResponse(c, deleteErr)
		return
	}

	list, getListerr := h.Usecase.GetRecordList(c.Request.Context(), uint(vehicleID))
	if getListerr != nil {
		response.ErrorResponse(c, getListerr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Detail

// Maintenance godoc
// @Summary      GetDetailList
// @Description  Get maintenance record's detail list
// @Tags         Maintenance
// @Param        maintenanceRecordID  query     uint  true  "Maintenance Record ID"
// @Success      200                  {object}  maintenanceRecordDetailResponse
// @Router       /api/maintenance/record/detail/ [get]
func (h *maintenanceHttpHandler) GetDetailList(c *gin.Context) {
	response := maintenanceRecordDetailResponse{Data: []maintenanceRecordDetail{}}

	maintenanceRecordIDString := c.Query("maintenanceRecordID")
	maintenanceRecordID, atoiErr := strconv.Atoi(maintenanceRecordIDString)
	if maintenanceRecordIDString == "" || atoiErr != nil {
		response.Error = "invalid refueling record id"
		c.AbortWithStatusJSON(400, response)
		return
	}

	list, getListErr := h.Usecase.GetDetailList(c.Request.Context(), uint(maintenanceRecordID))
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	for _, r := range list {
		response.Data = append(response.Data, maintenanceRecordDetail{
			Name:    r.Name,
			Value:   r.Value,
			Content: r.Content,
		})
	}
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Maintenance godoc
// @Summary      SetDetailList
// @Description  Set vehicle maintenance record's detail list
// @Tags         Maintenance
// @Param        maintenanceRecordID  query     uint                       true   "Maintenance Record ID"
// @Param        body                 body      []maintenanceRecordDetail  false  "post body"
// @Success      200                  {object}  maintenanceRecordDetailResponse
// @Router       /api/maintenance/record/detail/ [post]
func (h *maintenanceHttpHandler) SetDetailList(c *gin.Context) {
	response := maintenanceRecordDetailResponse{}

	maintenanceRecordIDString := c.Query("maintenanceRecordID")
	maintenanceRecordID, atoiErr := strconv.Atoi(maintenanceRecordIDString)
	if maintenanceRecordIDString == "" || atoiErr != nil {
		response.Error = "invalid refueling record id"
		c.AbortWithStatusJSON(400, response)
		return
	}

	var body []maintenanceRecordDetail
	getBodyErr := c.ShouldBindJSON(&body)
	if getBodyErr != nil {
		response.Error = getBodyErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}
	detailList := []entities.MaintenanceRecordDetail{}
	for _, detail := range body {
		detailList = append(detailList, entities.MaintenanceRecordDetail{
			Name:    detail.Name,
			Value:   detail.Value,
			Content: detail.Content,
		})
	}

	setErr := h.Usecase.SetDetailList(c.Request.Context(), uint(maintenanceRecordID), detailList)
	if setErr != nil {
		response.ErrorResponse(c, setErr)
		return
	}

	list, getListErr := h.Usecase.GetDetailList(c.Request.Context(), uint(maintenanceRecordID))
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	for _, r := range list {
		response.Data = append(response.Data, maintenanceRecordDetail{
			Name:    r.Name,
			Value:   r.Value,
			Content: r.Content,
		})
	}
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}
