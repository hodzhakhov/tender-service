package employee

import (
	"context"
	"database/sql"
	repoModel "tender-service/internal/repository/tables/tenders/public/model"
	"tender-service/internal/repository/tables/tenders/public/table"

	"github.com/go-jet/jet/v2/postgres"
)

type EmployeeRepo struct {
	db *sql.DB
}

func NewEmployeeRepo(db *sql.DB) *EmployeeRepo {
	return &EmployeeRepo{db: db}
}

func (h *EmployeeRepo) GetEmployeeByUsername(ctx context.Context, username string) (repoModel.Employee, error) {
	stmt := table.Employee.SELECT(table.Employee.AllColumns).
		WHERE(table.Employee.Username.EQ(postgres.String(username)))

	dest := repoModel.Employee{}

	err := stmt.QueryContext(ctx, h.db, &dest)
	if err != nil {
		return repoModel.Employee{}, err
	}

	return dest, nil
}
