package response

import "time"

type FamilyListResponseDTO struct {
	FlID       uint      `json:"fl_id"`
	CstID      uint      `json:"cst_id"`
	FlRelation string    `json:"fl_relation"`
	FlName     string    `json:"fl_name"`
	FlDOB      time.Time `json:"fl_dob"`
}
