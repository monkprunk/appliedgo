package merchant

import (
	"Advance-Golang-Programming/advanced/final/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-delve/delve/pkg/config"
	log "github.com/sirupsen/logrus"
)

type merchantService interface {
	Register(string, string) (int, error)
	Information(int) (Merchant, error)
	Add(int, string, int, int) (int, error)
	Find(int) ([]product.Product, error)
}

type Handler struct {
	conf        *config.Config
	merchantSrv merchantService
}

func NewHandler(srv merchantService) *Handler {
	return &Handler{merchantSrv: srv}
}

var runid = 0

func (h Handler) Register(c *gin.Context) {
	//repo := MemoryRepo{}
	//serv := NewHandler(repo)

	var item Merchant
	err := c.BindJSON(&item)
	if err != nil {
		log.Println("merchant register BindJSON error:", err)
		return
	}
	item.ID, err = h.merchantSrv.Register(item.Name, item.BankAccount)
	if err != nil {
		log.Println("merchant insert error:", err)
		return
	}
	log.Println("register merchant success")
	c.JSON(http.StatusOK, gin.H{"id": item.ID})
}

func (h Handler) Information(c *gin.Context) {
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("merchant information strconv:", err)
		c.JSON(http.StatusProcessing, err.Error())
		return
	}

	//repo := MemoryRepo{}
	//serv := NewHandler(repo)

	m, err := h.merchantSrv.Information(pId)
	if err != nil {
		log.Println("merchant information error:", err)
		c.JSON(http.StatusProcessing, err.Error())
		return
	}
	log.Printf("merchant: %+v", m)
	c.JSON(http.StatusOK, gin.H{"id": m.ID, "name": m.Name, "bankaccount": m.BankAccount})
}

func (h Handler) ListAllProduct(c *gin.Context) {
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("product information strconv:", err)
		c.JSON(http.StatusProcessing, err.Error())
		return
	}

	//repo := MemoryRepo{}
	//serv := NewHandler(repo)

	products, err := h.merchantSrv.Find(pId)
	if err != nil {
		log.Println("product information error:", err)
		c.JSON(http.StatusProcessing, err.Error())
		return
	}
	log.Printf("products: %+v", products)

	c.JSON(http.StatusOK, products)
}

var runidproduct = 0

func (h Handler) AddProduct(c *gin.Context) {
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("product information strconv:", err)
		c.JSON(http.StatusProcessing, err.Error())
		return
	}

	//repo := MemoryRepo{}
	//serv := NewHandler(repo)

	var item product.Product
	err = c.BindJSON(&item)
	if err != nil {
		log.Println("product register BindJSON error:", err)
		return
	}

	item.ID, err = h.merchantSrv.Add(pId, item.Name, item.Amount, item.Stock)
	if err != nil {
		log.Println("product register error:", err)
		return
	}
	log.Println("register product success")
	c.JSON(http.StatusOK, gin.H{"id": item.ID})
}

//==========================================================
