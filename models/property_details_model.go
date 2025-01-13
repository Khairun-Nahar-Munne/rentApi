package models


type PropertyDetailsResponse struct {
	Success  bool                `json:"success"`
	Property PropertyDetailsData `json:"property"`
}

type PropertyDetailsData struct {
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
	Description     string   `json:"description"`
	CityInTrans     string   `json:"city_in_trans"`
	Images          []string `json:"images"`
}
