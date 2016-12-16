package talk_training

import (
	"fmt"
	"log"
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

func InsertTalk(userID, productID, shopID int64, message string) error {
	query := fmt.Sprintf(
		`
        INSERT INTO ws_talk(
        talk_id,
        shop_id,
        product_id,
        user_id,
        status,
        message,
        total_comment,
        create_by,
        create_time
    )
    VALUES
        (
            nextval('ws_talk_talk_id_seq'),
            %d,
            %d,
            %d,
            1,
            '%s',
            0,
            1032,
            now()
        ) RETURNING talk_id
    `,
		shopID,
		productID,
		userID,
		message)
	var talkID int64
	err := MainDB.DB.QueryRow(query).Scan(&talkID)
	if err != nil {
		return err
	}
	log.Println(talkID)
	return nil
}
