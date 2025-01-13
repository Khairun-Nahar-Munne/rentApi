// controllers/property_controller.go
package controllers

import (
    "rentApi/models"
    "rentApi/services"
    "rentApi/utils"
    "sync"

    beego "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
    beego.Controller
}

// /v1/property/list
func (c *PropertyController) Get() {
    propertyService := services.PropertyService{}
    locations, err := propertyService.GetLocations()
    if err != nil {
        c.Data["json"] = map[string]interface{}{"error": err.Error()}
        c.ServeJSON()
        return
    }

    var response models.PropertyListResponse
    response.Success = true

    var wg sync.WaitGroup
    locationChan := make(chan models.LocationWithProperties, len(locations))

    for _, location := range locations {
        wg.Add(1)
        go func(location models.Location) {
            defer wg.Done()
            rentalProperties, err := propertyService.GetRentalProperties(int(location.Id))
            if err != nil {
                return
            }

            var properties []models.PropertyResponse
            var propertyWg sync.WaitGroup
            propertyChan := make(chan models.PropertyResponse, len(rentalProperties))

            for _, rentalProperty := range rentalProperties {
                propertyWg.Add(1)
                go func(rentalProperty models.RentalProperty) {
                    defer propertyWg.Done()
                    propertyDetails, err := propertyService.GetPropertyDetails(rentalProperty.Id)
                    if err != nil {
                        return
                    }

                    images := utils.PrepareImages(propertyDetails)
                    breadcrumbs := utils.PrepareBreadcrumbs(rentalProperty)
                    displayLocation := utils.SplitDisplayLocation(rentalProperty.DisplayLocation)
                    amenities := utils.SplitAmenities(rentalProperty.Amenities)

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
                        Type:            rentalProperty.Type,
                        Images:          images,
                    }
                    propertyChan <- property
                }(rentalProperty)
            }

            propertyWg.Wait()
            close(propertyChan)

            for property := range propertyChan {
                properties = append(properties, property)
            }

            locationWithProperties := models.LocationWithProperties{
                Id:         location.Id,
                DestId:     location.DestId,
                DestType:   location.DestType,
                Value:      location.Value,
                Properties: properties,
            }
            locationChan <- locationWithProperties
        }(location)
    }

    wg.Wait()
    close(locationChan)

    for locationWithProperties := range locationChan {
        response.Locations = append(response.Locations, locationWithProperties)
    }

    c.Data["json"] = response
    c.ServeJSON()
}