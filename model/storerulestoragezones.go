package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Reglas de Almacenaje - Zonas
type StoreRuleStorageZonesE struct {
	Uniqueid       int64                       `json:"uniqueid,omitempty"`
	Owner          NullInt32                   `json:"owner,omitempty"`
	Dispositivoid  NullInt32                   `json:"dispositivoid,omitempty"`
	Id             int32                       `json:"id,omitempty"`
	Sede           int32                       `json:"sede"`
	Flag1          string                      `json:"flag1,omitempty"`
	Flag2          string                      `json:"flag2,omitempty"`
	PersonaId      NullInt64                   `json:"personaid,omitempty"`
	TokendataId    NullString                  `json:"tokendataid,omitempty"`
	RuleStoraId    NullInt64                   `json:"rulestoraid,omitempty"`
	Secuencial     NullInt32                   `json:"secuencial,omitempty"`
	Orden          NullInt32                   `json:"orden,omitempty"`
	ZoneId         NullInt64                   `json:"zoneid,omitempty"`
	ZoneText       NullString                  `json:"zonetext,omitempty"`
	CategUbicaId   NullInt64                   `json:"categubicaid,omitempty"`
	CategUbicaText NullString                  `json:"categubicatext,omitempty"`
	Ruf1           NullString                  `json:"ruf1,omitempty"`
	Ruf2           NullString                  `json:"ruf2,omitempty"`
	Ruf3           NullString                  `json:"ruf3,omitempty"`
	Iv             NullString                  `json:"iv,omitempty"`
	Salt           NullString                  `json:"salt,omitempty"`
	Checksum       NullString                  `json:"checksum,omitempty"`
	FCreated       NullTime                    `json:"fcreated,omitempty"`
	FUpdated       NullTime                    `json:"fupdated,omitempty"`
	UCreated       NullString                  `json:"ucreated,omitempty"`
	UUpdated       NullString                  `json:"uupdated,omitempty"`
	Activo         int32                       `json:"activo,omitempty"`
	Estadoreg      int32                       `json:"estadoreg,omitempty"`
	TotalRecords   int64                       `json:"total_records,omitempty"`
	Actions        []*StoreRuleStorageActionsE `json:"actions:omitempty"`
}

func (e StoreRuleStorageZonesE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const queryListStoreRuleStorageZonesE = `select uniqueid, sede, flag1, flag2, orden, zoneid, zonetext, categubicatext, activo, estadoreg, total_records from store_rules_storage_zones_list( $1, $2)`
const queryLoadStoreRuleStorageZonesE = `select * from store_rules_storage_zones_list( $1, $2)`
const querySaveStoreRuleStorageZonesE = `SELECT store_rules_storage_zones_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreRuleStorageZonesE) GetAll(token string, filter string) ([]*StoreRuleStorageZonesE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreRuleStorageZonesE

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

	var lista []*StoreRuleStorageZonesE

	for rows.Next() {
		var rowdata StoreRuleStorageZonesE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Orden,
			&rowdata.ZoneId,
			&rowdata.ZoneText,
			&rowdata.CategUbicaText,
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
func (u *StoreRuleStorageZonesE) GetByUniqueid(token string, jsonText string) (*StoreRuleStorageZonesE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryLoadStoreRuleStorageZonesE

	var rowdata StoreRuleStorageZonesE
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
		&rowdata.RuleStoraId,
		&rowdata.Secuencial,
		&rowdata.Orden,
		&rowdata.ZoneId,
		&rowdata.ZoneText,
		&rowdata.CategUbicaId,
		&rowdata.CategUbicaText,
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
func (u *StoreRuleStorageZonesE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreRuleStorageZonesE
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
func (u *StoreRuleStorageZonesE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreRuleStorageZonesE
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
func (u *StoreRuleStorageZonesE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreRuleStorageZonesE
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
