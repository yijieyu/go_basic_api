package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
)

var (
	NotFoundError       = errors.New("not found")
	DuplicateEntryError = errors.New("Duplicate entry")
	LockErr             = errors.New("lock fail")
	UnlockErr           = errors.New("unlock fail")
)

func Error(err error) error {
	switch err {
	case sql.ErrNoRows:
		return NotFoundError
	}

	switch sqlErr := err.(type) {
	case *mysql.MySQLError:
		switch sqlErr.Number {
		// 主键或唯一索引冲突错误
		case 1062:
			return fmt.Errorf("Error %d: %w%s", sqlErr.Number, DuplicateEntryError, strings.TrimPrefix(sqlErr.Message, "Duplicate entry"))
		}
	}

	return err
}
