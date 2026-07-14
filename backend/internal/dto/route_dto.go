package dto

type CreateRouteRequest struct {
	OriginCity      string `json:"origin_city" binding:"required"`
	OriginPool      string `json:"origin_pool" binding:"required"`
	DestinationCity string `json:"destination_city" binding:"required"`
	DestinationPool string `json:"destination_pool" binding:"required"`
}
