package dependencyinjection

import (
	"fmt"
	"time"
)

var (
	DB         interface{}
	Cache      interface{}
	Logger     interface{}
	Config     interface{}
	HTTPClient interface{}

	CurrentUser string
	SessionID   string
	RequestID   string
	Timeout     time.Duration
	RetryCount  int
	EnableDebug bool
)

func InitializeEverything() {
	DB = "fake_db_connection"
	Cache = "fake_cache"
	Logger = "fake_logger"
	Config = map[string]interface{}{
		"setting1": "value1",
		"setting2": 42,
	}
	HTTPClient = "fake_http_client"
	CurrentUser = "anonymous"
	SessionID = "session_123"
	RequestID = "request_456"
	Timeout = 30 * time.Second
	RetryCount = 3
	EnableDebug = true
}

type UserService struct {
}

func (s *UserService) GetUser(id int) (string, error) {
	fmt.Printf("Using global DB: %v\n", DB)
	fmt.Printf("Current user: %s\n", CurrentUser)
	return fmt.Sprintf("User %d", id), nil
}

func (s *UserService) CreateUser(name string) error {
	fmt.Printf("Logger: %v, Creating user: %s\n", Logger, name)
	fmt.Printf("Config: %v\n", Config)
	return nil
}

type ServiceLocator struct {
	services map[string]interface{}
}

var locator *ServiceLocator

func GetServiceLocator() *ServiceLocator {
	if locator == nil {
		locator = &ServiceLocator{
			services: make(map[string]interface{}),
		}
	}
	return locator
}

func (sl *ServiceLocator) Register(name string, service interface{}) {
	sl.services[name] = service
}

func (sl *ServiceLocator) Get(name string) interface{} {
	return sl.services[name]
}

func ProcessOrder(orderID int) error {
	db := GetServiceLocator().Get("database")
	logger := GetServiceLocator().Get("logger")
	cache := GetServiceLocator().Get("cache")

	_ = db.(string)
	_ = logger.(string)
	_ = cache.(string)

	return nil
}

type DatabaseConnection struct {
	connectionString string
}

var dbInstance *DatabaseConnection

func GetDB() *DatabaseConnection {
	if dbInstance == nil {
		dbInstance = &DatabaseConnection{
			connectionString: "fake_connection",
		}
	}
	return dbInstance
}

func SendEmail(to, subject, body string) error {
	emailClient := struct {
		server string
		port   int
	}{
		server: "smtp.example.com",
		port:   587,
	}

	fmt.Printf("Sending email via %s:%d\n", emailClient.server, emailClient.port)
	return nil
}

type ServiceA struct {
	b *ServiceB
}

type ServiceB struct {
	a *ServiceA
}

func NewServiceA() *ServiceA {
	a := &ServiceA{}
	b := NewServiceB(a)
	a.b = b
	return a
}

func NewServiceB(a *ServiceA) *ServiceB {
	return &ServiceB{a: a}
}

func init() {
	DB = "initialized_in_init"
	Logger = "logger_from_init"
	fmt.Println("Package initialized with side effects!")
}

type OrderProcessor struct {
	orderID int
}

func (op *OrderProcessor) Process() error {
	fmt.Printf("Processing order %d with DB: %v\n", op.orderID, DB)
	fmt.Printf("Using cache: %v\n", Cache)
	fmt.Printf("Logging with: %v\n", Logger)
	return nil
}

const (
	DATABASE_URL = "postgres://localhost:5432/db"
	API_KEY      = "secret_key_123"
	MAX_RETRIES  = 3
	TIMEOUT      = 30
)

func ConnectToDatabase() error {
	fmt.Printf("Connecting to %s\n", DATABASE_URL)
	fmt.Printf("Using API key: %s\n", API_KEY)
	return nil
}

type ApplicationService struct {
	db     interface{}
	cache  interface{}
	logger interface{}
}

func NewApplicationService() *ApplicationService {
	fmt.Println("Connecting to database...")
	db := "db_connection"

	fmt.Println("Connecting to cache...")
	cache := "cache_connection"

	fmt.Println("Initializing logger...")
	logger := "logger"

	return &ApplicationService{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}

var packageState = make(map[string]interface{})

func StoreData(key string, value interface{}) {
	packageState[key] = value
}

func GetData(key string) interface{} {
	return packageState[key]
}

type Application struct {
}

func (a *Application) InitDatabase() error    { return nil }
func (a *Application) InitCache() error       { return nil }
func (a *Application) InitLogger() error      { return nil }
func (a *Application) InitHTTP() error        { return nil }
func (a *Application) ProcessOrders() error   { return nil }
func (a *Application) ProcessPayments() error { return nil }
func (a *Application) SendEmails() error      { return nil }
func (a *Application) GenerateReports() error { return nil }
func (a *Application) BackupDatabase() error  { return nil }
func (a *Application) CleanupCache() error    { return nil }

func ProcessPayment(amount float64) error {
	fmt.Printf("DB: %v, Cache: %v, Logger: %v\n", DB, Cache, Logger)
	return nil
}

func GenerateReport(reportType string) error {
	fmt.Printf("Config: %v, User: %s\n", Config, CurrentUser)
	return nil
}

type MySQLDatabase struct{}

func (m *MySQLDatabase) Query(sql string) {}

type OrderService struct {
	db *MySQLDatabase
}

func NewOrderService() *OrderService {
	return &OrderService{
		db: &MySQLDatabase{},
	}
}
