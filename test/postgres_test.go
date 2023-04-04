package test

import (
	"dcard_2023_bk/pkg/postgres"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertPage(t *testing.T) {
	r := setupRouter()
	r.GET("/mock", nil)
	defer teardown()

	expectId := "exist"

	expectData := "{}"

	mockPostgresServer.ExpectExec(regexp.QuoteMeta(`INSERT INTO "pages"("id", "data") VALUES($1, $2)`)).WithArgs(expectId, expectData).WillReturnResult(sqlmock.NewResult(1, 1))

	err := postgres.InsertPage(expectId, expectData)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, err, nil)
}

func TestGetDBPage(t *testing.T) {
	r := setupRouter()
	r.GET("/mock", nil)
	defer teardown()

	expectId := "exist"

	expectData := "{}"

	mockPostgresServer.ExpectQuery("^SELECT .*").WithArgs(expectId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "data"}).
			AddRow(expectId, expectData))

	data, err := postgres.GetDBPage(expectId)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, string(data), expectData)
}
