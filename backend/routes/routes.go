package routes

import (
	"pharmacy/controllers"
	"pharmacy/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/pharmacy")
	protected.Use(middleware.AuthMiddleware())
	{
		// Drugs
		protected.POST("/drugs", controllers.CreateDrug)
		protected.GET("/drugs", controllers.GetDrugs)
		protected.PUT("/drugs/:id", controllers.UpdateDrug)
		protected.DELETE("/drugs/:id", controllers.DeleteDrug)

		// Manufacturers
		protected.POST("/manufacturers", controllers.CreateManufacturer)
		protected.GET("/manufacturers", controllers.GetManufacturers)
		protected.PUT("/manufacturers/:id", controllers.UpdateManufacturer)
		protected.DELETE("/manufacturers/:id", controllers.DeleteManufacturer)

		// Shipments
		protected.POST("/shipments", controllers.CreateShipment)
		protected.GET("/shipments", controllers.GetShipments)
		protected.PUT("/shipments/:id", controllers.UpdateShipment)
		protected.DELETE("/shipments/:id", controllers.DeleteShipment)

		// ShipmentDetails
		protected.POST("/shipment-details", controllers.CreateShipmentDetail)
		protected.GET("/shipment-details", controllers.GetShipmentDetails)
		protected.PUT("/shipment-details/:id", controllers.UpdateShipmentDetail)
		protected.DELETE("/shipment-details/:id", controllers.DeleteShipmentDetail)

		// ShipmentStatus
		protected.POST("/shipment-status", controllers.CreateShipmentStatus)
		protected.GET("/shipment-status", controllers.GetShipmentStatuses)
		protected.PUT("/shipment-status/:id", controllers.UpdateShipmentStatus)
		protected.DELETE("/shipment-status/:id", controllers.DeleteShipmentStatus)

		// DrugForms
		protected.POST("/drug-forms", controllers.CreateDrugForm)
		protected.GET("/drug-forms", controllers.GetDrugForms)
		protected.PUT("/drug-forms/:id", controllers.UpdateDrugForm)
		protected.DELETE("/drug-forms/:id", controllers.DeleteDrugForm)

		// DrugTypes
		protected.POST("/drug-types", controllers.CreateDrugType)
		protected.GET("/drug-types", controllers.GetDrugTypes)
		protected.PUT("/drug-types/:id", controllers.UpdateDrugType)
		protected.DELETE("/drug-types/:id", controllers.DeleteDrugType)
	}
}
