package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Inventory Recuento
type StoreRecountE struct {
	Uniqueid      int64                  `json:"uniqueid,omitempty"`
	Owner         NullInt32              `json:"owner,omitempty"`
	Dispositivoid NullInt32              `json:"dispositivoid,omitempty"`
	Id            int32                  `json:"id,omitempty"`
	Sede          int32                  `json:"sede"`
	Flag1         string                 `json:"flag1,omitempty"`
	Flag2         string                 `json:"flag2,omitempty"`
	PersonaId     NullInt64              `json:"personaid,omitempty"`
	TokendataId   NullString             `json:"tokendataid,omitempty"`
	WarehouseId   NullInt64              `json:"warehouseid,omitempty"`
	Title         NullString             `json:"title,omitempty"`
	Notes         NullString             `json:"notes,omitempty"`
	FromDate      NullTime               `json:"fromdate,omitempty"`
	ToDate        NullTime               `json:"todate,omitempty"`
	Location      NullInt32              `json:"location,omitempty"`
	Closed        NullInt32              `json:"closed,omitempty"`
	Ruf1          NullString             `json:"ruf1,omitempty"`
	Ruf2          NullString             `json:"ruf2,omitempty"`
	Ruf3          NullString             `json:"ruf3,omitempty"`
	Iv            NullString             `json:"iv,omitempty"`
	Salt          NullString             `json:"salt,omitempty"`
	Checksum      NullString             `json:"checksum,omitempty"`
	FCreated      NullTime               `json:"fcreated,omitempty"`
	FUpdated      NullTime               `json:"fupdated,omitempty"`
	UCreated      NullString             `json:"ucreated,omitempty"`
	UUpdated      NullString             `json:"uupdated,omitempty"`
	Activo        int32                  `json:"activo,omitempty"`
	Estadoreg     int32                  `json:"estadoreg,omitempty"`
	TotalRecords  int64                  `json:"total_records,omitempty"`
	Boxes         []*StoreRecountBoxesE  `json:"boxes:omitempty"`
	Groups        []*StoreRecountGroupsE `json:"groups:omitempty"`
	Items         []*StoreRecountItemE   `json:"items:omitempty"`
}

func (e StoreRecountE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

func (e StoreRecountE) CreatedFormat() string {
	return e.FCreated.Time.Format("Jan 2006")
}

const queryListStoreRecountE = `select uniqueid, sede, flag1, flag2, warehouseid, title, fromdate, todate, location, closed, activo, estadoreg, total_records from store_recount_list( $1, $2)`
const queryLoadStoreRecountE = `select * from store_recount_list( $1, $2)`
const querySaveStoreRecountE = `SELECT store_recount_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreRecountE) GetAll(token string, filter string) ([]*StoreRecountE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreRecountE

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

	var lista []*StoreRecountE

	///uniqueid, sede, flag1, flag2, warehouseid, title, fromdate, todate, location, closed, activo, estadoreg, total_records
	for rows.Next() {
		var rowdata StoreRecountE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.WarehouseId,
			&rowdata.Title,
			&rowdata.FromDate,
			&rowdata.ToDate,
			&rowdata.Location,
			&rowdata.Closed,
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
func (u *StoreRecountE) GetByUniqueid(token string, jsonText string) (*StoreRecountE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryLoadStoreRecountE

	var rowdata StoreRecountE
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
		&rowdata.Title,
		&rowdata.Notes,
		&rowdata.FromDate,
		&rowdata.ToDate,
		&rowdata.Location,
		&rowdata.Closed,
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
func (u *StoreRecountE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreRecountE
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
func (u *StoreRecountE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreRecountE
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
func (u *StoreRecountE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreRecountE
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
