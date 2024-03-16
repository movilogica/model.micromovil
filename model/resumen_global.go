package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)

// Resumen
type ResumenGlobalE struct {
	Sede         int32       `json:"sede"`
	PersonaId    int64       `json:"personaid,omitempty"`
	Periodo      string      `json:"periodo,omitempty"`
	TokenDataId  string      `json:"tokendataid,omitempty"`
	Suscritos    NullInt64   `json:"suscritos,omitempty"`
	IncomeSubs   NullFloat64 `json:"incomesubs,omitempty"`
	IncomeServ   NullFloat64 `json:"incomeserv,omitempty"`
	Pendientes   NullFloat64 `json:"pendientes,omitempty"`
	TotalRecords int64       `json:"total_records"`
}

func (e ResumenGlobalE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

func (e ResumenGlobalE) PeriodoText() string {
	if len(e.Periodo) == 7 {
		year := e.Periodo[0:4]
		month, _ := strconv.Atoi(e.Periodo[5:])
		return fmt.Sprintf("%s-%s", time.Month(month), year)
	}
	return e.Periodo
}

func (e ResumenGlobalE) Incomes() float64 {
	return e.IncomeSubs.Float64 + e.IncomeServ.Float64
}

const querySelectResumenGlobal = `select * from resumen_global_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *ResumenGlobalE) GetAll(token string, filter string) ([]*ResumenGlobalE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectResumenGlobal

	// Se deseenvuelve el JSON del Filter para adicionar filtros
	var mapFilter map[string]interface{}
	json.Unmarshal([]byte(filter), &mapFilter)
	if mapFilter == nil {
		mapFilter = make(map[string]interface{})
	}
	// --- Adicion de filtros
	// mapFilter["tipo"] = tabla
	// Se empaqueta el JSON del Filter
	jsonFilter, err := json.Marshal(mapFilter)
	if err != nil {
		log.Println("Error convirtiendo el Filter")
	}
	log.Println("Where = " + string(jsonFilter))

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, token, string(jsonFilter))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []*ResumenGlobalE

	for rows.Next() {
		var rowdata ResumenGlobalE
		err := rows.Scan(
			&rowdata.Sede,
			&rowdata.PersonaId,
			&rowdata.Periodo,
			&rowdata.TokenDataId,
			&rowdata.Suscritos,
			&rowdata.IncomeSubs,
			&rowdata.IncomeServ,
			&rowdata.Pendientes,
			&rowdata.TotalRecords,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		lista = append(lista, &rowdata)
	}

	return lista, nil
}
