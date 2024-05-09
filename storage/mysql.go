package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"simple-api/models"
	"time"
)

type MysqlStore struct {
	db *sql.DB
}

func NewMysqlStore() (*MysqlStore, error) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB_ADDR"))
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MysqlStore{
		db: db,
	}, nil
}

func (s *MysqlStore) AddProduct(p *models.Product) error {
	q := `insert into product (name, price) values (?, ?)`

	result, err := s.db.Exec(
		q,
		p.Name,
		p.Price,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = id

	return nil
}

func (s *MysqlStore) GetProductByID(id int) (*models.Product, error) {
	q := `SELECT * FROM product WHERE id = ?`

	row := s.db.QueryRow(q, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var product models.Product
	if err := row.Scan(
		&product.ID,
		&product.Name,
		&product.Price,
		&product.Created,
		&product.Updated,
		&product.Deleted,
	); err != nil {
		return nil, err
	}

	return &product, nil
}
