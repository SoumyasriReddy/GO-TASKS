package prod

import (
	"Project2/model"
	"database/sql"
	"reflect"
	"testing"
)

func TestGetAllProducts(t *testing.T) {
	type args struct {
		db      *sql.DB
		product []model.Product
	}
	tests := []struct {
		name string
		args args
		want []model.Product
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAllProducts(tt.args.db, tt.args.product); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteProduct(t *testing.T) {
	type args struct {
		db *sql.DB
		ID string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteProduct(tt.args.db, tt.args.ID)
		})
	}
}

func TestCreateProduct(t *testing.T) {
	type args struct {
		db       *sql.DB
		products model.Product
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateProduct(tt.args.db, tt.args.products)
		})
	}
}

func TestUpdateProduct(t *testing.T) {
	type args struct {
		db       *sql.DB
		products model.Product
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateProduct(tt.args.db, tt.args.products)
		})
	}
}

func TestGetProduct(t *testing.T) {
	type args struct {
		db       *sql.DB
		products []model.Product
		ID       string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetProduct(tt.args.db, tt.args.products, tt.args.ID)
		})
	}
}
