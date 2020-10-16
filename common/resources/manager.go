package resources

import (
	_ "github.com/go-sql-driver/mysql"
)

func ReadResource(file string) (string, error) {
	c, err := FSByte(false, file)
	if err != nil {
		return "", err
	}
	return string(c), nil
}
