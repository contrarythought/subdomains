package json_structs

import "time"

type CRT []struct {
	Issuer_ca_id    int       `json:"issuer_ca_id"`
	Issuer_name     string    `json:"issuer_name"`
	Common_name     string    `json:"common_name"`
	Name_value      string    `json:"name_value"`
	Id              int       `json:"id"`
	Entry_timestamp time.Time `json:"entry_timestamp"`
	Not_before      time.Time `json:"not_before"`
	Not_after       time.Time `json:"not_after"`
	Serial_number   string    `json:"serial_number"`
}
