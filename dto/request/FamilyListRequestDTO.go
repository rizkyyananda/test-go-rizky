package request

type FamilyListRequestDTO struct {
	ID         uint   `json:"id"`
	CstID      uint   `json:"cst_id" binding:"required"`
	FlRelation string `json:"fl_relation" binding:"required"`
	FlName     string `json:"fl_name" binding:"required"`
	FlDOB      string `json:"fl_dob" binding:"required"`
}
