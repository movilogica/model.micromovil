package model

const Language_English = "en"
const Language_Spanish = "es"
const Language_Portuges = "po"
const Language_French = "fr"

const Cookie_App = "micromovil"
const Cookie_Language = "i18n_lan"
const Cookie_ShoppingCart = "shopping_cart"
const Cookie_ShippingInfo = "shipping_info"
const Cookie_PaymentInfo = "payment_info"

const URL_DATA_SERVICE = "URL_DATA_SERVICE"
const URL_RABBITMQ = "URL_RABBITMQ"
const URL_REDIS = "URL_REDIS"
const URL_LOGGER = "URL_LOGGER"

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
const TBL_STORE_PARAM_ZONES = "storeparamzones"
const TBL_STORE_STORES = "storestores"
const TBL_STORE_CATALOGS = "storecatalog"
const TBL_STORE_CATALOGS_STORES = "storecatalogstore"
const TBL_STORE_CATALOGS_CATEG_ITEMS = "storecatalogcatitem"
const TBL_STORE_CATALOGS_PERSONS_ROLES = "storecatalogpersrole"
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
const TBL_STORE_INVENTORY_STATUS = "storeinventstatus"
const TBL_STORE_INVENTORY_HISTORY = "storeinventhist"
const TBL_STORE_INVENTORY_KARDEX = "storeinventkardex"
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

// --- VIEWS
const VIEW_RESUMEN_GLOBAL = "resumenglobal"

// Constantes
type KeyPair struct {
	Key   string
	Value string
	Group string
}
type ConstantesE struct {
	ModoAlmacenaje   []KeyPair /// ONE_PALLET, MULTIPLE, RETAIL
	ProveedPref      []KeyPair /// MAIN, ALTERNAL
	PropositoPrecio  []KeyPair /// PURCHASE, DEPOSIT, COMPONENT_PRICE
	ReasonMov        []KeyPair /// VAR_FOUND, VAR_LOST, VAR_DAMAGED, VAR_STOLEN, VAR_INTEGR, VAR_SAMPLE
	StatusInventario []KeyPair /// INV_AVAILABLE, INV_ON_HOLD, INV_DEFECTIVE, INV_RETURNED
	StatusUbicacion  []KeyPair /// AVAILABLE, OCUPPIED, NOT_AVAILABLE, RESERVED
	TipoCategoria    []KeyPair /// BEST_SELLING, CATALOG, CROSS_SELL, GIFT_CARDS, GOOGLE_BASE, INDUSTRY, INTERNAL, MATERIALS, MIX_AND_MATCH, QUICK_ADD, SEARCH, TAX, USAGE
	TipoIden         []KeyPair /// SKU, EAN, HS CODE, ISBN, LIBRARY, MANUFACTURE MODEL, MODEL YEAR, UPCA, UPCE, OTHER
	TipoInventario   []KeyPair /// NON_SERIAL_INV_ITEM, SERIALIZED_INV_ITEM
	TipoKeyword      []KeyPair /// KEYWORD, TAG
	TipoMov          []KeyPair /// VENTAS, COMPRAS, AJUSTE, TRANSFERENCIA, DEVOLUCION
	TipoPrecio       []KeyPair /// DEFAULT_PRICE, BOX_PRICE, AVERAGE_COST, LIST_PRICE, PROMO_PRICE, SPECIAL_PROMO_PRICE, MINIMUM_PRICE, MAXIMUM_PRICE, COMPETITIVE_PRICE, MINIMUM_ORDER_PRICE, WHOLESALE_PRICE
	TipoProducto     []KeyPair /// FINISHED_GOOD, SERVICE_PRODUCT, DIGITAL_GOOD, ASSET_USAGE, RAW_MATERIAL
	TipoRol          []KeyPair /// ACCOUNT, ADDRESSEE, ADMINISTRATOR, AFFILLIATE, AGENT, APPROVER, ASSOCIATION, BILL_FROM_VENDOR, BILL_FROM_CUSTOMER, BUYER, CALENDAR, CARBON COPY, CARRIER, CASHIER, CLIENT, COMPETITOR, CONSUMER, CONTACT, CONTRACT, CUSTOMER, DISTRIBUTOR, EMPLOYEE, IMAGE_APPROVER, MANAGER, MANUFACTURER, OWNER, PARTNER, PERSON, PICKER, RECEIVER, REQUEST_MANAGER, REQUEST_ROLE, SALES_FORCES, SALES_REPRESENTATIVE, SHIP_FROM_VENDOR, SHIP_TO_CUSTOMER, SHIPMENT_CLERK, SHARE_HOLDER, SPONSOR, SPOUSE, STOCKER, SUPPLIER, VENDOR, WORKER
	TipoUbicacion    []KeyPair /// FLT_BULK, FLT_PICKLOC
	TipoUom          []KeyPair /// AREA, CURRENCY, DATA SIZE, DATA SPEED, DRY VOLUME, ENERGIA, LONGITUD, LIQUID VOLUME, TEMPERATURE, TIME/FREQUENCY, UNIT, WEIGHT, OTRO
	TipoZona         []KeyPair /// STORAGE, PICKING, REPOSITION
}

func (obj *ConstantesE) InitValues() {
	obj.ModoAlmacenaje = []KeyPair{
		KeyPair{Key: "ONE_PALLET", Value: "UN PALLET"},
		KeyPair{Key: "MULTIPLE", Value: "MULTIPLES"},
		KeyPair{Key: "RETAIL", Value: "RETAIL"},
	}
	obj.TipoUom = []KeyPair{
		KeyPair{Key: "UNIT", Value: "UNIDAD"},
		KeyPair{Key: "WEIGHT", Value: "PESO"},
		KeyPair{Key: "LONGITUD", Value: "LONGITUD"},
		KeyPair{Key: "AREA", Value: "AREA"},
		KeyPair{Key: "CURRENCY", Value: "CURRENCY"},
		KeyPair{Key: "DATA_SIZE", Value: "DATA SIZE"},
		KeyPair{Key: "DATA_SPEED", Value: "DATA SPEED"},
		KeyPair{Key: "DRY_VOLUME", Value: "DRY VOLUME"},
		KeyPair{Key: "ENERGIA", Value: "ENERGIA"},
		KeyPair{Key: "LIQUID_VOLUME", Value: "LIQUID VOLUME"},
		KeyPair{Key: "TEMPERATURE", Value: "TEMPERATURE"},
		KeyPair{Key: "TIME_FREQUENCY", Value: "TIME/FREQUENCY"},
	}
}
