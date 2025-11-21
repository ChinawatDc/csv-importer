package service

import (
	"csv-importer/internal/model"
	"database/sql"
)

func GetAllTasks(db *sql.DB) ([]model.Task, error) {
	rows, err := db.Query(`
		SELECT
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
		FROM barcode_tasks
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var t model.Task
		err := rows.Scan(
			&t.ProdOrderID,
			&t.SeqNr,
			&t.JobID,
			&t.OperationID,
			&t.ProdOrderStartDate,
			&t.CTSec,
			&t.OriginalQty,
			&t.RemainingQty,
			&t.LineSequence,
			&t.ProcessTime,
			&t.WorkingGroup,
			&t.WorkingStation,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}

func DeleteAll(db *sql.DB) (int, error) {
	res, err := db.Exec(`DELETE FROM barcode_tasks`)
	if err != nil {
		return 0, err
	}

	count, _ := res.RowsAffected()
	return int(count), nil
}
