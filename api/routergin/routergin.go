package routergin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/larikhide/urlshortener/api/handler"
)

type RouterGin struct {
	*gin.Engine
	hs *handler.Handlers
}

func NewRouterGin(hs *handler.Handlers) *RouterGin {

	r := gin.Default()
	ret := &RouterGin{
		hs: hs,
	}
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})
	r.POST("/create-short-url", ret.CreateShortUrl)
	r.GET("/:shortUrl", ret.HandleShortUrlRedirect)

	ret.Engine = r
	return ret
}

type URL handler.URL

func (r *RouterGin) CreateShortUrl(c *gin.Context) {
	url := URL{}
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL, err := r.hs.GenShortLink(c.Request.Context(), handler.URL(url))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "short url created successfully",
		"short_url": shortURL,
	})
}

func (r *RouterGin) HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl, err := r.hs.HandleShortUrlRedirect(c.Request.Context(), shortUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, initialUrl.LongURL)
}
