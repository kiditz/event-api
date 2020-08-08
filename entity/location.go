package entity

// Location used to save data of user location
type Location struct {
	ID                       uint    `json:"id" gorm:"primary_key"`
	Latitude                 float64 `json:"lat" validate:"required"`
	Longitude                float64 `json:"lng" validate:"required"`
	FormattedAddress         string  `json:"formatted_address" validate:"required"`
	Address                  *string `json:"address"`
	PhoneNumber              *string `json:"phone_number"`
	InternationalPhoneNumber *string `json:"international_phone_number"`
	PlaceID                  *string `json:"place_id"`
	Icon                     *string `json:"icon"`
	Name                     *string `json:"name"`
}
