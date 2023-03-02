package repositories

import (
	"backEnd/models"
	"time"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProduct() ([]models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product) (models.Product, error)
	GetProduct(ID int) (models.Product, error)
	DeleteProduct(product models.Product) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProduct() ([]models.Product, error) {
	var Products []models.Product
	err := r.db.Raw("SELECT * FROM products").Scan(&Products).Error

	return Products, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Exec("INSERT INTO products( name , price , description , stock , created_at , updated_at) VALUES (?,?,?,?,?,?)", product.Name, product.Price, product.Description, product.Stock, time.Now(), time.Now()).Error

	return product, err
}

func (r *repository) UpdateProduct(product models.Product) (models.Product, error) {
	// err := r.db.Raw("UPDATE products SET name=?, price=?, description=?, stock=?, WHERE id=?", product.Name, product.Price, product.Description, product.Stock, ID).Scan(&product).Error
	err := r.db.Save(&product).Error

	return product, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.Raw("SELECT * FROM products WHERE id=?", ID).Scan(&product).Error
	return product, err
}

func (r *repository) DeleteProduct(product models.Product) (models.Product, error) {
	var products models.Product
	err := r.db.Delete(&product).Error // Using Delete method

	return products, err
}
