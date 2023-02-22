package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/engine/models"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type repository struct {
	pg  *pgxpool.Pool
	log *zap.Logger
}

func NewRepository(pg *pgxpool.Pool, log *zap.Logger) *repository {
	return &repository{
		pg:  pg,
		log: log,
	}
}

func (r *repository) Get(ctx context.Context) (models.Response, error) {
	var accs []*models.AccountResponse
	err := pgxscan.Select(ctx, r.pg, &accs, "select name, balance, currency from account order by id desc")
	if err != nil {
		r.log.Error("Error fetch from database",
			zap.Error(err),
		)
		return models.Response{}, err
	}

	return models.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    accs,
	}, nil
}

func (r *repository) GetForEachRowGoroutines(ctx context.Context) (models.Response, error) {
	q := `select name, balance, currency from account order by id desc`
	rows, _ := r.pg.Query(ctx, q)
	wg := sync.WaitGroup{}

	var name, currency string
	var balance int64
	var accs []*models.AccountResponse
	// errChan := make(chan error)
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := pgx.ForEachRow(rows, []any{&name, &balance, &currency}, func() error {
			accs = append(accs, &models.AccountResponse{
				Name:     name,
				Balance:  balance,
				Currency: currency,
			})
			return nil
		})

		if err != nil {
		r.log.Error("Error fetch from database",
				zap.Error(err),
			)
		}
	}()
	wg.Wait()

	defer rows.Close()

	return models.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    &accs,
	}, nil
}

func (r *repository) GetForEachRow(ctx context.Context) (models.Response, error) {
	q := `select name, balance, currency from account order by id desc`
	rows, _ := r.pg.Query(ctx, q)

	var name, currency string
	var balance int64
	var accs []*models.AccountResponse

	_, err := pgx.ForEachRow(rows, []any{&name, &balance, &currency}, func() error {
		accs = append(accs, &models.AccountResponse{
			Name:     name,
			Balance:  balance,
			Currency: currency,
		})
		return nil
	})

	if err != nil {
		r.log.Error("Error fetch from database",
			zap.Error(err),
		)
	}

	defer rows.Close()

	return models.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    &accs,
	}, nil
}

func (r *repository) GetRow(ctx context.Context) (models.Response, error) {
	q := `select name, balance, currency from account order by id desc`
	rows, _ := r.pg.Query(ctx, q)

	accs, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[models.AccountResponse])
	if err != nil {
		r.log.Error("Error fetch from database",
			zap.Error(err),
		)
	}

	defer rows.Close()

	return models.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    &accs,
	}, nil
}

func (r *repository) Insert(ctx context.Context, req models.AccountRequest) (models.Response, error) {
	dt := time.Now()
	q := `insert into account (name, balance, currency, created_at) values ($1, $2, $3, $4);`
	a, err := r.pg.Exec(ctx, q, req.Name, req.Balance, req.Currency, dt)
	if err != nil {
		r.log.Error("Error input to database",
			zap.Error(err),
		)
		return models.Response{}, err
	}
	return models.Response{
		Status:  fiber.StatusOK,
		Message: "Success Insert",
		Data:    a,
	}, nil
}

func (r *repository) Update(ctx context.Context) (models.Response, error) {
	q := `update account set (balance, currency) = ($1, $2) where id = $3;`
	a, err := r.pg.Exec(ctx, q, 120000, "USD", 22)
	if err != nil {
		r.log.Error("Error update to database",
			zap.Error(err),
		)
		return models.Response{}, err
	}
	return models.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    a,
	}, nil
}

func (r *repository) Delete(ctx context.Context) (models.Response, error) {
	q := `delete from account where id = $1;`
	a, err := r.pg.Exec(ctx, q, 22)
	if err != nil {
		r.log.Error("Error remove in database",
			zap.Error(err),
		)
		return models.Response{}, err
	}
	return models.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    a,
	}, nil
}

func (r *repository) GetRowScan(ctx context.Context) (models.Response, error) {
	q := `select * from account`
	rows, err := r.pg.Query(ctx, q)
	if err != nil {
		r.log.Error("Error GetRowScan data",
			zap.Error(err),
		)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error rows: %v\n", err)
	}

	var accs []models.Account

	for rows.Next() {
		var acc models.Account
		if err := rows.Scan(&acc); err != nil {
			fmt.Printf("Error scan: %v\n", err)
		}
		accs = append(accs, acc)
	}

	defer rows.Close()

	return models.Response{
		Status:  fiber.StatusOK,
		Message: "Success",
		Data:    accs,
	}, nil
}
