package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Tipos de autos
type CredencialE struct {
	Uniqueid          int64      `json:"uniqueid,omitempty"`
	Owner             NullInt32  `json:"owner,omitempty"`
	Dispositivoid     int32      `json:"dispositivoid,omitempty"`
	Id                int32      `json:"id,omitempty"`
	Sede              int32      `json:"sede"`
	Flag1             string     `json:"flag1,omitempty"`
	Flag2             string     `json:"flag2,omitempty"`
	Nivel             NullInt32  `json:"nivel,omitempty"`
	Tipo              NullString `json:"tipo,omitempty"`
	Username          NullString `json:"username,omitempty"`
	Email             NullString `json:"email,omitempty"`
	Movil             NullString `json:"movil,omitempty"`
	Imei              NullString `json:"imei,omitempty"`
	Social            NullString `json:"social,omitempty"`
	Googleidtoken     NullString `json:"googleidtoken,omitempty"`
	Password          NullString `json:"password,omitempty"`
	ModoAutenticacion NullInt32  `json:"modoautenticacion,omitempty"`
	ModoAcceso        NullInt32  `json:"modoacceso,omitempty"`
	Changepwd         NullInt32  `json:"changepwd,omitempty"`
	NextChangepwdAt   NullTime   `json:"fnextchangepwd,omitempty"`
	LastAccessAt      NullTime   `json:"flastaccess,omitempty"`
	SedeActual        NullInt64  `json:"sedeactual,omitempty"`
	Welcome           NullInt32  `json:"welcome,omitempty"`
	WelcomeAt         NullTime   `json:"fwelcome,omitempty"`
	Notifier          NullInt32  `json:"notifier,omitempty"`
	NotifierAt        NullTime   `json:"fnotifier,omitempty"`
	Mailing           NullInt32  `json:"mailing,omitempty"`
	MailingAt         NullTime   `json:"fmailing,omitempty"`
	Pdf               NullInt32  `json:"pdf,omitempty"`
	PdfAt             NullTime   `json:"fpdf,omitempty"`
	Idioma            NullString `json:"idioma,omitempty"`
	Avatar            NullString `json:"avatar,omitempty"`
	TokenTerminal     NullString `json:"tokenterminal,omitempty"`
	Manual            NullInt32  `json:"manual,omitempty"`
	Ruf1              NullString `json:"ruf1,omitempty"`
	Ruf2              NullString `json:"ruf2,omitempty"`
	Ruf3              NullString `json:"ruf3,omitempty"`
	Iv                NullString `json:"iv,omitempty"`
	Salt              NullString `json:"salt,omitempty"`
	Checksum          NullString `json:"checksum,omitempty"`
	FCreated          NullTime   `json:"fcreated,omitempty"`
	FUpdated          NullTime   `json:"fupdated,omitempty"`
	Estadoreg         NullInt64  `json:"estadoreg,omitempty"`
	Activo            NullInt64  `json:"activo,omitempty"`
	TotalRecords      int64      `json:"total_records"`
}

func (e CredencialE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectCredencial = `select * from security_credenciales_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *CredencialE) GetAll(token string, filter string) ([]*CredencialE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCredencial

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

	var lista []*CredencialE

	for rows.Next() {
		var rowdata CredencialE
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Owner,
			&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Nivel,
			&rowdata.Tipo,
			&rowdata.Username,
			&rowdata.Email,
			&rowdata.Movil,
			&rowdata.Imei,
			&rowdata.Social,
			&rowdata.Googleidtoken,
			&rowdata.Password,
			&rowdata.ModoAutenticacion,
			&rowdata.ModoAcceso,
			&rowdata.Changepwd,
			&rowdata.NextChangepwdAt,
			&rowdata.LastAccessAt,
			&rowdata.SedeActual,
			&rowdata.Welcome,
			&rowdata.WelcomeAt,
			&rowdata.Notifier,
			&rowdata.NotifierAt,
			&rowdata.Mailing,
			&rowdata.MailingAt,
			&rowdata.Pdf,
			&rowdata.PdfAt,
			&rowdata.Idioma,
			&rowdata.Avatar,
			&rowdata.TokenTerminal,
			&rowdata.Manual,
			&rowdata.Ruf1,
			&rowdata.Ruf2,
			&rowdata.Ruf3,
			&rowdata.Iv,
			&rowdata.Salt,
			&rowdata.Checksum,
			&rowdata.FCreated,
			&rowdata.FUpdated,
			&rowdata.Estadoreg,
			&rowdata.Activo,
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

// GetByField returns one record by filter
/*func (u *SedeE) GetByField(fieldname string, value string, token string) (*SedeE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select uniqueid, owner, dispositivoid, id, sede, flag1, flag2,
			         code, descrip,
					 ruf1, ruf2, ruf3, fcreated, fupdated, estadoreg, activo, total_records
			    from sys_sedes_list( $1, $2)`

	rows, err := db.QueryContext(ctx, query, token, `{"tipo":"TYPE_OF_CARS"}`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result *Parameters

	if rows.Next() {
		var rowdata Parameters
		err := rows.Scan(
			&rowdata.Uniqueid,
			&rowdata.Sede,
			&rowdata.Flag1,
			&rowdata.Flag2,
			&rowdata.Tipo,
			&rowdata.Secuencial,
			&rowdata.Code,
			&rowdata.Descrip,
			&rowdata.CreatedAt,
			&rowdata.UpdatedAt,
			&rowdata.Activo,
			&rowdata.Estadoreg,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		result = &rowdata
	}

	return result, nil
}*/

// GetOne returns one user by id
func (u *CredencialE) GetByUniqueid(token string, uniqueid int) (*CredencialE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCredencial

	var rowdata CredencialE
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
		&rowdata.Nivel,
		&rowdata.Tipo,
		&rowdata.Username,
		&rowdata.Email,
		&rowdata.Movil,
		&rowdata.Imei,
		&rowdata.Social,
		&rowdata.Googleidtoken,
		&rowdata.Password,
		&rowdata.ModoAutenticacion,
		&rowdata.ModoAcceso,
		&rowdata.Changepwd,
		&rowdata.NextChangepwdAt,
		&rowdata.LastAccessAt,
		&rowdata.SedeActual,
		&rowdata.Welcome,
		&rowdata.WelcomeAt,
		&rowdata.Notifier,
		&rowdata.NotifierAt,
		&rowdata.Mailing,
		&rowdata.MailingAt,
		&rowdata.Pdf,
		&rowdata.PdfAt,
		&rowdata.Idioma,
		&rowdata.Avatar,
		&rowdata.TokenTerminal,
		&rowdata.Manual,
		&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,
		&rowdata.FCreated,
		&rowdata.FUpdated,
		&rowdata.Estadoreg,
		&rowdata.Activo,
		&rowdata.TotalRecords,
	)

	if err != nil {
		return nil, err
	}

	return &rowdata, nil
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *CredencialE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT security_credenciales_save($1, $2, $3)`
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
func (u *CredencialE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT security_credenciales_save($1, $2, $3)`
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
func (u *CredencialE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, 300)

	query := `SELECT security_credenciales_save($1, $2, $3)`
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
