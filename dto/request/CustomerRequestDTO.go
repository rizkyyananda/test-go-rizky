package request

type CustomerRequestDTO struct {
	ID            uint                   `json:"id"`
	NationalityID uint                   `json:"nationality_id" binding:"required"`
	CstName       string                 `json:"cst_name" binding:"required"`
	CstDOB        string                 `json:"cst_dob" binding:"required"`
	CstPhone      string                 `json:"cst_phone_number"`
	CstEmail      string                 `json:"cst_email"`
	Family        []FamilyListRequestDTO `json:"family"`
}
