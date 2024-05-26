package model

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

///public static int CAMPEONATO_INACTIVO = 0;
///public static int CAMPEONATO_ABIERTO = 1;
///public static int CAMPEONATO_CERRADO = 2;

// Torneos
type CoqTorneosE struct {
	Uniqueid             int64       `json:"uniqueid,omitempty"`
	Owner                NullInt32   `json:"owner,omitempty"`
	Dispositivoid        NullInt32   `json:"dispositivoid,omitempty"`
	Id                   int32       `json:"id,omitempty"`
	Sede                 int32       `json:"sede"`
	Flag1                string      `json:"flag1,omitempty"`
	Flag2                string      `json:"flag2,omitempty"`
	CountryCode          NullString  `json:"paiscode,omitempty"`
	Periodo              NullString  `json:"periodo,omitempty"`
	Nombre               NullString  `json:"nombre,omitempty"`
	Nroperacion          NullString  `json:"nroperacion"`
	Publico              int32       `json:"publico"`
	CodigoConsulta       NullString  `json:"codigo_consulta"`
	FApertura            NullTime    `json:"fecha_apertura,omitempty"`
	FLimiteEmpadrona     NullTime    `json:"fecha_limite_empadronamiento,omitempty"`
	FCierre              NullTime    `json:"fecha_cierre,omitempty"`
	PaisId               NullInt64   `json:"paisid,omitempty"`
	PaisText             NullString  `json:"paistext,omitempty"`
	CantidadMinAves      int32       `json:"cantidad_minima_aves"`
	CantidadMaxAves      int32       `json:"cantidad_maxima_aves"`
	CantidadFechas       int32       `json:"cantidad_fechas"`
	CantidadFechasDia    int32       `json:"cantidad_fechas_dias"`
	CantidadGallosFrente int32       `json:"cantidad_gallos_frente"`
	PesoMinAves          float64     `json:"peso_minimo_aves"`
	PesoMaxAves          float64     `json:"peso_maximo_aves"`
	PesoDiffCoteja       float64     `json:"peso_maximo_diff_coteja"`
	LimitePollonesSeg    int32       `json:"limite_pollones_seg"`
	MostarPollonesLimite int32       `json:"mostrar_pollones_limite"`
	TablasSegundos       int32       `json:"tablas_seg"`
	UrlLogotipo          NullString  `json:"url_logotipo,omitempty"`
	UrlColiseo           NullString  `json:"url_coliseo,omitempty"`
	SoloSocios           int32       `json:"solo_socios"`
	PesoAgil             int32       `json:"peso_registro_agil"`
	TorneoAnual          int32       `json:"torneo_anual"`
	RankingConsolidado   int32       `json:"ranking_consolidado"`
	RankingSoloPadres    int32       `json:"ranking_solo_padres"`
	RankingExclTardios   int32       `json:"ranking_excl_tardios"`
	RankingNoAcumulable  int32       `json:"ranking_no_acumulable"`
	Balanza              int32       `json:"balanza,omitempty"`
	ShowInicio           int32       `json:"showinicio,omitempty"`
	ValidarPagos         int32       `json:"validar_pagos_ind"`
	HoraPesoMax          NullTime    `json:"hora_peso_maxima,omitempty"`
	ValidarPlacas        int32       `json:"validar_placas_ind"`
	PuntosVictoria       int32       `json:"puntos_victoria"`
	PuntosEmpate         int32       `json:"puntos_empate"`
	PuntosPerdida        int32       `json:"puntos_perdida"`
	TorneoAnterior       NullInt64   `json:"campeonato_anterior"`
	QEmpadronaGalpones   int32       `json:"q_empadrona_galpones"`
	QEmpadronaAves       int32       `json:"q_empadrona_aves"`
	UltEmpadrona         NullString  `json:"ult_empadrona"`
	FUltEmpadrona        NullTime    `json:"fult_empadrona"`
	RegionId             NullInt64   `json:"regionid,omitempty"`
	RegionText           NullString  `json:"regiontext,omitempty"`
	DepartamentoId       NullInt64   `json:"departamentoid,omitempty"`
	DepartamentoText     NullString  `json:"departamentotext,omitempty"`
	ProvinciaId          NullInt64   `json:"provinciaid,omitempty"`
	ProvinciaText        NullString  `json:"provinciatext,omitempty"`
	DistritoId           NullInt64   `json:"distritoid,omitempty"`
	DistritoText         NullString  `json:"distritotext,omitempty"`
	Latitud              NullString  `json:"latitud,omitempty"`
	Longitud             NullString  `json:"longitud,omitempty"`
	FechaProxima         NullString  `json:"fecha_proxima,omitempty"` /// Analisis en base a las fechas del torneo
	EstadoId             int32       `json:"estado_id"`
	EstadoText           NullString  `json:"estadotext,omitempty"`
	Ruf1                 NullString  `json:"ruf1,omitempty"`
	Ruf2                 NullString  `json:"ruf2,omitempty"`
	Ruf3                 NullString  `json:"ruf3,omitempty"`
	Iv                   NullString  `json:"iv,omitempty"`
	Salt                 NullString  `json:"salt,omitempty"`
	Checksum             NullString  `json:"checksum,omitempty"`
	FCreated             NullTime    `json:"fcreated,omitempty"`
	FUpdated             NullTime    `json:"fupdated,omitempty"`
	Activo               int32       `json:"activo,omitempty"`
	Estadoreg            int32       `json:"estadoreg,omitempty"`
	Distance             NullFloat64 `json:"distance,omitempty"`
	TotalRecords         int64       `json:"total_records,omitempty"`
}

func (e CoqTorneosE) MarshalJSON() ([]byte, error) {
	return MarshalJSON_Not_Nulls(e)
}

const querySelectCoqTorneos = `select * from coq_torneos_list( $1, $2)`

//---------------------------------------------------------------------
//MySQL               PostgreSQL            Oracle
//=====               ==========            ======
//WHERE col = ?       WHERE col = $1        WHERE col = :col
//VALUES(?, ?, ?)     VALUES($1, $2, $3)    VALUES(:val1, :val2, :val3)
//---------------------------------------------------------------------

// GetAll returns a slice of all users, sorted by last name
func (u *CoqTorneosE) GetAll(token string, filter string) ([]*CoqTorneosE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCoqTorneos

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

	var lista []*CoqTorneosE

	for rows.Next() {
		var rowdata CoqTorneosE
		err := rows.Scan(
			&rowdata.Uniqueid,
			//&rowdata.Owner,
			//&rowdata.Dispositivoid,
			&rowdata.Id,
			&rowdata.Sede,
			//&rowdata.Flag1,
			//&rowdata.Flag2,
			&rowdata.Periodo,
			&rowdata.Nombre,
			&rowdata.Nroperacion,
			&rowdata.CodigoConsulta,
			&rowdata.FApertura,
			&rowdata.FCierre,
			&rowdata.FLimiteEmpadrona,
			&rowdata.HoraPesoMax,
			&rowdata.FechaProxima,
			&rowdata.CantidadMinAves,
			&rowdata.CantidadMaxAves,
			&rowdata.CantidadFechas,
			&rowdata.CantidadFechasDia,
			&rowdata.CantidadGallosFrente,
			&rowdata.PesoMinAves,
			&rowdata.PesoMaxAves,
			&rowdata.PesoDiffCoteja,
			&rowdata.PuntosVictoria,
			&rowdata.PuntosEmpate,
			&rowdata.PuntosPerdida,
			&rowdata.LimitePollonesSeg,
			&rowdata.MostarPollonesLimite,
			&rowdata.TablasSegundos,
			&rowdata.ValidarPagos,
			&rowdata.ValidarPlacas,
			&rowdata.EstadoId,
			&rowdata.EstadoText,
			&rowdata.QEmpadronaGalpones,
			&rowdata.QEmpadronaAves,
			&rowdata.UltEmpadrona,
			&rowdata.FUltEmpadrona,
			&rowdata.RankingNoAcumulable,
			&rowdata.RankingConsolidado,
			&rowdata.RankingSoloPadres,
			&rowdata.RankingExclTardios,
			&rowdata.SoloSocios,
			&rowdata.TorneoAnual,
			&rowdata.PesoAgil,
			&rowdata.UrlLogotipo,
			&rowdata.UrlColiseo,
			&rowdata.Balanza,
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
			&rowdata.CountryCode,
			&rowdata.Latitud,
			&rowdata.Longitud,
			&rowdata.ShowInicio,
			&rowdata.Publico,
			&rowdata.TorneoAnterior,
			&rowdata.Distance,
			/*&rowdata.Ruf1,
			&rowdata.Ruf2,
			&rowdata.Ruf3,
			&rowdata.Iv,
			&rowdata.Salt,
			&rowdata.Checksum,*/
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
func (u *CoqTorneosE) GetByUniqueid(token string, uniqueid int) (*CoqTorneosE, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := querySelectCoqTorneos

	var rowdata CoqTorneosE
	jsonText := fmt.Sprintf(`{"uniqueid":%d}`, uniqueid)
	row := db.QueryRowContext(ctx, query, token, jsonText)

	err := row.Scan(
		&rowdata.Uniqueid,
		//&rowdata.Owner,
		//&rowdata.Dispositivoid,
		&rowdata.Id,
		&rowdata.Sede,
		//&rowdata.Flag1,
		//&rowdata.Flag2,
		&rowdata.Periodo,
		&rowdata.Nombre,
		&rowdata.Nroperacion,
		&rowdata.CodigoConsulta,
		&rowdata.FApertura,
		&rowdata.FCierre,
		&rowdata.FLimiteEmpadrona,
		&rowdata.HoraPesoMax,
		&rowdata.FechaProxima,
		&rowdata.CantidadMinAves,
		&rowdata.CantidadMaxAves,
		&rowdata.CantidadFechas,
		&rowdata.CantidadFechasDia,
		&rowdata.CantidadGallosFrente,
		&rowdata.PesoMinAves,
		&rowdata.PesoMaxAves,
		&rowdata.PesoDiffCoteja,
		&rowdata.PuntosVictoria,
		&rowdata.PuntosEmpate,
		&rowdata.PuntosPerdida,
		&rowdata.LimitePollonesSeg,
		&rowdata.MostarPollonesLimite,
		&rowdata.TablasSegundos,
		&rowdata.ValidarPagos,
		&rowdata.ValidarPlacas,
		&rowdata.EstadoId,
		&rowdata.EstadoText,
		&rowdata.QEmpadronaGalpones,
		&rowdata.QEmpadronaAves,
		&rowdata.UltEmpadrona,
		&rowdata.FUltEmpadrona,
		&rowdata.RankingNoAcumulable,
		&rowdata.RankingConsolidado,
		&rowdata.RankingSoloPadres,
		&rowdata.RankingExclTardios,
		&rowdata.SoloSocios,
		&rowdata.TorneoAnual,
		&rowdata.PesoAgil,
		&rowdata.UrlLogotipo,
		&rowdata.UrlColiseo,
		&rowdata.Balanza,
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
		&rowdata.CountryCode,
		&rowdata.Latitud,
		&rowdata.Longitud,
		&rowdata.ShowInicio,
		&rowdata.Publico,
		&rowdata.TorneoAnterior,
		&rowdata.Distance,
		/*&rowdata.Ruf1,
		&rowdata.Ruf2,
		&rowdata.Ruf3,
		&rowdata.Iv,
		&rowdata.Salt,
		&rowdata.Checksum,*/
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
func (u *CoqTorneosE) Update(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT coq_torneos_save($1, $2, $3)`
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
func (u *CoqTorneosE) Delete(token string, data string, metricas string) (map[string]any, error) {
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

	query := `SELECT coq_torneos_save($1, $2, $3)`
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
func (u *CoqTorneosE) DeleteByID(token string, id int, metricas string) (map[string]any, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	jsonText := fmt.Sprintf(`{"id":%d, 
							  "uniqueid":%d, 
							  "estadoreg":%d
							  }`,
		id, id, 300)

	query := `SELECT coq_torneos_save($1, $2, $3)`
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
