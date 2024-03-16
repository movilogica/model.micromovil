package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Terminales
type DataTerminalE struct {
	Uniqueid       int64      `json:"uniqueid,omitempty"`
	Owner          NullInt32  `json:"owner,omitempty"`
	Dispositivoid  int32      `json:"dispositivoid,omitempty"`
	Id             int32      `json:"id,omitempty"`
	Sede           int32      `json:"sede"`
	Flag1          string     `json:"flag1,omitempty"`
	Flag2          string     `json:"flag2,omitempty"`
	PersonaId      NullInt64  `json:"personaid,omitempty"`
	Secuencial     NullInt32  `json:"secuencial"`
	Orden          NullInt32  `json:"orden"`
	TokenTerminal  NullString `json:"tokenterminal,omitempty"`
	Movil          NullString `json:"movil,omitempty"`
	Descrip        NullString `json:"descrip,omitempty"`
	RegisteredAt   NullTime   `json:"registered,omitempty"`
	TerminatedAt   NullTime   `json:"terminated,omitempty"`
	LeavedAt       NullTime   `json:"leaved,omitempty"`
	Validated      NullInt32  `json:"validated,omitempty"`
	FvalidatedAt   NullTime   `json:"fvalidated,omitempty"`
	Validatedby    NullString `json:"validatedby,omitempty"`
	StatusTerminal NullInt32  `json:"status_terminal,omitempty"`
	StatusDetail   NullString `json:"status_detail,omitempty"`
	StatusDateAt   NullTime   `json:"status_date,omitempty"`
	RemoteMachine  NullString `json:"remotemachine,omitempty"`
	RemoteHost     NullString `json:"remotehost,omitempty"`
	RemotePort     NullInt32  `json:"remoteport,omitempty"`
	HeaderData     NullString `json:"headerdata,omitempty"`
	HeaderChecksum NullString `json:"headerchecksum,omitempty"`
	Imei           NullString `json:"imei,omitempty"`
	Cellinfo       NullString `json:"cellinfo,omitempty"`
	Latitud        NullString `json:"latitud,omitempty"`
	Longitud       NullString `json:"longitud,omitempty"`
	Ruf1           NullString `json:"ruf1,omitempty"`
	Ruf2           NullString `json:"ruf2,omitempty"`
	Ruf3           NullString `json:"ruf3,omitempty"`
	Iv             NullString `json:"iv,omitempty"`
	Salt           NullString `json:"salt,omitempty"`
	Checksum       NullString `json:"checksum,omitempty"`
	FCreated       NullTime   `json:"fcreated,omitempty"`
	FUpdated       NullTime   `json:"fupdated,omitempty"`
	Activo         int32      `json:"activo,omitempty"`
	Estadoreg      int32      `json:"estadoreg,omitempty"`
	TotalRecords   int64      `json:"total_records"`
}

func (e DataTerminalE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectDataTermin = `select * from data_terminales_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *DataTerminalE) GetAll(token string, filter string) ([]*DataTerminalE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectDataTermin

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

	var lista []*DataTerminalE

	for rows.Next() {
		var rowdata DataTerminalE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.PersonaId,
			&rowdata.Secuencial,
			&rowdata.Orden,
			&rowdata.TokenTerminal,
			&rowdata.Movil,
			&rowdata.Descrip,
			&rowdata.RegisteredAt,
			&rowdata.TerminatedAt,
			&rowdata.LeavedAt,
			&rowdata.Validated,
			&rowdata.FvalidatedAt,
			&rowdata.Validatedby,
			&rowdata.StatusTerminal,
			&rowdata.StatusDetail,
			&rowdata.StatusDateAt,
			&rowdata.RemoteMachine,
			&rowdata.RemoteHost,
			&rowdata.RemotePort,
			&rowdata.HeaderData,
			&rowdata.HeaderChecksum,
			&rowdata.Imei,
			&rowdata.Cellinfo,
			&rowdata.Latitud,
			&rowdata.Longitud,
			&rowdata.Ruf1,
			&rowdata.Ruf2,
			&rowdata.Ruf3,
			&rowdata.Iv,
			&rowdata.Salt,
			&rowdata.Checksum,
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
func (u *DataTerminalE) GetByUniqueid(token string, uniqueid int) (*DataTerminalE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectDataTermin

	var rowdata DataTerminalE
	jsonText := fmt.Sprintf(`{"uniqueid":%d}`, uniqueid)
	row := db.QueryRowContext(ctx, query, token, jsonText)

	err := row.Scan(
		&rowdata.Uniqueid,
		&rowdata.Owner,
		&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		&rowdata.Flag1,
		&rowdata.Flag2,
		&rowdata.PersonaId,
		&rowdata.Secuencial,
		&rowdata.Orden,
		&rowdata.TokenTerminal,
		&rowdata.Movil,
		&rowdata.Descrip,
		&rowdata.RegisteredAt,
		&rowdata.TerminatedAt,
		&rowdata.LeavedAt,
		&rowdata.Validated,
		&rowdata.FvalidatedAt,
		&rowdata.Validatedby,
		&rowdata.StatusTerminal,
		&rowdata.StatusDetail,
		&rowdata.StatusDateAt,
		&rowdata.RemoteMachine,
		&rowdata.RemoteHost,
		&rowdata.RemotePort,
		&rowdata.HeaderData,
		&rowdata.HeaderChecksum,
		&rowdata.Imei,
		&rowdata.Cellinfo,
		&rowdata.Latitud,
		&rowdata.Longitud,
		&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,
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
func (u *DataTerminalE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT data_terminales_save($1, $2, $3)`
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
func (u *DataTerminalE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT data_terminales_save($1, $2, $3)`
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
func (u *DataTerminalE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT data_terminales_save($1, $2, $3)`
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
