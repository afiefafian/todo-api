package postgres

import (
	"fmt"
)

// queryErrHandling Handling specific query error
func queryErrHandling(err error) error {
	if err != nil {
		switch err.Error() {
		case "pg: no rows in result set":
			return nil
		default:
			fmt.Println(err.Error())
			return err
		}
	}

	return nil
}
