package dto

type CreateFleetRequest struct {
	Name        string `json:"name" binding:"required"`
	PlateNumber string `json:"plate_number" binding:"required"`
	Type        string `json:"type" binding:"required,oneof=Reguler Premium"`
	TotalSeats  int    `json:"total_seats" binding:"required,min=1"`
}

type UpdateFleetRequest struct {
	Name         *string `json:"name"`
	PlateNumber  *string `json:"plate_number"`
	Type         *string `json:"type" binding:"omitempty,oneof=Reguler Premium"`
	TotalSeats   *int    `json:"total_seats" binding:"omitempty,min=1"`
	ActiveStatus *bool   `json:"active_status"`
}
