package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Customer Suscritos a Cobros regulares
type BizSubscriptionStatusE struct {
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
	CarType             NullString  `json:"cartype,omitempty"`
	CarModel            NullString  `json:"carmodel,omitempty"`
	CarColor            NullString  `json:"carcolor,omitempty"`
	License             NullString  `json:"license,omitempty"`
	Movil               NullString  `json:"movil,omitempty"`
	LastPayment         NullTime    `json:"lastpayment,omitempty"`
	LastAmount          NullFloat64 `json:"lastamount,omitempty"`
	Balance             NullFloat64 `json:"balance,omitempty"`
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

func (e BizSubscriptionStatusE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectBizSubscripStatus = `select * from biz_subscriptions_status_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *BizSubscriptionStatusE) GetAll(token string, filter string) ([]*BizSubscriptionStatusE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizSubscripStatus

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

	var lista []*BizSubscriptionStatusE

	for rows.Next() {
		var rowdata BizSubscriptionStatusE
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
			&rowdata.CarType,
			&rowdata.CarModel,
			&rowdata.CarColor,
			&rowdata.License,
			&rowdata.Movil,
			&rowdata.LastPayment,
			&rowdata.LastAmount,
			&rowdata.Balance,
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
func (u *BizSubscriptionStatusE) GetByUniqueid(token string, uniqueid int) (*BizSubscriptionStatusE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizSubscripStatus

	var rowdata BizSubscriptionStatusE
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
		&rowdata.CarType,
		&rowdata.CarModel,
		&rowdata.CarColor,
		&rowdata.License,
		&rowdata.Movil,
		&rowdata.LastPayment,
		&rowdata.LastAmount,
		&rowdata.Balance,
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
