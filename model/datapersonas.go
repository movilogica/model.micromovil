package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Personas
type DataPersonasE struct {
	Uniqueid       int64       `json:"uniqueid,omitempty"`
	Owner          NullInt32   `json:"owner,omitempty"`
	Dispositivoid  NullInt32   `json:"dispositivoid,omitempty"`
	Id             int32       `json:"id,omitempty"`
	Sede           int32       `json:"sede"`
	Flag1          string      `json:"flag1,omitempty"`
	Flag2          string      `json:"flag2,omitempty"`
	Nroperacion    NullString  `json:"nroperacion,omitempty"`
	CredencialId   NullInt64   `json:"credencialid,omitempty"`
	Nickname       NullString  `json:"nickname,omitempty"`
	Nombres        NullString  `json:"nombres,omitempty"`
	Midlename      NullString  `json:"midlename,omitempty"`
	Apaterno       NullString  `json:"apaterno,omitempty"`
	Amaterno       NullString  `json:"amaterno,omitempty"`
	Movil          NullString  `json:"movil,omitempty"`
	Email          NullString  `json:"email,omitempty"`
	Avatar         NullString  `json:"avatar,omitempty"`
	MerchantId     NullString  `json:"merchantid,omitempty"`
	CountryIso2    NullString  `json:"country_iso_2,omitempty"`
	CountryIso3    NullString  `json:"country_iso_3,omitempty"`
	CountryIsoM49  NullString  `json:"country_iso_m49,omitempty"`
	CountryPrefijo NullString  `json:"country_prefijo,omitempty"`
	PaisId         NullInt64   `json:"paisid,omitempty"`
	PaisText       NullString  `json:"paistext,omitempty"`
	RoleTypeId     NullString  `json:"role_type_id,omitempty"`
	Clasificacion  NullInt32   `json:"clasificacion,omitempty"`
	Robot          NullString  `json:"robot,omitempty"`
	RobotAt        NullTime    `json:"frobot,omitempty"`
	Mailing        NullInt32   `json:"mailing,omitempty"`
	Notifier       NullInt32   `json:"notifier,omitempty"`
	MailingAt      NullTime    `json:"fmailing,omitempty"`
	NotifierAt     NullTime    `json:"fnotifier,omitempty"`
	AfiliacionAt   NullTime    `json:"fafiliacion,omitempty"`
	BajaAt         NullTime    `json:"fbaja,omitempty"`
	LastAccessAt   NullTime    `json:"flastaccess,omitempty"`
	LastMovementAt NullTime    `json:"flastmovement,omitempty"`
	LastTransactAt NullTime    `json:"flasttransact,omitempty"`
	StatusPersona  NullInt32   `json:"status_persona,omitempty"`
	StatusDetail   NullString  `json:"status_detail,omitempty"`
	StatusDateAt   NullTime    `json:"status_date,omitempty"`
	Distance       NullFloat64 `json:"distance,omitempty"`
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
	Detail         DataPersonaDetailE
	Addrs          []DataPersonaAddressE
	Ids            []DataPersonasIdE
	Medios         []DataPersonasMedioE
	Roles          []DataPersonasRolE
	Trusteds       []DataPersonaTrustedE
}

func (e DataPersonasE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

type PersonaInfoE struct {
	Info      *DataPersonasE         `json:"info"`
	Detail    *DataPersonaDetailE    `json:"detail"`
	Ids       []*DataPersonasIdE     `json:"ids"`
	Medios    []*DataPersonasMedioE  `json:"medios"`
	Roles     []*DataPersonasRolE    `json:"roles"`
	Address   []*DataPersonaAddressE `json:"address"`
	Trusted   []*DataPersonaTrustedE `json:"trusted"`
	Commerces []*DataComercioE       `json:"commerces"`
	Terminals []*DataTerminalE       `json:"terminales"`
}

func (e PersonaInfoE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectDataPer = `select * from data_personas_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *DataPersonasE) GetAll(token string, filter string) ([]*DataPersonasE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectDataPer

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

	var lista []*DataPersonasE

	for rows.Next() {
		var rowdata DataPersonasE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Nroperacion,
			&rowdata.CredencialId,
			&rowdata.Nickname,
			&rowdata.Nombres,
			&rowdata.Midlename,
			&rowdata.Apaterno,
			&rowdata.Amaterno,
			&rowdata.Movil,
			&rowdata.Email,
			&rowdata.Avatar,
			&rowdata.MerchantId,
			&rowdata.CountryIso2,
			&rowdata.CountryIso3,
			&rowdata.CountryIsoM49,
			&rowdata.CountryPrefijo,
			&rowdata.PaisId,
			&rowdata.PaisText,
			&rowdata.RoleTypeId,
			&rowdata.Clasificacion,
			&rowdata.Robot,
			&rowdata.RobotAt,
			&rowdata.Mailing,
			&rowdata.Notifier,
			&rowdata.MailingAt,
			&rowdata.NotifierAt,
			&rowdata.AfiliacionAt,
			&rowdata.BajaAt,
			&rowdata.LastAccessAt,
			&rowdata.LastMovementAt,
			&rowdata.LastTransactAt,
			&rowdata.StatusPersona,
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

// GetOne returns one user by id
func (u *DataPersonasE) GetByUniqueid(token string, uniqueid int) (*DataPersonasE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectDataPer

	var rowdata DataPersonasE
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
		&rowdata.Nroperacion,
		&rowdata.CredencialId,
		&rowdata.Nickname,
		&rowdata.Nombres,
		&rowdata.Midlename,
		&rowdata.Apaterno,
		&rowdata.Amaterno,
		&rowdata.Movil,
		&rowdata.Email,
		&rowdata.Avatar,
		&rowdata.MerchantId,
		&rowdata.CountryIso2,
		&rowdata.CountryIso3,
		&rowdata.CountryIsoM49,
		&rowdata.CountryPrefijo,
		&rowdata.PaisId,
		&rowdata.PaisText,
		&rowdata.RoleTypeId,
		&rowdata.Clasificacion,
		&rowdata.Robot,
		&rowdata.RobotAt,
		&rowdata.Mailing,
		&rowdata.Notifier,
		&rowdata.MailingAt,
		&rowdata.NotifierAt,
		&rowdata.AfiliacionAt,
		&rowdata.BajaAt,
		&rowdata.LastAccessAt,
		&rowdata.LastMovementAt,
		&rowdata.LastTransactAt,
		&rowdata.StatusPersona,
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

// GetOne returns one user by id
func (u *DataPersonasE) GetByCredencialid(sede int, credencialid int) (*DataPersonasE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select uniqueid, nickname from data_personas_list( $1, $2)`

	var rowdata DataPersonasE
	jsonText := fmt.Sprintf(`{"sede":%d, "credencialid":%d}`, sede, credencialid)
	row := db.QueryRowContext(ctx, query, "", jsonText)

	err := row.Scan(
		&rowdata.Uniqueid,
		&rowdata.Nickname,
	)

	if err != nil {
		return nil, err
	}

	return &rowdata, nil
}

// GetOne returns one user by id
func (u *DataPersonasE) GetPersonaOkById(token string, personaid int) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM data_personas_ok ($1, $2)`

	row := db.QueryRowContext(ctx, query, token, personaid)

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

func (u *DataPersonasE) GetPersonaOkByNroperacion(token string, tokendataid string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM data_personas_ok_by_nroperacion ($1, $2)`

	row := db.QueryRowContext(ctx, query, token, tokendataid)

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

func (u *DataPersonasE) GetPersonaStatusById(token string, personaid int) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT * FROM data_personas_status ($1, $2)`

	row := db.QueryRowContext(ctx, query, token, personaid)

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
func (u *DataPersonasE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT data_personas_save($1, $2, $3)`

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
func (u *DataPersonasE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT data_personas_save($1, $2, $3)`
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
func (u *DataPersonasE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT data_personas_save($1, $2, $3)`
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
