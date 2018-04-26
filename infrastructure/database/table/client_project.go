package table

import (
  "fmt"

  "github.com/devishot/grpc-go-time_tracking/infrastructure/database"
)

const ClientProjectTableName = "client_project"

type ClientProjectTable struct {
	DB   *database.DB
	Name string
}

type ClientProjectRow struct {
	ID          string
	Description string
}

func NewClientProjectTable(db *database.DB) (*ClientProjectTable, error) {
	t := &ClientProjectTable{DB: db, Name: ClientProjectTableName}

	err := t.createTable()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *ClientProjectTable) createTable() error {
	const q = `
CREATE TABLE IF NOT EXISTS client_project (
	id uuid PRIMARY KEY,
	description text NOT NULL
)`
	if _, err := t.DB.Conn.Exec(q); err != nil {
		return fmt.Errorf("when: create table | table: ProjectTable | error: %s", err.Error())
	}
	return nil
}

func (t *ClientProjectTable) Insert(row ClientProjectRow) (newRow ClientProjectRow, err error) {
	const q = `
INSERT INTO client_project (
	id, description
)
VALUES (
	$1, $2
)
RETURNING
	id, description
`
	err = t.DB.Conn.QueryRow(q, row.ID, row.Description).
		Scan(&newRow.ID, &newRow.Description)
	if err != nil {
		return newRow, fmt.Errorf("when: insert row | table: ProjectTable | error: %s", err.Error())
	}

	return
}

func (t *ClientProjectTable) Delete(id string) (err error) {
	const q = `
DELETE FROM client_project
WHERE id = $1
`
	if _, err := t.DB.Conn.Exec(q, id); err != nil {
		return fmt.Errorf("when: delete row | table: ProjectTable | error: %s", err.Error())
	}
	return nil
}

func (t *ClientProjectTable) FindByID(id string) (newRow ClientProjectRow, err error) {
	const q = `
SELECT
	id, description
FROM
	client_project
WHERE
	id = $1
`
	err = t.DB.Conn.QueryRow(q, id).
		Scan(&newRow.ID, &newRow.Description)
	if err != nil {
		return newRow, fmt.Errorf("when: find by id | table: ProjectTable | error: %s", err.Error())
	}

	return
}
