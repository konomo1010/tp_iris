package model

import (
	"log"
	db "tp_iris/databases"
)

// MarketEmotionSQL example.

type ME_Field struct {
	Date string `json:"date"`
	Lbs  string `json:"lbs"`
	Bd   string `json:"bd"`
	Sjzt string `json:"sjzt"`
	Sjdt string `json:"sjdt"`
}

func (this *ME_Field) Queryall() (emotion_data []ME_Field) {
	emotion_data = make([]ME_Field, 0)
	rows, err := db.MySql.Query("select * from me_2024")
	var res ME_Field
	if err != nil {
		log.Fatalln("db.MySql.Query : ", err)
	}

	defer rows.Close()

	for rows.Next() {

		rows.Scan(&res.Date, &res.Lbs, &res.Bd, &res.Sjzt, &res.Sjdt)
		emotion_data = append(emotion_data, res)
	}

	return emotion_data
}
