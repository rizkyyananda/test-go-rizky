package response

type NationalityResponseDTO struct {
	NationalityID   uint   `json:"nationality_id"`
	NationalityName string `json:"nationality_name"`
	NationalityCode string `json:"nationality_code"`
}
