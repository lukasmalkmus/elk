// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func URLParamInt(r *http.Request, p string) (int, error) {
	id, err = strconv.Atoi(chi.URLParam(r, p))
	if err != nil {
		return 0, err
	}

	return id, nil
}
