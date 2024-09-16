package employee

import "context"

type EmployeeService struct {
	employeeRepo *EmployeeRepo
}

func NewEmployeeService(employeeRepo *EmployeeRepo) *EmployeeService {
	return &EmployeeService{
		employeeRepo: employeeRepo,
	}
}

func (h *EmployeeService) GetEmployeeByUsername(ctx context.Context, username string) (Employee, error) {
	tx, err := h.employeeRepo.db.BeginTx(ctx, nil)
	if err != nil {
		return Employee{}, err
	}

	defer tx.Rollback()

	res, err := h.employeeRepo.GetEmployeeByUsername(ctx, username)
	if err != nil {
		return Employee{}, err
	}

	employee := Employee{
		ID:        res.ID.String(),
		Username:  res.Username,
		FirstName: *res.FirstName,
		LastName:  *res.LastName,
		CreatedAt: *res.CreatedAt,
		UpdatedAt: *res.UpdatedAt,
	}

	return employee, nil
}
