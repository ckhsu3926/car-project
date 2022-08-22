package http

import (
	"car-record/entities"
	"car-record/middleware"
	"car-record/tools"

	"github.com/gin-gonic/gin"
)

type userGasStationHttpHandler struct {
	UUsecase entities.UserGasStationUsecase
}

func NewUserGasStationHttpHandler(r *gin.RouterGroup, uu entities.UserGasStationUsecase) {
	handler := &userGasStationHttpHandler{
		UUsecase: uu,
	}

	r.POST("/add", handler.Add)
	r.GET("/list", handler.GetList)
	r.DELETE("/delete", handler.Delete)
}

type addBody struct {
	Name string `json:"name" binding:"required"`
}
type deleteBody struct {
	ID uint `json:"id" binding:"required"`
}
type addResponse struct {
	tools.GinResponse
	Data []entities.UserGasStation `json:"data"`
}

// UserGasStation godoc
// @Summary      Add
// @Description  Create user's station record
// @Tags         userGasStation
// @Param        body  body      addBody  false  "post body"
// @Success      200  {object}  addResponse
// @Router       /api/user/gas/station/add [post]
func (h *userGasStationHttpHandler) Add(c *gin.Context) {
	response := addResponse{}

	user, getMiddlewareAuthorizeErr := middleware.GetMiddlewareAuthorize(c)
	if getMiddlewareAuthorizeErr != nil {
		response.Error = getMiddlewareAuthorizeErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	var body addBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		response.Error = err.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	addErr := h.UUsecase.Add(c.Request.Context(), user.ID, body.Name)
	if addErr != nil {
		response.ErrorResponse(c, addErr)
		return
	}

	list, getListErr := h.UUsecase.GetList(c.Request.Context(), user.ID)
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// UserGasStation godoc
// @Summary      GetList
// @Description  Get user's station record
// @Tags         userGasStation
// @Success      200   {object}  addResponse
// @Router       /api/user/gas/station/list [get]
func (h *userGasStationHttpHandler) GetList(c *gin.Context) {
	response := addResponse{}

	user, getMiddlewareAuthorizeErr := middleware.GetMiddlewareAuthorize(c)
	if getMiddlewareAuthorizeErr != nil {
		response.Error = getMiddlewareAuthorizeErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	list, getListErr := h.UUsecase.GetList(c.Request.Context(), user.ID)
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// UserGasStation godoc
// @Summary      Delete
// @Description  Delete user's station record
// @Tags         userGasStation
// @Param        body  body      deleteBody  false  "post body"
// @Success      200   {object}  addResponse
// @Router       /api/user/gas/station/delete [delete]
func (h *userGasStationHttpHandler) Delete(c *gin.Context) {
	response := addResponse{}

	user, getMiddlewareAuthorizeErr := middleware.GetMiddlewareAuthorize(c)
	if getMiddlewareAuthorizeErr != nil {
		response.Error = getMiddlewareAuthorizeErr.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	var body deleteBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		response.Error = err.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	deleteErr := h.UUsecase.Delete(c.Request.Context(), user.ID, body.ID)
	if deleteErr != nil {
		response.ErrorResponse(c, deleteErr)
		return
	}

	list, getListErr := h.UUsecase.GetList(c.Request.Context(), user.ID)
	if getListErr != nil {
		response.ErrorResponse(c, getListErr)
		return
	}

	response.Data = list
	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}
