package rest_api

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/volatiletech/sqlboiler/boil"
	"io"
	"net/http"
)

func unMarshalIntoModel(r io.Reader, i SqlBoilerModel) error {
	d := json.NewDecoder(r)
	d.DisallowUnknownFields()
	if err := d.Decode(i); err != nil {
		return err
	}
	return nil
}

func writeRequestError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusBadRequest)
	r := map[string]interface{}{
		"error": e.Error(),
	}
	str, _ := json.Marshal(r)
	w.Write(str)
}

type SqlBoilerModel interface {
	InsertG(ctx context.Context, columns boil.Columns) error
}

func ModelCreate(i SqlBoilerModel, r io.Reader) error {
	if err := unMarshalIntoModel(r, i); err != nil {
		return err
	}
	if err := i.InsertG(context.Background(), boil.Blacklist("id")); err != nil {
		return err
	}
	return nil
}

func ModelCreateAndMarshal(m SqlBoilerModel, b io.Reader) ([]byte, error) {
	if err := ModelCreate(m, b); err != nil {
		return nil, err
	}
	if str, err := json.Marshal(m); err != nil {
		return nil, err
	} else {
		return str, nil
	}
}

func CreateHandler(m SqlBoilerModel, w http.ResponseWriter, r *http.Request) {
	if b, err := ModelCreateAndMarshal(m, r.Body); err != nil {
		writeRequestError(w, err)
	} else {
		w.WriteHeader(201)
		w.Write(b)
	}
}

func HandleFind(m SqlBoilerModel, e error, w http.ResponseWriter) {
	switch e {
	case nil:
		b, _ := json.Marshal(m)
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	case sql.ErrNoRows:
		w.WriteHeader(http.StatusNotFound)
		r, _ := json.Marshal(map[string]interface{}{
			"error": "Not found.",
		})
		w.Write(r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		r, _ := json.Marshal(map[string]interface{}{
			"error": e.Error(),
		})
		w.Write(r)
	}
}