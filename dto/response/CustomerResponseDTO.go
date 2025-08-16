package response

import "time"

type CustomerResponseDTO struct {
	CstID         uint                    `json:"cst_id"`
	NationalityID uint                    `json:"nationality_id"`
	CstName       string                  `json:"cst_name"`
	CstDOB        time.Time               `json:"cst_dob"`
	CstPhone      string                  `json:"cst_phone_number"`
	CstEmail      string                  `json:"cst_email"`
	Nationality   *NationalityResponseDTO `json:"nationality,omitempty"`
	FamilyList    []FamilyListResponse    `json:"family_list,omitempty"`
}
