package model

import "encoding/json"

const Language_English = "en"
const Language_Spanish = "es"
const Language_Portuges = "po"
const Language_French = "fr"

const MENU_HOME = ""
const MENU_PARAMETERS = "parameters"
const MENU_MANAGEMENT = "management"

const Cookie_App = "micromovil"
const Cookie_Language = "i18n_lan"
const Cookie_SessionInfo = "session_info"
const Cookie_ShoppingCart = "shopping_cart"
const Cookie_ShippingInfo = "shipping_info"
const Cookie_PaymentInfo = "payment_info"

const APP_VERSION = "APP_VERSION"
const APP_CRM = "CRM"
const APP_STORE = "Store"
const APP_BALANCE = "Balance"
const APP_MICROFINANZAS = "Microfinanzas"
const APP_TORNEOS = "Torneos"
const APP_INVENTORY = "Inventory"
const APP_PEDIDOS = "Pedidos"

const URL_DATA_SERVICE = "URL_DATA_SERVICE"
const URL_RABBITMQ = "URL_RABBITMQ"
const URL_REDIS = "URL_REDIS"
const URL_LOGGER = "URL_LOGGER"
const URL_AUTH_RPC = "URL_AUTH_RPC"
const URL_BROKER_RPC = "URL_BROKER_RPC"

const ROLE_GUEST = "GUEST"
const ROLE_OPERATOR = "OPERATOR"
const ROLE_MANAGER = "MANAGER"

// --- Utilizados como campos clave para valores en session.cache
const MEM_ORDERS_OU_ITEMS = "orders-ou-items"
const MEM_ORDERS_OU_PEND = "orders-ou-pend"
const MEM_ORDERS_IN_PEND = "orders-in-pend"

// --- TABLAS SEGURIDAD
const TBL_SEDES = "sede"
const TBL_CREDENCIAL = "credencial"
const TBL_OPCIONES = "opciones"
const TBL_ROLES = "roles"
const TBL_ROLES_OPCIONES = "rolesopciones"
const TBL_ROLES_USUARIOS = "rolesusuarios"
const TBL_USUARIOS_ROLES = "usuariosroles"
const TBL_USUARIOS_SEDES = "usuariossedes"
const TBL_USUARIO_INFO = "usuarioinfo"
const TBL_USUARIO_LABEL = "usuariolabel"
const TBL_SESION_DATA = "sesiondata"

// --- TABLAS COMUNES
const TBL_PARAMETROS = "param" // Utilizado como referencia en algunos casos se debe indicar en Action del Request la 'tabla destino'
const TBL_CIIU = "ciiu"
const TBL_DIVISAS = "divisas"
const TBL_PAISES = "paises"
const TBL_REGIONES = "regiones"
const TBL_TIPO_DOCUMENTO = "tipodoc"
const TBL_ZIPCODE = "zipcode"
const TBL_BANCOS = "bancos"
const TBL_CARD_BRANDS = "cardbrands"
const TBL_DUTIES = "duties"
const TBL_DUTIES_DETAIL = "dutiesdetail"
const TBL_DUTIES_SCOPE = "dutiesscope"
const TBL_GIROS = "giros"
const TBL_TIPO_PRODUCTO = "tipoprod"
const TBL_HOLIDAYS = "holidays"
const TBL_TASA_CAMBIO = "tasacambio"

// --- PERSONAS
const TBL_CUSTOMER_INFO = "customerinfo"
const TBL_DATA_PERSONAS = "datapers"
const TBL_DATA_PERSONAS_OK = "datapersok"
const TBL_DATA_PERSONAS_STATUS = "datapersstatus"
const TBL_DATA_PERSONAS_DETAIL = "datapersdetail"
const TBL_DATA_PERSONAS_ID = "datapersid"
const TBL_DATA_PERSONAS_ADDR = "datapersaddr"
const TBL_DATA_PERSONAS_MEDIOS = "datapersmedios"
const TBL_DATA_PERSONAS_ROLES = "datapersroles"
const TBL_DATA_PERSONAS_TRUSTED = "dataperstrusted"

// --- COMERCIOS
const TBL_DATA_COMERCIOS = "datacomer"
const TBL_DATA_COMERCIOS_MINIMAL = "datacomerminimal"
const TBL_DATA_COMERCIOS_PARAM = "datacomerparam"
const TBL_DATA_COMERCIOS_PERSO = "datacomerperso"
const TBL_DATA_COMERCIOS_ROLES = "datacomerroles"
const TBL_DATA_COMERCIOS_BITACORA = "datacomerbitacora"

// --- TERMINALES
const TBL_DATA_TERMINALES = "dataterminal"

// --- MOVIMIENTOS
const TBL_MOV_PERSONAS = "movpers"
const TBL_MOV_TERMINAL = "movterminal"

// --- BIZ
const TBL_BIZ_PERSONAS = "bizpers"
const TBL_BIZ_PERSONAS_MEDIOS = "bizpersmedios"
const TBL_BIZ_PERSONAS_VEHICU = "bizpersvehicu"
const TBL_BIZ_SUBSCRIP = "bizsubscrip"
const TBL_BIZ_SUBSCRIP_STATUS = "bizsubscripstatus"
const TBL_BIZ_SUBSCRIP_PAYOFF = "bizsubscrippayoff"
const TBL_BIZ_PARAM_INVOICE_TYPE = "bizparaminvtype"
const TBL_BIZ_PARAM_INVOICE_ITEM_TYPE = "bizparaminvitemtype"
const TBL_BIZ_NOTES_TEXT = "biznotestext"
const TBL_BIZ_NOTES_VOICE = "biznotesvoice"
const TBL_BIZ_INVOICE_CAB = "bizinvoicecab"
const TBL_BIZ_INVOICE_DET = "bizinvoicedet"
const TBL_BIZ_KARDEX = "bizkardex"
const TBL_BIZ_KARDEX_PERSONA = "bizkardexpers"
const TBL_BIZ_INVOICE_APL = "bizinvapl"
const TBL_BIZ_INVOICE_STATUS = "bizinvstatus"
const TBL_BIZ_PAYMENT = "bizpay"
const TBL_BIZ_PAYMENT_APL = "bizpayapl"
const TBL_BIZ_PAYMENT_STATUS = "bizpaystatus"

// --- STORE
const TBL_STORE_PARAMETROS = "storeparametros"
const TBL_STORE_PARAM_CATEG_ITEM = "storeparamcatitem"
const TBL_STORE_PARAM_CATEG_ITEM_ATTR = "storeparamcatitemattr"
const TBL_STORE_PARAM_CATEG_ITEM_PERS = "storeparamcatitempers"
const TBL_STORE_PARAM_CATEG_UBICA = "storeparamcatubica"
const TBL_STORE_PARAM_CATEG_PICK = "storeparamcatpick"
const TBL_STORE_PARAM_CATEG_REPO = "storeparamcatrepo"
const TBL_STORE_PARAM_STORAGE_TYPE = "storeparamstotype"
const TBL_STORE_PARAM_ORDER_TYPE = "storeparamordertype"
const TBL_STORE_PARAM_TRANSACT_TYPE = "storeparamtransacttype"
const TBL_STORE_PARAM_ZONES = "storeparamzones"
const TBL_STORE_STORES = "storestores"
const TBL_STORE_CATALOGS = "storecatalog"
const TBL_STORE_CATALOGS_STORES = "storecatalogstore"
const TBL_STORE_CATALOGS_CATEG_ITEMS = "storecatalogcatitem"
const TBL_STORE_CATALOGS_PERSONS_ROLES = "storecatalogpersrole"
const TBL_STORE_ORDERS = "storeorders"
const TBL_STORE_ORDERS_INFO = "storeordersinfo"
const TBL_STORE_ORDERS_ITEMS = "storeordersitems"
const TBL_STORE_ORDERS_STATUS = "storeordersstatus"
const TBL_STORE_ORDERS_ACTIONS = "storeordersactions"
const TBL_STORE_ORDERS_DYE = "storedyeorders"
const TBL_STORE_ORDERS_DYE_INFO = "storedyeordersinfo"
const TBL_STORE_ORDERS_DYE_ITEMS = "storedyeordersitems"
const TBL_STORE_ORDERS_DYE_STATUS = "storedyeordersstatus"
const TBL_STORE_PRODUCTS = "storeproduct"
const TBL_STORE_PRODUCTS_KEYWORDS = "storeproductkeyword"
const TBL_STORE_PRODUCTS_IDEN = "storeproductiden"
const TBL_STORE_PRODUCTS_PRICES = "storeproductprices"
const TBL_STORE_PRODUCTS_SUPPLIERS = "storeproductsuppliers"
const TBL_STORE_PRODUCTS_CATEG_ITEMS = "storeproductcatitem"
const TBL_STORE_PRODUCTS_ATTR = "storeproductattr"
const TBL_STORE_PRODUCTS_ASSO = "storeproductasso"
const TBL_STORE_PRODUCTS_PERS = "storeproductpers"
const TBL_STORE_PRODUCTS_TIPOSPROD = "storeproducttiposprod"
const TBL_STORE_PRODUCTS_LOTES = "storeproductlotes"
const TBL_STORE_WAREHOUSE = "storewarehouse"
const TBL_STORE_WAREHOUSE_STAGES = "storewarehousestages"
const TBL_STORE_WAREHOUSE_LOCATIONS = "storewarehouseloca"
const TBL_STORE_WAREHOUSE_ZONES = "storewarehousezones"
const TBL_STORE_RULES_STORA = "storerulestora"
const TBL_STORE_RULES_STORA_ZONES = "storerulestorazones"
const TBL_STORE_RULES_STORA_ACTIONS = "storerulestoraactions"
const TBL_STORE_RULES_PICK = "storerulepick"
const TBL_STORE_RULES_PICK_ZONES = "storerulepickzones"
const TBL_STORE_RULES_PICK_ACTIONS = "storerulepickactions"
const TBL_STORE_RULES_REPO = "storerulerepo"
const TBL_STORE_RULES_REPO_ZONES = "storerulerepozones"
const TBL_STORE_RULES_REPO_ACTIONS = "storerulerepoactions"
const TBL_STORE_INVENTORY = "storeinventory"
const TBL_STORE_INVENTORY_EXTENDED = "storeinventoryextended"
const TBL_STORE_INVENTORY_BY_PROD = "storeinventorybyprod"
const TBL_STORE_INVENTORY_BY_PRESENTA = "storeinventorybypresenta"
const TBL_STORE_INVENTORY_BY_DIVISION = "storeinventorybydivision"
const TBL_STORE_INVENTORY_BY_DOCUMENT = "storeinventorydocuments"
const TBL_STORE_INVENTORY_STATUS = "storeinventstatus"
const TBL_STORE_INVENTORY_HISTORY = "storeinventhist"
const TBL_STORE_INVENTORY_KARDEX = "storeinventkardex"
const TBL_STORE_INVENTORY_STOCK_MINIMO = "storeinventstockmin"
const TBL_STORE_INVENTORY_TRANSACTIONS = "storeinventtransact"
const TBL_STORE_INVENTORY_TRANSACTIONS_DET = "storeinventtransactdet"
const TBL_STORE_RECOUNT = "storerecount"
const TBL_STORE_RECOUNT_GROUPS = "storerecountgroups"
const TBL_STORE_RECOUNT_BOXES = "storerecountboxes"
const TBL_STORE_RECOUNT_ITEMS = "storerecountitems"

// --- GALLOS
const TBL_COQ_FICHA_CRIADORES = "coqfichacriador"
const TBL_COQ_FICHA_GALPONES = "coqfichagalpon"
const TBL_COQ_FICHA_AVES = "coqfichaaves"
const TBL_COQ_FICHA_SKU = "coqfichasku"
const TBL_COQ_FICHA_FOTOS = "coqfichafotos"
const TBL_COQ_FICHA_CONSANGUINEA = "coqfichaconsanguinea"
const TBL_COQ_CRIADORES = "coqcriador"
const TBL_COQ_GALPONES = "coqgalpon"
const TBL_COQ_EMPADRONADOR = "coqempadrona"
const TBL_COQ_FECHA_GALPON = "coqfechagalpon"
const TBL_COQ_FECHA_TORNEO = "coqfechatorneo"
const TBL_COQ_GALPON_TORNEO = "coqgalpontorneo"
const TBL_COQ_PELEA = "coqpelea"
const TBL_COQ_PESAJE = "coqpesaje"
const TBL_COQ_REPORTE = "coqreporte"
const TBL_COQ_RANKING = "coqranking"
const TBL_COQ_SORTEO = "coqsorteo"
const TBL_COQ_TORNEOS = "coqtorneos"

// --- PROCEDIMIENTOS
const PROCEDURE_BIZ_FULLDATA = "bizpersfulldata"
const PROCEDURE_BIZ_PAYMENT = "bizpayment"
const PROCEDURE_BIZ_BILL = "bizbill"
const PROCEDURE_BIZ_SUBS_BILLS = "bizsubsbills"
const PROCEDURE_BIZ_ANULAR_BILL = "bizanularbill"
const PROCEDURE_REGISTER_ORDERS = "register_orders"
const PROCEDURE_UPDATE_ORDERS = "update_orders"
const PROCEDURE_RECEIVED_ORDERS = "received_orders"
const PROCEDURE_PICKUP_ORDERS = "pickup_orders"
const PROCEDURE_REGISTER_RECEIVED_ORDERS = "register_received_orders"
const PROCEDURE_REGISTER_PICKUP_ORDERS = "register_pickup_orders"
const PROCEDURE_CHANGE_STATUS_ORDERS = "change_status_orders"
const PROCEDURE_GENERATE_SKUS = "generate_skus"

const PROCEDURE_DYE_REGISTER = "register_dyeorders"
const PROCEDURE_DYE_UPDATE = "update_dyeorders"
const PROCEDURE_DYE_RECEIVED = "received_dyeorders"
const PROCEDURE_DYE_REGISTER_RECEIVED = "register_received_dyeorders"
const PROCEDURE_DYE_CHANGE_STATUS = "change_status_dyeorders"

const PROCEDURE_TRANSFER_STOCK = "transfer_stock"
const PROCEDURE_ADJUST_STOCK = "adjust_stock"
const PROCEDURE_UPDATE_INFO_STOCK = "update_info_stock"

// --- VIEWS
const VIEW_RESUMEN_GLOBAL = "resumenglobal"
const STORE_RESUMEN_GLOBAL = "storeresumenglobal"

// --- METODOS
const METHOD_ORDER_NEXT_NUMBER = "ordernextnumber"
const METHOD_PRODUCT_NEXT_NUMBER = "ordernextnumber"
const METHOD_DYEORDER_NEXT_NUMBER = "dyenextnumber"

// Constantes
type KeyPair struct {
	Key   string
	Value string
	Group string
}
type ConstantesE struct {
	ModoAlmacenaje     []KeyPair /// ONE_PALLET, MULTI_PALLET, RETAIL
	MultiAlmacenaje    []KeyPair /// ONE_PRODUCT, MULTI_PRODUCTS, NON_PRODUCTS
	ProveedPref        []KeyPair /// MAIN, ALTERNAL
	PropositoPrecio    []KeyPair /// PURCHASE, DEPOSIT, COMPONENT_PRICE
	ReasonMov          []KeyPair /// VAR_FOUND, VAR_LOST, VAR_DAMAGED, VAR_STOLEN, VAR_INTEGR, VAR_SAMPLE
	StatusInventario   []KeyPair /// INV_AVAILABLE, INV_ON_HOLD, INV_AT_DYE, INV_DEFECTIVE, INV_RETURNED
	StatusOrder        []KeyPair /// SCHEDULED, IN_PROCESS, ON_HOLD, CANCELLED, PROCESSED
	StatusUbicacion    []KeyPair /// AVAILABLE, OCUPPIED, NOT_AVAILABLE, RESERVED
	TipoCategoria      []KeyPair /// BEST_SELLING, CATALOG, CROSS_SELL, GIFT_CARDS, GOOGLE_BASE, INDUSTRY, INTERNAL, MATERIALS, MIX_AND_MATCH, QUICK_ADD, SEARCH, TAX, USAGE
	TipoCanal          []KeyPair /// MOBILE, PHONE, EMAIL, WEBSITE, FACEBOOK, INSTAGRAM, TWITTER, FAX
	TipoDomicilio      []KeyPair /// HOME, OFFICE, MAILING, DELIVERY, CONTACT
	TipoIden           []KeyPair /// SKU, EAN, HS CODE, ISBN, LIBRARY, MANUFACTURE MODEL, MODEL YEAR, UPCA, UPCE, OTHER
	TipoInventario     []KeyPair /// NON_SERIAL_INV_ITEM, SERIALIZED_INV_ITEM
	TipoKeyword        []KeyPair /// KEYWORD, TAG
	TipoMov            []KeyPair /// VENTAS, COMPRAS, AJUSTE, TRANSFERENCIA, DEVOLUCION
	TipoOrdenWarehouse []KeyPair /// INCOMING_SHIPMENT, OUTGOING_SHIPMENT, OUTGOING_PURCHASE_RETURN, INCOMING_PURCHASE_SHIPMENT, INCOMING_SALES_RETURN, OUTGOING_SALES_SHIPMENT, TRANSFER, DROP_SHIPMENT
	TipoPedido         []KeyPair /// GENERAL
	TipoPrecio         []KeyPair /// DEFAULT_PRICE, BOX_PRICE, AVERAGE_COST, LIST_PRICE, PROMO_PRICE, SPECIAL_PROMO_PRICE, MINIMUM_PRICE, MAXIMUM_PRICE, COMPETITIVE_PRICE, MINIMUM_ORDER_PRICE, WHOLESALE_PRICE
	ClaseProducto      []KeyPair /// FINISHED_GOOD, SERVICE_PRODUCT, DIGITAL_GOOD, ASSET_USAGE, SUPPLY, RAW_MATERIAL
	TipoRol            []KeyPair /// ACCOUNT, ADDRESSEE, ADMINISTRATOR, AFFILLIATE, AGENT, APPROVER, ASSOCIATION, BILL_FROM_VENDOR, BILL_FROM_CUSTOMER, BUYER, CALENDAR, CARBON COPY, CARRIER, CASHIER, CLIENT, COMPETITOR, CONSUMER, CONTACT, CONTRACT, CUSTOMER, DISTRIBUTOR, EMPLOYEE, IMAGE_APPROVER, MANAGER, MANUFACTURER, OWNER, PARTNER, PERSON, PICKER, RECEIVER, REQUEST_MANAGER, REQUEST_ROLE, SALES_FORCES, SALES_REPRESENTATIVE, SHIP_FROM_VENDOR, SHIP_TO_CUSTOMER, SHIPMENT_CLERK, SHARE_HOLDER, SPONSOR, SPOUSE, STOCKER, SUPPLIER, VENDOR, WORKER
	TipoTransaccion    []KeyPair /// ADJUSTMENT, TRANSFER
	TipoUbicacion      []KeyPair /// FLT_BULK, FLT_PICKLOC
	TipoUom            []KeyPair /// AREA, CURRENCY, DATA SIZE, DATA SPEED, DRY VOLUME, ENERGIA, LONGITUD, LIQUID VOLUME, TEMPERATURE, TIME/FREQUENCY, UNIT, WEIGHT, OTRO
	TipoZona           []KeyPair /// STORAGE, PICKING, REPOSITION
	RolesPersonal      []KeyPair /// STORAGE, PICKING, REPOSITION
}

func (obj *ConstantesE) PublicAccess() []KeyPair {
	return []KeyPair{
		{Key: "/datapers", Value: TBL_DATA_PERSONAS},
		{Key: "/datapers.do", Value: TBL_DATA_PERSONAS},
		{Key: "/sede", Value: TBL_SEDES},
		{Key: "/sede.do", Value: TBL_SEDES},
		{Key: "/credencial", Value: TBL_CREDENCIAL},
		{Key: "/credencial.do", Value: TBL_CREDENCIAL},
		{Key: TBL_OPCIONES, Value: TBL_OPCIONES},
		{Key: "/roles", Value: TBL_ROLES},
		{Key: "/roles.do", Value: TBL_ROLES},
		{Key: TBL_ROLES_OPCIONES, Value: TBL_ROLES_OPCIONES},
		{Key: "/credencial/roles/", Value: TBL_ROLES_USUARIOS},
		{Key: "/credencial.do/roles/", Value: TBL_ROLES_USUARIOS},
		{Key: TBL_USUARIOS_ROLES, Value: TBL_USUARIOS_ROLES},
		{Key: TBL_USUARIOS_SEDES, Value: TBL_USUARIOS_SEDES},
		{Key: TBL_USUARIO_INFO, Value: TBL_USUARIO_INFO},
		{Key: TBL_USUARIO_LABEL, Value: TBL_USUARIO_LABEL},
		{Key: TBL_SESION_DATA, Value: TBL_SESION_DATA},
		{Key: "/paises", Value: TBL_PAISES},
		{Key: "/paises.do", Value: TBL_PAISES},
		{Key: "/regiones", Value: TBL_REGIONES},
		{Key: "/regiones.do", Value: TBL_REGIONES},
		{Key: "/departamentos", Value: TBL_REGIONES},
		{Key: "/departamentos.do", Value: TBL_REGIONES},
		{Key: "/provincias", Value: TBL_REGIONES},
		{Key: "/provincias.do", Value: TBL_REGIONES},
		{Key: "/distritos", Value: TBL_REGIONES},
		{Key: "/distritos.do", Value: TBL_REGIONES},
		{Key: "/zipcodes", Value: TBL_ZIPCODE},
		{Key: "/zipcodes.do", Value: TBL_ZIPCODE},
		{Key: "/param/variable", Value: "variable"},
		{Key: "/param.do/variable", Value: "variable"},
	}
}

func (obj *ConstantesE) InitValues() {
	obj.ModoAlmacenaje = []KeyPair{
		{Key: "ONE_PALLET", Value: "UN PALLET"},
		{Key: "MULTI_PALLET", Value: "MULTI PALLET"},
		{Key: "RETAIL", Value: "RETAIL"},
	}
	obj.MultiAlmacenaje = []KeyPair{
		{Key: "ONE_PRODUCT", Value: "UN PRODUCTO"},
		{Key: "MULTI_PRODUCTS", Value: "MULTI PRODUCTOS"},
		{Key: "NON_PRODUCTS", Value: "SIN PRODUCTOS"},
	}
	obj.ProveedPref = []KeyPair{
		{Key: "MAIN", Value: "PRINCIPAL"},
		{Key: "ALTERNAL", Value: "ALTERNO"},
	}
	obj.PropositoPrecio = []KeyPair{
		{Key: "PURCHASE", Value: "COMPRA"},
		{Key: "DEPOSIT", Value: "DEPOSITO"},
		{Key: "COMPONENT_PRICE", Value: "PRECIO COMPONENTE"},
	}
	obj.ReasonMov = []KeyPair{ ///VAR_FOUND, VAR_LOST, VAR_DAMAGED, VAR_STOLEN, VAR_INTEGR, VAR_SAMPLE
		{Key: "VAR_DISTRIB", Value: "DISTRIBUCION"},
		{Key: "VAR_WRONG", Value: "MAL UBICADO"},
		{Key: "VAR_FOUND", Value: "ENCONTRADO"},
		{Key: "VAR_LOST", Value: "PERDIDA"},
		{Key: "VAR_DAMAGED", Value: "DAÑADO"},
		{Key: "VAR_STOLEN", Value: "ROBO"},
		{Key: "VAR_INTEGR", Value: "INTEGRIDAD"},
		{Key: "VAR_SAMPLE", Value: "MUESTRA"},
	}
	obj.StatusInventario = []KeyPair{ /// INV_AVAILABLE, INV_ON_HOLD, INV_DEFECTIVE, INV_RETURNED
		{Key: "INV_AVAILABLE", Value: "DISPONIBLE"},
		{Key: "INV_AT_DYE", Value: "TINTORERIA/EMBELL"},
		{Key: "INV_ON_HOLD", Value: "EN ESPERA"},
		{Key: "INV_DEFECTIVE", Value: "DEFECTUOSA"},
		{Key: "INV_RETURNED", Value: "RETORNADO"},
	}
	obj.StatusOrder = []KeyPair{ /// SCHEDULED, IN_PROCESS, ON_HOLD, CANCELLED, PROCESSED
		{Key: "SCHEDULED", Value: "AGENDADA"},    /// Por defecto esta en agenda
		{Key: "IN_PROCESS", Value: "EN PROCESO"}, /// En proceso
		{Key: "ON_HOLD", Value: "EN ESPERA"},     /// Es cuando necesita aprobacion
		{Key: "PARTIAL_RCV", Value: "RECEP.PARCIAL"},
		{Key: "PARTIAL_PCK", Value: "PICK.PARCIAL"},
		{Key: "AT_DYE_HOUSE", Value: "EN TINTORERIA"},
		{Key: "CANCELLED", Value: "CANCELADA"},
		{Key: "PROCESSED", Value: "PROCESADA"},
	}
	obj.StatusUbicacion = []KeyPair{ /// AVAILABLE, OCUPPIED, NOT_AVAILABLE, RESERVED
		{Key: "AVAILABLE", Value: "DISPONIBLE"},
		{Key: "OCUPPIED", Value: "OCUPADO"},
		{Key: "NOT_AVAILABLE", Value: "NO DISPONIBLE"},
		{Key: "RESERVED", Value: "RESERVADO"},
	}
	obj.TipoCanal = []KeyPair{ /// MOBILE, PHONE, EMAIL, WEBSITE, FACEBOOK, INSTAGRAM, TWITTER, FAX
		{Key: "MOBILE", Value: "TELEFONO MOVIL"},
		{Key: "PHONE", Value: "TELEFONO FIJO"},
		{Key: "EMAIL", Value: "EMAIL"},
		{Key: "WEBSITE", Value: "WEBSITE"},
		{Key: "FACEBOOK", Value: "FACEBOOK"},
		{Key: "INSTAGRAM", Value: "INSTAGRAM"},
		{Key: "TWITTER", Value: "TWITTER"},
		{Key: "FAX", Value: "FAX"},
	}

	obj.TipoCategoria = []KeyPair{ /// BEST_SELLING, CATALOG, CROSS_SELL, GIFT_CARDS, GOOGLE_BASE, INDUSTRY, INTERNAL, MATERIALS, MIX_AND_MATCH, QUICK_ADD, SEARCH, TAX, USAGE
		{Key: "BEST_SELLING", Value: "LA MAS VENDIDA"},
		{Key: "CATALOG", Value: "CATALOGO"},
		{Key: "CROSS_SELL", Value: "VENTA CRUZADA"},
		{Key: "GIFT_CARDS", Value: "TARJETA DE REGALO"},
		{Key: "GOOGLE_BASE", Value: "GOOGLE BASE"},
		{Key: "INDUSTRY", Value: "INDUSTRIAL"},
		{Key: "INTERNAL", Value: "USO INTERNO"},
		{Key: "MATERIALS", Value: "MATERIALES"},
		{Key: "MIX_AND_MATCH", Value: "MEZCLAR Y COMBINAR"},
		{Key: "QUICK_ADD", Value: "REGISTRO RAPIDO"},
		{Key: "SEARCH", Value: "BUSQUEDA"},
		{Key: "TAX", Value: "IMPUESTO"},
		{Key: "USAGE", Value: "DE USO"},
	}
	obj.TipoDomicilio = []KeyPair{ /// HOME, OFFICE, MAILING, DELIVERY, CONTACT
		{Key: "HOME", Value: "DOMICILIO"},
		{Key: "OFFICE", Value: "TRABAJO"},
		{Key: "MAILING", Value: "CORREO POSTAL"},
		{Key: "DELIVERY", Value: "DELIVERY"},
		{Key: "CONTACT", Value: "CONTACTO"},
	}
	obj.TipoIden = []KeyPair{ /// SKU, EAN, HS CODE, ISBN, LIBRARY, MANUFACTURE MODEL, MODEL YEAR, UPCA, UPCE, OTHER
		{Key: "SKU", Value: "SKU"},
		{Key: "EAN", Value: "EAN"},
		{Key: "HS_CODE", Value: "HS CODE"},
		{Key: "ISBN", Value: "ISBN"},
		{Key: "LIBRARY", Value: "BIBLIOTECA"},
		{Key: "MANUFACTURE_MODEL", Value: "MODELO DE FABRICACION"},
		{Key: "MODEL", Value: "MODELO"},
		{Key: "YEAR", Value: "AÑO"},
		{Key: "UPCA", Value: "UPCA"},
		{Key: "UPCE", Value: "UPCE"},
		{Key: "OTHER", Value: "OTRO"},
	}
	obj.TipoInventario = []KeyPair{ /// NON_SERIAL_INV_ITEM, SERIALIZED_INV_ITEM
		{Key: "NON_SERIAL_INV_ITEM", Value: "SIN SERIE"},
		{Key: "SERIALIZED_INV_ITEM", Value: "CON SERIE"},
	}
	obj.TipoKeyword = []KeyPair{ /// KEYWORD, TAG
		{Key: "KEYWORD", Value: "KEYWORD"},
		{Key: "TAG", Value: "TAG"},
	}
	obj.TipoMov = []KeyPair{ /// VENTAS, COMPRAS, AJUSTE, TRANSFERENCIA, DEVOLUCION
		{Key: "SALES", Value: "VENTAS"},           /// =1 OU
		{Key: "PURCHASE", Value: "COMPRAS"},       /// =2 IN
		{Key: "MANUFACTURE", Value: "PRODUCCION"}, /// =2 IN
		{Key: "ADJUST", Value: "AJUSTE"},          /// BOTH
		{Key: "TRANSFER", Value: "TRANSFERENCIA"}, /// BOTH
		{Key: "ROLLBACK", Value: "REVERTIDO"},     /// BOTH
		{Key: "RETURN", Value: "DEVOLUCION"},      /// BOTH
	}
	obj.TipoOrdenWarehouse = []KeyPair{ /// INCOMING_SHIPMENT, OUTGOING_SHIPMENT, OUTGOING_PURCHASE_RETURN, INCOMING_PURCHASE_SHIPMENT, INCOMING_SALES_RETURN, OUTGOING_SALES_SHIPMENT, TRANSFER, DROP_SHIPMENT
		{Key: "INCOMING_SHIPMENT", Value: "INGRESO ALMACEN"},
		{Key: "OUTGOING_SHIPMENT", Value: "SALIDA ENTREGA"},
		{Key: "OUTGOING_PURCH_RETU", Value: "SALIDA RETORNO/COMPRA"},
		{Key: "INCOMING_SALE_RETU", Value: "INGRESO RETORNO/VENTA"},
		{Key: "OUTGOING_SALE_SHIP", Value: "SALIDA ENTREGA/VENTA"},
		{Key: "DROP_SHIPMENT", Value: "ENTREGA DIRECTA"},
	}

	obj.TipoPedido = []KeyPair{ /// GENERAL
		{Key: "GENERAL", Value: "GENERAL"},
	}

	obj.TipoPrecio = []KeyPair{ /// DEFAULT_PRICE, BOX_PRICE, AVERAGE_COST, LIST_PRICE, PROMO_PRICE, SPECIAL_PROMO_PRICE, MINIMUM_PRICE, MAXIMUM_PRICE, COMPETITIVE_PRICE, MINIMUM_ORDER_PRICE, WHOLESALE_PRICE
		{Key: "DEFAULT_PRICE", Value: "PRECIO POR DEFECTO"},
		{Key: "BOX_PRICE", Value: "PRECIO CAJA"},
		{Key: "AVERAGE_COST", Value: "COSTO PROMEDIO"},
		{Key: "LIST_PRICE", Value: "PRECIO LISTA"},
		{Key: "PROMO_PRICE", Value: "PRECIO PROMOCION"},
		{Key: "SPECIAL_PROMO_PRICE", Value: "PROMOCION ESPECIAL"},
		{Key: "MINIMUM_PRICE", Value: "PRECIO MINIMO"},
		{Key: "MAXIMUM_PRICE", Value: "PRECIO MAXIMO"},
		{Key: "COMPETITIVE_PRICE", Value: "PRECIO COMPETITIVO"},
		{Key: "MINIMUM_ORDER_PRICE", Value: "PRECIO MINIMO PEDIDO"},
		{Key: "WHOLESALE_PRICE", Value: "PRECIO AL POR MAYOR"},
	}
	obj.ClaseProducto = []KeyPair{ /// FINISHED_GOOD, SERVICE_PRODUCT, DIGITAL_GOOD, ASSET_USAGE, RAW_MATERIAL
		{Key: "FINISHED_GOOD", Value: "PRODUCTO FINAL"},
		{Key: "SERVICE_PRODUCT", Value: "SERVICIO/PRODUCTO"},
		{Key: "DIGITAL_GOOD", Value: "BIEN DIGITAL"},
		{Key: "ASSET_USAGE", Value: "USADO COMO ACTIVO"},
		{Key: "SUPPLY", Value: "INSUMO"},
		{Key: "RAW_MATERIAL", Value: "CRUDO"},
	}
	obj.TipoRol = []KeyPair{ /// ACCOUNT, ADDRESSEE, ADMINISTRATOR, AFFILLIATE, AGENT, APPROVER, ASSOCIATION, BILL_FROM_VENDOR, BILL_FROM_CUSTOMER, BUYER, CALENDAR, CARBON COPY, CARRIER, CASHIER, CLIENT, COMPETITOR, CONSUMER, CONTACT, CONTRACT, CUSTOMER, DISTRIBUTOR, EMPLOYEE, IMAGE_APPROVER, MANAGER, MANUFACTURER, OWNER, PARTNER, PERSON, PICKER, RECEIVER, REQUEST_MANAGER, REQUEST_ROLE, SALES_FORCE, SALES_REPRESENTATIVE, SHIP_FROM_VENDOR, SHIP_TO_CUSTOMER, SHIPMENT_CLERK, SHARE_HOLDER, SPONSOR, SPOUSE, STOCKER, SUPPLIER, VENDOR, WORKER
		{Key: "SUPPLIER", Value: "ABASTECEDOR"},
		{Key: "SHARE_HOLDER", Value: "ACCIONISTA"},
		{Key: "ADMINISTRATOR", Value: "ADMINISTRADOR"},
		{Key: "AFFILLIATE", Value: "AFILIADO"},
		{Key: "AGENT", Value: "AGENTE"},
		{Key: "APPROVER", Value: "APROBADOR"},
		{Key: "IMAGE_APPROVER", Value: "APROBADOR DE IMAGEN"},
		{Key: "ASSOCIATION", Value: "ASOCIACION"},
		{Key: "CASHIER", Value: "CAJERO"},
		{Key: "CALENDAR", Value: "CALENDARIO"},
		{Key: "CUSTOMER", Value: "CLIENTE BIENES/SERV"},
		{Key: "CLIENT", Value: "CLIENTE SERV"},
		{Key: "COMPETITOR", Value: "COMPETIDOR"},
		{Key: "BUYER", Value: "COMPRADOR"},
		{Key: "CONSUMER", Value: "CONSUMIDOR"},
		{Key: "CONTACT", Value: "CONTACTO"},
		{Key: "ACCOUNT", Value: "CONTADOR"},
		{Key: "CONTRACT", Value: "CONTRATO"},
		{Key: "SPOUSE", Value: "CONYUGE"},
		{Key: "CARBON_COPY", Value: "COPIA DE CARBON"},
		{Key: "ADDRESSEE", Value: "DESTINATARIO"},
		{Key: "DISTRIBUTOR", Value: "DISTRIBUIDOR"},
		{Key: "EMPLOYE", Value: "EMPLEADO"},
		{Key: "SHIPMENT_CLERK", Value: "EMPLEADO DE ENVIO"},
		{Key: "SHIP_TO_CUSTOMER", Value: "ENVIO AL CLIENTE"},
		{Key: "SHIP_FROM_VENDOR", Value: "ENVIO DESDE PROVEEDOR"},
		{Key: "MANUFACTURER", Value: "FABRICANTE"},
		{Key: "BILL_FROM_CUSTOMER", Value: "FACTURA DEL CLIENTE"},
		{Key: "BILL_FROM_VENDOR", Value: "FACTURA DEL PROVEEDOR"},
		{Key: "SALES_FORCE", Value: "FUERZA DE VENTAS"},
		{Key: "MANAGER", Value: "GERENTE"},
		{Key: "STOCKER", Value: "GESTOR ALMACEN"},
		{Key: "WORKER", Value: "OBRERO"},
		{Key: "REQUEST_MANAGER", Value: "PEDIDOS MANAGER"},
		{Key: "REQUEST_ROLE", Value: "PEDIDOS USUARIO"},
		{Key: "PERSON", Value: "PERSONA"},
		{Key: "PICKER", Value: "PICKEADOR"},
		{Key: "OWNER", Value: "PROPIETARIO"},
		{Key: "VENDOR", Value: "PROVEEDOR"},
		{Key: "RECEIVER", Value: "RECEPTOR"},
		{Key: "SALES_REPRESENTATIVE", Value: "REPRESENTANTE VENTAS"},
		{Key: "PARTNER", Value: "SOCIO"},
		{Key: "SPONSOR", Value: "SPONSOR"},
		{Key: "CARRIER", Value: "TRANSPORTISTA"},
	}

	obj.TipoTransaccion = []KeyPair{ /// ADJUSTMENT, TRANSFER
		{Key: "ADJUSTMENT", Value: "AJUSTE"},
		{Key: "TRANSFER", Value: "TRANSFERENCIA"},
	}

	obj.TipoUbicacion = []KeyPair{ /// FLT_BULK, FLT_PICKLOC
		{Key: "FLT_BULK", Value: "VENTA A GRANEL"},
		{Key: "FLT_PICKLOC", Value: "PICKING/RECOGIDA"},
	}

	obj.TipoUom = []KeyPair{
		{Key: "UNIT", Value: "UNIDAD"},
		{Key: "WEIGHT", Value: "PESO"},
		{Key: "LONGITUD", Value: "LONGITUD"},
		{Key: "AREA", Value: "AREA"},
		{Key: "CURRENCY", Value: "CURRENCY"},
		{Key: "DATA_SIZE", Value: "DATA SIZE"},
		{Key: "DATA_SPEED", Value: "DATA SPEED"},
		{Key: "DRY_VOLUME", Value: "DRY VOLUME"},
		{Key: "ENERGIA", Value: "ENERGIA"},
		{Key: "LIQUID_VOLUME", Value: "LIQUID VOLUME"},
		{Key: "TEMPERATURE", Value: "TEMPERATURE"},
		{Key: "TIME_FREQUENCY", Value: "TIME/FREQUENCY"},
	}
	obj.TipoZona = []KeyPair{ /// STORAGE, PICKING, REPOSITION
		{Key: "STORAGE", Value: "ALMACENAJE"},
		{Key: "PICKING", Value: "PICKING"},
		{Key: "REPOSITION", Value: "REPOSICION"},
	}
	obj.RolesPersonal = []KeyPair{
		{Key: ROLE_GUEST, Value: "INVITADO"},
		{Key: ROLE_OPERATOR, Value: "OPERADOR"},
		{Key: ROLE_MANAGER, Value: "ADMINISTRADOR"},
	}
}

func (obj *ConstantesE) JsonText(keyvalue []KeyPair) string {
	mapData := make(map[string]string)
	for _, e := range keyvalue {
		mapData[e.Key] = e.Value
	}
	jsonBinary, _ := json.Marshal(mapData)
	return string(jsonBinary)
}

func (obj *ConstantesE) Text(key string) string {
	for _, e := range obj.ModoAlmacenaje {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.MultiAlmacenaje {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.ProveedPref {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.PropositoPrecio {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.ReasonMov {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.StatusInventario {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.StatusOrder {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.StatusUbicacion {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoCanal {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoCategoria {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoDomicilio {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoIden {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoInventario {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoKeyword {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoMov {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoPrecio {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.ClaseProducto {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoRol {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoUbicacion {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoUom {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoZona {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.RolesPersonal {
		if e.Key == key {
			return e.Value
		}
	}
	return key
}
