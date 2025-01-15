// controllers/property_controller.go
package controllers

import (
    "rentApi/models"
    "rentApi/services"
    "rentApi/utils"

    beego "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
    beego.Controller
}

// /v1/property/list
func (c *PropertyController) Get() {
    propertyService := services.PropertyService{}
    
    // Get locations
    locations, err := propertyService.GetLocations()
    if err != nil {
        c.Data["json"] = map[string]interface{}{"error": err.Error()}
        c.ServeJSON()
        return
    }

    var response models.PropertyListResponse
    response.Success = true
    response.Locations = make([]models.LocationWithProperties, 0)

    // Process each location serially
    for _, location := range locations {
        // Get rental properties for current location
        rentalProperties, err := propertyService.GetRentalProperties(int(location.Id))
        if err != nil {
            continue
        }

        properties := make([]models.PropertyResponse, 0)

        // Process each rental property serially
        for _, rentalProperty := range rentalProperties {
            // Get property details
            propertyDetails, err := propertyService.GetPropertyDetails(rentalProperty.Id)
            if err != nil {
                continue
            }

            // Prepare property data
            images := utils.PrepareImages(propertyDetails)
            breadcrumbs := utils.PrepareBreadcrumbs(rentalProperty)
            displayLocation := utils.SplitDisplayLocation(rentalProperty.DisplayLocation)
            amenities := utils.SplitAmenities(rentalProperty.Amenities)

            // Create property response
            property := models.PropertyResponse{
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
                Type:           rentalProperty.Type,
                Images:         images,
            }

            properties = append(properties, property)
        }

        // Create location with properties
        locationWithProperties := models.LocationWithProperties{
            Id:         location.Id,
            DestId:     location.DestId,
            DestType:   location.DestType,
            Value:      location.Value,
            Properties: properties,
        }

        response.Locations = append(response.Locations, locationWithProperties)
    }

    c.Data["json"] = response
    c.ServeJSON()
}