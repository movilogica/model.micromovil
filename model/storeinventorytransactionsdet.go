package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Transacciones - items
type StoreInventoryTransactionsDetE struct {
	Uniqueid        int64       `json:"uniqueid,omitempty"`
	Owner           NullInt32   `json:"owner,omitempty"`
	Dispositivoid   NullInt32   `json:"dispositivoid,omitempty"`
	Id              int32       `json:"id,omitempty"`
	Sede            int32       `json:"sede"`
	Flag1           string      `json:"flag1,omitempty"`
	Flag2           string      `json:"flag2,omitempty"`
	PersonaId       NullInt64   `json:"personaid,omitempty"`
	TokendataId     NullString  `json:"tokendataid,omitempty"`
	TransactId      NullInt64   `json:"transactid,omitempty"`
	NroperacionMask NullString  `json:"nroperacionmask,omitempty"`
	Numero          NullInt64   `json:"numero,omitempty"`
	Fecha           NullTime    `json:"fecha,omitempty"`
	TipoMov         NullString  `json:"tipomov,omitempty"`
	Secuencial      NullInt32   `json:"secuencial,omitempty"`
	WarehouseId     NullInt64   `json:"warehouseid,omitempty"`
	WarehouseText   NullString  `json:"warehousetext,omitempty"`
	LocationId      NullInt64   `json:"locationid,omitempty"`
	LocationSeqId   NullString  `json:"locationseqid,omitempty"`
	ToWarehouseId   NullInt64   `json:"towarehouseid,omitempty"`
	ToWarehouseText NullString  `json:"towarehousetext,omitempty"`
	ToLocationId    NullInt64   `json:"tolocationid,omitempty"`
	ToLocationSeqId NullString  `json:"tolocationseqid,omitempty"`
	ProductId       NullInt64   `json:"productid,omitempty"`
	ProductText     NullString  `json:"producttext,omitempty"`
	Lotid           NullString  `json:"lotid,omitempty"`
	SkuNumber       NullString  `json:"skunumber,omitempty"`
	BinNumber       NullString  `json:"binnumber,omitempty"`
	SerialNumber    NullString  `json:"serialnumber,omitempty"`
	UDisplay        NullString  `json:"udisplay,omitempty"`
	Uom             NullString  `json:"uom,omitempty"`
	Quom            NullFloat64 `json:"quom,omitempty"`
	Quantity        NullFloat64 `json:"quantity,omitempty"`
	QTotal          NullFloat64 `json:"qtotal,omitempty"`
	QWeight         NullFloat64 `json:"qweight,omitempty"`
	UomWeight       NullString  `json:"uomweight,omitempty"`
	Xs              NullInt64   `json:"xs,omitempty"`
	S               NullInt64   `json:"s,omitempty"`
	M               NullInt64   `json:"m,omitempty"`
	L               NullInt64   `json:"l,omitempty"`
	Xl              NullInt64   `json:"xl,omitempty"`
	Xxl             NullInt64   `json:"xxl,omitempty"`
	Xxxl            NullInt64   `json:"xxxl,omitempty"`
	Os              NullInt64   `json:"os,omitempty"`
	Total           NullFloat64 `json:"total,omitempty"`
	Ruf1            NullString  `json:"ruf1,omitempty"`
	Ruf2            NullString  `json:"ruf2,omitempty"`
	Ruf3            NullString  `json:"ruf3,omitempty"`
	Iv              NullString  `json:"iv,omitempty"`
	Salt            NullString  `json:"salt,omitempty"`
	Checksum        NullString  `json:"checksum,omitempty"`
	FCreated        NullTime    `json:"fcreated,omitempty"`
	FUpdated        NullTime    `json:"fupdated,omitempty"`
	UCreated        NullString  `json:"ucreated,omitempty"`
	UUpdated        NullString  `json:"uupdated,omitempty"`
	Activo          int32       `json:"activo,omitempty"`
	Estadoreg       int32       `json:"estadoreg,omitempty"`
	TotalRecords    int64       `json:"total_records,omitempty"`
}

func (e StoreInventoryTransactionsDetE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

// La lista de items puede incluir columnas relacionadas de otras tablas
const queryListStoreInventTransDetE = `select uniqueid, sede, flag1, flag2, transactid, nroperacionmask, numero, fecha, tipomov, secuencial, warehousetext, locationseqid, towarehousetext, tolocationseqid, producttext, udisplay, uom, quom, quantity, qtotal, qweight, uomweight, total, activo, estadoreg, total_records from store_inventory_transact_det_list( $1, $2)`
const queryLoadStoreInventTransDetE = `select * from store_inventory_transact_det_list( $1, $2)`
const querySaveStoreInventTransDetE = `SELECT store_inventory_transact_det_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreInventoryTransactionsDetE) GetAll(token string, filter string) ([]*StoreInventoryTransactionsDetE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreInventTransDetE

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

	var lista []*StoreInventoryTransactionsDetE

	/// `uniqueid, sede, flag1, flag2, transactid, secuencial, warehousetext, locationseqid, towarehousetext, tolocationseqid, producttext,
	//   udisplay, uom, quom, quantity, qtotal, qweight, uomweight, total, activo, estadoreg, total_records
	for rows.Next() {
		var rowdata StoreInventoryTransactionsDetE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.TransactId,
			&rowdata.NroperacionMask,
			&rowdata.Numero,
			&rowdata.Fecha,
			&rowdata.TipoMov,
			&rowdata.Secuencial,
			&rowdata.WarehouseText,
			&rowdata.LocationSeqId,
			&rowdata.ToWarehouseText,
			&rowdata.ToLocationSeqId,
			&rowdata.ProductText,
			&rowdata.UDisplay,
			&rowdata.Uom,
			&rowdata.Quom,
			&rowdata.Quantity,
			&rowdata.QTotal,
			&rowdata.QWeight,
			&rowdata.UomWeight,
			&rowdata.Total,
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
	log.Printf("Resultado items = %d records\r\n", len(lista))

	return lista, nil
}

// GetOne returns one user by id
func (u *StoreInventoryTransactionsDetE) GetByUniqueid(token string, jsonText string) (*StoreInventoryTransactionsDetE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryLoadStoreInventTransDetE

	var rowdata StoreInventoryTransactionsDetE
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
		&rowdata.TransactId,
		&rowdata.NroperacionMask,
		&rowdata.Numero,
		&rowdata.Fecha,
		&rowdata.TipoMov,
		&rowdata.Secuencial,
		&rowdata.WarehouseId,
		&rowdata.WarehouseText,
		&rowdata.LocationId,
		&rowdata.LocationSeqId,
		&rowdata.ToWarehouseId,
		&rowdata.ToWarehouseText,
		&rowdata.ToLocationId,
		&rowdata.ToLocationSeqId,
		&rowdata.ProductId,
		&rowdata.ProductText,
		&rowdata.Lotid,
		&rowdata.SkuNumber,
		&rowdata.BinNumber,
		&rowdata.SerialNumber,
		&rowdata.UDisplay,
		&rowdata.Uom,
		&rowdata.Quom,
		&rowdata.Quantity,
		&rowdata.QTotal,
		&rowdata.QWeight,
		&rowdata.UomWeight,
		&rowdata.Xs,
		&rowdata.S,
		&rowdata.M,
		&rowdata.L,
		&rowdata.Xl,
		&rowdata.Xxl,
		&rowdata.Xxxl,
		&rowdata.Os,
		&rowdata.Total,
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
func (u *StoreInventoryTransactionsDetE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreInventTransDetE
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
func (u *StoreInventoryTransactionsDetE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreInventTransDetE
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
func (u *StoreInventoryTransactionsDetE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreInventTransDetE
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

func (u *StoreInventoryTransactionsDetE) UMedidaText() string {
	umedidatext := ""
	/// 20Kgm/Rollo
	if u.Uom != u.UDisplay {
		if u.Quom.Float64 > 1 {
			umedidatext = fmt.Sprintf("%.2f%s//%s", u.Quom.Float64, u.Uom.String, u.UDisplay.String)
		} else {
			umedidatext = u.UDisplay.String
		}
	} else {
		if u.Quom.Float64 > 1 {
			umedidatext = fmt.Sprintf("%.2fUN//%s", u.Quom.Float64, u.UDisplay.String)
		} else {
			umedidatext = u.UDisplay.String
		}
	}
	if umedidatext == "" {
		return "-"
	} else {
		return umedidatext
	}
}
