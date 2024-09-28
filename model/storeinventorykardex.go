package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Inventory Kardex
type StoreInventoryKardexE struct {
	Uniqueid      int64       `json:"uniqueid,omitempty"`
	Owner         NullInt32   `json:"owner,omitempty"`
	Dispositivoid NullInt32   `json:"dispositivoid,omitempty"`
	Id            int32       `json:"id,omitempty"`
	Sede          int32       `json:"sede"`
	Flag1         string      `json:"flag1,omitempty"`
	Flag2         string      `json:"flag2,omitempty"`
	PersonaId     NullInt64   `json:"personaid,omitempty"`
	TokendataId   NullString  `json:"tokendataid,omitempty"`
	WarehouseId   NullInt64   `json:"warehouseid,omitempty"`
	Fecha         NullTime    `json:"fecha,omitempty"`
	ProductId     NullInt64   `json:"productid,omitempty"`
	Io            NullString  `json:"io,omitempty"`
	TipoMov       NullString  `json:"tipomov,omitempty"`
	ReasonEnumId  NullString  `json:"reasonenumid,omitempty"`
	UDisplay      NullString  `json:"udisplay,omitempty"`
	Quantity      NullFloat64 `json:"quantity,omitempty"`
	Uom           NullString  `json:"uom,omitempty"`
	Entrada       NullFloat64 `json:"entrada,omitempty"`
	Salida        NullFloat64 `json:"salida,omitempty"`
	Balance       NullFloat64 `json:"balance,omitempty"`
	ReceiptId     NullInt64   `json:"receiptid,omitempty"`
	ReceiptText   NullString  `json:"receipttext,omitempty"`
	OrderId       NullInt64   `json:"orderid,omitempty"`
	ShiptmentId   NullInt64   `json:"shiptmentid,omitempty"`
	DocumentText  NullString  `json:"documenttext,omitempty"`
	Notas         NullString  `json:"notas,omitempty"`
	Ruf1          NullString  `json:"ruf1,omitempty"`
	Ruf2          NullString  `json:"ruf2,omitempty"`
	Ruf3          NullString  `json:"ruf3,omitempty"`
	Iv            NullString  `json:"iv,omitempty"`
	Salt          NullString  `json:"salt,omitempty"`
	Checksum      NullString  `json:"checksum,omitempty"`
	FCreated      NullTime    `json:"fcreated,omitempty"`
	FUpdated      NullTime    `json:"fupdated,omitempty"`
	UCreated      NullString  `json:"ucreated,omitempty"`
	UUpdated      NullString  `json:"uupdated,omitempty"`
	Activo        int32       `json:"activo,omitempty"`
	Estadoreg     int32       `json:"estadoreg,omitempty"`
	TotalRecords  int64       `json:"total_records,omitempty"`
}

func (e StoreInventoryKardexE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

func (e StoreInventoryKardexE) CreatedFormat() string {
	return e.FCreated.Time.Format("Jan 2006")
}

const queryListStoreInventoryKardexE = `select uniqueid, sede, flag1, flag2, fecha, productid, io, tipomov, udisplay, quantity, uom, entrada, salida, balance, fcreated, activo, estadoreg, total_records from store_inventory_kardex_list( $1, $2)`
const queryLoadStoreInventoryKardexE = `select * from store_inventory_kardex_list( $1, $2)`
const querySaveStoreInventoryKardexE = `SELECT store_inventory_kardex_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreInventoryKardexE) GetAll(token string, filter string) ([]*StoreInventoryKardexE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreInventoryKardexE

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

	var lista []*StoreInventoryKardexE

	///`select uniqueid, sede, flag1, flag2, fecha, productid, io, tipomov, udisplay, quantity, uom,
	///        entrada, salida, balance, fcreated, activo, estadoreg, total_records from store_inventory_kardex_list( $1, $2)`

	for rows.Next() {
		var rowdata StoreInventoryKardexE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Fecha,
			&rowdata.ProductId,
			&rowdata.Io,
			&rowdata.TipoMov,
			&rowdata.UDisplay,
			&rowdata.Quantity,
			&rowdata.Uom,
			&rowdata.Entrada,
			&rowdata.Salida,
			&rowdata.Balance,
			&rowdata.FCreated,
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
func (u *StoreInventoryKardexE) GetByUniqueid(token string, jsonText string) (*StoreInventoryKardexE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryLoadStoreInventoryKardexE

	var rowdata StoreInventoryKardexE
	log.Println("Where = " + string(jsonText))
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
		&rowdata.Fecha,
		&rowdata.ProductId,
		&rowdata.Io,
		&rowdata.TipoMov,
		&rowdata.ReasonEnumId,
		&rowdata.UDisplay,
		&rowdata.Quantity,
		&rowdata.Uom,
		&rowdata.Entrada,
		&rowdata.Salida,
		&rowdata.Balance,
		&rowdata.ReceiptId,
		&rowdata.ReceiptText,
		&rowdata.OrderId,
		&rowdata.ShiptmentId,
		&rowdata.DocumentText,
		&rowdata.Notas,
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
func (u *StoreInventoryKardexE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreInventoryKardexE
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
func (u *StoreInventoryKardexE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreInventoryKardexE
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
func (u *StoreInventoryKardexE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreInventoryKardexE
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
