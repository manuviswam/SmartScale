package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"errors"
	"fmt"

	"github.com/manuviswam/SmartScale/handlers"
	m "github.com/manuviswam/SmartScale/models"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type StubEmployeeGetter struct {}
func (eg *StubEmployeeGetter) GetEmployeeFromInternalNumber(n int64) (m.Employee, error) {
	return m.Employee{ 
		EmpId: "12345",
		EmployeeName: "Name",
		InternalNumber: string(n),
	}, nil
}

type StubEmployeeGetterFail struct {}
func (eg *StubEmployeeGetterFail) GetEmployeeFromInternalNumber(n int64) (m.Employee, error) {
	return m.Employee{}, errors.New("failed")
}

func TestSaveWeightSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()
    rows := sqlmock.NewRows([]string{"weight", "recorded_at"}).
        AddRow(1.2, "2016-08-05T11:26:00.579423Z").
        AddRow(2.1, "2016-08-05T11:27:00.579423Z")
    mock.ExpectExec(`INSERT INTO weights\(empid,weight,recorded_at\) VALUES\(\$1,\$2,\$3\)`).WithArgs("12345",52.5,sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
    mock.ExpectQuery(`WITH t AS \(SELECT weight,recorded_at FROM weights WHERE empid=\$1 ORDER BY recorded_at DESC LIMIT 10\) SELECT \* FROM t ORDER BY recorded_at ASC`).WithArgs("12345").WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/api/weight?weight=52.5&internalNumber=1111", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.SaveWeight(db, &StubEmployeeGetter{}))
    handler.ServeHTTP(rr, req)

    assert := assert.New(t)
    assert.Equal(http.StatusOK, rr.Code)
    err = mock.ExpectationsWereMet()
    assert.Nil(err)
}

func TestSaveWeightFailsForInvalidWeight(t *testing.T) {
	db, _, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

	req, err := http.NewRequest("GET", "/api/weight?weight=foo&internalNumber=1111", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.SaveWeight(db, &StubEmployeeGetter{}))
    handler.ServeHTTP(rr, req)

    assert := assert.New(t)
    assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestSaveWeightFailsForInvalidInternalNumber(t *testing.T) {
	db, _, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

	req, err := http.NewRequest("GET", "/api/weight?weight=52.5&internalNumber=foo", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.SaveWeight(db, &StubEmployeeGetter{}))
    handler.ServeHTTP(rr, req)

    assert := assert.New(t)
    assert.Equal(http.StatusBadRequest, rr.Code)
}

func TestSaveWeightFailsWhenGetEmployeeFails(t *testing.T) {
	db, _, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

	req, err := http.NewRequest("GET", "/api/weight?weight=52.5&internalNumber=1111", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.SaveWeight(db, &StubEmployeeGetterFail{}))
    handler.ServeHTTP(rr, req)

    assert := assert.New(t)
    assert.Equal(http.StatusInternalServerError, rr.Code)
}

func TestSaveWeightFailsWhenDBOperationFails(t *testing.T) {
	db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()
    mock.ExpectExec(`INSERT INTO weights\(empid,weight,recorded_at\) VALUES\(\$1,\$2,\$3\)`).WithArgs("12345",52.5,sqlmock.AnyArg()).WillReturnError(fmt.Errorf("some error"))

	req, err := http.NewRequest("GET", "/api/weight?weight=52.5&internalNumber=1111", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.SaveWeight(db, &StubEmployeeGetter{}))
    handler.ServeHTTP(rr, req)

    assert := assert.New(t)
    assert.Equal(http.StatusInternalServerError, rr.Code)
    err = mock.ExpectationsWereMet()
    assert.Nil(err)
}