package dto

type CreateFleetRequest struct {
	Name        string `json:"name" binding:"required"`
	PlateNumber string `json:"plate_number" binding:"required"`
	Type        string `json:"type" binding:"required,oneof=Reguler Premium"`
	TotalSeats  int    `json:"total_seats" binding:"required,min=1"`
}
