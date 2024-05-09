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

func (s *MysqlStore) AddProduct(product *models.Product) error {
	q := `INSERT INTO product (name, price) VALUES (?, ?)`

	result, err := s.db.Exec(
		q,
		product.Name,
		product.Price,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	product.ID = id

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

func (s *MysqlStore) AddClient(client *models.Client) error {
	q := `INSERT INTO client (email, password) VALUES (?, ?)`

	result, err := s.db.Exec(q, client.Email, client.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	client.ID = id
	return nil
}

func (s *MysqlStore) GetClientByID(id int) (*models.Client, error) {
	q := `SELECT * FROM client WHERE id = ?`

	row := s.db.QueryRow(q, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var client models.Client
	if err := row.Scan(
		&client.ID,
		&client.Email,
		&client.Password,
		&client.Created,
		&client.Updated,
		&client.Deleted,
	); err != nil {
		return nil, err
	}

	return &client, nil
}

func (s *MysqlStore) GetClientByEmail(email string) (*models.Client, error) {
	q := `SELECT * FROM client WHERE email = ?`

	row := s.db.QueryRow(q, email)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var client models.Client
	if err := row.Scan(
		&client.ID,
		&client.Email,
		&client.Password,
		&client.Created,
		&client.Updated,
		&client.Deleted,
	); err != nil {
		return nil, err
	}

	return &client, nil
}
