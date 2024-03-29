package http

import (
	"car-record/entities"
	"car-record/middleware"
	"car-record/tools"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

type vehicleHttpHandler struct {
	VUsecase entities.VehicleUsecase
}

func NewVehicleHttpHandler(r *gin.RouterGroup, vu entities.VehicleUsecase) {
	handler := &vehicleHttpHandler{
		VUsecase: vu,
	}

	r.GET("/", handler.Get)
	r.POST("/add", handler.Add)
	r.PUT("/edit", handler.Edit)
	r.GET("/list", handler.GetList)
	r.DELETE("/delete", handler.Delete)
}

type getDetailResponse struct {
	tools.GinResponse
	Data entities.VehicleDetail `json:"data"`
}

type addBody struct {
	Name    string `json:"name" binding:"required"`
	Licence string `json:"license" binding:"required"`
	Company string `json:"company"`
	Model   string `json:"model"`
}

type editBody struct {
	ID                  uint            `json:"id"`
	Name                string          `json:"name"`
	License             string          `json:"license"`
	Company             string          `json:"company"`
	Model               string          `json:"model"`
	EngineDisplacement  decimal.Decimal `json:"engineDisplacement"`
	EngineNumber        string          `json:"engineNumber"`
	DefaultOctaneNumber int             `json:"defaultOctaneNumber"`
	Purchase            int             `json:"purchase"`
	PurchaseDate        string          `json:"purchaseDate"`
	PurchaseLocation    string          `json:"purchaseLocation"`
	PurchaseMileage     decimal.Decimal `json:"purchaseMileage"`
	Sold                int             `json:"sold"`
	SoldDate            string          `json:"soldDate"`
	SoldMileage         decimal.Decimal `json:"soldMileage"`
	MileageReset        decimal.Decimal `json:"mileageReset"`
}

type getListResponse struct {
	tools.GinResponse
	Data []entities.Vehicle `json:"data"`
}

// Vehicle Get godoc
// @Summary      Get Vehicle
// @Description  Get User's Vehicle Detail
// @Tags         Vehicle
// @Param        vehicleID  query     uint  true  "Vehicle ID"
// @Success      200        {object}  getDetailResponse
// @Router       /api/vehicle/ [get]
func (h *vehicleHttpHandler) Get(c *gin.Context) {
	response := getDetailResponse{}

	vehicleIDString := c.Query("vehicleID")
	vehicleID, atoiErr := strconv.Atoi(vehicleIDString)
	if vehicleIDString == "" || atoiErr != nil {
		response.Error = "invalid vehicle id"
		c.AbortWithStatusJSON(400, response)
		return
	}

	detail, getDetailErr := h.VUsecase.Get(c.Request.Context(), uint(vehicleID))
	if getDetailErr != nil {
		response.ErrorResponse(c, getDetailErr)
		return
	}

	response.Data = detail
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Vehicle Add godoc
// @Summary      Add
// @Description  Create vehicle
// @Tags         Vehicle
// @Param        body  body      addBody  false  "post body"
// @Success      200   {object}  getListResponse
// @Router       /api/vehicle/add [post]
func (h *vehicleHttpHandler) Add(c *gin.Context) {
	response := getListResponse{}

	user, getMiddlewareAuthorizeErr := middleware.GetMiddlewareAuthorize(c)
	if getMiddlewareAuthorizeErr != nil {
		response.Error = getMiddlewareAuthorizeErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	var body addBody
	getBodyErr := c.ShouldBindJSON(&body)
	if getBodyErr != nil {
		response.Error = getBodyErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	addErr := h.VUsecase.Add(c.Request.Context(), user.ID, body.Name, body.Licence, body.Company, body.Model)
	if addErr != nil {
		response.ErrorResponse(c, addErr)
		return
	}

	list, getListErr := h.VUsecase.GetList(c.Request.Context(), user.ID)
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Vehicle Edit godoc
// @Summary      Edit
// @Description  Edit vehicle detail
// @Tags         Vehicle
// @Param        body  body      editBody  false  "put body"
// @Success      200   {object}  getDetailResponse
// @Router       /api/vehicle/edit [put]
func (h *vehicleHttpHandler) Edit(c *gin.Context) {
	response := getDetailResponse{}

	user, getMiddlewareAuthorizeErr := middleware.GetMiddlewareAuthorize(c)
	if getMiddlewareAuthorizeErr != nil {
		response.Error = getMiddlewareAuthorizeErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	var body editBody
	getBodyErr := c.ShouldBindJSON(&body)
	if getBodyErr != nil {
		response.Error = getBodyErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	detail := entities.VehicleDetail{
		ID: body.ID,
		// UserID:              body.UserID,
		Name:                body.Name,
		License:             body.License,
		Company:             body.Company,
		Model:               body.Model,
		EngineDisplacement:  body.EngineDisplacement,
		EngineNumber:        body.EngineNumber,
		DefaultOctaneNumber: body.DefaultOctaneNumber,
		Purchase:            body.Purchase,
		PurchaseDate:        body.PurchaseDate,
		PurchaseLocation:    body.PurchaseLocation,
		PurchaseMileage:     body.PurchaseMileage,
		Sold:                body.Sold,
		SoldDate:            body.SoldDate,
		SoldMileage:         body.SoldMileage,
		MileageReset:        body.MileageReset,
	}

	editErr := h.VUsecase.Edit(c.Request.Context(), user.ID, detail)
	if editErr != nil {
		response.ErrorResponse(c, editErr)
		return
	}

	response.Data = detail
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Vehicle GetList godoc
// @Summary      Get List
// @Description  Get User's Vehicle List
// @Tags         Vehicle
// @Success      200        {object}  getListResponse
// @Router       /api/vehicle/list [get]
func (h *vehicleHttpHandler) GetList(c *gin.Context) {
	response := getListResponse{}

	user, getMiddlewareAuthorizeErr := middleware.GetMiddlewareAuthorize(c)
	if getMiddlewareAuthorizeErr != nil {
		response.Error = getMiddlewareAuthorizeErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	list, getListErr := h.VUsecase.GetList(c.Request.Context(), user.ID)
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Vehicle Delete godoc
// @Summary      Delete
// @Description  Delete vehicle
// @Tags         Vehicle
// @Param        vehicleID  query     uint  true  "Vehicle ID"
// @Success      200  {object}  getListResponse
// @Router       /api/vehicle/delete [delete]
func (h *vehicleHttpHandler) Delete(c *gin.Context) {
	response := getListResponse{}

	vehicleIDString := c.Query("vehicleID")
	vehicleID, atoiErr := strconv.Atoi(vehicleIDString)
	if vehicleIDString == "" || atoiErr != nil {
		response.Error = "invalid vehicle id"
		c.AbortWithStatusJSON(400, response)
		return
	}

	user, getMiddlewareAuthorizeErr := middleware.GetMiddlewareAuthorize(c)
	if getMiddlewareAuthorizeErr != nil {
		response.Error = getMiddlewareAuthorizeErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	deleteErr := h.VUsecase.Delete(c.Request.Context(), uint(vehicleID))
	if deleteErr != nil {
		response.ErrorResponse(c, deleteErr)
		return
	}

	list, getListErr := h.VUsecase.GetList(c.Request.Context(), user.ID)
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}
