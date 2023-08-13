package Database

import (
	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	Link string `json:"link"`
}

func Find(token string) (*Tag, error) {
	query := "SELECT link FROM url WHERE token = ?"
	var tag Tag
	err := db.QueryRow(query, token).Scan(&tag.Link)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}
