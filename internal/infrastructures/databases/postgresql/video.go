package postgresql

import (
	"database/sql"

	"canty/internal/core/entities"
)

type PostgresVideoRepository struct {
	DB *sql.DB
}

func NewPostgresVideoRepository(db *sql.DB) *PostgresVideoRepository {
	return &PostgresVideoRepository{
		DB: db,
	}
}

func (repo *PostgresVideoRepository) Create(video *entities.Video) error {
	query := `INSERT INTO videos (id, title, description, url) VALUES ($1, $2, $3, $4)`
	_, err := repo.DB.Exec(query, video.ID, video.Title, video.Description, video.URL)
	return err
}

func (repo *PostgresVideoRepository) Read(id string) (*entities.Video, error) {
	query := `SELECT id, title, description, url FROM videos WHERE id = $1`
	row := repo.DB.QueryRow(query, id)
	video := &entities.Video{}
	err := row.Scan(&video.ID, &video.Title, &video.Description, &video.URL)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (repo *PostgresVideoRepository) Update(video *entities.Video) error {
	query := `UPDATE videos SET title = $2, description = $3, url = $4 WHERE id = $1`
	_, err := repo.DB.Exec(query, video.ID, video.Title, video.Description, video.URL)
	return err
}

func (repo *PostgresVideoRepository) Delete(id string) error {
	query := `DELETE FROM videos WHERE id = $1`
	_, err := repo.DB.Exec(query, id)
	return err
}
