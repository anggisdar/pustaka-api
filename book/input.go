package book

import (
	"encoding/json"
)

// post & binding json
type BookInput struct { // struct
	Title    string      `json:"title" binding:"required"`        // apabila required tidak terpenuhi maka sever akan berhenti
	Price    json.Number `json:"price" binding:"required,number"` // int "required,number" // json.Number support int "10" "bukanini"
	Subtitle string      `json:"sub_title"`                       // alias
}
