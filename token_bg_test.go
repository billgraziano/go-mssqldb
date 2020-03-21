package mssql

import (
	"database/sql"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrinting(t *testing.T) {
	assert := assert.New(t)
	var err error
	query := url.Values{}
	query.Add("app name", "MyAppName")
	query.Add("log", "3")

	u := &url.URL{
		Scheme: "sqlserver",
		//User:   url.UserPassword(username, password),
		//Host:   fmt.Sprintf("%s:%d", hostname, port),
		Host:     "D40",
		Path:     "SQL2017", // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())
	assert.NoError(err)

	_, err = db.Exec("PRINT 'This line was printed'; ")
	assert.NoError(err)

	_, err = db.Exec("ZYZ")
	assert.Error(err)

	_, err = db.Exec("PRINT 'Starting #1...'; WAITFOR DELAY '00:00:02'; PRINT 'Done';")
	_, err = db.Exec("RAISERROR('Starting #2 (RAISERROR)...', 0, 1) WITH NOWAIT; WAITFOR DELAY '00:00:02'; PRINT 'Done';")
	assert.NoError(err)

	_, err = db.Exec("SET XACT_ABORT ON; PRINT 'Before the error...'; RAISERROR('From RAISERROR!', 16, 1); PRINT 'After the error';")
	assert.Error(err)

	_, err = db.Exec("PRINT 'one'\r\nGO\r\nTWO\r\nGO\r\n")
	assert.Error(err)
}
