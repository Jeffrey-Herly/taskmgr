package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewsHandler struct{}

func NewViewsHandler() *ViewsHandler {
	return &ViewsHandler{}
}

// RenderInitialPage merender tampilan halaman initial.
func (h *ViewsHandler) RenderInitialPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Title": "Initial Page",
	})
}

// RenderLoginPage merender tampilan halaman login.
func (h *ViewsHandler) RenderLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"Title": "Login Page",
	})
}

// RenderDashboardPage merender tampilan halaman dashboard.
func (h *ViewsHandler) RenderDashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
		"Title": "Dashboard",
	})
}

// RenderProductDetailPage merender tampilan halaman detail produk.
func (h *ViewsHandler) RenderProductDetailPage(c *gin.Context) {
	id := c.Param("id")
	c.HTML(http.StatusOK, "product_detail.tmpl", gin.H{
		"Title":     "Product Detail",
		"ProductID": id,
	})
}

func (h *ViewsHandler) RenderHelloPage(c *gin.Context) {
	// Tidak wajib mengirim data (nil atau gin.H{}) jika hanya HTML statis
	c.HTML(http.StatusOK, "hello.html", nil)
}
