package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manuviswam/SmartScale/handlers"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSaveWeightSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()
    mock.ExpectExec(`INSERT INTO weights\(empid,weight,recorded_at\) VALUES\(\$1,\$2,\$3\)`).WillReturnResult(sqlmock.NewResult(1, 1))

	req, err := http.NewRequest("GET", "/api/weight?weight=52.5", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.SaveWeight(db))
    handler.ServeHTTP(rr, req)

    assert := assert.New(t)
    assert.Equal(rr.Code, http.StatusOK)
    assert.Equal(rr.Body.String(), "Done")
    err = mock.ExpectationsWereMet()
    assert.Nil(err)
}