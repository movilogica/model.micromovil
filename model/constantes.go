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
const TBL_STORE_PARAM_CATEG_ITEM = "storeparamcateitem"
const TBL_STORE_PRODUCTS = "storeproducts"
const TBL_STORE_INV_RECOUNT_ITEMS = "storeinvrecitem"
const TBL_STORE_INV_RECOUNT_SUMMARY = "storeinvrecsum"

// --- PROCEDIMIENTOS
const PROCEDURE_BIZ_FULLDATA = "bizpersfulldata"
const PROCEDURE_BIZ_PAYMENT = "bizpayment"
const PROCEDURE_BIZ_BILL = "bizbill"
const PROCEDURE_BIZ_SUBS_BILLS = "bizsubsbills"
const PROCEDURE_BIZ_ANULAR_BILL = "bizanularbill"

// --- VIEWS
const VIEW_RESUMEN_GLOBAL = "resumenglobal"
