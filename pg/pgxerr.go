package pg

import (
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func TranslateReadError(err error) error {
	if err == nil {
		return nil
	}

	if pgxscan.NotFound(err) {
		return NewErrorWithCodeMessage(ErrorEmptyCode, err.Error())
	}

	return NewErrorWithCodeMessage(ErrorInternalCode, err.Error())
}

func TranslateReadErrorFunc() func(error) error {
	return func(err error) error {
		return TranslateReadError(err)
	}
}

func TranslateWriteError(err error) error {
	if err == nil {
		return nil
	}

	pgError, pgOK := err.(*pgconn.PgError)
	if pgOK && pgError.Code == pgerrcode.UniqueViolation {
		return NewErrorWithCodeMessage(ErrorUniqueCode, err.Error())
	}

	if pgOK && pgerrcode.IsIntegrityConstraintViolation(pgError.Code) {
		return NewErrorWithCodeMessage(ErrorIntegrityCode, err.Error())
	}

	return NewErrorWithCodeMessage(ErrorInternalCode, err.Error())
}

func TranslateWriteErrorFunc() func(error) error {
	return func(err error) error {
		return TranslateWriteError(err)
	}
}
