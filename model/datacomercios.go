package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Comercios
type DataComercioE struct {
	Uniqueid        int64       `json:"uniqueid,omitempty"`
	Owner           NullInt32   `json:"owner,omitempty"`
	Dispositivoid   int32       `json:"dispositivoid,omitempty"`
	Id              int32       `json:"id,omitempty"`
	Sede            int32       `json:"sede"`
	Flag1           string      `json:"flag1,omitempty"`
	Flag2           string      `json:"flag2,omitempty"`
	PersonaId       NullInt64   `json:"personaid,omitempty"`
	Secuencial      NullInt32   `json:"secuencial"`
	Orden           NullInt32   `json:"orden"`
	TokenDataId     NullString  `json:"tokendataid,omitempty"`
	TradeName       NullString  `json:"tradename,omitempty"`
	ShortName       NullString  `json:"shortname,omitempty"`
	Tipo            NullString  `json:"tipo,omitempty"`
	DocId           NullString  `json:"doc_id,omitempty"`
	Clasificacion   NullInt32   `json:"clasificacion,omitempty"`
	EmailBills      NullString  `json:"email_bills,omitempty"`
	TipoBarcode     NullString  `json:"tipo_barcode,omitempty"`
	TipoCommerce    NullString  `json:"tipo_commerce,omitempty"`
	Contacto        NullString  `json:"contacto,omitempty"`
	Movil           NullString  `json:"movil,omitempty"`
	GiroId          NullInt64   `json:"giroid,omitempty"`
	GiroText        NullString  `json:"girotext,omitempty"`
	Ciiu            NullInt64   `json:"ciiu,omitempty"`
	CiiuText        NullString  `json:"ciiutext,omitempty"`
	AddressId       NullInt64   `json:"addressid,omitempty"`
	Domicilio       NullString  `json:"domicilio,omitempty"`
	City            NullString  `json:"city,omitempty"`
	Zipcode         NullString  `json:"zipcode,omitempty"`
	Latitud         NullString  `json:"latitud,omitempty"`
	Longitud        NullString  `json:"longitud,omitempty"`
	FstartedAt      NullTime    `json:"fstarted,omitempty"`
	Validated       NullInt32   `json:"validated,omitempty"`
	FvalidatedAt    NullTime    `json:"fvalidated,omitempty"`
	Validatedby     NullString  `json:"validatedby,omitempty"`
	FterminatedAt   NullTime    `json:"fterminated,omitempty"`
	FlastAccessAt   NullTime    `json:"flastaccess,omitempty"`
	FlastMovementAt NullTime    `json:"flastmovement,omitempty"`
	FlastTransactAt NullTime    `json:"flasttransact,omitempty"`
	StatusComercio  NullInt32   `json:"status_comercio,omitempty"`
	StatusDetail    NullString  `json:"status_detail,omitempty"`
	StatusDateAt    NullTime    `json:"status_date,omitempty"`
	Distance        NullFloat64 `json:"distance,omitempty"`
	Ruf1            NullString  `json:"ruf1,omitempty"`
	Ruf2            NullString  `json:"ruf2,omitempty"`
	Ruf3            NullString  `json:"ruf3,omitempty"`
	Iv              NullString  `json:"iv,omitempty"`
	Salt            NullString  `json:"salt,omitempty"`
	Checksum        NullString  `json:"checksum,omitempty"`
	FCreated        NullTime    `json:"fcreated,omitempty"`
	FUpdated        NullTime    `json:"fupdated,omitempty"`
	Estadoreg       NullInt64   `json:"estadoreg,omitempty"`
	Activo          NullInt64   `json:"activo,omitempty"`
	TotalRecords    int64       `json:"total_records"`
}

type DataComercioMinimalE struct {
	Uniqueid       int64       `json:"uniqueid,omitempty"`
	Sede           int32       `json:"sede"`
	Flag1          string      `json:"flag1,omitempty"`
	Flag2          string      `json:"flag2,omitempty"`
	PersonaId      NullInt64   `json:"personaid,omitempty"`
	TokenDataId    NullString  `json:"tokendataid,omitempty"`
	Orden          NullInt32   `json:"orden"`
	TradeName      NullString  `json:"tradename,omitempty"`
	Domicilio      NullString  `json:"domicilio,omitempty"`
	City           NullString  `json:"city,omitempty"`
	Zipcode        NullString  `json:"zipcode,omitempty"`
	Latitud        NullString  `json:"latitud,omitempty"`
	Longitud       NullString  `json:"longitud,omitempty"`
	Distance       NullFloat64 `json:"distance,omitempty"`
	LocationType   NullString  `json:"location_type,omitempty"`
	ParkingLotName NullString  `json:"parking_lot_name,omitempty"`
	ParkingSpots   NullString  `json:"parking_spots,omitempty"`
	TemsUseTicket  NullString  `json:"terms_use_ticket,omitempty"`
	LocationPhoto  NullString  `json:"location_photo,omitempty"`
	FterminatedAt  NullTime    `json:"fterminated,omitempty"`
	Estadoreg      NullInt64   `json:"estadoreg,omitempty"`
	Activo         NullInt64   `json:"activo,omitempty"`
	TotalRecords   int64       `json:"total_records"`
}

func (e DataComercioE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

func (e DataComercioMinimalE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectDataComer = `select * from data_comercios_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *DataComercioE) GetAll(token string, filter string) ([]*DataComercioE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectDataComer

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

	var lista []*DataComercioE

	for rows.Next() {
		var rowdata DataComercioE
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
			&rowdata.TokenDataId,
			&rowdata.TradeName,
			&rowdata.ShortName,
			&rowdata.Tipo,
			&rowdata.DocId,
			&rowdata.Clasificacion,
			&rowdata.EmailBills,
			&rowdata.TipoBarcode,
			&rowdata.TipoCommerce,
			&rowdata.Contacto,
			&rowdata.Movil,
			&rowdata.GiroId,
			&rowdata.GiroText,
			&rowdata.Ciiu,
			&rowdata.CiiuText,
			&rowdata.AddressId,
			&rowdata.Domicilio,
			&rowdata.City,
			&rowdata.Zipcode,
			&rowdata.Latitud,
			&rowdata.Longitud,
			&rowdata.FstartedAt,
			&rowdata.Validated,
			&rowdata.FvalidatedAt,
			&rowdata.Validatedby,
			&rowdata.FterminatedAt,
			&rowdata.FlastAccessAt,
			&rowdata.FlastMovementAt,
			&rowdata.FlastTransactAt,
			&rowdata.StatusComercio,
			&rowdata.StatusDetail,
			&rowdata.StatusDateAt,
			&rowdata.Distance,
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

func (u *DataComercioE) GetAllWithParams(token string, filter string) ([]*DataComercioMinimalE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select * from data_comercios_minimal_list($1, $2)`

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

	var lista []*DataComercioMinimalE

	for rows.Next() {
		var rowdata DataComercioMinimalE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.PersonaId,
			&rowdata.TokenDataId,
			&rowdata.Orden,
			&rowdata.TradeName,
			&rowdata.Domicilio,
			&rowdata.City,
			&rowdata.Zipcode,
			&rowdata.Latitud,
			&rowdata.Longitud,
			&rowdata.Distance,
			&rowdata.LocationType,
			&rowdata.ParkingLotName,
			&rowdata.ParkingSpots,
			&rowdata.TemsUseTicket,
			&rowdata.LocationPhoto,
			&rowdata.FterminatedAt,
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
func (u *DataComercioE) GetByUniqueid(token string, uniqueid int) (*DataComercioE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectDataComer

	var rowdata DataComercioE
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
		&rowdata.TokenDataId,
		&rowdata.TradeName,
		&rowdata.ShortName,
		&rowdata.Tipo,
		&rowdata.DocId,
		&rowdata.Clasificacion,
		&rowdata.EmailBills,
		&rowdata.TipoBarcode,
		&rowdata.TipoCommerce,
		&rowdata.Contacto,
		&rowdata.Movil,
		&rowdata.GiroId,
		&rowdata.GiroText,
		&rowdata.Ciiu,
		&rowdata.CiiuText,
		&rowdata.AddressId,
		&rowdata.Domicilio,
		&rowdata.City,
		&rowdata.Zipcode,
		&rowdata.Latitud,
		&rowdata.Longitud,
		&rowdata.FstartedAt,
		&rowdata.Validated,
		&rowdata.FvalidatedAt,
		&rowdata.Validatedby,
		&rowdata.FterminatedAt,
		&rowdata.FlastAccessAt,
		&rowdata.FlastMovementAt,
		&rowdata.FlastTransactAt,
		&rowdata.StatusComercio,
		&rowdata.StatusDetail,
		&rowdata.StatusDateAt,
		&rowdata.Distance,
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
func (u *DataComercioE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT data_comercios_save($1, $2, $3)`
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
func (u *DataComercioE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT data_comercios_save($1, $2, $3)`
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
func (u *DataComercioE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT data_comercios_save($1, $2, $3)`
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
