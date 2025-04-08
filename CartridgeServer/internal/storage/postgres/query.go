package postgres

import (
	"CartridgeServer/internal/models"
	"context"
	"fmt"
)

/*
ReadCartridges выполняет SELECT запрос к базе данных
Билдит SQL запрос основываясь на параметры offset, limit, name
в случае если параметр name = "", то запрос будет без фильтрации
отдает массив картриджей и ошибку
TODO: возможно вынести ошибки в отдельный файл и возвращать их. Но не уверен, ведь в любом случае будет возвращаться 500
*/
func (d *Database) ReadCartridges(offset int, limit int, name string) ([]models.CartridgeModel, error) {
	var (
		cartridges = make([]models.CartridgeModel, 0)
		cartridge  = models.CartridgeModel{}
		args       []interface{}
		index      = 1
	)

	query := "SELECT * FROM cartridges"
	if name != "" {
		query += fmt.Sprintf(" WHERE name = $%d", index)
		args = append(args, name)
		index++
	}

	query += fmt.Sprintf(" LIMIT $%d", index)
	args = append(args, limit)
	index++

	if offset > 0 {
		query += fmt.Sprintf(" OFFSET $2")
	}

	rows, err := d.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("reading cartridges error: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err = rows.Scan(
			&cartridge.Name, &cartridge.Parameters,
			&cartridge.Status, &cartridge.ReceivedFrom,
			&cartridge.ReceivedFromSubdivisionDate, &cartridge.SendToRefillingDate,
			&cartridge.ReceivedFromRefillingDate, &cartridge.SendTo,
			&cartridge.SendToSubdivisionDate, &cartridge.BarcodeNumber); err != nil {
			return nil, fmt.Errorf("reading rows error: %v", err)
		}
		cartridges = append(cartridges, cartridge)
	}

	return cartridges, nil
}
