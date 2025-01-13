// controllers/property_details_controller.go
package controllers

import (
    "rentApi/models"
    "rentApi/services"
    "rentApi/utils"

    "github.com/beego/beego/v2/server/web"
)

type PropertyDetailsController struct {
    web.Controller
}

func (c *PropertyDetailsController) Get() {
    // Get property_id from query parameters
    propertyId, err := c.GetInt64("property_id")
    if err != nil {
        c.Data["json"] = map[string]interface{}{"error": "Invalid property_id"}
        c.ServeJSON()
        return
    }

    propertyService := services.PropertyService{}

    rentalProperty, err := propertyService.GetRentalProperty(propertyId)
    if err != nil {
        c.Data["json"] = map[string]interface{}{"error": err.Error()}
        c.ServeJSON()
        return
    }

    propertyDetails, err := propertyService.GetPropertyDetails(rentalProperty.Id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{"error": err.Error()}
        c.ServeJSON()
        return
    }

    images := utils.PrepareImages(propertyDetails)
    breadcrumbs := utils.PrepareBreadcrumbs(rentalProperty)
    displayLocation := utils.SplitDisplayLocation(rentalProperty.DisplayLocation)
    amenities := utils.SplitAmenities(rentalProperty.Amenities)

    property := models.PropertyDetailsData{
        Id:              rentalProperty.Id,
        PropertyId:      rentalProperty.PropertyId,
        PropertySlugId:  rentalProperty.PropertySlugId,
        HotelName:       rentalProperty.HotelName,
        Bedrooms:        rentalProperty.Bedrooms,
        Bathrooms:       rentalProperty.Bathrooms,
        GuestCount:      rentalProperty.GuestCount,
        Rating:          rentalProperty.Rating,
        ReviewCount:     rentalProperty.ReviewCount,
        Price:           rentalProperty.Price,
        Breadcrumbs:     breadcrumbs,
        DisplayLocation: displayLocation,
        Amenities:       amenities,
        Type:            rentalProperty.Type,
        CityInTrans:     propertyDetails[0].CityInTrans,
    }

    if len(propertyDetails) > 0 {
        property.Description = propertyDetails[0].Description
    }

    property.Images = images

    response := models.PropertyDetailsResponse{
        Success:  true,
        Property: property,
    }

    c.Data["json"] = response
    c.ServeJSON()
}