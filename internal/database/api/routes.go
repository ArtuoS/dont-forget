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
	engine.Run()
}

func (r *Router) routeItems(engine *gin.Engine) {
	engine.GET("/items", r.handleGetItems)
	engine.POST("/items", r.handlePostItems)
	engine.GET("/items/:guid", r.handleGetItem)
	engine.DELETE("/items/:guid", r.handleDeleteItem)
	engine.PUT("/items", r.handleUpdateItem)
}

func (r *Router) handleGetItems(c *gin.Context) {
	items, err := r.ItemRepository.GetAll()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, items)
}

func (r *Router) handleGetItem(c *gin.Context) {
	guid := c.Param("guid")
	item, err := r.ItemRepository.Get(guid)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, item)
}

func (r *Router) handlePostItems(c *gin.Context) {
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

func (r *Router) handleDeleteItem(c *gin.Context) {
	guid := c.Param("guid")
	err := r.ItemRepository.Delete(guid)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, nil)
}

func (r *Router) handleUpdateItem(c *gin.Context) {
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
