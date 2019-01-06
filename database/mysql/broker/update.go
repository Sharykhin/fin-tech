package broker

import (
	"context"
	"fmt"
	"strings"

	"github.com/Sharykhin/fin-tech/request"
)

func (s Storage) Update(ctx context.Context, ID int64, ubr request.UpdateBrokerRequest) error {
	query := "UPDATE brokers SET "
	var replacements []interface{}

	if ubr.Role.Valid {
		query += "role=?, "
		replacements = append(replacements, ubr.Role)
	}

	if ubr.Position.Valid {
		query += "position=?, "
		replacements = append(replacements, ubr.Position)
	}

	if len(replacements) == 0 {
		return nil
	}

	query = strings.TrimRight(query, ", ")
	query += " WHERE user_id=?"
	replacements = append(replacements, ID)
	fmt.Println(query, replacements)
	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("could not create a prepared statement: %v", err)
	}

	result, err := stmt.Exec(replacements...)
	if err != nil {
		return fmt.Errorf("could not execute a prepated statement: %v", err)
	}

	fmt.Println("update broker succeed")
	lastID, err := result.LastInsertId()
	fmt.Println(lastID, err)
	rowsAffected, err := result.RowsAffected()
	fmt.Println(rowsAffected, err)

	return nil
}
