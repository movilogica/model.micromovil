package model

import (
	"context"
	"log"
)

// Items Recount Summary
type StoreInventRecountSummaryE struct {
	Sede        int32       `json:"sede"`
	PersonaId   NullInt64   `json:"personaid,omitempty"`
	TokendataId NullString  `json:"tokendataid,omitempty"`
	WarehouseId NullInt64   `json:"warehouseid,omitempty"`
	RecountId   NullInt64   `json:"recountid,omitempty"`
	GroupText   NullString  `json:"grouptext,omitempty"`
	PersonText  NullString  `json:"persontext,omitempty"`
	StyleCode   NullString  `json:"stylecode,omitempty"`
	Qpersons    NullInt64   `json:"qpersons,omitempty"`
	Qitems      NullInt64   `json:"qitems,omitempty"`
	Qboxes      NullInt64   `json:"qboxes,omitempty"`
	Qstyles     NullInt64   `json:"qstyles,omitempty"`
	Percent     NullFloat64 `json:"percent,omitempty"`
}

func (e StoreInventRecountSummaryE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

func (u *StoreInventRecountSummaryE) GetSummaryGlobal(token string, filter string) ([]*StoreInventRecountSummaryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM view_sto_inv_rec_summary_global`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*StoreInventRecountSummaryE

	for rows.Next() {
		var rowdata StoreInventRecountSummaryE
		err := rows.Scan(
			&rowdata.Sede,
			&rowdata.PersonaId,
			&rowdata.TokendataId,
			&rowdata.WarehouseId,
			&rowdata.RecountId,
			&rowdata.Qitems,
			&rowdata.Qboxes,
			&rowdata.Qstyles,
			&rowdata.Qpersons,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		lista = append(lista, &rowdata)
	}

	return lista, nil
}

func (u *StoreInventRecountSummaryE) GetSummaryGroup(token string, filter string) ([]*StoreInventRecountSummaryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM view_sto_inv_rec_summary_groups`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*StoreInventRecountSummaryE

	for rows.Next() {
		var rowdata StoreInventRecountSummaryE
		err := rows.Scan(
			&rowdata.Sede,
			&rowdata.PersonaId,
			&rowdata.TokendataId,
			&rowdata.WarehouseId,
			&rowdata.RecountId,
			&rowdata.GroupText,
			&rowdata.Qpersons,
			&rowdata.Qitems,
			&rowdata.Qboxes,
			&rowdata.Qstyles,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		lista = append(lista, &rowdata)
	}

	return lista, nil
}

func (u *StoreInventRecountSummaryE) GetSummaryPersons(token string, filter string) ([]*StoreInventRecountSummaryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM view_sto_inv_rec_summary_persons`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*StoreInventRecountSummaryE

	for rows.Next() {
		var rowdata StoreInventRecountSummaryE
		err := rows.Scan(
			&rowdata.Sede,
			&rowdata.PersonaId,
			&rowdata.TokendataId,
			&rowdata.WarehouseId,
			&rowdata.RecountId,
			&rowdata.PersonText,
			&rowdata.Qstyles,
			&rowdata.Qitems,
			&rowdata.Qboxes,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		lista = append(lista, &rowdata)
	}

	return lista, nil
}

func (u *StoreInventRecountSummaryE) GetSummaryStyles(token string, filter string) ([]*StoreInventRecountSummaryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM view_sto_inv_rec_summary_styles`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*StoreInventRecountSummaryE

	for rows.Next() {
		var rowdata StoreInventRecountSummaryE
		err := rows.Scan(
			&rowdata.Sede,
			&rowdata.PersonaId,
			&rowdata.TokendataId,
			&rowdata.WarehouseId,
			&rowdata.RecountId,
			&rowdata.StyleCode,
			&rowdata.Qitems,
			&rowdata.Qboxes,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		lista = append(lista, &rowdata)
	}

	return lista, nil
}
