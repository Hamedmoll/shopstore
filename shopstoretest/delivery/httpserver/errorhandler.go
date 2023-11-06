package httpserver

import (
	"net/http"
	"shopstoretest/pkg/richerror"
)

func kindToHttpCode(kind richerror.Kind) int {
	switch kind {

	case richerror.KindUnauthorized:

		return http.StatusUnauthorized

	case richerror.KindNotUnique:

		return http.StatusBadRequest

	case richerror.KindInvalid:

		return http.StatusBadRequest

	case richerror.KindUnexpected:

		return http.StatusInternalServerError

	case richerror.KindNotFound:

		return http.StatusNotFound

	case richerror.KindDontHaveCredit:

		return http.StatusBadRequest

	case richerror.KindForbidden:

		return http.StatusForbidden

	default:

		return http.StatusBadRequest
	}
}

func errorCodeAndMessage(err error) (int, string) {
	switch err.(type) {

	case richerror.RichError:
		r := err.(richerror.RichError)

		switch r.Kind() {

		case richerror.KindUnexpected:

			return http.StatusInternalServerError, "Internal Server Error"
		}

		return kindToHttpCode(r.Kind()), r.Message()

	default:

		return http.StatusBadRequest, "Bad Request"
	}

}
