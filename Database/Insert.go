package Database

import (
	"fmt"
)

func Insert(link, token string) error {
	query := "INSERT INTO url (link, token) VALUES (?, ?)"
	result, err := db.Exec(query, link, token)
	if err != nil {
		return err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("%d Added!\n", rowCount)
	return nil
}
