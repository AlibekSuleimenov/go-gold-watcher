package repository

// Package repository provides an interface and methods for interacting with a holdings repository.
import (
	"errors"
	"time"
)

var (
	errUpdateFailed = errors.New("update failed") // errUpdateFailed is returned when an update operation fails.
	errDeleteFailed = errors.New("delete failed") // errDeleteFailed is returned when a delete operation fails.
)

// Repository defines the methods that a repository must implement.
type Repository interface {
	Migrate() error                                    // Migrate performs any necessary database migrations.
	InsertHolding(holding Holdings) (*Holdings, error) // InsertHolding inserts a new holding into the repository and returns the inserted holding.
	AllHoldings() ([]Holdings, error)                  // AllHoldings returns all holdings stored in the repository.
	GetHoldingByID(id int) (*Holdings, error)          // GetHoldingByID retrieves a holding by its ID from the repository.
	UpdateHolding(id int64, updated Holdings) error    // UpdateHolding updates a holding in the repository by its ID with the provided updated holding information.
	DeleteHolding(id int64) error                      // DeleteHolding deletes a holding from the repository by its ID.
}

// Holdings represents a holding entity with its attributes.
type Holdings struct {
	ID            int64     `json:"id"`
	Amount        int       `json:"amount"`
	PurchaseDate  time.Time `json:"purchase_date"`
	PurchasePrice int       `json:"purchase_price"`
}
