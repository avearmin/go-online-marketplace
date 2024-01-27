package api

import (
	"time"

	"github.com/avearmin/gorage-sale/internal/database"
	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int32     `json:"price"`
	Sold        bool      `json:"sold"`
	SellerID    uuid.UUID `json:"seller_id"`
}

func dbItemToJSONItem(item database.Item) Item {
	return Item{
		ID:          item.ID,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		Name:        item.Name,
		Description: item.Description,
		Price:       item.Price,
		Sold:        item.Sold,
		SellerID:    item.SellerID,
	}
}
