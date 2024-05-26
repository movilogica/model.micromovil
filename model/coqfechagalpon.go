package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Fecha Por Galpon
type CoqFechaGalponE struct {
	Uniqueid         int64       `json:"uniqueid,omitempty"`
	Owner            NullInt32   `json:"owner,omitempty"`
	Dispositivoid    NullInt32   `json:"dispositivoid,omitempty"`
	Id               int32       `json:"id,omitempty"`
	Sede             int32       `json:"sede"`
	Flag1            string      `json:"flag1,omitempty"`
	Flag2            string      `json:"flag2,omitempty"`
	CountryCode      NullString  `json:"countrycode,omitempty"`
	Campeonatoid     NullInt64   `json:"campeonato_id,omitempty"`
	Galponid         NullInt64   `json:"galpon_id,omitempty"`
	GalponText       NullString  `json:"galpon_text,omitempty"`
	GalponPadre      NullInt64   `json:"galpon_padre,omitempty"`
	Urbaniza         NullString  `json:"urbaniza,omitempty"`
	CriadorText      NullString  `json:"criador_text,omitempty"`
	Movil            NullString  `json:"movil,omitempty"`
	Email            NullString  `json:"email,omitempty"`
	FechaNumeroid    NullInt64   `json:"fecha_numero_id,omitempty"`
	FechaPelea       NullString  `json:"fecha_pelea,omitempty"`
	FechaCerrada     NullInt32   `json:"fecha_cerrada_ind,omitempty"`
	ExcluirRanking   NullInt32   `json:"excluir_ranking,omitempty"`
	MontoPactada     NullFloat64 `json:"monto_pactada,omitempty"`
	MontoEspuelas    NullFloat64 `json:"monto_espuelas,omitempty"`
	MontoInscripcion NullFloat64 `json:"monto_inscripcion,omitempty"`
	MontoCuota       NullFloat64 `json:"monto_cuota,omitempty"`
	MontoPenalidad   NullFloat64 `json:"monto_penalidad,omitempty"`
	MontoOtros       NullFloat64 `json:"monto_otros,omitempty"`
	OtrosText        NullString  `json:"otros_text,omitempty"`
	Estadoid         NullInt64   `json:"estado_id,omitempty"`
	EstadoText       NullString  `json:"estado_text,omitempty"`
	UAutoriza        NullString  `json:"uautoriza,omitempty"`
	FAutoriza        NullString  `json:"fautoriza,omitempty"`
	Nroperacion      NullString  `json:"nroperacion,omitempty"`
	Ruf1             NullString  `json:"ruf1,omitempty"`
	Ruf2             NullString  `json:"ruf2,omitempty"`
	Ruf3             NullString  `json:"ruf3,omitempty"`
	Iv               NullString  `json:"iv,omitempty"`
	Salt             NullString  `json:"salt,omitempty"`
	Checksum         NullString  `json:"checksum,omitempty"`
	FCreated         NullTime    `json:"fcreated,omitempty"`
	FUpdated         NullTime    `json:"fupdated,omitempty"`
	Activo           int32       `json:"activo,omitempty"`
	Estadoreg        int32       `json:"estadoreg,omitempty"`
	TotalRecords     int64       `json:"total_records"`
}

func (e CoqFechaGalponE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectCoqFechaGalpon = `select * from coq_fecha_por_galpon_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *CoqFechaGalponE) GetAll(token string, filter string) ([]*CoqFechaGalponE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCoqFechaGalpon

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

	var lista []*CoqFechaGalponE

	for rows.Next() {
		var rowdata CoqFechaGalponE
		err := rows.Scan(
			&rowdata.Uniqueid,
			//&rowdata.Owner,
			//&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			//&rowdata.Flag1,
			//&rowdata.Flag2,
			&rowdata.CountryCode,
			&rowdata.Campeonatoid,
			&rowdata.Galponid,
			&rowdata.GalponText,
			&rowdata.GalponPadre,
			&rowdata.Urbaniza,
			&rowdata.CriadorText,
			&rowdata.Movil,
			&rowdata.Email,
			&rowdata.FechaNumeroid,
			&rowdata.FechaPelea,
			&rowdata.FechaCerrada,
			&rowdata.ExcluirRanking,
			&rowdata.MontoPactada,
			&rowdata.MontoEspuelas,
			&rowdata.MontoInscripcion,
			&rowdata.MontoCuota,
			&rowdata.MontoPenalidad,
			&rowdata.MontoOtros,
			&rowdata.OtrosText,
			&rowdata.Estadoid,
			&rowdata.EstadoText,
			&rowdata.UAutoriza,
			&rowdata.FAutoriza,
			&rowdata.Nroperacion,
			/*&rowdata.Ruf1,
			&rowdata.Ruf2,
			&rowdata.Ruf3,
			&rowdata.Iv,
			&rowdata.Salt,
			&rowdata.Checksum,*/
			&rowdata.FCreated,
			&rowdata.FUpdated,
			&rowdata.Activo,
			&rowdata.Estadoreg,
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

// GetOne returns one user by id
func (u *CoqFechaGalponE) GetByUniqueid(token string, uniqueid int) (*CoqFechaGalponE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCoqFechaGalpon

	var rowdata CoqFechaGalponE
	jsonText := fmt.Sprintf(`{"uniqueid":%d}`, uniqueid)
	row := db.QueryRowContext(ctx, query, token, jsonText)

	err := row.Scan(
		&rowdata.Uniqueid,
		//&rowdata.Owner,
		//&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		//&rowdata.Flag1,
		//&rowdata.Flag2,
		&rowdata.CountryCode,
		&rowdata.Campeonatoid,
		&rowdata.Galponid,
		&rowdata.GalponText,
		&rowdata.GalponPadre,
		&rowdata.Urbaniza,
		&rowdata.CriadorText,
		&rowdata.Movil,
		&rowdata.Email,
		&rowdata.FechaNumeroid,
		&rowdata.FechaPelea,
		&rowdata.FechaCerrada,
		&rowdata.ExcluirRanking,
		&rowdata.MontoPactada,
		&rowdata.MontoEspuelas,
		&rowdata.MontoInscripcion,
		&rowdata.MontoCuota,
		&rowdata.MontoPenalidad,
		&rowdata.MontoOtros,
		&rowdata.OtrosText,
		&rowdata.Estadoid,
		&rowdata.EstadoText,
		&rowdata.UAutoriza,
		&rowdata.FAutoriza,
		&rowdata.Nroperacion,
		/*&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,*/
		&rowdata.FCreated,
		&rowdata.FUpdated,
		&rowdata.Activo,
		&rowdata.Estadoreg,
		&rowdata.TotalRecords,
	)

	if err != nil {
		return nil, err
	}

	return &rowdata, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *CoqFechaGalponE) Update(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Se deseenvuelve el JSON del Data para adicionar filtros
	var mapData map[string]interface{}
	json.Unmarshal([]byte(data), &mapData)
	if mapData == nil {
		mapData = make(map[string]interface{})
	}
	// --- Adicion de filtro de tipos de carros
	// mapData["tipo"] = tabla

	// Se empaqueta el JSON del Data
	jsonData, err := json.Marshal(mapData)
	if err != nil {
		log.Println("Error convirtiendo el Dato")
		return nil, err
	}
	log.Println("Data = " + string(jsonData))

	query := `SELECT coq_fecha_por_galpon_save($1, $2, $3)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.QueryContext(ctx, token, string(jsonData), metricas)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var uniqueid int64

	if result.Next() {
		err := result.Scan(&uniqueid)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = uniqueid

	return retorno, nil
}

// Delete deletes one user from the database, by User.ID
func (u *CoqFechaGalponE) Delete(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// Se deseenvuelve el JSON del Data para adicionar filtros
	var mapData map[string]interface{}
	json.Unmarshal([]byte(data), &mapData)
	if mapData == nil {
		mapData = make(map[string]interface{})
	}
	// --- Adicion de estado de eliminacion de record
	mapData["estadoreg"] = 300

	// Se empaqueta el JSON del Data
	jsonData, err := json.Marshal(mapData)
	if err != nil {
		log.Println("Error convirtiendo el Dato")
		return nil, err
	}
	log.Println("Data = " + string(jsonData))

	query := `SELECT coq_fecha_por_galpon_save($1, $2, $3)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.QueryContext(ctx, token, string(jsonData), metricas)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var uniqueid int64

	if result.Next() {
		err := result.Scan(&uniqueid)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = uniqueid

	return retorno, nil
}

// DeleteByID deletes one user from the database, by ID
func (u *CoqFechaGalponE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"id":%d, 
							  "uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, id, 300)

	query := `SELECT coq_fecha_por_galpon_save($1, $2, $3)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	result, err := stmt.QueryContext(ctx, token, jsonText, metricas)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var uniqueid int64

	if result.Next() {
		err := result.Scan(&uniqueid)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = uniqueid

	return retorno, nil
}
