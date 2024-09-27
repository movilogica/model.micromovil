package model

import (
	"context"
	"encoding/json"
	"log"
)

// Ranking
type CoqRankingE struct {
	Orden        NullInt32  `json:"orden,omitempty"`
	Campeonatoid NullInt64  `json:"campeonato_id,omitempty"`
	Galponid     NullInt64  `json:"galpon_id,omitempty"`
	GalponText   NullString `json:"galpon_nombre,omitempty"`
	Pendientes   NullInt32  `json:"pendientes,omitempty"`
	Finalizadas  NullInt32  `json:"finalizadas,omitempty"`
	Ganadas      NullInt32  `json:"ganadas,omitempty"`
	Empatadas    NullInt32  `json:"empatadas,omitempty"`
	Perdidas     NullInt32  `json:"perdidas,omitempty"`
	Tiempo       NullInt32  `json:"tiempo,omitempty"`
	Total        NullInt32  `json:"total,omitempty"`
	Ruf1         NullString `json:"ruf1,omitempty"`
	Ruf2         NullString `json:"ruf2,omitempty"`
	Ruf3         NullString `json:"ruf3,omitempty"`
	Iv           NullString `json:"iv,omitempty"`
	Salt         NullString `json:"salt,omitempty"`
	Checksum     NullString `json:"checksum,omitempty"`
	FCreated     NullTime   `json:"fcreated,omitempty"`
	FUpdated     NullTime   `json:"fupdated,omitempty"`
	UCreated     NullString `json:"ucreated,omitempty"`
	UUpdated     NullString `json:"uupdated,omitempty"`
	Activo       int32      `json:"activo,omitempty"`
	Estadoreg    int32      `json:"estadoreg,omitempty"`
	TotalRecords int64      `json:"total_records,omitempty"`
}

func (e CoqRankingE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectCoqRanking = `select * from coq_ranking_peleas_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *CoqRankingE) GetAll(token string, filter string) ([]*CoqRankingE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCoqRanking

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

	var lista []*CoqRankingE

	for rows.Next() {
		var rowdata CoqRankingE
		err := rows.Scan(
			&rowdata.Orden,
			&rowdata.Campeonatoid,
			&rowdata.Galponid,
			&rowdata.GalponText,
			&rowdata.Pendientes,
			&rowdata.Finalizadas,
			&rowdata.Ganadas,
			&rowdata.Empatadas,
			&rowdata.Perdidas,
			&rowdata.Tiempo,
			&rowdata.Total,
			/*&rowdata.Ruf1,
			&rowdata.Ruf2,
			&rowdata.Ruf3,
			&rowdata.Iv,
			&rowdata.Salt,
			&rowdata.Checksum,
			&rowdata.FCreated,
			&rowdata.FUpdated,
			&rowdata.Activo,
			&rowdata.Estadoreg,*/
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
