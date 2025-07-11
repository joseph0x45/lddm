package repo

import (
	"server/models"

	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

func (r *ProductRepo) InsertProduct(product *models.Product) error {
	const query = `
    insert into products (
      id, name, variant, price,
      image, description
    )
    values (
      :id, :name, :variant, :price,
      :image, :description
    );
  `
	_, err := r.db.NamedExec(query, product)
	return err
}

func (r *ProductRepo) GetAllProducts() ([]models.Product, error) {
	products := make([]models.Product, 0)
	const query = `select * from products`
	err := r.db.Select(&products, query)
	return products, err
}

func (r *ProductRepo) UpdateProduct(productData models.ProductUpdateData) error {
	panic("Not implemented")
}
