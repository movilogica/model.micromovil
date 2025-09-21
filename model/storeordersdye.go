package model

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

// Ordenes de Dye House
type StoreOrdersDyeE struct {
	Uniqueid            int64                     `json:"uniqueid,omitempty"`
	Owner               NullInt32                 `json:"owner,omitempty"`
	Dispositivoid       NullInt32                 `json:"dispositivoid,omitempty"`
	Id                  int32                     `json:"id,omitempty"`
	Sede                int32                     `json:"sede"`
	Flag1               string                    `json:"flag1,omitempty"`
	Flag2               string                    `json:"flag2,omitempty"`
	PersonaId           NullInt64                 `json:"personaid,omitempty"`
	TokendataId         NullString                `json:"tokendataid,omitempty"`
	BizPersonaId        NullInt64                 `json:"bizpersonaid,omitempty"`
	Nroperacion         NullString                `json:"nroperacion,omitempty"`
	NroperacionMask     NullString                `json:"nroperacionmask,omitempty"`
	DyeOrder            NullInt32                 `json:"dyeorder,omitempty"`
	CutTextil           NullString                `json:"cuttextil,omitempty"`
	OrdenEntrada        NullString                `json:"ordenentrada,omitempty"`
	DivisionCode        NullString                `json:"divisioncode,omitempty"`
	JobNumber           NullString                `json:"jobnumber,omitempty"`
	PersonaText         NullString                `json:"personatext,omitempty"`
	Email               NullString                `json:"email,omitempty"`
	Movil               NullString                `json:"movil,omitempty"`
	Priority            NullInt32                 `json:"priority,omitempty"`
	Fecha               NullTime                  `json:"fecha,omitempty"`
	FDyeHouse           NullTime                  `json:"fdyehouse,omitempty"`
	QDiasEstimated      NullInt32                 `json:"qdiasestimated,omitempty"`
	FEstimated          NullTime                  `json:"festimated,omitempty"`
	FDelivered          NullTime                  `json:"fdelivered,omitempty"`
	FProcessed          NullTime                  `json:"fprocessed,omitempty"`
	FCancelled          NullTime                  `json:"fcancelled,omitempty"`
	StatusId            NullString                `json:"statusid,omitempty"`
	DivisaId            NullInt64                 `json:"divisaid,omitempty"`
	DivisaText          NullString                `json:"divisatext,omitempty"`
	DivisaSimbolo       NullString                `json:"divisasimbolo,omitempty"`
	DivisaDecimal       NullInt32                 `json:"divisadecimal,omitempty"`
	TasaVenta           NullFloat64               `json:"tasaventa,omitempty"`
	TasaCompra          NullFloat64               `json:"tasacompra,omitempty"`
	ColorCode           NullString                `json:"colorcode,omitempty"`
	Udisplay            NullString                `json:"udisplay,omitempty"`
	Uom                 NullString                `json:"uom,omitempty"`
	ItemsText           NullString                `json:"itemstext,omitempty"`
	TotaLineas          NullInt32                 `json:"totalineas,omitempty"`
	TotalItems          NullFloat64               `json:"totalitems,omitempty"`
	TotalItemsReceived  NullFloat64               `json:"totalitemsreceived,omitempty"`
	TotalQtotal         NullFloat64               `json:"totalqtotal,omitempty"`
	TotalQtotalReceived NullFloat64               `json:"totalqtotalreceived,omitempty"`
	TotalGrossWeight    NullFloat64               `json:"totalgrossweight,omitempty"`
	TotalNetWeight      NullFloat64               `json:"totalnetweight,omitempty"`
	TotalOrden          NullFloat64               `json:"totalorden,omitempty"`
	MermaPorc           NullFloat64               `json:"mermaporc,omitempty"`
	MermaWeight         NullFloat64               `json:"mermaweight,omitempty"`
	MermaExceden        NullFloat64               `json:"mermaexceden,omitempty"`
	PrecioPromedio      NullFloat64               `json:"precioprom,omitempty"`
	Comentarios         NullString                `json:"comentarios,omitempty"`
	NotesToDyehouse     NullString                `json:"notestodyehouse,omitempty"`
	NotesFromFactory    NullString                `json:"notesfromfactory,omitempty"`
	InternalNotes       NullString                `json:"internalnotes,omitempty"`
	NotifyVendor        NullInt32                 `json:"notifyvendor,omitempty"`
	Fnotifyvendor       NullTime                  `json:"fnotifyvendor,omitempty"`
	Latitud             NullFloat64               `json:"latitud,omitempty"`
	Longitud            NullFloat64               `json:"longitud,omitempty"`
	Checkin             NullInt32                 `json:"checkin,omitempty"`
	QPendientes         NullInt32                 `json:"qpendientes,omitempty"`
	Ruf1                NullString                `json:"ruf1,omitempty"`
	Ruf2                NullString                `json:"ruf2,omitempty"`
	Ruf3                NullString                `json:"ruf3,omitempty"`
	Iv                  NullString                `json:"iv,omitempty"`
	Salt                NullString                `json:"salt,omitempty"`
	Checksum            NullString                `json:"checksum,omitempty"`
	FCreated            NullTime                  `json:"fcreated,omitempty"`
	FUpdated            NullTime                  `json:"fupdated,omitempty"`
	UCreated            NullString                `json:"ucreated,omitempty"`
	UUpdated            NullString                `json:"uupdated,omitempty"`
	Activo              int32                     `json:"activo,omitempty"`
	Estadoreg           int32                     `json:"estadoreg,omitempty"`
	TotalRecords        int64                     `json:"total_records,omitempty"`
	Items               []*StoreOrdersDyeItemsE   `json:"items:omitempty"`
	Statuses            []*StoreOrdersDyeStatusE  `json:"statuses:omitempty"`
	Actions             []*StoreOrdersDyeActionsE `json:"actions:omitempty"`
}

func (e StoreOrdersDyeE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const queryListStoreOrdersDyeE = `select uniqueid, sede, flag1, flag2, bizpersonaid, personatext, dyeorder, cuttextil, ordenentrada, fecha, fdyehouse, festimated, fprocessed, fcancelled, statusid, nroperacionmask, divisasimbolo, colorcode, udisplay, uom, itemstext, totalqtotal, totalitems, totalgrossweight, totalnetweight, totalorden, mermaporc, mermaweight, mermaexceden, qpendientes, activo, estadoreg, total_records from store_orders_dye_list( $1, $2)`
const queryLoadStoreOrdersDyeE = `select * from store_orders_dye_list( $1, $2)`
const querySaveStoreOrdersDyeE = `SELECT store_orders_dye_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreOrdersDyeE) GetAll(token string, filter string) ([]*StoreOrdersDyeE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreOrdersDyeE

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

	var lista []*StoreOrdersDyeE

	/// uniqueid, sede, flag1, flag2, bizpersonaid, personatext, dyeorder, cuttextil, ordenentrada, fecha, fdyehouse, festimated, fprocessed,
	//  fcancelled, statusid, nroperacionmask, divisasimbolo, itemstext, totalqtotal, totalitems, totalgrossweight, totalnetweight, totalorden,
	//  mermaporc, mermaweight, mermaexceden, qpendientes, activo, estadoreg, total_records
	for rows.Next() {
		var rowdata StoreOrdersDyeE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.BizPersonaId,
			&rowdata.PersonaText,
			&rowdata.DyeOrder,
			&rowdata.CutTextil,
			&rowdata.OrdenEntrada,
			&rowdata.Fecha,
			&rowdata.FDyeHouse,
			&rowdata.FEstimated,
			&rowdata.FProcessed,
			&rowdata.FCancelled,
			&rowdata.StatusId,
			&rowdata.NroperacionMask,
			&rowdata.DivisaSimbolo,
			&rowdata.ColorCode,
			&rowdata.Udisplay,
			&rowdata.Uom,
			&rowdata.ItemsText,
			&rowdata.TotalQtotal,
			&rowdata.TotalItems,
			&rowdata.TotalGrossWeight,
			&rowdata.TotalNetWeight,
			&rowdata.TotalOrden,
			&rowdata.MermaPorc,
			&rowdata.MermaWeight,
			&rowdata.MermaExceden,
			&rowdata.QPendientes,
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
func (u *StoreOrdersDyeE) GetByUniqueid(token string, jsonText string) (*StoreOrdersDyeE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	/// Adicionar '@fulldata=1'
	var mapData map[string]any
	err := json.Unmarshal([]byte(jsonText), &mapData)
	if err != nil {
		return nil, err
	}
	mapData["@fulldata"] = 1
	jsonBinary, _ := json.Marshal(mapData)

	query := queryLoadStoreOrdersDyeE

	var rowdata StoreOrdersDyeE
	log.Printf("[%s] Where = %s\n", query, string(jsonBinary))
	row := db.QueryRowContext(ctx, query, token, jsonBinary)

	err = row.Scan(
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
		&rowdata.NroperacionMask,
		&rowdata.DyeOrder,
		&rowdata.CutTextil,
		&rowdata.OrdenEntrada,
		&rowdata.DivisionCode,
		&rowdata.JobNumber,
		&rowdata.PersonaText,
		&rowdata.Email,
		&rowdata.Movil,
		&rowdata.Priority,
		&rowdata.Fecha,
		&rowdata.FDyeHouse,
		&rowdata.QDiasEstimated,
		&rowdata.FEstimated,
		&rowdata.FDelivered,
		&rowdata.FProcessed,
		&rowdata.FCancelled,
		&rowdata.StatusId,
		&rowdata.DivisaId,
		&rowdata.DivisaText,
		&rowdata.DivisaSimbolo,
		&rowdata.DivisaDecimal,
		&rowdata.TasaVenta,
		&rowdata.TasaCompra,
		&rowdata.ColorCode,
		&rowdata.Udisplay,
		&rowdata.Uom,
		&rowdata.ItemsText,
		&rowdata.TotaLineas,
		&rowdata.TotalItems,
		&rowdata.TotalQtotal,
		&rowdata.TotalItemsReceived,
		&rowdata.TotalQtotalReceived,
		&rowdata.TotalGrossWeight,
		&rowdata.TotalNetWeight,
		&rowdata.TotalOrden,
		&rowdata.MermaPorc,
		&rowdata.MermaWeight,
		&rowdata.MermaExceden,
		&rowdata.PrecioPromedio,
		&rowdata.Comentarios,
		&rowdata.NotesToDyehouse,
		&rowdata.NotesFromFactory,
		&rowdata.InternalNotes,
		&rowdata.NotifyVendor,
		&rowdata.Fnotifyvendor,
		&rowdata.Latitud,
		&rowdata.Longitud,
		&rowdata.Checkin,
		&rowdata.QPendientes,
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

// Get next order number
func (u *StoreOrdersDyeE) GetNextOrderNumber(token string, jsonText string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "select (COALESCE(MAX(dyeorder),1) + 1)::varchar from store_orders_dye"
	var mapData map[string]string
	err := json.Unmarshal([]byte(jsonText), &mapData)
	if err != nil {
		fmt.Println(err)
	}
	sede, _ := strconv.Atoi(mapData["sede"])
	query += fmt.Sprintf(" WHERE sede = %d", sede)
	tokendataid := mapData["tokendataid"]
	query += fmt.Sprintf(" AND tokendataid = '%s'", tokendataid)

	log.Printf("Query = %s\n", query)
	row := db.QueryRowContext(ctx, query)
	var resultado string
	err = row.Scan(&resultado)
	if err != nil {
		return "", err
	}

	return resultado, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *StoreOrdersDyeE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreOrdersDyeE
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
func (u *StoreOrdersDyeE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreOrdersDyeE
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
func (u *StoreOrdersDyeE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreOrdersDyeE
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
 * Estos procedimientos registran datos en multiples tablas; transaccionalmente.-
 *********/
func (u *StoreOrdersDyeE) RegisterOrder(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("RegisterDyeOrder [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL register_dye_orders($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	///result, err := stmt.ExecContext(ctx, token, data, metricas, &v_uniqueid)
	err = stmt.QueryRowContext(ctx, token, data, metricas, &v_uniqueid).Scan(&v_uniqueid)
	if err != nil {
		return nil, err
	}
	///defer result.Close()

	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid

	log.Printf("RegisterDyeOrder [ID]=%v\n", v_uniqueid)

	return retorno, nil
}

func (u *StoreOrdersDyeE) UpdateOrder(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("UpdateDyeOrder [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL update_dye_orders($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	///result, err := stmt.ExecContext(ctx, token, data, metricas, &v_uniqueid)
	err = stmt.QueryRowContext(ctx, token, data, metricas, &v_uniqueid).Scan(&v_uniqueid)
	if err != nil {
		return nil, err
	}
	///defer result.Close()

	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid

	log.Printf("UpdateDyeOrder [ID]=%v\n", v_uniqueid)

	return retorno, nil
}

func (u *StoreOrdersDyeE) ChangeStatusOrder(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("ChangeStatusDyeOrder [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL changestatus_dye_orders($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	err = stmt.QueryRowContext(ctx, token, data, metricas, &v_uniqueid).Scan(&v_uniqueid)
	if err != nil {
		return nil, err
	}

	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid

	log.Printf("ChangeStatusDyeOrder [ID]=%v\n", v_uniqueid)

	return retorno, nil
}

func (u *StoreOrdersDyeE) ReceiverOrder(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("ReceiverDyeOrder [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL received_dye_orders($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	///_, err = stmt.ExecContext(ctx, token, data, metricas, &v_uniqueid)
	err = stmt.QueryRowContext(ctx, token, data, metricas, &v_uniqueid).Scan(&v_uniqueid)
	if err != nil {
		return nil, err
	}
	///defer result.Close()

	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid

	log.Printf("RegisterDyeOrder [ID]=%v\n", v_uniqueid)

	return retorno, nil
}

func (u *StoreOrdersDyeE) RegisterAndReceiverOrder(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("RegisterAndReceiverDyeOrder [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL register_received_dye_orders($1, $2, $3, $4, $5, $6)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	var v_uniqueid float64
	var v_received int32
	var v_message string
	err = stmt.QueryRowContext(ctx, token, data, metricas, &v_uniqueid, &v_received, &v_message).Scan(&v_uniqueid, &v_received, &v_message)
	if err != nil {
		return nil, err
	}
	retorno := make(map[string]any)
	retorno["uniqueid"] = v_uniqueid
	retorno["received"] = v_received
	retorno["message"] = v_message

	log.Printf("RegisterAndReceiverDyeOrder [ID]=%f [received]=%v [message]=%s\n", v_uniqueid, v_received, v_message)

	return retorno, nil
}

func (u *StoreOrdersDyeE) PayloadData(plantilla string) string {
	log.Println("PayloadData " + plantilla)
	mapa, _ := structToMap(u)
	for key, value := range mapa {
		///log.Printf("PayloadData %s=%v\n", key, value)
		plantilla = strings.Replace(plantilla, fmt.Sprintf("{{%s}}", key), fmt.Sprintf("%v", value), -1)
	}
	return plantilla
}

func (u *StoreOrdersDyeE) BuildPdf(writer io.Writer) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	// First page: manual local link
	pdf.SetFont("Helvetica", "", 12)
	_, lineHt := pdf.GetFontSize()
	link := pdf.AddLink()
	pdf.WriteLinkID(lineHt, "here", link)
	pdf.SetFont("", "", 0)
	// Second page: image link and basic HTML with link
	pdf.AddPage()
	pdf.SetLink(link, 0, -1)
	urlLogo := "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQcj-igRNigzyZ751AM6mSaC8omjnJWyJegdQ&s"
	httpimg.Register(pdf, urlLogo, "")
	///name, x, y, w, h
	pdf.Image(urlLogo, 5, 15, 50, 0, false, "", 0, "http://www.fpdf.org")
	pdf.SetLeftMargin(45)
	pdf.SetFontSize(14)
	_, lineHt = pdf.GetFontSize()
	ht := pdf.PointConvert(lineHt)
	pdf.Ln(2 * ht)
	for j, item := range u.Items {
		pdf.Write(lineHt, fmt.Sprintf("%d", j))
		pdf.SetX(40)
		pdf.Write(lineHt, item.ProductCode.String)
		pdf.SetX(70)
		pdf.Write(lineHt, item.ProductText.String)
		if item.Lotid.Valid {
			pdf.SetX(70)
			pdf.Write(lineHt, fmt.Sprintf("Lote %s", item.Lotid.String))
		}
		if item.Roadmap.Valid {
			pdf.SetX(70)
			pdf.Write(lineHt, fmt.Sprintf("SKU %s", item.Roadmap.String))
		}
		pdf.SetX(70)
		pdf.Write(lineHt, item.UMedidaText())
		pdf.SetX(70)
		pdf.Write(lineHt, fmt.Sprintf("%.2f", item.Quantity.Float64))
		pdf.SetX(70)
		pdf.Ln(1)
	}
	err := pdf.Output(writer)
	return err
}
