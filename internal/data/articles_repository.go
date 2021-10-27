package data

import (
	pb "api-grpc-articles-videogame/proto"
	"context"
	"log"
	"time"
)

type ArticlesRepository struct {
	Data *Data
}

/*
	GET ONE

*/

func (pr *ArticlesRepository) GetOne(ctx context.Context, id uint32) (*pb.Article, error) {
	q := `SELECT id, name, price, description,user_id, created_at, updated_at FROM articles WHERE id = $1;`

	row := pr.Data.DB.QueryRowContext(ctx, q, id)
	var art pb.Article

	var created_at time.Time
	var updated_at time.Time

	err := row.Scan(&art.Id, &art.Name, &art.Price, &art.Description, &art.UserId, &created_at, &updated_at)

	art.CreatedAt = created_at.String()
	art.UpdatedAt = updated_at.String()

	if err != nil {
		log.Printf("Error %v", err.Error())

		return nil, err
	}

	return &art, nil
}

/*
	GET ALL
*/

func (pr *ArticlesRepository) GetAll(ctx context.Context) ([]*pb.Article, error) {
	q := `SELECT id, name, price, description, user_id, created_at, updated_at FROM articles; `

	rows, err := pr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var articles []*pb.Article
	for rows.Next() {
		var p pb.Article

		var created_at time.Time
		var updated_at time.Time

		rows.Scan(&p.Id, &p.Name, &p.Price, &p.Description, &p.UserId, &created_at, &updated_at)

		p.CreatedAt = created_at.String()
		p.UpdatedAt = updated_at.String()

		articles = append(articles, &p)
	}

	return articles, nil
}

/*
	BY ID USER
*/

func (pr *ArticlesRepository) GetByUser(ctx context.Context, userID uint32) ([]*pb.Article, error) {
	q := `SELECT id, name, price, description, user_id, created_at, updated_at FROM articles WHERE user_id = $1;`

	rows, err := pr.Data.DB.QueryContext(ctx, q, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var articles []*pb.Article
	for rows.Next() {
		var p pb.Article
		rows.Scan(&p.Id, &p.Name, &p.Price, &p.Description, &p.UserId, &p.CreatedAt, &p.UpdatedAt)
		articles = append(articles, &p)
	}

	return articles, nil
}

/*
	INSERTAR
*/

func (pr *ArticlesRepository) Create(ctx context.Context, p *pb.CreateArticlerRequest) (*pb.CreateArticlerRequest, error) {
	q := `
    INSERT INTO ARTICLES (name, price, description, user_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id;
    `

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, p.Name, p.Price, p.Description, p.UserId, time.Now(), time.Now())

	err = row.Scan(&p.Id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

/*
	UPDATE
*/

func (pr *ArticlesRepository) Update(ctx context.Context, id uint32, p *pb.CreateArticlerRequest) (*pb.CreateArticlerRequest, error) {
	q := `UPDATE articles set name=$1, price=$2, description=$3, updated_at=$4 WHERE id=$5; `

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, p.Name, p.Price, p.Description, time.Now(), id,
	)

	if err != nil {
		return nil, err
	}

	return p, nil
}

/*
	DELETE
*/

func (pr *ArticlesRepository) Delete(ctx context.Context, id uint32) error {
	q := `DELETE FROM ARTICLES WHERE id=$1;`

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
