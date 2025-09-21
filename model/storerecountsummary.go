package model

import (
	"context"
	"log"
)

// Items Recount Summary
type StoreRecountSummaryE struct {
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

func (e StoreRecountSummaryE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

func (e StoreRecountSummaryE) InitValues() {
	e.Sede = 0
	e.PersonaId.Value(0)
	e.TokendataId.Value("")
	e.WarehouseId.Value(0)
	e.RecountId.Value(0)
	e.GroupText.Value("")
	e.PersonText.Value("")
	e.StyleCode.Value("")
	e.Qpersons.Value(0)
	e.Qitems.Value(0)
	e.Qboxes.Value(0)
	e.Qstyles.Value(0)
	e.Percent.Value(0)
}

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

func (u *StoreRecountSummaryE) GetSummaryGlobal(token string, filter string) ([]*StoreRecountSummaryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM view_recount_summary_global`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*StoreRecountSummaryE

	for rows.Next() {
		var rowdata StoreRecountSummaryE
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

func (u *StoreRecountSummaryE) GetSummaryGroup(token string, filter string) ([]*StoreRecountSummaryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM view_recount_summary_groups`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*StoreRecountSummaryE

	for rows.Next() {
		var rowdata StoreRecountSummaryE
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

func (u *StoreRecountSummaryE) GetSummaryPersons(token string, filter string) ([]*StoreRecountSummaryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM view_recount_summary_persons`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*StoreRecountSummaryE

	for rows.Next() {
		var rowdata StoreRecountSummaryE
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

func (u *StoreRecountSummaryE) GetSummaryStyles(token string, filter string) ([]*StoreRecountSummaryE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM view_recount_summary_styles`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*StoreRecountSummaryE

	for rows.Next() {
		var rowdata StoreRecountSummaryE
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
