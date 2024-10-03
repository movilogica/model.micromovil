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
	ModoAlmacenaje   []KeyPair /// ONE_PALLET, MULTI_PALLET, RETAIL
	MultiAlmacenaje  []KeyPair /// ONE_PRODUCT, MULTI_PRODUCTS, NON_PRODUCTS
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
	obj.MultiAlmacenaje = []KeyPair{
		KeyPair{Key: "ONE_PRODUCT", Value: "UN PRODUCTO"},
		KeyPair{Key: "MULTI_PRODUCTS", Value: "MULTI PRODUCTOS"},
		KeyPair{Key: "NON_PRODUCTS", Value: "SIN PRODUCTOS"},
	}
	obj.ProveedPref = []KeyPair{
		KeyPair{Key: "MAIN", Value: "PRINCIPAL"},
		KeyPair{Key: "ALTERNAL", Value: "ALTERNO"},
	}
	obj.PropositoPrecio = []KeyPair{
		KeyPair{Key: "PURCHASE", Value: "COMPRA"},
		KeyPair{Key: "DEPOSIT", Value: "DEPOSITO"},
		KeyPair{Key: "COMPONENT_PRICE", Value: "PRECIO COMPONENTE"},
	}
	obj.ReasonMov = []KeyPair{ ///VAR_FOUND, VAR_LOST, VAR_DAMAGED, VAR_STOLEN, VAR_INTEGR, VAR_SAMPLE
		KeyPair{Key: "VAR_FOUND", Value: "ENCONTRADO"},
		KeyPair{Key: "VAR_LOST", Value: "PERDIDA"},
		KeyPair{Key: "VAR_DAMAGED", Value: "DAÑADO"},
		KeyPair{Key: "VAR_STOLEN", Value: "ROBO"},
		KeyPair{Key: "VAR_INTEGR", Value: "INTEGRIDAD"},
		KeyPair{Key: "VAR_SAMPLE", Value: "MUESTRA"},
	}
	obj.StatusInventario = []KeyPair{ /// INV_AVAILABLE, INV_ON_HOLD, INV_DEFECTIVE, INV_RETURNED
		KeyPair{Key: "INV_AVAILABLE", Value: "DISPONIBLE"},
		KeyPair{Key: "INV_ON_HOLD", Value: "EN ESPERA"},
		KeyPair{Key: "INV_DEFECTIVE", Value: "DEFECTUOSA"},
		KeyPair{Key: "INV_RETURNED", Value: "RETORNADO"},
	}
	obj.StatusUbicacion = []KeyPair{ /// AVAILABLE, OCUPPIED, NOT_AVAILABLE, RESERVED
		KeyPair{Key: "AVAILABLE", Value: "DISPONIBLE"},
		KeyPair{Key: "OCUPPIED", Value: "OCUPADO"},
		KeyPair{Key: "NOT_AVAILABLE", Value: "NO DISPONIBLE"},
		KeyPair{Key: "RESERVED", Value: "RESERVADO"},
	}
	obj.TipoCategoria = []KeyPair{ /// BEST_SELLING, CATALOG, CROSS_SELL, GIFT_CARDS, GOOGLE_BASE, INDUSTRY, INTERNAL, MATERIALS, MIX_AND_MATCH, QUICK_ADD, SEARCH, TAX, USAGE
		KeyPair{Key: "BEST_SELLING", Value: "LA MAS VENDIDA"},
		KeyPair{Key: "CATALOG", Value: "CATALOGO"},
		KeyPair{Key: "CROSS_SELL", Value: "VENTA CRUZADA"},
		KeyPair{Key: "GIFT_CARDS", Value: "TARJETA DE REGALO"},
		KeyPair{Key: "GOOGLE_BASE", Value: "GOOGLE BASE"},
		KeyPair{Key: "INDUSTRY", Value: "INDUSTRIAL"},
		KeyPair{Key: "INTERNAL", Value: "USO INTERNO"},
		KeyPair{Key: "MATERIALS", Value: "MATERIALES"},
		KeyPair{Key: "MIX_AND_MATCH", Value: "MEZCLAR Y COMBINAR"},
		KeyPair{Key: "QUICK_ADD", Value: "REGISTRO RAPIDO"},
		KeyPair{Key: "SEARCH", Value: "BUSQUEDA"},
		KeyPair{Key: "TAX", Value: "IMPUESTO"},
		KeyPair{Key: "USAGE", Value: "DE USO"},
	}
	obj.TipoIden = []KeyPair{ /// SKU, EAN, HS CODE, ISBN, LIBRARY, MANUFACTURE MODEL, MODEL YEAR, UPCA, UPCE, OTHER
		KeyPair{Key: "SKU", Value: "SKU"},
		KeyPair{Key: "EAN", Value: "EAN"},
		KeyPair{Key: "HS_CODE", Value: "HS CODE"},
		KeyPair{Key: "ISBN", Value: "ISBN"},
		KeyPair{Key: "LIBRARY", Value: "BIBLIOTECA"},
		KeyPair{Key: "MANUFACTURE_MODEL", Value: "MODELO DE FABRICACION"},
		KeyPair{Key: "MODEL", Value: "MODELO"},
		KeyPair{Key: "YEAR", Value: "AÑO"},
		KeyPair{Key: "UPCA", Value: "UPCA"},
		KeyPair{Key: "UPCE", Value: "UPCE"},
		KeyPair{Key: "OTHER", Value: "OTRO"},
	}
	obj.TipoInventario = []KeyPair{ /// NON_SERIAL_INV_ITEM, SERIALIZED_INV_ITEM
		KeyPair{Key: "NON_SERIAL_INV_ITEM", Value: "SIN SERIE"},
		KeyPair{Key: "SERIALIZED_INV_ITEM", Value: "CON SERIE"},
	}
	obj.TipoKeyword = []KeyPair{ /// KEYWORD, TAG
		KeyPair{Key: "KEYWORD", Value: "KEYWORD"},
		KeyPair{Key: "TAG", Value: "TAG"},
	}
	obj.TipoMov = []KeyPair{ /// VENTAS, COMPRAS, AJUSTE, TRANSFERENCIA, DEVOLUCION
		KeyPair{Key: "VENTAS", Value: "VENTAS"},
		KeyPair{Key: "COMPRAS", Value: "COMPRAS"},
		KeyPair{Key: "AJUSTE", Value: "AJUSTE"},
		KeyPair{Key: "TRANSFERENCIA", Value: "TRANSFERENCIA"},
		KeyPair{Key: "DEVOLUCION", Value: "DEVOLUCION"},
	}
	obj.TipoPrecio = []KeyPair{ /// DEFAULT_PRICE, BOX_PRICE, AVERAGE_COST, LIST_PRICE, PROMO_PRICE, SPECIAL_PROMO_PRICE, MINIMUM_PRICE, MAXIMUM_PRICE, COMPETITIVE_PRICE, MINIMUM_ORDER_PRICE, WHOLESALE_PRICE
		KeyPair{Key: "DEFAULT_PRICE", Value: "PRECIO POR DEFECTO"},
		KeyPair{Key: "BOX_PRICE", Value: "PRECIO DE CAJA"},
		KeyPair{Key: "AVERAGE_COST", Value: "COSTO PROMEDIO"},
		KeyPair{Key: "LIST_PRICE", Value: "PRECIO DE LISTA"},
		KeyPair{Key: "PROMO_PRICE", Value: "PRECIO DE PROMOCION"},
		KeyPair{Key: "SPECIAL_PROMO_PRICE", Value: "PROMOCION ESPECIAL"},
		KeyPair{Key: "MINIMUM_PRICE", Value: "PRECIO MINIMO"},
		KeyPair{Key: "MAXIMUM_PRICE", Value: "PRECIO MAXIMO"},
		KeyPair{Key: "COMPETITIVE_PRICE", Value: "PRECIO COMPETITIVO"},
		KeyPair{Key: "MINIMUM_ORDER_PRICE", Value: "PRECIO MINIMO DE PEDIDO"},
		KeyPair{Key: "WHOLESALE_PRICE", Value: "PRECIO AL POR MAYOR"},
	}
	obj.TipoProducto = []KeyPair{ /// FINISHED_GOOD, SERVICE_PRODUCT, DIGITAL_GOOD, ASSET_USAGE, RAW_MATERIAL
		KeyPair{Key: "FINISHED_GOOD", Value: "PRODUCTO FINAL"},
		KeyPair{Key: "SERVICE_PRODUCT", Value: "SERVICIO/PRODUCTO"},
		KeyPair{Key: "DIGITAL_GOOD", Value: "BIEN DIGITAL"},
		KeyPair{Key: "ASSET_USAGE", Value: "USADO COMO ACTIVO"},
		KeyPair{Key: "RAW_MATERIAL", Value: "INSUMO"},
	}
	obj.TipoRol = []KeyPair{ /// ACCOUNT, ADDRESSEE, ADMINISTRATOR, AFFILLIATE, AGENT, APPROVER, ASSOCIATION, BILL_FROM_VENDOR, BILL_FROM_CUSTOMER, BUYER, CALENDAR, CARBON COPY, CARRIER, CASHIER, CLIENT, COMPETITOR, CONSUMER, CONTACT, CONTRACT, CUSTOMER, DISTRIBUTOR, EMPLOYEE, IMAGE_APPROVER, MANAGER, MANUFACTURER, OWNER, PARTNER, PERSON, PICKER, RECEIVER, REQUEST_MANAGER, REQUEST_ROLE, SALES_FORCE, SALES_REPRESENTATIVE, SHIP_FROM_VENDOR, SHIP_TO_CUSTOMER, SHIPMENT_CLERK, SHARE_HOLDER, SPONSOR, SPOUSE, STOCKER, SUPPLIER, VENDOR, WORKER
		KeyPair{Key: "SUPPLIER", Value: "ABASTECEDOR"},
		KeyPair{Key: "SHARE_HOLDER", Value: "ACCIONISTA"},
		KeyPair{Key: "ADMINISTRATOR", Value: "ADMINISTRADOR"},
		KeyPair{Key: "AFFILLIATE", Value: "AFILIADO"},
		KeyPair{Key: "AGENT", Value: "AGENTE"},
		KeyPair{Key: "APPROVER", Value: "APROBADOR"},
		KeyPair{Key: "IMAGE_APPROVER", Value: "APROBADOR DE IMAGEN"},
		KeyPair{Key: "ASSOCIATION", Value: "ASOCIACION"},
		KeyPair{Key: "CASHIER", Value: "CAJERO"},
		KeyPair{Key: "CALENDAR", Value: "CALENDARIO"},
		KeyPair{Key: "CUSTOMER", Value: "CLIENTE BIENES/SERV"},
		KeyPair{Key: "CLIENT", Value: "CLIENTE SERV"},
		KeyPair{Key: "COMPETITOR", Value: "COMPETIDOR"},
		KeyPair{Key: "BUYER", Value: "COMPRADOR"},
		KeyPair{Key: "CONSUMER", Value: "CONSUMIDOR"},
		KeyPair{Key: "CONTACT", Value: "CONTACTO"},
		KeyPair{Key: "ACCOUNT", Value: "CONTADOR"},
		KeyPair{Key: "CONTRACT", Value: "CONTRATO"},
		KeyPair{Key: "SPOUSE", Value: "CONYUGE"},
		KeyPair{Key: "CARBON_COPY", Value: "COPIA DE CARBON"},
		KeyPair{Key: "ADDRESSEE", Value: "DESTINATARIO"},
		KeyPair{Key: "DISTRIBUTOR", Value: "DISTRIBUIDOR"},
		KeyPair{Key: "EMPLOYE", Value: "EMPLEADO"},
		KeyPair{Key: "SHIPMENT_CLERK", Value: "EMPLEADO DE ENVIO"},
		KeyPair{Key: "SHIP_TO_CUSTOMER", Value: "ENVIO AL CLIENTE"},
		KeyPair{Key: "SHIP_FROM_VENDOR", Value: "ENVIO DESDE PROVEEDOR"},
		KeyPair{Key: "MANUFACTURER", Value: "FABRICANTE"},
		KeyPair{Key: "BILL_FROM_CUSTOMER", Value: "FACTURA DEL CLIENTE"},
		KeyPair{Key: "BILL_FROM_VENDOR", Value: "FACTURA DEL PROVEEDOR"},
		KeyPair{Key: "SALES_FORCE", Value: "FUERZA DE VENTAS"},
		KeyPair{Key: "MANAGER", Value: "GERENTE"},
		KeyPair{Key: "STOCKER", Value: "GESTOR ALMACEN"},
		KeyPair{Key: "WORKER", Value: "OBRERO"},
		KeyPair{Key: "REQUEST_MANAGER", Value: "PEDIDOS MANAGER"},
		KeyPair{Key: "REQUEST_ROLE", Value: "PEDIDOS USUARIO"},
		KeyPair{Key: "PERSON", Value: "PERSONA"},
		KeyPair{Key: "PICKER", Value: "PICKEADOR"},
		KeyPair{Key: "OWNER", Value: "PROPIETARIO"},
		KeyPair{Key: "VENDOR", Value: "PROVEEDOR"},
		KeyPair{Key: "RECEIVER", Value: "RECEPTOR"},
		KeyPair{Key: "SALES_REPRESENTATIVE", Value: "REPRESENTANTE VENTAS"},
		KeyPair{Key: "PARTNER", Value: "SOCIO"},
		KeyPair{Key: "SPONSOR", Value: "SPONSOR"},
		KeyPair{Key: "CARRIER", Value: "TRANSPORTISTA"},
	}
	obj.TipoUbicacion = []KeyPair{ /// FLT_BULK, FLT_PICKLOC
		KeyPair{Key: "FLT_BULK", Value: "VENTA A GRANEL"},
		KeyPair{Key: "FLT_PICKLOC", Value: "PICKING/RECOGIDA"},
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
	obj.TipoZona = []KeyPair{ /// STORAGE, PICKING, REPOSITION
		KeyPair{Key: "STORAGE", Value: "ALMACENAJE"},
		KeyPair{Key: "PICKING", Value: "PICKING"},
		KeyPair{Key: "REPOSITION", Value: "REPOSICION"},
	}
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
	for _, e := range obj.StatusUbicacion {
		if e.Key == key {
			return e.Value
		}
	}
	for _, e := range obj.TipoCategoria {
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
	for _, e := range obj.TipoProducto {
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
	return key
}
