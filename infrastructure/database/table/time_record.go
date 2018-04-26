package table

import (
	"fmt"
	"time"

	"github.com/devishot/grpc-go-time_tracking/infrastructure/database"
)

const TimeRecordTableName = "time_record"

type TimeRecordTable struct {
	DB   *database.DB
	Name string
}

type TimeRecordRow struct {
	ID          string
	Amount      int32
	Timestamp   time.Time
	Description string
	UserID      string //TODO: not implemented; add reference in DB
	ProjectID   string //TODO: not implemented; add reference type
}

func NewTimeRecordTable(db *database.DB) (*TimeRecordTable, error) {
	t := &TimeRecordTable{DB: db, Name: TimeRecordTableName}

	err := t.createTable()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *TimeRecordTable) createTable() error {
	const q = `
CREATE TABLE IF NOT EXISTS time_record (
	id uuid PRIMARY KEY,
	amount INT NOT NULL,
	timestamp timestamptz DEFAULT current_timestamp,
	description text NOT NULL
)`
	if _, err := t.DB.Conn.Exec(q); err != nil {
		return fmt.Errorf("when: create table | table: TimeRecordTable | error: %s", err.Error())
	}
	return nil
}

func (t *TimeRecordTable) Insert(row TimeRecordRow) (newRow TimeRecordRow, err error) {
	const q = `
INSERT INTO time_record (
	id, amount, description
)
VALUES (
	$1, $2, $3
)
RETURNING
	id, amount, timestamp, description
`
	err = t.DB.Conn.QueryRow(q, row.ID, row.Amount, row.Description).
		Scan(&newRow.ID, &newRow.Amount, &newRow.Timestamp, &newRow.Description)
	if err != nil {
		return newRow, fmt.Errorf("when: insert row | table: TimeRecordTable | error: %s", err.Error())
	}

	return
}

func (t *TimeRecordTable) Delete(id string) (err error) {
	const q = `
DELETE FROM time_record
WHERE id = $1
`
	if _, err := t.DB.Conn.Exec(q, id); err != nil {
		return fmt.Errorf("when: delete row | table: TimeRecordTable | error: %s", err.Error())
	}
	return nil
}

func (t *TimeRecordTable) FindByID(id string) (newRow TimeRecordRow, err error) {
	const q = `
SELECT
	id, amount, timestamp, description
FROM
	time_record
WHERE
	id = $1
`
	err = t.DB.Conn.QueryRow(q, id).
		Scan(&newRow.ID, &newRow.Amount, &newRow.Timestamp, &newRow.Description)
	if err != nil {
		return newRow, fmt.Errorf("when: find by id | table: TimeRecordTable | error: %s", err.Error())
	}

	return
}
