package search

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	store elasticStore
}

func NewHandler(store elasticStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) SearchContent(c *gin.Context) {
	var req SearchRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.String(http.StatusBadRequest, "Invalid request: %s", err.Error())
		return
	}

	query, err := setQuery(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	res, err := h.store.GetContentByKeyword("content_search", query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	re, err := convertResponse(res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, re)

}
