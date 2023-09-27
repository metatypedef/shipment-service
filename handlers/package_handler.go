package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	dto "shipment-service/dto"
	services "shipment-service/services"
)

// Calculate - Take items count and calculate packages count
func Calculate(c *gin.Context) {
	var itemsCountReq dto.PackagingCalcRequest

	if err := c.BindJSON(&itemsCountReq); err != nil {
		fmt.Println("Error: json parsing problem")
		return
	}

	var result, err = services.CalculatePackagesFor(itemsCountReq.ItemsCount)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "failed",
			"error": err.Error()})
		return
	}

	var resp = dto.PackagingCalcResponse{}
	for key := range result {
		var respItem = dto.PackagingCalcRespItems{PackageSize: key,
			PackageCount: result[key]}
		resp.Data = append(resp.Data, respItem)
	}

	c.IndentedJSON(http.StatusOK, resp)
}

// AddPackageSize - Add package size
func AddPackageSize(c *gin.Context) {
	var request dto.AddPackageSizeRequest

	if err := c.BindJSON(&request); err != nil {
		fmt.Println("Error: json parsing problem")
		return
	}

	var _, err = services.AddPackageSize(request.PackageSize)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "failed",
			"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"status": "ok"})
}

// RemovePackageSize - Remove package size
func RemovePackageSize(c *gin.Context) {
	var request dto.RemovePackageSizeRequest

	if err := c.BindJSON(&request); err != nil {
		fmt.Println("Error: json parsing problem")
		return
	}

	var _, err = services.RemovePackageSize(request.PackageSize)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "failed",
			"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}
