package repository

import (
	"database/sql"
	"errors"
	"time"
)

// SQLiteRepository represents a SQLite database repository for holdings.
type SQLiteRepository struct {
	Conn *sql.DB
}

// NewSQLiteRepository creates a new instance of SQLiteRepository.
func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

// Migrate creates the holdings table if it does not exist.
func (repo *SQLiteRepository) Migrate() error {
	query := `create table if not exists holdings(
                  id integer primary key autoincrement,
                  amount real not null,
                  purchase_data integer not null,
                  purchase_price integer not null)`

	_, err := repo.Conn.Exec(query)
	return err
}

// InsertHolding inserts a new holding into the database.
func (repo *SQLiteRepository) InsertHolding(holding Holdings) (*Holdings, error) {
	stmt := "insert into holdings (amount, purchase_date, purchase_price) values (?, ?, ?)"

	res, err := repo.Conn.Exec(stmt, holding.Amount, holding.PurchaseDate.Unix(), holding.PurchasePrice)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	holding.ID = id

	return &holding, nil
}

// AllHoldings retrieves all holdings from the database.
func (repo *SQLiteRepository) AllHoldings() ([]Holdings, error) {
	query := "select id, amount, purchase_date, purchase_price from holdings order by purchase_date"

	rows, err := repo.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []Holdings

	for rows.Next() {
		var h Holdings
		var unixTime int64

		err := rows.Scan(
			&h.ID,
			&h.Amount,
			&unixTime,
			&h.PurchasePrice,
		)
		if err != nil {
			return nil, err
		}

		h.PurchaseDate = time.Unix(unixTime, 0)
		all = append(all, h)
	}

	return all, nil
}

// GetHoldingByID retrieves a holding by its ID from the database.
func (repo *SQLiteRepository) GetHoldingByID(id int) (*Holdings, error) {
	row := repo.Conn.QueryRow("select id amount, purchase_data, purchase_price from holdings where id = ?", id)

	var h Holdings
	var unixTime int64

	err := row.Scan(
		&h.ID,
		&h.Amount,
		&unixTime,
		&h.PurchasePrice,
	)
	if err != nil {
		return nil, err
	}

	h.PurchaseDate = time.Unix(unixTime, 0)

	return &h, nil
}

// UpdateHolding updates a holding in the database.
func (repo *SQLiteRepository) UpdateHolding(id int64, updated Holdings) error {
	if id == 0 {
		return errors.New("invalid updated ID")
	}

	stmt := "update holdings set amount = ?, purchase_data = ?, purchase_price = ? where id = ?"

	res, err := repo.Conn.Exec(stmt, updated.Amount, updated.PurchaseDate.Unix(), updated.PurchasePrice, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}

// DeleteHolding deletes a holding from the database.
func (repo *SQLiteRepository) DeleteHolding(id int64) error {
	res, err := repo.Conn.Exec("delete from holdings where id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errUpdateFailed
	}

	return nil
}
