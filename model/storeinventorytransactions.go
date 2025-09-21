package model

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/httpimg"
)

// Lista de transacciones (TRANSFERENCIAS/AJUSTES)
type StoreInventoryTransactionsE struct {
	Uniqueid         int64                `json:"uniqueid,omitempty"`
	Owner            NullInt32            `json:"owner,omitempty"`
	Dispositivoid    NullInt32            `json:"dispositivoid,omitempty"`
	Id               int32                `json:"id,omitempty"`
	Sede             int32                `json:"sede"`
	Flag1            string               `json:"flag1,omitempty"`
	Flag2            string               `json:"flag2,omitempty"`
	PersonaId        NullInt64            `json:"personaid,omitempty"`
	TokendataId      NullString           `json:"tokendataid,omitempty"`
	Nroperacion      NullString           `json:"nroperacion,omitempty"`
	NroperacionMask  NullString           `json:"nroperacionmask,omitempty"`
	Numero           NullInt64            `json:"numero,omitempty"`
	DocumentText     NullString           `json:"documenttext,omitempty"`
	Fecha            NullTime             `json:"fecha,omitempty"`
	TipoMov          NullString           `json:"tipomov,omitempty"`
	TipoTransactId   NullInt64            `json:"tipotransactid,omitempty"`
	TipoTransactText NullString           `json:"tipotransacttext,omitempty"`
	ReasonEnumId     NullString           `json:"reasonenumid,omitempty"`
	TotalText        NullString           `json:"totaltext,omitempty"`
	Motivo           NullString           `json:"motivo,omitempty"`
	Notas            NullString           `json:"notas,omitempty"`
	Ruf1             NullString           `json:"ruf1,omitempty"`
	Ruf2             NullString           `json:"ruf2,omitempty"`
	Ruf3             NullString           `json:"ruf3,omitempty"`
	Iv               NullString           `json:"iv,omitempty"`
	Salt             NullString           `json:"salt,omitempty"`
	Checksum         NullString           `json:"checksum,omitempty"`
	FCreated         NullTime             `json:"fcreated,omitempty"`
	FUpdated         NullTime             `json:"fupdated,omitempty"`
	UCreated         NullString           `json:"ucreated,omitempty"`
	UUpdated         NullString           `json:"uupdated,omitempty"`
	Activo           int32                `json:"activo,omitempty"`
	Estadoreg        int32                `json:"estadoreg,omitempty"`
	TotalRecords     int64                `json:"total_records,omitempty"`
	Items            []*StoreOrdersItemsE `json:"items:omitempty"`
}

func (e StoreInventoryTransactionsE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const queryListStoreInventTransE = `select uniqueid, sede, flag1, flag2, fecha, tipomov, nroperacionmask, numero, totaltext, activo, estadoreg, total_records from store_inventory_transact_list( $1, $2)`
const queryLoadStoreInventTransE = `select * from store_inventory_transact_list( $1, $2)`
const querySaveStoreInventTransE = `SELECT store_inventory_transact_save($1, $2, $3)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *StoreInventoryTransactionsE) GetAll(token string, filter string) ([]*StoreInventoryTransactionsE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := queryListStoreInventTransE

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

	var lista []*StoreInventoryTransactionsE

	/// uniqueid, sede, flag1, flag2, fecha, tipomov, nroperacionmask, numero, totaltext, activo, estadoreg, total_records
	for rows.Next() {
		var rowdata StoreInventoryTransactionsE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Fecha,
			&rowdata.TipoMov,
			&rowdata.NroperacionMask,
			&rowdata.Numero,
			&rowdata.TotalText,
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
func (u *StoreInventoryTransactionsE) GetByUniqueid(token string, jsonText string) (*StoreInventoryTransactionsE, error) {
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

	query := queryLoadStoreInventTransE

	var rowdata StoreInventoryTransactionsE
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
		&rowdata.Nroperacion,
		&rowdata.NroperacionMask,
		&rowdata.Numero,
		&rowdata.DocumentText,
		&rowdata.Fecha,
		&rowdata.TipoMov,
		&rowdata.TipoTransactId,
		&rowdata.TipoTransactText,
		&rowdata.ReasonEnumId,
		&rowdata.TotalText,
		&rowdata.Motivo,
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
func (u *StoreInventoryTransactionsE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreInventTransE
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
func (u *StoreInventoryTransactionsE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := querySaveStoreInventTransE
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
func (u *StoreInventoryTransactionsE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := querySaveStoreInventTransE
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
func (u *StoreInventoryTransactionsE) Transferencia(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("Transferencia [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL transfer_stock($1, $2, $3, $4)`
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

	log.Printf("Transferencia [ID]=%v\n", v_uniqueid)

	return retorno, nil
}

func (u *StoreInventoryTransactionsE) AjustarStock(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("AjustarStock [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL ajuste_stock($1, $2, $3, $4)`
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

	log.Printf("AjustarStock [ID]=%v\n", v_uniqueid)

	return retorno, nil
}

func (u *StoreInventoryTransactionsE) BuildPdf(writer io.Writer) error {
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
		if item.SkuNumber.Valid {
			pdf.SetX(70)
			pdf.Write(lineHt, fmt.Sprintf("SKU %s", item.SkuNumber.String))
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
