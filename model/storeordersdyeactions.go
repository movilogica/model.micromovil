package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Ordenes Dye House - Actions
type StoreOrdersDyeActionsE struct {
	Uniqueid       int64       `json:"uniqueid,omitempty"`
	Owner          NullInt32   `json:"owner,omitempty"`
	Dispositivoid  NullInt32   `json:"dispositivoid,omitempty"`
	Id             int32       `json:"id,omitempty"`
	Sede           int32       `json:"sede"`
	Flag1          string      `json:"flag1,omitempty"`
	Flag2          string      `json:"flag2,omitempty"`
	PersonaId      NullInt64   `json:"personaid,omitempty"`
	TokendataId    NullString  `json:"tokendataid,omitempty"`
	OrderId        NullInt64   `json:"orderid,omitempty"`
	OrderitemId    NullInt64   `json:"orderitemid,omitempty"`
	Nroperacion    NullString  `json:"nroperacion,omitempty"`
	Fecha          NullTime    `json:"fecha,omitempty"`
	ActionText     NullString  `json:"actiontext,omitempty"`
	Secuencial     NullInt32   `json:"secuencial,omitempty"`
	WarehouseId    NullInt64   `json:"warehouseid,omitempty"`
	LocationId     NullInt64   `json:"locationid,omitempty"`
	LocationSeqId  NullString  `json:"locationseqid,omitempty"`
	HistoricoId    NullInt64   `json:"historicoid,omitempty"`
	ProductId      NullInt64   `json:"productid,omitempty"`
	ProductCode    NullString  `json:"productcode,omitempty"`
	ProductText    NullString  `json:"producttext,omitempty"`
	Lotid          NullString  `json:"lotid,omitempty"`
	LotExpired     NullTime    `json:"lotexpired,omitempty"`
	SkuNumber      NullString  `json:"skunumber,omitempty"`
	SerialNumber   NullString  `json:"serialnumber,omitempty"`
	Containerid    NullString  `json:"containerid,omitempty"`
	BinNumber      NullString  `json:"binnumber,omitempty"`
	Softidentifier NullString  `json:"softidentifier,omitempty"`
	Barcodebox     NullString  `json:"barcodebox,omitempty"`
	Barcodeitem    NullString  `json:"barcodeitem,omitempty"`
	UnitPrice      NullFloat64 `json:"unitprice,omitempty"`
	UDisplay       NullString  `json:"udisplay,omitempty"`
	Uom            NullString  `json:"uom,omitempty"`
	Quom           NullFloat64 `json:"quom,omitempty"`
	Quantity       NullFloat64 `json:"quantity,omitempty"`
	Xs             NullInt64   `json:"xs,omitempty"`
	S              NullInt64   `json:"s,omitempty"`
	M              NullInt64   `json:"m,omitempty"`
	L              NullInt64   `json:"l,omitempty"`
	Xl             NullInt64   `json:"xl,omitempty"`
	Xxl            NullInt64   `json:"xxl,omitempty"`
	Xxxl           NullInt64   `json:"xxxl,omitempty"`
	Os             NullInt64   `json:"os,omitempty"`
	QTotal         NullFloat64 `json:"qtotal,omitempty"`
	QWeight        NullFloat64 `json:"qweight,omitempty"`
	Totalunits     NullFloat64 `json:"totalunits,omitempty"`
	Totalprice     NullFloat64 `json:"totalprice,omitempty"`
	Notes          NullString  `json:"notes,omitempty"`
	Ruf1           NullString  `json:"ruf1,omitempty"`
	Ruf2           NullString  `json:"ruf2,omitempty"`
	Ruf3           NullString  `json:"ruf3,omitempty"`
	Iv             NullString  `json:"iv,omitempty"`
	Salt           NullString  `json:"salt,omitempty"`
	Checksum       NullString  `json:"checksum,omitempty"`
	FCreated       NullTime    `json:"fcreated,omitempty"`
	FUpdated       NullTime    `json:"fupdated,omitempty"`
	UCreated       NullString  `json:"ucreated,omitempty"`
	UUpdated       NullString  `json:"uupdated,omitempty"`
	Activo         int32       `json:"activo,omitempty"`
	Estadoreg      int32       `json:"estadoreg,omitempty"`
	TotalRecords   int64       `json:"total_records,omitempty"`
}

func (e StoreOrdersDyeActionsE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const queryListStoreOrdersDyeActionsE = `select uniqueid, sede, flag1, flag2, fecha, actiontext, warehouseid, locationseqid, historicoid, secuencial, productid, productcode, producttext, lotid, lotexpired, skunumber, binnumber, serialnumber, unitprice, udisplay, uom, quom, quantity, qtotal, qweight, totalunits, totalprice, activo, estadoreg, total_records from store_orders_dye_actions_list( $1, $2)`
const queryLoadStoreOrdersDyeActionsE = `select * from store_orders_dye_actions_list( $1, $2)`
const querySaveStoreOrdersDyeActionsE = `SELECT store_orders_dye_actions_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreOrdersDyeActionsE) GetAll(token string, filter string) ([]*StoreOrdersDyeActionsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreOrdersDyeActionsE

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

	var lista []*StoreOrdersDyeActionsE

	/// `select uniqueid, sede, flag1, flag2, fecha, actiontext, warehouseid, locationseqid,
	///         secuencial, productid, productcode, producttext, lotid, lotexpired, skunumber,
	///         binnumber, serialnumber, unitprice, udisplay, uom, quom, quantity, qtotal, qweight,
	///         totalunits, totalprice, activo, estadoreg, total_records
	for rows.Next() {
		var rowdata StoreOrdersDyeActionsE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Fecha,
			&rowdata.ActionText,
			&rowdata.WarehouseId,
			&rowdata.LocationSeqId,
			&rowdata.HistoricoId,
			&rowdata.Secuencial,
			&rowdata.ProductId,
			&rowdata.ProductCode,
			&rowdata.ProductText,
			&rowdata.Lotid,
			&rowdata.LotExpired,
			&rowdata.SkuNumber,
			&rowdata.BinNumber,
			&rowdata.SerialNumber,
			&rowdata.UnitPrice,
			&rowdata.UDisplay,
			&rowdata.Uom,
			&rowdata.Quom,
			&rowdata.Quantity,
			&rowdata.QTotal,
			&rowdata.QWeight,
			&rowdata.Totalunits,
			&rowdata.Totalprice,
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
	log.Printf("Resultado = %d records\r\n", len(lista))

	return lista, nil
}

// GetOne returns one user by id
func (u *StoreOrdersDyeActionsE) GetByUniqueid(token string, jsonText string) (*StoreOrdersDyeActionsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryLoadStoreOrdersDyeActionsE

	var rowdata StoreOrdersDyeActionsE
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
		&rowdata.OrderId,
		&rowdata.OrderitemId,
		&rowdata.Nroperacion,
		&rowdata.Fecha,
		&rowdata.ActionText,
		&rowdata.Secuencial,
		&rowdata.WarehouseId,
		&rowdata.LocationId,
		&rowdata.LocationSeqId,
		&rowdata.HistoricoId,
		&rowdata.ProductId,
		&rowdata.ProductCode,
		&rowdata.ProductText,
		&rowdata.Lotid,
		&rowdata.LotExpired,
		&rowdata.SkuNumber,
		&rowdata.SerialNumber,
		&rowdata.Containerid,
		&rowdata.BinNumber,
		&rowdata.Softidentifier,
		&rowdata.Barcodebox,
		&rowdata.Barcodeitem,
		&rowdata.UnitPrice,
		&rowdata.UDisplay,
		&rowdata.Uom,
		&rowdata.Quom,
		&rowdata.Quantity,
		&rowdata.QTotal,
		&rowdata.QWeight,
		&rowdata.Xs,
		&rowdata.S,
		&rowdata.M,
		&rowdata.L,
		&rowdata.Xl,
		&rowdata.Xxl,
		&rowdata.Xxxl,
		&rowdata.Os,
		&rowdata.Totalunits,
		&rowdata.Totalprice,
		&rowdata.Notes,
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
func (u *StoreOrdersDyeActionsE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreOrdersDyeActionsE
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
func (u *StoreOrdersDyeActionsE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreOrdersDyeActionsE
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
func (u *StoreOrdersDyeActionsE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreOrdersDyeActionsE
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
