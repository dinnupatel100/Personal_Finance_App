package app

import (
	"reflect"
	"testing"

	"github.com/personal-finance-app/db"
	"github.com/personal-finance-app/domain"
)

func TestNewService(t *testing.T) {
	type args struct {
		str db.Storer
	}
	tests := []struct {
		name string
		args args
		want Service
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Signup(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		u User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			if err := s.Signup(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("service.Signup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Login(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		l Login
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			if err := s.Login(tt.args.l); (err != nil) != tt.wantErr {
				t.Errorf("service.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_Search(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		tag string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got, err := s.Search(tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_AddBudget(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		b Budget
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			if err := s.AddBudget(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("service.AddBudget() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_AddCategory(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		c Category
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			if err := s.AddCategory(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("service.AddCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetAllBudgets(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Budget
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got, err := s.GetAllBudgets()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetAllBudgets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetAllBudgets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_DeleteBudget(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		b Budget
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			if err := s.DeleteBudget(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteBudget() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetBudgetById(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Budget
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got, err := s.GetBudgetById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetBudgetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetBudgetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetTransactionData(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got, err := s.GetTransactionData()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetTransactionData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetTransactionData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetBudgetData(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got, err := s.GetBudgetData()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetBudgetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetBudgetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_UpdateBudget(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		b Budget
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			if err := s.UpdateBudget(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateBudget() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_AddTransaction(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		t Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			if err := s.AddTransaction(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("service.AddTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_UpdateTransaction(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		t Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			if err := s.UpdateTransaction(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("service.UpdateTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_DeleteTransaction(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		t Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			if err := s.DeleteTransaction(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("service.DeleteTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetTransactionById(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got, err := s.GetTransactionById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetTransactionById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetTransactionById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetAllTransactions(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got, err := s.GetAllTransactions()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetAllTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetAllTransactions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetTransactionByCategory(t *testing.T) {
	type fields struct {
		store db.Storer
	}
	type args struct {
		category string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				store: tt.fields.store,
			}
			got, err := s.GetTransactionByCategory(tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetTransactionByCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.GetTransactionByCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}
