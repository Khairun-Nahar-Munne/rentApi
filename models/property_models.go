// Struct to hold the response for /v1/property/list
package models

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
