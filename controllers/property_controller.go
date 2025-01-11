package controllers

import (
	"fmt"
	"strings"

	"rentApi/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	beego.Controller
}

// Struct to hold the response for /v1/property/list
type PropertyListResponse struct {
	Success   bool                     `json:"success"`
	Locations []LocationWithProperties `json:"locations"`
}

type LocationWithProperties struct {
	Id         int64              `json:"id"`
	DestId     string             `json:"dest_id"`
	DestType   string             `json:"dest_type"`
	Value      string             `json:"value"`
	Properties []PropertyResponse `json:"properties"`
}

type PropertyResponse struct {
	Id              int64    `json:"id"`
	PropertyId      int64    `json:"property_id"`
	PropertySlugId  string   `json:"property_slug_id"`
	HotelName       string   `json:"hotel_name"`
	Bedrooms        int      `json:"bedrooms"`
	Bathrooms       int      `json:"bathrooms"`
	GuestCount      int      `json:"guest_count"`
	Rating          float64  `json:"rating"`
	ReviewCount     int      `json:"review_count"`
	Price           string   `json:"price"`
	Breadcrumbs     []string `json:"breadcrumbs"`
	DisplayLocation []string `json:"display_location"`
	Amenities       []string `json:"amenities"`
	Type            string   `json:"type"`
	Images          []string `json:"images"`
}

// /v1/property/list
func (c *PropertyController) Get() {
	o := orm.NewOrm()
	var locations []models.Location
	_, err := o.QueryTable(new(models.Location)).All(&locations)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": fmt.Sprintf("Error fetching locations: %v", err)}
		c.ServeJSON()
		return
	}

	var response PropertyListResponse
	response.Success = true

	for _, location := range locations {
		// Fetch properties associated with this location
		var rentalProperties []models.RentalProperty
		_, err := o.QueryTable(new(models.RentalProperty)).Filter("Location__Id", location.Id).RelatedSel().All(&rentalProperties)
		if err != nil {
			fmt.Printf("Error fetching properties for location %d: %v\n", location.Id, err)
			continue
		}

		var properties []PropertyResponse
		for _, rentalProperty := range rentalProperties {
			// Fetch property details
			var propertyDetails []models.PropertyDetails
			_, err := o.QueryTable(new(models.PropertyDetails)).Filter("RentalProperty__Id", rentalProperty.Id).All(&propertyDetails)
			if err != nil {
				fmt.Printf("Error fetching property details for property %d: %v\n", rentalProperty.Id, err)
				continue
			}

			// Prepare images list
			var images []string
			for _, detail := range propertyDetails {
				if detail.ImageUrl1 != "" {
					images = append(images, detail.ImageUrl1)
				}
				if detail.ImageUrl2 != "" {
					images = append(images, detail.ImageUrl2)
				}
				if detail.ImageUrl3 != "" {
					images = append(images, detail.ImageUrl3)
				}
				if detail.ImageUrl4 != "" {
					images = append(images, detail.ImageUrl4)
				}
				if detail.ImageUrl5 != "" {
					images = append(images, detail.ImageUrl5)
				}
			}

			// Prepare breadcrumbs
			var breadcrumbs []string
			if rentalProperty.Breadcrumb1 != "" {
				breadcrumbs = append(breadcrumbs, rentalProperty.Breadcrumb1)
			}
			if rentalProperty.Breadcrumb2 != "" {
				breadcrumbs = append(breadcrumbs, rentalProperty.Breadcrumb2)
			}
			if rentalProperty.Breadcrumb3 != "" {
				breadcrumbs = append(breadcrumbs, rentalProperty.Breadcrumb3)
			}
			if rentalProperty.Breadcrumb4 != "" {
				breadcrumbs = append(breadcrumbs, rentalProperty.Breadcrumb4)
			}

			// Split DisplayLocation into parent-child relationship
			displayLocation := strings.Split(rentalProperty.DisplayLocation, ",")
			amenities := strings.Split(rentalProperty.Amenities, ", ")
			// Prepare property response
			property := PropertyResponse{
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
			properties = append(properties, property)
		}

		// Prepare location with properties response
		locationWithProperties := LocationWithProperties{
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

