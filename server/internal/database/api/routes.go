package api

import (
	"github.com/ArtuoS/dont-forget/internal/database/repository"
	"github.com/ArtuoS/dont-forget/internal/entity"
	"github.com/gin-gonic/gin"
)

type Router struct {
	ItemRepository *repository.ItemRepository
}

func NewRouter(itemRepositorysitory *repository.ItemRepository) *Router {
	return &Router{
		ItemRepository: itemRepositorysitory,
	}
}

func (r *Router) StartRouting() {
	engine := gin.Default()
	r.routeItems(engine)
	engine.Run("localhost:8081")
}

func (r *Router) routeItems(engine *gin.Engine) {
	engine.GET("/items", r.handleGetAll)
	engine.POST("/items", r.handleCreate)
	engine.GET("/items/:guid", r.handleGetByGuid)
	engine.DELETE("/items/:guid", r.handleDelete)
	engine.PUT("/items", r.handleUpdate)
}

func setupCors(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (r *Router) handleGetAll(c *gin.Context) {
	setupCors(c)

	items, err := r.ItemRepository.GetAll()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, items)
}

func (r *Router) handleGetByGuid(c *gin.Context) {
	setupCors(c)

	guid := c.Param("guid")
	item, err := r.ItemRepository.Get(guid)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, item)
}

func (r *Router) handleCreate(c *gin.Context) {
	setupCors(c)

	var item entity.Item
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(500, err)
		return
	}

	item.GenerateGuid()
	err = item.IsValid()
	if err != nil {
		c.JSON(500, err)
		return
	}

	err = r.ItemRepository.Save(&item)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, item)
}

func (r *Router) handleDelete(c *gin.Context) {
	setupCors(c)

	guid := c.Param("guid")
	err := r.ItemRepository.Delete(guid)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, nil)
}

func (r *Router) handleUpdate(c *gin.Context) {
	setupCors(c)

	var item entity.Item
	err := c.BindJSON(&item)
	if err != nil {
		c.JSON(500, err)
		return
	}

	err = item.IsValid()
	if err != nil {
		c.JSON(500, err)
		return
	}

	err = r.ItemRepository.Update(&item)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, item)
}
