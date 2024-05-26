package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Personas Address
type DataPersonaAddressE struct {
	Uniqueid         int64       `json:"uniqueid,omitempty"`
	Owner            NullInt32   `json:"owner,omitempty"`
	Dispositivoid    NullInt32   `json:"dispositivoid,omitempty"`
	Id               int32       `json:"id,omitempty"`
	Sede             int32       `json:"sede"`
	Flag1            string      `json:"flag1,omitempty"`
	Flag2            string      `json:"flag2,omitempty"`
	PersonaId        NullInt64   `json:"personaid,omitempty"`
	Secuencial       NullInt32   `json:"secuencial"`
	Orden            NullInt32   `json:"orden"`
	TokenAddress     NullString  `json:"tokenaddress,omitempty"`
	AddressTypeId    NullString  `json:"address_type_id,omitempty"`
	FullAddress      NullString  `json:"full_address,omitempty"`
	StreetAddress    NullString  `json:"street_address,omitempty"`
	NumberAddress    NullString  `json:"number_address,omitempty"`
	CityAddress      NullString  `json:"city_address,omitempty"`
	Zipcode          NullString  `json:"zipcode,omitempty"`
	ZipcodeText      NullString  `json:"zipcodetext,omitempty"`
	RegionId         NullInt64   `json:"regionid,omitempty"`
	RegionText       NullString  `json:"regiontext,omitempty"`
	DepartamentoId   NullInt64   `json:"departamentoid,omitempty"`
	DepartamentoText NullString  `json:"departamentotext,omitempty"`
	ProvinciaId      NullInt64   `json:"provinciaid,omitempty"`
	ProvinciaText    NullString  `json:"provinciatext,omitempty"`
	DistritoId       NullInt64   `json:"distritoid,omitempty"`
	DistritoText     NullString  `json:"distritotext,omitempty"`
	PaisId           NullInt64   `json:"paisid,omitempty"`
	PaisText         NullString  `json:"paistext,omitempty"`
	Latitud          NullString  `json:"latitud,omitempty"`
	Longitud         NullString  `json:"longitud,omitempty"`
	IssuedAt         NullTime    `json:"issued,omitempty"`
	ExpiredAt        NullTime    `json:"expired,omitempty"`
	Notes            NullString  `json:"notes,omitempty"`
	Validated        NullInt32   `json:"validated,omitempty"`
	FvalidatedAt     NullTime    `json:"fvalidated,omitempty"`
	ValidatedBy      NullString  `json:"validatedby,omitempty"`
	StatusAddress    NullInt32   `json:"status_address,omitempty"`
	StatusDetail     NullString  `json:"status_detail,omitempty"`
	StatusDateAt     NullTime    `json:"status_date,omitempty"`
	Foremost         NullInt32   `json:"foremost,omitempty"`
	Distance         NullFloat64 `json:"distance,omitempty"`
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
	TotalRecords     int64       `json:"total_records,omitempty"`
}

func (e DataPersonaAddressE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectDataPerAddress = `select * from data_personas_address_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *DataPersonaAddressE) GetAll(token string, filter string) ([]*DataPersonaAddressE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectDataPerAddress

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

	var lista []*DataPersonaAddressE

	for rows.Next() {
		var rowdata DataPersonaAddressE
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
			&rowdata.TokenAddress,
			&rowdata.AddressTypeId,
			&rowdata.FullAddress,
			&rowdata.StreetAddress,
			&rowdata.NumberAddress,
			&rowdata.CityAddress,
			&rowdata.Zipcode,
			&rowdata.ZipcodeText,
			&rowdata.RegionId,
			&rowdata.RegionText,
			&rowdata.DepartamentoId,
			&rowdata.DepartamentoText,
			&rowdata.ProvinciaId,
			&rowdata.ProvinciaText,
			&rowdata.DistritoId,
			&rowdata.DistritoText,
			&rowdata.PaisId,
			&rowdata.PaisText,
			&rowdata.Latitud,
			&rowdata.Longitud,
			&rowdata.IssuedAt,
			&rowdata.ExpiredAt,
			&rowdata.Notes,
			&rowdata.Validated,
			&rowdata.FvalidatedAt,
			&rowdata.ValidatedBy,
			&rowdata.StatusAddress,
			&rowdata.StatusDetail,
			&rowdata.StatusDateAt,
			&rowdata.Foremost,
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

// GetOne returns one user by id
func (u *DataPersonaAddressE) GetByUniqueid(token string, uniqueid int) (*DataPersonaAddressE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectDataPerAddress

	var rowdata DataPersonaAddressE
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
		&rowdata.TokenAddress,
		&rowdata.AddressTypeId,
		&rowdata.FullAddress,
		&rowdata.StreetAddress,
		&rowdata.NumberAddress,
		&rowdata.CityAddress,
		&rowdata.Zipcode,
		&rowdata.ZipcodeText,
		&rowdata.RegionId,
		&rowdata.RegionText,
		&rowdata.DepartamentoId,
		&rowdata.DepartamentoText,
		&rowdata.ProvinciaId,
		&rowdata.ProvinciaText,
		&rowdata.DistritoId,
		&rowdata.DistritoText,
		&rowdata.PaisId,
		&rowdata.PaisText,
		&rowdata.Latitud,
		&rowdata.Longitud,
		&rowdata.IssuedAt,
		&rowdata.ExpiredAt,
		&rowdata.Notes,
		&rowdata.Validated,
		&rowdata.FvalidatedAt,
		&rowdata.ValidatedBy,
		&rowdata.StatusAddress,
		&rowdata.StatusDetail,
		&rowdata.StatusDateAt,
		&rowdata.Foremost,
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
func (u *DataPersonaAddressE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT data_personas_address_save($1, $2, $3)`
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
func (u *DataPersonaAddressE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT data_personas_address_save($1, $2, $3)`
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
func (u *DataPersonaAddressE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT data_personas_address_save($1, $2, $3)`
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
