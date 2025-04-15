package postgres

import (
	"DbService/internal/models"
	"DbService/internal/storage"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"log"
)

const (
	SEND_TO_REFILLING = "отправлено на заправку"
	UNIQUE_CODE_PG    = "23505"
)

func (d *Database) CreateCartridge(cartridge *models.CreateCartridge) error {
	var (
		opts = pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}
		pgErr = &pgconn.PgError{}
	)

	tx, err := d.db.BeginTx(context.Background(), opts)
	if err != nil {
		log.Println("failed to begin tx: ", err)
		return storage.ErrBeginTxn
	}

	defer func() {
		_ = tx.Rollback(context.Background())
	}()

	_, err = tx.Exec(context.Background(),
		"INSERT INTO cartridges"+
			"(name, status, received_from, received_from_subdivision_date, send_to_refilling_date, barcode_number) "+
			"VALUES($1, $2, $3, $4, $5, $6)", cartridge.Name, SEND_TO_REFILLING, cartridge.ReceivedFrom, cartridge.Timestamp, cartridge.Timestamp, cartridge.BarcodeNumber)
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == UNIQUE_CODE_PG {
			log.Println("failed to insert into cartridges: ", err)
			return storage.ErrUniqueBarcode
		}
		return fmt.Errorf("create cartridge %s: %w", cartridge.Name, err)
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Println("failed to commit tx: ", err)
		return storage.ErrCommitTxn
	}
	return nil
}

func (d *Database) UpdateCartridgeReceiveStatus(cartridge *models.UpdateCartridgeReceive) error {
	var (
		opts = pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}
	)

	tx, err := d.db.BeginTx(context.Background(), opts)
	if err != nil {
		log.Println("failed to begin tx: ", err)
		return storage.ErrBeginTxn
	}

	defer func() {
		_ = tx.Rollback(context.Background())
	}()

	ctag, err := tx.Exec(context.Background(),
		"UPDATE cartridges SET status = $1, received_from_refilling_date = $2 WHERE barcode_number = $3",
		cartridge.NewStatus, cartridge.Timestamp, cartridge.BarcodeNumber)
	if err != nil {
		return fmt.Errorf("update cartridge receive :%w", err)
	}

	if ctag.RowsAffected() == 0 {
		log.Println("cartridge not found")
		return storage.ErrNotFound
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Println("failed to commit tx: ", err)
		return storage.ErrCommitTxn
	}
	return nil
}

func (d *Database) UpdateCartridgeSendStatus(cartridge *models.UpdateCartridgeSend) error {
	var (
		opts = pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}
	)

	tx, err := d.db.BeginTx(context.Background(), opts)
	if err != nil {
		log.Println("failed to begin tx: ", err)
		return storage.ErrBeginTxn
	}

	defer func() {
		_ = tx.Rollback(context.Background())
	}()

	ctag, err := tx.Exec(context.Background(),
		"UPDATE cartridges SET status = $1, send_to_subdivision_date = $2, send_to = $3 WHERE barcode_number = $4",
		cartridge.NewStatus, cartridge.Timestamp, cartridge.SendTo, cartridge.BarcodeNumber)
	if err != nil {
		return fmt.Errorf("update cartridge send :%w", err)
	}
	if ctag.RowsAffected() == 0 {
		log.Println("cartridge not found")
		return storage.ErrNotFound
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Println("failed to commit tx: ", err)
		return storage.ErrCommitTxn
	}
	return nil
}

func (d *Database) DeleteCartridge(barcodeNumber int) error {
	var (
		opts = pgx.TxOptions{
			IsoLevel: pgx.ReadCommitted,
		}
	)

	tx, err := d.db.BeginTx(context.Background(), opts)
	if err != nil {
		log.Println("failed to begin tx: ", err)
		return storage.ErrBeginTxn
	}

	defer func() {
		_ = tx.Rollback(context.Background())
	}()

	ctag, err := tx.Exec(context.Background(), "DELETE FROM cartridges WHERE barcode_number = $1", barcodeNumber)
	if err != nil {
		return fmt.Errorf("delete cartridge %s: %w", barcodeNumber, err)
	}

	if ctag.RowsAffected() == 0 {
		log.Println("cartridge not found")
		return storage.ErrNotFound
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Println("failed to commit tx: ", err)
		return storage.ErrCommitTxn
	}
	return nil
}
