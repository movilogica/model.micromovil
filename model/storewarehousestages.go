package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Almacenes Stages
type StoreWarehouseStagesE struct {
	Uniqueid      int64      `json:"uniqueid,omitempty"`
	Owner         NullInt32  `json:"owner,omitempty"`
	Dispositivoid NullInt32  `json:"dispositivoid,omitempty"`
	Id            int32      `json:"id,omitempty"`
	Sede          int32      `json:"sede"`
	Flag1         string     `json:"flag1,omitempty"`
	Flag2         string     `json:"flag2,omitempty"`
	PersonaId     NullInt64  `json:"personaid,omitempty"`
	TokendataId   NullString `json:"tokendataid,omitempty"`
	WarehouseId   NullInt64  `json:"warehouseid,omitempty"`
	Secuencial    NullInt32  `json:"secuencial,omitempty"`
	StageId       NullString `json:"stageid,omitempty"`
	Descrip       NullString `json:"descrip,omitempty"`
	Ruf1          NullString `json:"ruf1,omitempty"`
	Ruf2          NullString `json:"ruf2,omitempty"`
	Ruf3          NullString `json:"ruf3,omitempty"`
	Iv            NullString `json:"iv,omitempty"`
	Salt          NullString `json:"salt,omitempty"`
	Checksum      NullString `json:"checksum,omitempty"`
	FCreated      NullTime   `json:"fcreated,omitempty"`
	FUpdated      NullTime   `json:"fupdated,omitempty"`
	UCreated      NullString `json:"ucreated,omitempty"`
	UUpdated      NullString `json:"uupdated,omitempty"`
	Activo        int32      `json:"activo,omitempty"`
	Estadoreg     int32      `json:"estadoreg,omitempty"`
	TotalRecords  int64      `json:"total_records,omitempty"`
}

func (e StoreWarehouseStagesE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const queryListStoreWarehouseStagesE = `select uniqueid, sede, flag1, flag2, secuencial, stageid, descrip, activo, estadoreg, total_records from store_warehouse_stages_list( $1, $2)`
const queryLoadStoreWarehouseStagesE = `select * from store_warehouse_stages_list( $1, $2)`
const querySaveStoreWarehouseStagesE = `SELECT store_warehouse_stages_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreWarehouseStagesE) GetAll(token string, filter string) ([]*StoreWarehouseStagesE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreWarehouseStagesE

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

	var lista []*StoreWarehouseStagesE

	///uniqueid, sede, flag1, flag2, secuencial, stageid, descrip, activo, estadoreg, total_records
	for rows.Next() {
		var rowdata StoreWarehouseStagesE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Secuencial,
			&rowdata.StageId,
			&rowdata.Descrip,
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
func (u *StoreWarehouseStagesE) GetByUniqueid(token string, jsonText string) (*StoreWarehouseStagesE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryLoadStoreWarehouseStagesE

	var rowdata StoreWarehouseStagesE
	log.Printf("[%s] Where = %s\n", query, string(jsonText))
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
		&rowdata.TokendataId,
		&rowdata.WarehouseId,
		&rowdata.Secuencial,
		&rowdata.StageId,
		&rowdata.Descrip,
		&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,
		&rowdata.FCreated,
		&rowdata.FUpdated,
		&rowdata.UCreated,
		&rowdata.UUpdated,
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
func (u *StoreWarehouseStagesE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreWarehouseStagesE
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
func (u *StoreWarehouseStagesE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreWarehouseStagesE
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
func (u *StoreWarehouseStagesE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreWarehouseStagesE
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
