package models

import "github.com/sujeet-crossml/GoLang_Backend_Project/internal/config"

type Order struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	ItemName string  `json:"item_name"`
	Amount   float64 `json:"amount"`
}

func GetOrdersByUserID(userID int) ([]Order, error) {
	rows, err := config.DB.Query("SELECT id, user_id, item_name, amount FROM orders WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.ItemName, &o.Amount); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}
