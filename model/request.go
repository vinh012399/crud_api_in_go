package model

type BookRequest struct {
	Title      *string  `json:"title" bson:"title"`
	Author     *string  `json:"author" bson:"author"`
	Price      *float64 `json:"price" bson:"price"`
	Is_Deleted *bool    `json:"is_deleted" bson:"is_deleted"`
	Created_By *string  `json:"created_by" bson:"created_by"`
	Updated_By *string  `json:"updated_by" bson:"updated_by"`
}
