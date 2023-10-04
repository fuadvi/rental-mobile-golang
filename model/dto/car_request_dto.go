package dto

type CarRequestDto struct {
	LEASETYPEID int    `json:"lease_type_id"`
	TITLE       string `json:"title" validate:"required"`
	PRICE       int    `json:"price" validate:"required,number"`
	DURATION    string `json:"duration" validate:"required"`
	IMGURL      string `json:"img_url"`
	DESCRIPTION string `json:"description" validate:"required"`
	PASSENGER   int8   `json:"passenger" validate:"required,number"`
	LUGGAGE     int8   `json:"luggage" validate:"required,number"`
	CARTYPE     string `json:"car_type" validate:"required"`
	ISDRIVER    bool   `json:"is_driver" validate:"required,boolean"`
}
