package zipcode

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type order struct {
	id  int
	zip string
}

func (o order) String() string {
	return fmt.Sprintf("Order ID: %d, Zipcode: %s", o.id, o.zip)
}

func resetOrders(ctx context.Context, db *sql.DB) error {
	// Drop previous table of same name if one exists.
	_, err := db.ExecContext(ctx, "DROP TABLE IF EXISTS orders")
	if err != nil {
		return err
	}
	log.Println("Finished dropping table (if it existed)")
	// Create table.
	_, err = db.ExecContext(ctx, "CREATE TABLE orders (id serial PRIMARY KEY, zipcode VARCHAR(50));")
	if err != nil {
		return err
	}
	log.Println("Finished creating table")
	return nil
}

func listOrders(ctx context.Context, db *sql.DB) ([]order, error) {
	log.Println("Listing Orders")
	rows, err := db.QueryContext(ctx, "SELECT * from orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []order
	for rows.Next() {
		var o order
		if err := rows.Scan(&o.id, &o.zip); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func insertOrders(ctx context.Context, db *sql.DB, zips ...string) error {
	const stmt = "INSERT INTO orders (zipcode) VALUES ($1)"
	s, err := db.Prepare(stmt)
	if err != nil {
		return err
	}
	defer s.Close()
	for _, zip := range zips {
		log.Println("Inserting order", zip)
		_, err = s.ExecContext(ctx, zip)
		if err != nil {
			return err
		}
	}
	log.Printf("Added %d Order(s)\n", len(zips))
	return nil
}

func orderCountByCity(ctx context.Context, db *sql.DB) (map[string]int, error) {
	os, err := listOrders(ctx, db)
	if err != nil {
		return nil, err
	}
	count := make(map[string]int)
	for _, o := range os {
		city, err := zipToCity(o.zip)
		if err != nil {
			return nil, fmt.Errorf("could not get city for order %d: %s", o.id, err.Error())
		}
		count[city] = count[city] + 1
	}
	return count, nil
}
