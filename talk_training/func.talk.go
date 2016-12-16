package talk_training

import (
	"fmt"
)

func GetTalks(productID int64) ([]Talks, error) {
	talks := []Talks{}
	query := fmt.Sprintf(
		`
        SELECT
            talk_id,
            product_id,
            message,
            create_time
        FROM
            ws_talk
        WHERE
            product_id = %d
        `, productID)

	rows, err := MainDB.Query(query)
	if err != nil {
		return talks, err
	}
	defer rows.Close()
	for rows.Next() {
		t := Talks{}
		errScan := rows.Scan(&t.ID,
			&t.ProductID,
			&t.Message,
			&t.CreateTime)
		if errScan != nil {
			return talks, errScan
		}
		talks = append(talks, t)
	}
	return talks, nil
}
