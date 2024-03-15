package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Customer Medios
type BizPaymentE struct {
	Uniqueid            int64                `json:"uniqueid,omitempty"`
	Owner               NullInt32            `json:"owner,omitempty"`
	Dispositivoid       int32                `json:"dispositivoid,omitempty"`
	Id                  int32                `json:"id,omitempty"`
	Sede                int32                `json:"sede"`
	Flag1               string               `json:"flag1,omitempty"`
	Flag2               string               `json:"flag2,omitempty"`
	PersonaId           NullInt64            `json:"personaid,omitempty"`
	TokendataId         NullString           `json:"tokendataid,omitempty"`
	BizPersonaId        NullInt64            `json:"bizpersonaid,omitempty"`
	Nroperacion         NullString           `json:"nroperacion,omitempty"`
	Maskoperacion       NullString           `json:"maskoperacion,omitempty"`
	Numero              NullString           `json:"numero,omitempty"`
	NumeroText          NullString           `json:"numerotext,omitempty"`
	Fecha               NullTime             `json:"fecha,omitempty"`
	PaymentTypeId       NullString           `json:"paymenttypeid,omitempty"`
	PaymentMethodTypeId NullString           `json:"paymentmethodtypeid,omitempty"`
	PaymentText         NullString           `json:"paymenttext,omitempty"`
	BizPersonaText      NullString           `json:"bizpersonatext,omitempty"`
	BizPersonaRuc       NullString           `json:"bizpersonaruc,omitempty"`
	StatusId            NullString           `json:"statusid,omitempty"`
	Secuencial          NullInt32            `json:"secuencial,omitempty"`
	Amount              NullFloat64          `json:"amount,omitempty"`
	DivisaId            NullInt64            `json:"divisaid,omitempty"`
	DivisaText          NullString           `json:"divisatext,omitempty"`
	DivisaSimbolo       NullString           `json:"divisasimbolo,omitempty"`
	DivisaDecimal       NullInt32            `json:"divisadecimal,omitempty"`
	TasaVenta           NullFloat64          `json:"tasaventa,omitempty"`
	TasaCompra          NullFloat64          `json:"tasacompra,omitempty"`
	PaymentRefNum       NullString           `json:"paymentrefnum,omitempty"`
	BancoId             NullInt64            `json:"bancoid,omitempty"`
	BancoText           NullString           `json:"bancotext,omitempty"`
	MarcaTarjeta        NullString           `json:"marcatarjeta,omitempty"`
	UltimosDigitos      NullString           `json:"ultimosdigitos,omitempty"`
	Fdiferida           NullTime             `json:"fdiferida,omitempty"`
	Ruf1                NullString           `json:"ruf1,omitempty"`
	Ruf2                NullString           `json:"ruf2,omitempty"`
	Ruf3                NullString           `json:"ruf3,omitempty"`
	Iv                  NullString           `json:"iv,omitempty"`
	Salt                NullString           `json:"salt,omitempty"`
	Checksum            NullString           `json:"checksum,omitempty"`
	FCreated            NullTime             `json:"fcreated,omitempty"`
	FUpdated            NullTime             `json:"fupdated,omitempty"`
	UCreated            NullString           `json:"ucreated,omitempty"`
	UUpdated            NullString           `json:"uupdated,omitempty"`
	Estadoreg           NullInt64            `json:"estadoreg,omitempty"`
	Activo              NullInt64            `json:"activo,omitempty"`
	TotalRecords        int64                `json:"total_records"`
	Status              []*BizPaymentStatusE `json:"status:omitempty"`
}

type BizPaymentStatusE struct {
	Uniqueid      int64      `json:"uniqueid,omitempty"`
	Owner         NullInt32  `json:"owner,omitempty"`
	Dispositivoid int32      `json:"dispositivoid,omitempty"`
	Id            int32      `json:"id,omitempty"`
	Sede          int32      `json:"sede"`
	Flag1         string     `json:"flag1,omitempty"`
	Flag2         string     `json:"flag2,omitempty"`
	PersonaId     NullInt64  `json:"personaid,omitempty"`
	TokendataId   NullString `json:"tokendataid,omitempty"`
	BizPersonaId  NullInt64  `json:"bizpersonaid,omitempty"`
	PaymentId     NullString `json:"paymentid,omitempty"`
	StatusId      NullString `json:"statusid,omitempty"`
	StatusDetail  NullString `json:"statusdetail,omitempty"`
	Ruf1          NullString `json:"ruf1,omitempty"`
	Ruf2          NullString `json:"ruf2,omitempty"`
	Ruf3          NullString `json:"ruf3,omitempty"`
	Iv            NullString `json:"iv,omitempty"`
	Salt          NullString `json:"salt,omitempty"`
	Checksum      NullString `json:"checksum,omitempty"`
	FCreated      NullTime   `json:"fcreated,omitempty"`
	FUpdated      NullTime   `json:"fupdated,omitempty"`
	Estadoreg     NullInt64  `json:"estadoreg,omitempty"`
	Activo        NullInt64  `json:"activo,omitempty"`
	TotalRecords  int64      `json:"total_records"`
}

func (e BizPaymentE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}
func (e BizPaymentStatusE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectBizPayment = `select * from biz_payment_list( $1, $2)`
const querySelectBizPaymentStatus = `select * from biz_payment_status_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *BizPaymentE) GetAll(token string, filter string) ([]*BizPaymentE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizPayment

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

	var lista []*BizPaymentE

	for rows.Next() {
		var rowdata BizPaymentE
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
			&rowdata.BizPersonaId,
			&rowdata.Nroperacion,
			&rowdata.Maskoperacion,
			&rowdata.Numero,
			&rowdata.NumeroText,
			&rowdata.Fecha,
			&rowdata.PaymentTypeId,
			&rowdata.PaymentMethodTypeId,
			&rowdata.PaymentText,
			&rowdata.BizPersonaText,
			&rowdata.BizPersonaRuc,
			&rowdata.StatusId,
			&rowdata.Secuencial,
			&rowdata.Amount,
			&rowdata.DivisaId,
			&rowdata.DivisaText,
			&rowdata.DivisaSimbolo,
			&rowdata.DivisaDecimal,
			&rowdata.TasaVenta,
			&rowdata.TasaCompra,
			&rowdata.PaymentRefNum,
			&rowdata.BancoId,
			&rowdata.BancoText,
			&rowdata.MarcaTarjeta,
			&rowdata.UltimosDigitos,
			&rowdata.Fdiferida,
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
			log.Println("Error scanning", err)
			return nil, err
		}

		lista = append(lista, &rowdata)
	}

	return lista, nil
}

// GetOne returns one user by id
func (u *BizPaymentE) GetByUniqueid(token string, uniqueid int) (*BizPaymentE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizPayment

	var rowdata BizPaymentE
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
		&rowdata.BizPersonaId,
		&rowdata.Nroperacion,
		&rowdata.Maskoperacion,
		&rowdata.Numero,
		&rowdata.NumeroText,
		&rowdata.Fecha,
		&rowdata.PaymentTypeId,
		&rowdata.PaymentMethodTypeId,
		&rowdata.PaymentText,
		&rowdata.BizPersonaText,
		&rowdata.BizPersonaRuc,
		&rowdata.StatusId,
		&rowdata.Secuencial,
		&rowdata.Amount,
		&rowdata.DivisaId,
		&rowdata.DivisaText,
		&rowdata.DivisaSimbolo,
		&rowdata.DivisaDecimal,
		&rowdata.TasaVenta,
		&rowdata.TasaCompra,
		&rowdata.PaymentRefNum,
		&rowdata.BancoId,
		&rowdata.BancoText,
		&rowdata.MarcaTarjeta,
		&rowdata.UltimosDigitos,
		&rowdata.Fdiferida,
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

	jsonText = fmt.Sprintf(`{"paymentid":%d}`, uniqueid)

	/// Se adicionan los status del invoice
	statusE := BizPaymentStatusE{}
	status, err := statusE.GetAll(token, jsonText)
	if err == nil {
		rowdata.Status = status
	}

	return &rowdata, nil
}

func (u *BizPaymentStatusE) GetAll(token string, filter string) ([]*BizPaymentStatusE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizPaymentStatus

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

	var lista []*BizPaymentStatusE

	for rows.Next() {
		var rowdata BizPaymentStatusE
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
			&rowdata.BizPersonaId,
			&rowdata.PaymentId,
			&rowdata.StatusId,
			&rowdata.StatusDetail,
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

func (u *BizPaymentStatusE) GetByUniqueid(token string, uniqueid int) (*BizPaymentStatusE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizPaymentStatus

	var rowdata BizPaymentStatusE
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
		&rowdata.BizPersonaId,
		&rowdata.PaymentId,
		&rowdata.StatusId,
		&rowdata.StatusDetail,
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

	/// Se adicionan los items del invoice

	return &rowdata, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *BizPaymentE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT biz_payment_save($1, $2, $3)`
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

func (u *BizPaymentStatusE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT biz_payment_status_save($1, $2, $3)`
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
func (u *BizPaymentE) Delete(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT biz_payment_save($1, $2, $3)`
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

// DeleteByID deletes one user from the database, by ID
func (u *BizPaymentStatusE) Delete(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT biz_payment_status_save($1, $2, $3)`
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

/*****
 * Este procedimiento registra datos en multiples tablas transaccionales.-
 *********/
func (u *BizPaymentE) RegisterPayment(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("RegisterPayment [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL register_payment($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	_, err = stmt.ExecContext(ctx, token, data, metricas, &v_uniqueid)
	if err != nil {
		return nil, err
	}
	///defer result.Close()

	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid

	log.Printf("RegisterPayment [ID]=%v\n", v_uniqueid)

	return retorno, nil
}
