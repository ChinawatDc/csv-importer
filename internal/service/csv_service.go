package service

import (
	"csv-importer/internal/model"
	"database/sql"
	"encoding/csv"
	"io"
	"mime/multipart"
	"strings"
)

func ImportCSV(file multipart.File, db *sql.DB) (int, error) {
	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	_, err := reader.Read()
	if err != nil {
		return 0, err
	}

	batchSize := 5000
	batch := make([]model.Task, 0, batchSize)
	total := 0

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return total, err
		}

		for i := range record {
			record[i] = strings.TrimSpace(record[i])
		}

		if len(record) == 13 && record[12] == "" {
			record = record[:12]
		}

		if len(record) < 12 {
			continue
		}

		if strings.EqualFold(record[0], "Production order ID") {
			continue
		}

		batch = append(batch, model.Task{
			ProdOrderID:        record[0],
			SeqNr:              record[1],
			JobID:              record[2],
			OperationID:        record[3],
			ProdOrderStartDate: record[4],
			CTSec:              record[5],
			OriginalQty:        record[6],
			RemainingQty:       record[7],
			LineSequence:       record[8],
			ProcessTime:        record[9],
			WorkingGroup:       record[10],
			WorkingStation:     record[11],
		})

		if len(batch) >= batchSize {
			inserted, err := insertBatch(db, batch)
			if err != nil {
				return total, err
			}
			total += inserted
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		inserted, err := insertBatch(db, batch)
		if err != nil {
			return total, err
		}
		total += inserted
	}

	return total, nil
}

func insertBatch(db *sql.DB, batch []model.Task) (int, error) {
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	stmt, err := tx.Prepare(`
        INSERT INTO barcode_tasks (
            prod_order_id,
            seq_nr,
            job_id,
            operation_id,
            prod_order_start_date,
            ct_sec,
            original_qty,
            remaining_qty,
            line_sequence,
            process_time,
            working_group,
            working_station
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	count := 0

	for _, t := range batch {
		_, err := stmt.Exec(
			t.ProdOrderID,
			t.SeqNr,
			t.JobID,
			t.OperationID,
			t.ProdOrderStartDate,
			t.CTSec,
			t.OriginalQty,
			t.RemainingQty,
			t.LineSequence,
			t.ProcessTime,
			t.WorkingGroup,
			t.WorkingStation,
		)

		if err != nil {
			tx.Rollback()
			return count, err
		}
		count++
	}

	return count, tx.Commit()
}
