package model

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func NewDataAccess(dbPool *sql.DB) DataAccess {
	db = dbPool

	return DataAccess{
		Parametros:                 ParametersE{},
		Sede:                       SedeE{},
		Usuario:                    UsuarioE{},
		Credencial:                 CredencialE{},
		Opciones:                   OpcionesE{},
		Roles:                      RolesE{},
		RolesOpciones:              RolesOpcionesE{},
		RolesUsuarios:              RolesUsuariosE{},
		UsuariosRoles:              UsuariosRolesE{},
		UsuariosSedes:              UsuariosSedesE{},
		SesionData:                 SesionDataE{},
		Ciiu:                       CiiuE{},
		Divisa:                     DivisaE{},
		Pais:                       PaisE{},
		Region:                     RegionE{},
		TipoDocumento:              TipoDocumentoE{},
		ZipCode:                    ZipCodeE{},
		Banco:                      BancosE{},
		CardBrand:                  CardBrandE{},
		Dutie:                      DutieE{},
		DutieDetail:                DutieDetailE{},
		DutieScope:                 DutieScopeE{},
		Giros:                      GirosE{},
		Holiday:                    HolidayE{},
		TasaCambio:                 TasaCambioE{},
		DataPersonas:               DataPersonasE{},
		DataPersonaDetail:          DataPersonaDetailE{},
		DataPersonaAddress:         DataPersonaAddressE{},
		DataPersonaTrusted:         DataPersonaTrustedE{},
		DataPersonasMedio:          DataPersonasMedioE{},
		DataPersonasRol:            DataPersonasRolE{},
		DataPersonasId:             DataPersonasIdE{},
		DataComercio:               DataComercioE{},
		DataComercioParam:          DataComercioParamE{},
		DataComercioPersonal:       DataComercioPersonalE{},
		DataComercioRole:           DataComercioRoleE{},
		DataTerminal:               DataTerminalE{},
		CustomerInfo:               CustomerInfoE{},
		MovPersonas:                MovPersonasE{},
		MovTerminales:              MovTerminalesE{},
		BizPersonas:                BizPersonasE{},
		BizPersonasMedio:           BizPersonasMedioE{},
		BizPersonasVehicles:        BizPersonasVehiclesE{},
		BizPersonasAddress:         BizPersonasAddressE{},
		BizPersonasBitacora:        BizPersonasBitacoraE{},
		BizNotesText:               BizNotesTextE{},
		BizParamInvoiceType:        BizParamInvoiceTypeE{},
		BizParamInvoiceItemType:    BizParamInvoiceItemTypeE{},
		BizInvoice:                 BizInvoiceE{},
		BizInvoiceDet:              BizInvoiceDetE{},
		BizInvoiceStatus:           BizInvoiceStatusE{},
		BizPayment:                 BizPaymentE{},
		BizPaymentStatus:           BizPaymentStatusE{},
		BizKardex:                  BizKardexE{},
		BizSubscriptions:           BizSubscriptionsE{},
		BizSubscriptionsPay:        BizSubscriptionsPayE{},
		StoreProduct:               StoreProductE{},
		StoreParamCategItem:        StoreParamCategItemE{},
		StoreInventRecountItemE:    StoreInventRecountItemE{},
		StoreInventRecountSummaryE: StoreInventRecountSummaryE{},
		PersonaInfo:                PersonaInfoE{},
		TipoProducto:               TipoProductoE{},
		CoqFichaCriador:            CoqFichaCriadorE{},
		CoqFichaGalpon:             CoqFichaGalponE{},
		CoqCriador:                 CoqCriadorE{},
		CoqGalpon:                  CoqGalponE{},
		CoqGalponTorneo:            CoqGalponTorneoE{},
		CoqFechaGalpon:             CoqFechaGalponE{},
		CoqFechaTorneo:             CoqFechaTorneoE{},
		CoqEmpadronador:            CoqEmpadronadorE{},
		CoqPelea:                   CoqPeleaE{},
		CoqPesaje:                  CoqPesajeE{},
		CoqRanking:                 CoqRankingE{},
		CoqReference:               CoqReferenceE{},
		CoqSorteo:                  CoqSorteoE{},
		CoqTorneos:                 CoqTorneosE{},
	}
}

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type DataAccess struct {
	Parametros                 ParametersE
	Sede                       SedeE
	Usuario                    UsuarioE
	Credencial                 CredencialE
	Opciones                   OpcionesE
	Roles                      RolesE
	RolesOpciones              RolesOpcionesE
	RolesUsuarios              RolesUsuariosE
	UsuariosRoles              UsuariosRolesE
	UsuariosSedes              UsuariosSedesE
	SesionData                 SesionDataE
	Ciiu                       CiiuE
	Divisa                     DivisaE
	Pais                       PaisE
	Region                     RegionE
	TipoDocumento              TipoDocumentoE
	ZipCode                    ZipCodeE
	Banco                      BancosE
	CardBrand                  CardBrandE
	Dutie                      DutieE
	DutieDetail                DutieDetailE
	DutieScope                 DutieScopeE
	Giros                      GirosE
	Holiday                    HolidayE
	TasaCambio                 TasaCambioE
	DataPersonas               DataPersonasE
	DataPersonaDetail          DataPersonaDetailE
	DataPersonaAddress         DataPersonaAddressE
	DataPersonaTrusted         DataPersonaTrustedE
	DataPersonasMedio          DataPersonasMedioE
	DataPersonasRol            DataPersonasRolE
	DataPersonasId             DataPersonasIdE
	DataComercio               DataComercioE
	DataComercioParam          DataComercioParamE
	DataComercioPersonal       DataComercioPersonalE
	DataComercioRole           DataComercioRoleE
	DataTerminal               DataTerminalE
	PersonaInfo                PersonaInfoE
	MovPersonas                MovPersonasE
	MovTerminales              MovTerminalesE
	BizPersonas                BizPersonasE
	BizPersonasMedio           BizPersonasMedioE
	BizPersonasVehicles        BizPersonasVehiclesE
	BizPersonasAddress         BizPersonasAddressE
	BizPersonasBitacora        BizPersonasBitacoraE
	BizNotesText               BizNotesTextE
	BizParamInvoiceType        BizParamInvoiceTypeE
	BizParamInvoiceItemType    BizParamInvoiceItemTypeE
	BizInvoice                 BizInvoiceE
	BizInvoiceDet              BizInvoiceDetE
	BizInvoiceStatus           BizInvoiceStatusE
	BizPayment                 BizPaymentE
	BizPaymentStatus           BizPaymentStatusE
	BizKardex                  BizKardexE
	BizSubscriptions           BizSubscriptionsE
	BizSubscriptionsStatus     BizSubscriptionStatusE
	BizSubscriptionsPay        BizSubscriptionsPayE
	StoreProduct               StoreProductE
	StoreParamCategItem        StoreParamCategItemE
	StoreInventRecountItemE    StoreInventRecountItemE
	StoreInventRecountSummaryE StoreInventRecountSummaryE
	CustomerInfo               CustomerInfoE
	ResumenGlobal              ResumenGlobalE
	TipoProducto               TipoProductoE
	CoqFichaCriador            CoqFichaCriadorE
	CoqFichaGalpon             CoqFichaGalponE
	CoqCriador                 CoqCriadorE
	CoqGalpon                  CoqGalponE
	CoqGalponTorneo            CoqGalponTorneoE
	CoqFechaGalpon             CoqFechaGalponE
	CoqFechaTorneo             CoqFechaTorneoE
	CoqEmpadronador            CoqEmpadronadorE
	CoqPelea                   CoqPeleaE
	CoqPesaje                  CoqPesajeE
	CoqRanking                 CoqRankingE
	CoqReference               CoqReferenceE
	CoqSorteo                  CoqSorteoE
	CoqTorneos                 CoqTorneosE
}
