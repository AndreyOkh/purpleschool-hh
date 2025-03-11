package vacancy

import "time"

type VacancyCreateForm struct {
	Role     string
	Company  string
	Type     string
	Salary   string
	Location string
	Email    string
}
type Vacancy struct {
	Id          int       `db:"id"`
	Email       string    `db:"email"`
	Role        string    `db:"role"`
	Company     string    `db:"company"`
	Salary      string    `db:"salary"`
	CompanyType string    `db:"type"`
	Location    string    `db:"location"`
	CreatedAt   time.Time `db:"created_at"`
}
