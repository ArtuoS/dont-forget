package repository

import (
	"database/sql"
	"log"

	"github.com/ArtuoS/dont-forget/internal/entity"
)

type ItemRepository struct {
	Db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{Db: db}
}

func (r *ItemRepository) Save(item *entity.Item) error {
	log.Println("Saving item with Guid: ", item.Guid)

	stmt, err := r.Db.Prepare("INSERT INTO items(guid, name, description, start_date, end_date) values(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Guid, item.Name, item.Description, item.StartDate, item.EndDate)
	if err != nil {
		return err
	}
	log.Println("Saved item with Guid: ", item.Guid)
	return nil
}

func (r *ItemRepository) Get(guid string) (*entity.Item, error) {
	log.Println("Getting item with Guid: ", guid)

	stmt, err := r.Db.Prepare("SELECT guid, name, description, start_date, end_date FROM items WHERE guid = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(guid)
	var item entity.Item
	err = row.Scan(&item.Guid, &item.Name, &item.Description, &item.StartDate, &item.EndDate)
	if err != nil {
		return nil, err
	}
	log.Println("Got item with Guid: ", item.Guid)
	return &item, nil
}

func (r *ItemRepository) GetAll() ([]entity.Item, error) {
	log.Println("Getting all items")

	rows, err := r.Db.Query("SELECT guid, name, description, start_date, end_date FROM items")
	if err != nil {
		return nil, err
	}
	var items []entity.Item
	for rows.Next() {
		var item entity.Item
		err = rows.Scan(&item.Guid, &item.Name, &item.Description, &item.StartDate, &item.EndDate)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	log.Println("Got all items")
	return items, nil
}

func (r *ItemRepository) Delete(guid string) error {
	log.Println("Deleting item with Guid: ", guid)

	stmt, err := r.Db.Prepare("DELETE FROM items WHERE guid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(guid)
	if err != nil {
		return err
	}

	log.Println("Deleted item with Guid: ", guid)
	return nil
}

func (r *ItemRepository) Update(item *entity.Item) error {
	log.Println("Updating item with Guid: ", item.Guid)

	stmt, err := r.Db.Prepare("UPDATE items SET name = ?, description = ?, start_date = ?, end_date = ? WHERE guid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.Name, item.Description, item.StartDate, item.EndDate, item.Guid)
	if err != nil {
		return err
	}

	log.Println("Updated item with Guid: ", item.Guid)
	return nil
}
