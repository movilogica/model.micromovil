package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Almacenes Locaciones
type StoreWarehouseLocationsE struct {
	Uniqueid           int64      `json:"uniqueid,omitempty"`
	Owner              NullInt32  `json:"owner,omitempty"`
	Dispositivoid      NullInt32  `json:"dispositivoid,omitempty"`
	Id                 int32      `json:"id,omitempty"`
	Sede               int32      `json:"sede"`
	Flag1              string     `json:"flag1,omitempty"`
	Flag2              string     `json:"flag2,omitempty"`
	PersonaId          NullInt64  `json:"personaid,omitempty"`
	TokendataId        NullString `json:"tokendataid,omitempty"`
	WarehouseId        NullInt64  `json:"warehouseid,omitempty"`
	LocationSeqId      NullString `json:"locationseqid,omitempty"`
	LocationTypeEnumId NullString `json:"locationtypeenumid,omitempty"`
	LocationStatusId   NullString `json:"locationstatusid,omitempty"`
	StageId            NullString `json:"stageid,omitempty"`
	StorageTypeId      NullString `json:"storagetypeid,omitempty"`
	AreaId             NullString `json:"areaid,omitempty"`
	AisleId            NullString `json:"aisleid,omitempty"`
	SectionId          NullString `json:"sectionid,omitempty"`
	LevelId            NullString `json:"levelid,omitempty"`
	PositionId         NullString `json:"positionid,omitempty"`
	Permanent          NullInt32  `json:"permanent,omitempty"`
	Ruf1               NullString `json:"ruf1,omitempty"`
	Ruf2               NullString `json:"ruf2,omitempty"`
	Ruf3               NullString `json:"ruf3,omitempty"`
	Iv                 NullString `json:"iv,omitempty"`
	Salt               NullString `json:"salt,omitempty"`
	Checksum           NullString `json:"checksum,omitempty"`
	FCreated           NullTime   `json:"fcreated,omitempty"`
	FUpdated           NullTime   `json:"fupdated,omitempty"`
	UCreated           NullString `json:"ucreated,omitempty"`
	UUpdated           NullString `json:"uupdated,omitempty"`
	Activo             int32      `json:"activo,omitempty"`
	Estadoreg          int32      `json:"estadoreg,omitempty"`
	TotalRecords       int64      `json:"total_records,omitempty"`
}

func (e StoreWarehouseLocationsE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const queryListStoreWarehouseLocationsE = `select * from store_warehouse_locations_list( $1, $2)`
const querySaveStoreWarehouseLocationsE = `SELECT store_warehouse_locations_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreWarehouseLocationsE) GetAll(token string, filter string) ([]*StoreWarehouseLocationsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreWarehouseLocationsE

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

	var lista []*StoreWarehouseLocationsE

	for rows.Next() {
		var rowdata StoreWarehouseLocationsE
		err := rows.Scan(
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
			&rowdata.LocationSeqId,
			&rowdata.LocationTypeEnumId,
			&rowdata.LocationStatusId,
			&rowdata.StageId,
			&rowdata.StorageTypeId,
			&rowdata.AreaId,
			&rowdata.AisleId,
			&rowdata.SectionId,
			&rowdata.LevelId,
			&rowdata.PositionId,
			&rowdata.Permanent,
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
func (u *StoreWarehouseLocationsE) GetByUniqueid(token string, uniqueid int) (*StoreWarehouseLocationsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreWarehouseLocationsE

	var rowdata StoreWarehouseLocationsE
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
		&rowdata.TokendataId,
		&rowdata.WarehouseId,
		&rowdata.LocationSeqId,
		&rowdata.LocationTypeEnumId,
		&rowdata.LocationStatusId,
		&rowdata.StageId,
		&rowdata.StorageTypeId,
		&rowdata.AreaId,
		&rowdata.AisleId,
		&rowdata.SectionId,
		&rowdata.LevelId,
		&rowdata.PositionId,
		&rowdata.Permanent,
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
func (u *StoreWarehouseLocationsE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreWarehouseLocationsE
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
func (u *StoreWarehouseLocationsE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreWarehouseLocationsE
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
func (u *StoreWarehouseLocationsE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreWarehouseLocationsE
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