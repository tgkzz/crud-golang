package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func UpdateDb(db *sql.DB, login string, fields map[string]string) error {
	if len(fields) == 0 {
		return errors.New("no fields to update")
	}

	var (
		setString []string
		args      []interface{}
	)

	i := 1
	for field, value := range fields {
		if value == "" {
			continue
		}
		if field == "age" {
			_, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
		}
		setString = append(setString, fmt.Sprintf("%s = $%d", field, i))
		args = append(args, value)
		i++
	}

	query := fmt.Sprintf(
		"UPDATE users SET %s WHERE login = $%d", strings.Join(setString, ", "),
		i,
	)

	args = append(args, login)

	_, err := db.Exec(query, args...)

	return err
}
