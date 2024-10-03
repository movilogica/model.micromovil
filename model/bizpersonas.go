package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Customers
type BizPersonasE struct {
	Uniqueid        int64       `json:"uniqueid,omitempty"`
	Owner           NullInt32   `json:"owner,omitempty"`
	Dispositivoid   NullInt32   `json:"dispositivoid,omitempty"`
	Id              int32       `json:"id,omitempty"`
	Sede            int32       `json:"sede"`
	Flag1           string      `json:"flag1,omitempty"`
	Flag2           string      `json:"flag2,omitempty"`
	PersonaId       NullInt64   `json:"personaid,omitempty"`
	Nroperacion     NullString  `json:"nroperacion,omitempty"`
	Maskoperacion   NullString  `json:"maskoperacion,omitempty"`
	Nickname        NullString  `json:"nickname,omitempty"`
	Nombres         NullString  `json:"nombres,omitempty"`
	Midlename       NullString  `json:"midlename,omitempty"`
	Apaterno        NullString  `json:"apaterno,omitempty"`
	Amaterno        NullString  `json:"amaterno,omitempty"`
	Movil           NullString  `json:"movil,omitempty"`
	Email           NullString  `json:"email,omitempty"`
	Phone           NullString  `json:"phone,omitempty"`
	Avatar          NullString  `json:"avatar,omitempty"`
	TipodocId       NullInt64   `json:"tipodocid,omitempty"`
	TipodocCode     NullString  `json:"tipodoccode,omitempty"`
	TipodocText     NullString  `json:"tipodoctext,omitempty"`
	Numerodoc       NullString  `json:"numerodoc,omitempty"`
	Documento       NullString  `json:"documento,omitempty"`
	RoleTypeId      NullString  `json:"roletypeid,omitempty"`
	Clasificacion   NullInt32   `json:"clasificacion,omitempty"`
	Privado         NullInt32   `json:"privado,omitempty"`
	HourValue       NullFloat64 `json:"hourvalue,omitempty"`
	Mailing         NullInt32   `json:"mailing,omitempty"`
	Notifier        NullInt32   `json:"notifier,omitempty"`
	FMailing        NullTime    `json:"fmailing,omitempty" time_format:"sql_datetime"`
	FNotifier       NullTime    `json:"fnotifier,omitempty" time_format:"sql_datetime"`
	FMembership     NullTime    `json:"fmembership,omitempty" time_format:"sql_datetime"`
	FResignation    NullTime    `json:"fresignation,omitempty" time_format:"sql_datetime"`
	FLastAccess     NullTime    `json:"flastaccess,omitempty" time_format:"sql_datetime"`
	FLastMovement   NullTime    `json:"flastmovement,omitempty" time_format:"sql_datetime"`
	FLastTransact   NullTime    `json:"flasttransact,omitempty" time_format:"sql_datetime"`
	FLastOrder      NullTime    `json:"flastorder,omitempty" time_format:"sql_datetime"`
	FSubscription   NullTime    `json:"fsubscription,omitempty" time_format:"sql_datetime"`
	FUnsubscription NullTime    `json:"funsubscription,omitempty" time_format:"sql_datetime"`
	Suscribed       NullInt32   `json:"suscribed,omitempty"`
	Balance         NullFloat64 `json:"balance,omitempty"`
	StatusPersona   NullInt32   `json:"statuspersona,omitempty"`
	StatusDetail    NullString  `json:"statusdetail,omitempty"`
	StatusDateAt    NullTime    `json:"statusdate,omitempty"`
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
	Addrs           []*BizPersonasAddressE
	Bitacora        []*BizPersonasBitacoraE
	Medios          []*BizPersonasMedioE
	Vehicles        []*BizPersonasVehiclesE
}

func (e BizPersonasE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

func (e BizPersonasE) Fullname() string {
	return fmt.Sprintf("%s %s %s", e.Nombres.String, e.Midlename.String, e.Apaterno.String)
}

type CustomerInfoE struct {
	Info          *BizPersonasE           `json:"info"`
	Medios        []*BizPersonasMedioE    `json:"medios"`
	Vehicles      []*BizPersonasVehiclesE `json:"vehicles"`
	Notes         []*BizNotesTextE        `json:"notes"`
	Subscriptions []*BizSubscriptionsE    `json:"subscriptions"`
	Payments      []*BizPaymentE          `json:"payments"`
	Invoices      []*BizInvoiceE          `json:"invoices"`
}

func (e CustomerInfoE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectBizPer = `select * from biz_personas_list( $1, $2)`
const querySelectBizPerMinimal = `select uniqueid, sede, flag1, flag2, personaid, nroperacion, maskoperacion, nombres, apaterno, movil, email, phone, role_type_id, privado, suscribed, balance, activo, estadoreg, total_records from biz_personas_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *BizPersonasE) GetAll(token string, filter string) ([]*BizPersonasE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizPer

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

	var lista []*BizPersonasE

	for rows.Next() {
		var rowdata BizPersonasE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.PersonaId,
			&rowdata.Nroperacion,
			&rowdata.Maskoperacion,
			&rowdata.Nickname,
			&rowdata.Nombres,
			&rowdata.Midlename,
			&rowdata.Apaterno,
			&rowdata.Amaterno,
			&rowdata.Movil,
			&rowdata.Email,
			&rowdata.Phone,
			&rowdata.Avatar,
			&rowdata.TipodocId,
			&rowdata.TipodocCode,
			&rowdata.TipodocText,
			&rowdata.Numerodoc,
			&rowdata.Documento,
			&rowdata.RoleTypeId,
			&rowdata.Clasificacion,
			&rowdata.Privado,
			&rowdata.HourValue,
			&rowdata.Mailing,
			&rowdata.Notifier,
			&rowdata.FMailing,
			&rowdata.FNotifier,
			&rowdata.FMembership,
			&rowdata.FResignation,
			&rowdata.FLastAccess,
			&rowdata.FLastMovement,
			&rowdata.FLastTransact,
			&rowdata.FLastOrder,
			&rowdata.FSubscription,
			&rowdata.FUnsubscription,
			&rowdata.Suscribed,
			&rowdata.Balance,
			&rowdata.StatusPersona,
			&rowdata.StatusDetail,
			&rowdata.StatusDateAt,
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

// GetAll: uniqueid, sede, flag1, flag2, personaid, tokendataid, nroperacion, maskoperacion, nickname, role_type_id, suscribed, balance, activo, estadoreg
func (u *BizPersonasE) LookingFor(token string, filter string) ([]*BizPersonasE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizPerMinimal

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

	var lista []*BizPersonasE

	for rows.Next() {
		var rowdata BizPersonasE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.PersonaId,
			&rowdata.Nroperacion,
			&rowdata.Maskoperacion,
			&rowdata.Nombres,
			&rowdata.Apaterno,
			&rowdata.Movil,
			&rowdata.Email,
			&rowdata.Phone,
			&rowdata.RoleTypeId,
			&rowdata.Privado,
			&rowdata.Suscribed,
			&rowdata.Balance,
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
func (u *BizPersonasE) GetByUniqueid(token string, uniqueid int) (*BizPersonasE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectBizPer

	var rowdata BizPersonasE
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
		&rowdata.Nroperacion,
		&rowdata.Maskoperacion,
		&rowdata.Nickname,
		&rowdata.Nombres,
		&rowdata.Midlename,
		&rowdata.Apaterno,
		&rowdata.Amaterno,
		&rowdata.Movil,
		&rowdata.Email,
		&rowdata.Phone,
		&rowdata.Avatar,
		&rowdata.TipodocId,
		&rowdata.TipodocCode,
		&rowdata.TipodocText,
		&rowdata.Numerodoc,
		&rowdata.Documento,
		&rowdata.RoleTypeId,
		&rowdata.Clasificacion,
		&rowdata.Privado,
		&rowdata.HourValue,
		&rowdata.Mailing,
		&rowdata.Notifier,
		&rowdata.FMailing,
		&rowdata.FNotifier,
		&rowdata.FMembership,
		&rowdata.FResignation,
		&rowdata.FLastAccess,
		&rowdata.FLastMovement,
		&rowdata.FLastTransact,
		&rowdata.FLastOrder,
		&rowdata.FSubscription,
		&rowdata.FUnsubscription,
		&rowdata.Suscribed,
		&rowdata.Balance,
		&rowdata.StatusPersona,
		&rowdata.StatusDetail,
		&rowdata.StatusDateAt,
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

// GetOne returns one user by id
func (u *BizPersonasE) GetCustomerOkById(token string, customerid int) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM biz_personas_ok ($1, $2)`

	row := db.QueryRowContext(ctx, query, token, customerid)

	var p_success int32
	err := row.Scan(
		&p_success,
	)

	if err != nil {
		return nil, err
	}

	retorno := make(map[string]any)
	retorno["success"] = p_success

	return retorno, nil
}

func (u *BizPersonasE) GetCustomerStatusById(token string, customerid int) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM biz_personas_status ($1, $2)`

	row := db.QueryRowContext(ctx, query, token, customerid)

	var p_status NullInt32
	var p_activo NullInt32
	var p_estadoreg NullInt32

	err := row.Scan(
		&p_status,
		&p_activo,
		&p_estadoreg,
	)

	if err != nil {
		return nil, err
	}

	retorno := make(map[string]any)
	retorno["status"] = p_status.Int32
	retorno["activo"] = p_activo.Int32
	retorno["estadoreg"] = p_estadoreg.Int32

	return retorno, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *BizPersonasE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT biz_personas_save($1, $2, $3)`

	log.Printf("%s [Data = %s]", query, string(jsonData))

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
func (u *BizPersonasE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT biz_personas_save($1, $2, $3)`
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
func (u *BizPersonasE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT biz_personas_save($1, $2, $3)`
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
func (u *BizPersonasE) RegisterCustomer(token string, data string, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	log.Printf("RegisterCustomer [data]=%s [metricas]=%s\n", data, metricas)

	query := `CALL register_customer($1, $2, $3, $4)`
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

	log.Printf("RegisterCustomer [ID]=%v\n", v_uniqueid)

	return retorno, nil
}
