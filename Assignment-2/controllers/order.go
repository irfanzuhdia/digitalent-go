package controllers

import (
	"assignment-2/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (c *Controller) CreateOrder(ctx *gin.Context) {
	var newOrder models.Order

	// Attempt to parse JSON data from the request body
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Parsing JSON Data: %v", err),
		})
		return
	}

	// Set the ordered_at field to the current time
	newOrder.OrderedAt = time.Now()

	// Attempt to create the order in the database
	if err := c.DB.Create(&newOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Creating Order Data: %v", err),
		})
		return
	}

	// Respond with the created order
	ctx.JSON(http.StatusCreated, gin.H{
		"result": newOrder,
	})
}

func (c *Controller) GetOrders(ctx *gin.Context) {
	var (
		orders []models.Order
	)

	if err := c.DB.Model(&models.Order{}).Preload("Items").Order("id asc").Find(&orders).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Getting Order Data: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": orders,
	})
}

func (c *Controller) UpdateOrder(ctx *gin.Context) {
	// Extract order ID from the request URL
	orderId := ctx.Param("orderId")

	// Create a variable to hold the updated order data
	var updatedOrder models.Order

	// Attempt to parse JSON data from the request body
	if err := ctx.ShouldBindJSON(&updatedOrder); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Parsing JSON Data: %v", err),
		})
		return
	}

	// Check if the order with the given ID exists in the database
	var existingOrder models.Order
	if err := c.DB.First(&existingOrder, "id = ?", orderId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v Not Found", orderId),
		})
		return
	}

	// Update the order data in the database
	if err := c.DB.Model(&models.Order{}).Where("id = ?", orderId).Updates(updatedOrder).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Updating Order Data: %v", err.Error()),
		})
		return
	}

	// Respond with a success message
	ctx.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("Order with id %v Has Been Successfully Updated", orderId),
	})
}

func (c *Controller) DeleteOrder(ctx *gin.Context) {
	var orderId = ctx.Param("orderId")

	// Check if the order exists
	var existingOrder models.Order
	if err := c.DB.First(&existingOrder, "id = ?", orderId).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"result": fmt.Sprintf("Order with id %v Not Found", orderId),
		})
		return
	}

	// Delete associated items first
	if err := c.DB.Where("order_id = ?", orderId).Delete(&models.Item{}).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Deleting Items Associated with Order: %v", err.Error()),
		})
		return
	}

	// Delete the order after deleting associated items
	if err := c.DB.Where("id = ?", orderId).Delete(&models.Order{}).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"result": fmt.Sprintf("Error Deleting Order: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": fmt.Sprintf("Order with id %v Has Been Successfully Deleted", orderId),
	})
}
