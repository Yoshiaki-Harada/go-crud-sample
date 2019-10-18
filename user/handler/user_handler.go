package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"web-app/models"
	"web-app/user"
)

type JsonErrorResponse struct {
	Message string `json:"message"`
}

type JsonIdrResponse struct {
	Id uint `json:"id"`
}
type userHandler struct {
	use user.Usecase
}

func NewUserHandler(r *gin.Engine, u user.Usecase) {
	h := &userHandler{u}
	r.GET("/hello", h.Hello)
	r.GET("/users/:id", h.GetUser)
	r.GET("/users", h.GetUsers)
	r.PUT("/users/:id", h.Update)
	r.POST("/users", h.Create)
	r.DELETE("/users/:id", h.Delete)
}

func (h *userHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"message": "hello"})
}

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.use.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, JsonErrorResponse{err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *userHandler) GetUser(c *gin.Context) {
	sid := c.Param("id")
	iid, er := strconv.Atoi(sid)
	if er != nil {
		c.JSON(http.StatusBadRequest, JsonErrorResponse{er.Error()})
		return
	}
	uid := uint(iid)
	u, err := h.use.GetById(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, JsonErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, JsonIdrResponse{Id: u.ID})
}

func (h *userHandler) Delete(c *gin.Context) {
	u := models.User{}
	sid := c.Param("id")
	iid, er := strconv.Atoi(sid)
	if er != nil {
		c.JSON(http.StatusBadRequest, JsonErrorResponse{Message: er.Error()})
		return
	}
	u.ID = uint(iid)
	id, err := h.use.Delete(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, JsonErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, JsonIdrResponse{id})
}

func (h *userHandler) Create(c *gin.Context) {
	nu := models.User{}
	er := c.BindJSON(&nu)
	if er != nil {
		c.JSON(http.StatusBadRequest, JsonErrorResponse{Message: er.Error()})
		return
	}
	u, err := h.use.Create(nu)
	if err != nil {
		c.JSON(http.StatusBadRequest, JsonErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusBadRequest, u)
}

func (h *userHandler) Update(c *gin.Context) {
	nu := models.User{}
	e := c.BindJSON(&nu)
	if e != nil {
		c.JSON(http.StatusBadRequest, JsonErrorResponse{Message: e.Error()})
		return
	}
	sid := c.Param("id")
	iid, er := strconv.Atoi(sid)
	if er != nil {
		c.JSON(http.StatusBadRequest, JsonErrorResponse{Message: er.Error()})
		return
	}
	nu.ID = uint(iid)
	u, err := h.use.Update(nu)
	if err != nil {
		c.JSON(http.StatusBadRequest, JsonErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, u)
}
