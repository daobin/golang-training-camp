package dto

type Dto struct {
	ID int `json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn int `json:"deleted_on"`
}
