package http

import (
	"car-record/entities"
	"car-record/tools"

	"github.com/gin-gonic/gin"
)

type userHttpHandler struct {
	UUsecase entities.UserUsecase
}

func NewUserHttpHandler(r *gin.RouterGroup, uu entities.UserUsecase) {
	handler := &userHttpHandler{
		UUsecase: uu,
	}

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.POST("/logout", handler.Logout)
}

type registerBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}
type registerResponse struct {
	tools.GinResponse
	Data string `json:"data"`
}

// Register godoc
// @Summary      Register
// @Description  Create account and get login token
// @Tags         user
// @Param        body  body      registerBody  false  "post body"
// @Success      200   {object}  registerResponse
// @Router       /api/register [post]
func (h *userHttpHandler) Register(c *gin.Context) {
	response := registerResponse{}

	var body registerBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		response.Error = err.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	response.Data, err = h.UUsecase.Register(c.Request.Context(), body.Username, body.Password)
	if err != nil {
		response.ErrorResponse(c, err)
		return
	}

	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type loginResponse struct {
	tools.GinResponse
	Data string `json:"data"`
}

// Login godoc
// @Summary      Login
// @Description  Login and get login token
// @Tags         user
// @Param        body  body      registerBody  false  "post body"
// @Success      200   {object}  registerResponse
// @Router       /api/login [post]
func (h *userHttpHandler) Login(c *gin.Context) {
	response := loginResponse{}

	var body loginBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		response.Error = err.Error()
		c.AbortWithStatusJSON(400, response)
		return
	}

	response.Data, err = h.UUsecase.Login(c.Request.Context(), body.Username, body.Password)
	if err != nil {
		response.ErrorResponse(c, err)
		return
	}

	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}

// Logout godoc
// @Summary      Logout
// @Description  get by header token and remove db token
// @Tags         user
// @Success      200
// @Router       /api/logout [post]
func (h *userHttpHandler) Logout(c *gin.Context) {
	response := loginResponse{}

	token := c.Request.Header.Get("token")
	if token == "" {
		c.AbortWithStatus(200)
		return
	}

	err := h.UUsecase.Logout(c.Request.Context(), token)
	if err != nil {
		response.ErrorResponse(c, err)
		return
	}

	response.Result = 1
	c.AbortWithStatusJSON(200, response)
}
