package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Customer Suscritos a Cobros regulares
type BizSubscriptionsE struct {
	Uniqueid            int64       `json:"uniqueid,omitempty"`
	Owner               NullInt32   `json:"owner,omitempty"`
	Dispositivoid       int32       `json:"dispositivoid,omitempty"`
	Id                  int32       `json:"id,omitempty"`
	Sede                int32       `json:"sede"`
	Flag1               string      `json:"flag1,omitempty"`
	Flag2               string      `json:"flag2,omitempty"`
	PersonaId           NullInt64   `json:"personaid,omitempty"`
	TokendataId         NullString  `json:"tokendataid,omitempty"`
	BizPersonaId        NullInt64   `json:"bizpersonaid,omitempty"`
	BizPersonaText      NullString  `json:"bizpersonatext,omitempty"`
	BizPersonaLabel     NullString  `json:"bizpersonalabel,omitempty"`
	ComercioText        NullString  `json:"comerciotext,omitempty"`
	Issued              NullTime    `json:"issued,omitempty"`
	Expired             NullTime    `json:"expired,omitempty"`
	TipoIntervalo       NullInt32   `json:"tipointervalo,omitempty"`
	Frecuencia          NullInt32   `json:"frecuencia,omitempty"`
	DiaCorte            NullInt32   `json:"diacorte,omitempty"`
	Calculate           NullInt32   `json:"calculate,omitempty"`
	PaymentMethodTypeId NullString  `json:"paymentmethodtypeid,omitempty"`
	InvoiceTypeId       NullString  `json:"invoicetypeid,omitempty"`
	ProductId           NullInt64   `json:"productid,omitempty"`
	ProductText         NullString  `json:"producttext,omitempty"`
	CategoryText        NullString  `json:"categitemtext,omitempty"`
	DivisaId            NullInt64   `json:"divisaid,omitempty"`
	DivisaText          NullString  `json:"divisatext,omitempty"`
	DivisaSimbolo       NullString  `json:"divisasimbolo,omitempty"`
	DivisaDecimal       NullInt32   `json:"divisadecimal,omitempty"`
	Amount              NullFloat64 `json:"amount,omitempty"`
	StatusId            NullString  `json:"status_id,omitempty"`
	SmsReminder         NullInt32   `json:"smsreminder,omitempty"`
	EmailReminder       NullInt32   `json:"emailreminder,omitempty"`
	Notes               NullString  `json:"notes,omitempty"`
	Ruf1                NullString  `json:"ruf1,omitempty"`
	Ruf2                NullString  `json:"ruf2,omitempty"`
	Ruf3                NullString  `json:"ruf3,omitempty"`
	Iv                  NullString  `json:"iv,omitempty"`
	Salt                NullString  `json:"salt,omitempty"`
	Checksum            NullString  `json:"checksum,omitempty"`
	FCreated            NullTime    `json:"fcreated,omitempty"`
	FUpdated            NullTime    `json:"fupdated,omitempty"`
	Activo              int32       `json:"activo,omitempty"`
	Estadoreg           int32       `json:"estadoreg,omitempty"`
	TotalRecords        int64       `json:"total_records"`
}

func (e BizSubscriptionsE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

func (e BizSubscriptionsE) IssuedText() string {
	return e.Issued.Time.Format("Dec-2000")
}

const querySelectBizSubscrip = `select * from biz_subscriptions_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *BizSubscriptionsE) GetAll(token string, filter string) ([]*BizSubscriptionsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizSubscrip

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

	var lista []*BizSubscriptionsE

	for rows.Next() {
		var rowdata BizSubscriptionsE
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
			&rowdata.BizPersonaText,
			&rowdata.BizPersonaLabel,
			&rowdata.ComercioText,
			&rowdata.Issued,
			&rowdata.Expired,
			&rowdata.TipoIntervalo,
			&rowdata.Frecuencia,
			&rowdata.DiaCorte,
			&rowdata.Calculate,
			&rowdata.PaymentMethodTypeId,
			&rowdata.InvoiceTypeId,
			&rowdata.ProductId,
			&rowdata.ProductText,
			&rowdata.CategoryText,
			&rowdata.DivisaId,
			&rowdata.DivisaText,
			&rowdata.DivisaSimbolo,
			&rowdata.DivisaDecimal,
			&rowdata.Amount,
			&rowdata.StatusId,
			&rowdata.SmsReminder,
			&rowdata.EmailReminder,
			&rowdata.Notes,
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
func (u *BizSubscriptionsE) GetByUniqueid(token string, uniqueid int) (*BizSubscriptionsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizSubscrip

	var rowdata BizSubscriptionsE
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
		&rowdata.BizPersonaText,
		&rowdata.BizPersonaLabel,
		&rowdata.ComercioText,
		&rowdata.Issued,
		&rowdata.Expired,
		&rowdata.TipoIntervalo,
		&rowdata.Frecuencia,
		&rowdata.DiaCorte,
		&rowdata.Calculate,
		&rowdata.PaymentMethodTypeId,
		&rowdata.InvoiceTypeId,
		&rowdata.ProductId,
		&rowdata.ProductText,
		&rowdata.CategoryText,
		&rowdata.DivisaId,
		&rowdata.DivisaText,
		&rowdata.DivisaSimbolo,
		&rowdata.DivisaDecimal,
		&rowdata.Amount,
		&rowdata.StatusId,
		&rowdata.SmsReminder,
		&rowdata.EmailReminder,
		&rowdata.Notes,
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
func (u *BizSubscriptionsE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT biz_subscriptions_save($1, $2, $3)`
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
func (u *BizSubscriptionsE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT biz_subscriptions_save($1, $2, $3)`
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
func (u *BizSubscriptionsE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT biz_subscriptions_save($1, $2, $3)`
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
func (u *BizSubscriptionsE) GenerateBills(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("GenerateBills [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL subscriptions_bills($1, $2, $3, $4)`
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
	retorno["total"] = v_uniqueid

	log.Printf("GenerateBills [Total]=%v\n", v_uniqueid)

	return retorno, nil
}
