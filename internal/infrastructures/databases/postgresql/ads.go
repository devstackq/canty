package postgresql

import (
	"database/sql"

	"canty/internal/core/entities"
)

type PostgresAdvertisementRepository struct {
	DB *sql.DB
}

func NewPostgresAdvertisementRepository(db *sql.DB) *PostgresAdvertisementRepository {
	return &PostgresAdvertisementRepository{
		DB: db,
	}
}

func (repo *PostgresAdvertisementRepository) Create(ad *entities.Advertisement) error {
	query := `INSERT INTO advertisements (id, title, content, url) VALUES ($1, $2, $3, $4)`
	_, err := repo.DB.Exec(query, ad.ID, ad.Title, ad.Content, ad.URL)
	return err
}

func (repo *PostgresAdvertisementRepository) Read(id string) (*entities.Advertisement, error) {
	query := `SELECT id, title, content, url FROM advertisements WHERE id = $1`
	row := repo.DB.QueryRow(query, id)
	ad := &entities.Advertisement{}
	err := row.Scan(&ad.ID, &ad.Title, &ad.Content, &ad.URL)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (repo *PostgresAdvertisementRepository) Update(ad *entities.Advertisement) error {
	query := `UPDATE advertisements SET title = $2, content = $3, url = $4 WHERE id = $1`
	_, err := repo.DB.Exec(query, ad.ID, ad.Title, ad.Content, ad.URL)
	return err
}

func (repo *PostgresAdvertisementRepository) Delete(id string) error {
	query := `DELETE FROM advertisements WHERE id = $1`
	_, err := repo.DB.Exec(query, id)
	return err
}
