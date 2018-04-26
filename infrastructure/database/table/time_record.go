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
	UserID      string
	ProjectID   string
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
	description text NOT NULL,
  client_project_id uuid REFERENCES client_project(id) ON DELETE CASCADE,
  user_id uuid REFERENCES app_user(id) ON DELETE CASCADE
)`
	if _, err := t.DB.Conn.Exec(q); err != nil {
		return fmt.Errorf("when: create table | table: TimeRecordTable | error: %s", err.Error())
	}
	return nil
}

func (t *TimeRecordTable) Insert(row TimeRecordRow) (newRow TimeRecordRow, err error) {
	const q = `
INSERT INTO time_record (
	id, amount, description, client_project_id, user_id
)
VALUES (
	$1, $2, $3, $4, $5
)
RETURNING
	id, amount, timestamp, description, client_project_id, user_id
`
	err = t.DB.Conn.QueryRow(q, row.ID, row.Amount, row.Description, row.ProjectID, row.UserID).
		Scan(&newRow.ID, &newRow.Amount, &newRow.Timestamp, &newRow.Description, &newRow.ProjectID, &newRow.UserID)
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
	id, amount, timestamp, description, client_project_id, user_id
FROM
	time_record
WHERE
	id = $1
`
	err = t.DB.Conn.QueryRow(q, id).
		Scan(&newRow.ID, &newRow.Amount, &newRow.Timestamp, &newRow.Description, &newRow.ProjectID, &newRow.UserID)
	if err != nil {
		return newRow, fmt.Errorf("when: find by id | table: TimeRecordTable | error: %s", err.Error())
	}

	return
}

func (t *TimeRecordTable) FindByUserID(uid string) (rows []TimeRecordRow, err error) {
  const q = `
SELECT
	id, amount, timestamp, description, client_project_id, user_id
FROM
	time_record
WHERE
	user_id = $1
`
  iterator, err := t.DB.Conn.Query(q, uid)
  if err != nil {
    return rows, fmt.Errorf("when: query FindByUserID(uid: %s) | table: TimeRecordTable | error: %s", uid, err.Error())
  }
  defer iterator.Close()

  for iterator.Next() {
    var row = TimeRecordRow{}
    err = iterator.Scan(&row.ID, &row.Amount, &row.Timestamp, &row.Description, &row.ProjectID, &row.UserID)
    if err != nil {
      return rows, fmt.Errorf("when: scanning FindByUserID(uid: %s)| table: TimeRecordTable | error: %s", uid, err.Error())
    }

    rows = append(rows, row)
  }

  if err = iterator.Err(); err != nil {
    return rows, fmt.Errorf("when: while iterating FindByUserID(uid: %s)| table: TimeRecordTable | error: %s", uid, err.Error())
  }

  return
}

func (t *TimeRecordTable) FindByProjectID(pid string) (rows []TimeRecordRow, err error) {
  const q = `
SELECT
	id, amount, timestamp, description, client_project_id, user_id
FROM
	time_record
WHERE
	client_project_id = $1
`
  iterator, err := t.DB.Conn.Query(q, pid)
  if err != nil {
    return rows, fmt.Errorf("when: query FindByProjectID(uid: %s) | table: TimeRecordTable | error: %s", pid, err.Error())
  }
  defer iterator.Close()

  for iterator.Next() {
    var row = TimeRecordRow{}
    err = iterator.Scan(&row.ID, &row.Amount, &row.Timestamp, &row.Description, &row.ProjectID, &row.UserID)
    if err != nil {
      return rows, fmt.Errorf("when: scanning FindByProjectID(uid: %s)| table: TimeRecordTable | error: %s", pid, err.Error())
    }

    rows = append(rows, row)
  }

  if err = iterator.Err(); err != nil {
    return rows, fmt.Errorf("when: while iterating FindByProjectID(uid: %s)| table: TimeRecordTable | error: %s", pid, err.Error())
  }

  return
}
