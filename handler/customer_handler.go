package handler

import "github.com/gin-gonic/gin"

type CustomerHandler struct {
}

// GetCustomer godoc
// @Summary Get Customer
// @Description  Get customer by ID
// @Security ApiKeyAuth
// @ID GetCustomer
// @Accept json
// @Produce json
// @Param id path int true "ID of customer to be gotten"
// @Success 200 {object} model.CustomerResponse "OK"
// @Success 400 {object} model.Response "Bad Request"
// @Success 401 {object} model.Response "Unauthorized"
// @Success 409 {object} model.Response "Conflict"
// @Success 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/customers/:id [get]
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	//Logic goes here
}

// ListCustomers godoc
// @Summary List Customers
// @Description List all customers
// @Security ApiKeyAuth
// @ID ListCustomers
// @Accept json
// @Produce json
// @Success 200 {object} model.CustomersResponse "OK"
// @Success 400 {object} model.Response "Bad Request"
// @Success 401 {object} model.Response "Unauthorized"
// @Success 409 {object} model.Response "Conflict"
// @Success 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/customers [get]
func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	//Logic goes here
}

// CreateCustomer godoc
// @Summary Create Customer
// @Description Create new customer
// @Security ApiKeyAuth
// @ID CreateCustomer
// @Accept json
// @Produce json
// @Param Customer body model.CustomerForCreate true "Customer data to be created"
// @Success 200 {object} model.Response "OK"
// @Success 400 {object} model.Response "Bad Request"
// @Success 401 {object} model.Response "Unauthorized"
// @Success 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/customers [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	//Logic goes here
}

// DeleteCustomer godoc
// @Summary Delete Customer
// @Description Delete customer by ID
// @Security ApiKeyAuth
// @ID DeleteCustomer
// @Accept json
// @Produce json
// @Param id path int true "ID of customer to be deleted"
// @Success 200 {object} model.Response "OK"
// @Success 400 {object} model.Response "Bad Request"
// @Success 401 {object} model.Response "Unauthorized"
// @Success 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/customers/:id [delete]
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	//Logic goes here
}

// UpdateCustomer godoc
// @Summary Update Customer
// @Description Update customer by ID
// @Security ApiKeyAuth
// @ID UpdateCustomer
// @Accept json
// @Produce json
// @Param id path int true "ID of customer to be updated"
// @Param Customer body model.CustomerForUpdate true "Customer data to be updated"
// @Success 200 {object} model.Response "OK"
// @Success 400 {object} model.Response "Bad Request"
// @Success 401 {object} model.Response "Unauthorized"
// @Success 500 {object} model.Response "Internal Server Error"
// @Router /api/v1/customers/:id [patch]
func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	//Logic goes here
}
