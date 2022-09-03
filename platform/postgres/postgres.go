package postgres

import (
	"database/sql"

	"github.com/ellipizle/golang-mvc/model"
	"github.com/ellipizle/golang-mvc/repository"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func New(db *sql.DB) repository.Repository {
	return &Postgres{db}
}

func (ps *Postgres) AddEmployee(employee model.Employee) (*model.Employee, error) {
	const query = `INSERT INTO employee (name, email, salary) VALUES ($1, $2, $3)`
	tx, err := ps.db.Begin()
	if err != nil {
		return &employee, err
	}

	_, err = tx.Exec(query, employee.Name, employee.Email, employee.Salary)
	if err != nil {
		tx.Rollback()
		return &employee, err
	}

	tx.Commit()
	return &employee, nil
}

func (ps *Postgres) EditEmployee(employee model.Employee) (*model.Employee, error) {
	const query = `UPDATE employee SET name=$2, email=$3, salary=$4 WHERE id=$1`
	tx, err := ps.db.Begin()
	if err != nil {
		return &employee, err
	}

	_, err = tx.Exec(query, employee.Name, employee.Email, employee.Salary, employee.Id)
	if err != nil {
		tx.Rollback()
		return &employee, err
	}

	tx.Commit()
	return &employee, nil

}
func (ps *Postgres) GetEmployee(id int) (*model.Employee, error) {
	employee := &model.Employee{}
	// rows, err := ps.db.Query(`SELECT * FROM "employee"`)
	const query = `SELECT * FROM employee WHERE id = $1`
	err := ps.db.QueryRow(query, id).Scan(&employee)

	if err != nil {
		return employee, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var roll int

		err = rows.Scan(&name, &roll)
	}
}
func (ps *Postgres) DeleteEmployee(id int) error {

	const query = `DELETE FROM employee WHERE id =$1`
	tx, err := ps.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}
func (ps *Postgres) GetAllEmployee() ([]*model.Employee, error) {

}
