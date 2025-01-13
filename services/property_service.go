// services/property_service.go
package services

import (
    "fmt"
    "rentApi/models"

    "github.com/beego/beego/v2/client/orm"
)

type PropertyService struct{}

func (s *PropertyService) GetLocations() ([]models.Location, error) {
    o := orm.NewOrm()
    var locations []models.Location
    _, err := o.QueryTable(new(models.Location)).All(&locations)
    if err != nil {
        return nil, fmt.Errorf("error fetching locations: %v", err)
    }
    return locations, nil
}

func (s *PropertyService) GetRentalProperties(locationID int) ([]models.RentalProperty, error) {
    o := orm.NewOrm()
    var rentalProperties []models.RentalProperty
    _, err := o.QueryTable(new(models.RentalProperty)).Filter("Location__Id", locationID).RelatedSel().All(&rentalProperties)
    if err != nil {
        return nil, fmt.Errorf("error fetching properties for location %d: %v", locationID, err)
    }
    return rentalProperties, nil
}

func (s *PropertyService) GetPropertyDetails(propertyID int64) ([]models.PropertyDetails, error) {
    o := orm.NewOrm()
    var propertyDetails []models.PropertyDetails
    _, err := o.QueryTable(new(models.PropertyDetails)).Filter("RentalProperty__Id", propertyID).All(&propertyDetails)
    if err != nil {
        return nil, fmt.Errorf("error fetching property details for property %d: %v", propertyID, err)
    }
    return propertyDetails, nil
}
func (s *PropertyService) GetRentalProperty(propertyId int64) (models.RentalProperty, error) {
    o := orm.NewOrm()
    var rentalProperty models.RentalProperty
    err := o.QueryTable(new(models.RentalProperty)).Filter("PropertyId", propertyId).RelatedSel().One(&rentalProperty)
    if err != nil {
        return rentalProperty, fmt.Errorf("error fetching property: %v", err)
    }
    return rentalProperty, nil
}