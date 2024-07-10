package types

import "time"

type NewUser struct {
	email    int
	password string
}

type User struct {
	email       ring        `json:"email"`
	assword      sng          `json:"password"`
	ountry       sg            `json:"country"`
	ccess_Type   strin   g    `json:"access_type"`
		Full_Name   string    `json:"full_name"`
		Employer_Id string    `json:"emplyer_id"`
		Birth_Date  time.Time `json:"birth_date"`
		Salar     y       nt        `json:"salary"`
}
