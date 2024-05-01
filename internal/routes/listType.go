package routes

type List struct {
	ID int32 `json:"ID"`
	BorrowerName string `json:"BorrowerName"`
	Books Book `json:"Books"`
}
