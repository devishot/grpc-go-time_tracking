package table

import (
  "fmt"

  "github.com/devishot/grpc-go-time_tracking/infrastructure/database"
)

const UserTableName = "app_user"

type UserTable struct {
	DB   *database.DB
	Name string
}

type UserRow struct {
	ID          string
	FirstName   string
	LastName    string
}

func NewUserTable(db *database.DB) (*UserTable, error) {
	t := &UserTable{DB: db, Name: UserTableName}

	err := t.createTable()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *UserTable) createTable() error {
	const q = `
CREATE TABLE IF NOT EXISTS app_user (
	id uuid PRIMARY KEY,
	first_name text NOT NULL,
  last_name text NOT NULL
)`
	if _, err := t.DB.Conn.Exec(q); err != nil {
		return fmt.Errorf("when: create table | table: UserTable | error: %s", err.Error())
	}
	return nil
}

func (t *UserTable) Insert(row UserRow) (newRow UserRow, err error) {
	const q = `
INSERT INTO app_user (
	id, first_name, last_name
)
VALUES (
	$1, $2, $3
)
RETURNING
	id, first_name, last_name
`
	err = t.DB.Conn.QueryRow(q, row.ID, row.FirstName).
		Scan(&newRow.ID, &newRow.FirstName, &newRow.LastName)
	if err != nil {
		return newRow, fmt.Errorf("when: insert row | table: UserTable | error: %s", err.Error())
	}

	return
}

func (t *UserTable) Delete(id string) (err error) {
	const q = `
DELETE FROM app_user
WHERE id = $1
`
	if _, err := t.DB.Conn.Exec(q, id); err != nil {
		return fmt.Errorf("when: delete row | table: UserTable | error: %s", err.Error())
	}
	return nil
}

func (t *UserTable) FindByID(id string) (newRow UserRow, err error) {
	const q = `
SELECT
	id, first_name, last_name
FROM
	app_user
WHERE
	id = $1
`
	err = t.DB.Conn.QueryRow(q, id).
		Scan(&newRow.ID, &newRow.FirstName, &newRow.LastName)
	if err != nil {
		return newRow, fmt.Errorf("when: find by id | table: UserTable | error: %s", err.Error())
	}

	return
}
