package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/jmoiron/sqlx"
)

type ProductRepositoryImpl struct {
	db *sqlx.DB
}

func (pr ProductRepositoryImpl) CreateProduct(ctx context.Context, product *models.ProductRequest) (*models.ProductResponse, error) {
	tx, err := pr.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
		}
	}()
	fmt.Println(product)

	query := `INSERT INTO product (name, price, currency, description, image_url, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	var productID string
	err = tx.QueryRowContext(ctx, query, product.Name, product.Price, product.Currency, product.Description, product.ImageURL, time.Now(), time.Now()).Scan(&productID)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	var sizeIDs []string
	rows, err := tx.QueryContext(ctx, "SELECT id FROM size")
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var sizeID string
		if err := rows.Scan(&sizeID); err != nil {
			_ = tx.Rollback()
			return nil, err
		}
		sizeIDs = append(sizeIDs, sizeID)
	}

	for _, sizeID := range sizeIDs {
		_, err = tx.ExecContext(ctx, "INSERT INTO product_size (product_id, size_id) VALUES ($1, $2)", productID, sizeID)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &models.ProductResponse{ID: productID}, nil
}

func (pr ProductRepositoryImpl) UpdateProduct(ctx context.Context, productID string, product *models.ProductRequest) error {
	tx, err := pr.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
		}
	}()

	query := `UPDATE product SET name=$1, price=$2, currency=$3, description=$4, image_url=$5, updated_at=$6 WHERE id=$7`
	_, err = tx.ExecContext(ctx, query, product.Name, product.Price, product.Currency, product.Description, product.ImageURL, time.Now(), productID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM product_size WHERE product_id=$1", productID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	for _, sizeID := range product.SizeIDs {
		_, err = tx.ExecContext(ctx, "INSERT INTO product_size (product_id, size_id) VALUES ($1, $2)", productID, sizeID)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (pr ProductRepositoryImpl) DeleteProduct(ctx context.Context, productID string) error {
	tx, err := pr.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
		}
	}()

	_, err = tx.ExecContext(ctx, "DELETE FROM product WHERE id=$1", productID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM product_size WHERE product_id=$1", productID)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (pr ProductRepositoryImpl) GetAllProducts(ctx context.Context, foodType string, page, limit int) ([]*models.Product, error) {
	offset := (page - 1) * limit

	query := `
        SELECT p.id, p.name, p.price, p.currency, p.description, p.image_url, p.created_at, p.updated_at,
               s.id as size_id, s.size_name
        FROM product p
        LEFT JOIN product_size ps ON p.id = ps.product_id
        LEFT JOIN size s ON ps.size_id = s.id
        
    `

	// if foodType != "" {
	// 	query += ` AND p.food_type = $1`
	// }

	query += ` ORDER BY p.id LIMIT $2 OFFSET $3`

	rows, err := pr.db.QueryContext(ctx, query, foodType, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productsMap := make(map[string]*models.Product)
	for rows.Next() {
		var (
			productID       string
			productName     string
			productPrice    int
			productCurrency string
			productDesc     string
			productImageURL string
			createdAt       time.Time
			updatedAt       time.Time
			sizeID          sql.NullString
			sizeName        sql.NullString
		)
		if err := rows.Scan(&productID, &productName, &productPrice, &productCurrency, &productDesc, &productImageURL, &createdAt, &updatedAt, &sizeID, &sizeName); err != nil {
			return nil, err
		}

		product, ok := productsMap[productID]
		if !ok {
			product = &models.Product{
				ID:          productID,
				Name:        productName,
				Price:       productPrice,
				Currency:    productCurrency,
				Description: productDesc,
				ImageURL:    productImageURL,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
				Sizes:       []*models.Size{},
			}
			productsMap[productID] = product
		}

		if sizeID.Valid && sizeName.Valid {
			product.Sizes = append(product.Sizes, &models.Size{
				ID:       sizeID.String,
				SizeName: sizeName.String,
			})
		}
	}

	products := make([]*models.Product, 0, len(productsMap))
	for _, product := range productsMap {
		products = append(products, product)
	}

	return products, nil
}
func (pr ProductRepositoryImpl) GetProductByID(ctx context.Context, productID string) (*models.Product, error) {
	query := `SELECT * FROM product WHERE id=$1`
	var product models.Product
	err := pr.db.GetContext(ctx, &product, query, productID)
	if err != nil {
		return nil, err
	}

	sizeQuery := `
        SELECT s.id, s.size_name
        FROM size s
        JOIN product_size ps ON s.id = ps.size_id
        WHERE ps.product_id = $1
    `
	rows, err := pr.db.QueryxContext(ctx, sizeQuery, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Kumpulkan semua ukuran terkait ke dalam produk
	for rows.Next() {
		var size models.Size
		if err := rows.StructScan(&size); err != nil {
			return nil, err
		}
		product.Sizes = append(product.Sizes, &size)
	}

	return &product, nil
}
func NewProductRepositoryImpl(db *sqlx.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db}
}
