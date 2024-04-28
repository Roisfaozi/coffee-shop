package repository

import (
	"context"
	"github.com/Roisfaozi/coffee-shop/config"
	"math"
	"time"

	"github.com/Roisfaozi/coffee-shop/internal/models"
	"github.com/jmoiron/sqlx"
)

type ProductRepositoryImpl struct {
	db *sqlx.DB
}

func NewProductRepositoryImpl(db *sqlx.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db}
}

func (pr *ProductRepositoryImpl) CreateProduct(ctx context.Context, product *models.ProductRequest) (*config.Result, error) {
	tx, err := pr.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
		}
	}()

	query := `INSERT INTO product (name, price, currency, description, image_url, category, created_at, updated_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var productID string
	err = tx.QueryRowContext(ctx, query, product.Name, product.Price, product.Currency, product.Description, product.ImageURL, product.Category, time.Now(), time.Now()).Scan(&productID)
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

	return &config.Result{
		Data:    &models.ProductResponse{ID: productID},
		Message: "Product created successfully",
	}, nil
}

func (pr *ProductRepositoryImpl) UpdateProduct(ctx context.Context, productID string, product *models.ProductRequest) error {
	tx, err := pr.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
		}
	}()

	query := `UPDATE product SET name=$1, price=$2, currency=$3, description=$4, image_url=$5, category=$6, updated_at=$7 WHERE id=$8`
	_, err = tx.ExecContext(ctx, query, product.Name, product.Price, product.Currency, product.Description, product.ImageURL, product.Category, time.Now(), productID)
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

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (pr *ProductRepositoryImpl) DeleteProduct(ctx context.Context, productID string) error {
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

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
func (pr *ProductRepositoryImpl) GetAllProducts(ctx context.Context, foodType string, page, limit int) (*config.Result, error) {
	offset := (page - 1) * limit

	productQuery := `
        SELECT p.id, p.name, p.price, p.currency, p.description, p.image_url, p.category, p.created_at, p.updated_at
        FROM product p
        WHERE p.category ILIKE '%' || $1 || '%'
        ORDER BY p.id LIMIT $2 OFFSET $3
    `

	sizeQuery := `
        SELECT ps.product_id, s.id as size_id, s.size_name
        FROM product_size ps
        JOIN size s ON ps.size_id = s.id
    `

	rows, err := pr.db.QueryContext(ctx, productQuery, foodType, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	productMap := make(map[string]*models.Product)
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Currency, &product.Description, &product.ImageURL, &product.Category, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, &product)
		productMap[product.ID] = &product
	}

	rows, err = pr.db.QueryContext(ctx, sizeQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var productID, sizeID, sizeName string
		if err := rows.Scan(&productID, &sizeID, &sizeName); err != nil {
			return nil, err
		}

		if product, ok := productMap[productID]; ok {
			product.Sizes = append(product.Sizes, &models.Size{
				ID:       sizeID,
				SizeName: sizeName,
			})
		}
	}

	// Calculate total count of products (assuming it's not affected by pagination)
	totalCountQuery := `SELECT COUNT(*) FROM product WHERE category ILIKE '%' || $1 || '%'`
	var totalCount int
	if err := pr.db.QueryRowContext(ctx, totalCountQuery, foodType).Scan(&totalCount); err != nil {
		return nil, err
	}

	// Calculate metadata for pagination
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	var next, prev interface{}
	if page < totalPages {
		next = page + 1
	}
	if page > 1 {
		prev = page - 1
	}

	meta := &config.Metas{
		Total: totalCount,
		Next:  next,
		Prev:  prev,
	}

	return &config.Result{Data: products, Meta: meta}, nil
}

func (pr *ProductRepositoryImpl) GetProductByID(ctx context.Context, productID string) (*config.Result, error) {
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

	for rows.Next() {
		var size models.Size
		if err := rows.StructScan(&size); err != nil {
			return nil, err
		}
		product.Sizes = append(product.Sizes, &size)
	}

	return &config.Result{Data: product}, nil
}
