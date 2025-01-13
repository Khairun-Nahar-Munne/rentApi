// utils/property_utils.go
package utils

import (
	"rentApi/models"
	"strings"
)

func PrepareImages(propertyDetails []models.PropertyDetails) []string {
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
    return images
}

func PrepareBreadcrumbs(rentalProperty models.RentalProperty) []string {
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
    return breadcrumbs
}

func SplitDisplayLocation(displayLocation string) []string {
    return strings.Split(displayLocation, ",")
}

func SplitAmenities(amenities string) []string {
    return strings.Split(amenities, ", ")
}