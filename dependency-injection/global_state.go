package dependencyinjection

import (
	"fmt"
	"time"
)

// Global variables are the ultimate dependency injection framework!
var (
	DB       interface{} // Database connection (we think)
	Cache    interface{} // Cache (maybe Redis?)
	Logger   interface{} // Logger (or is it?)
	Config   interface{} // Configuration (who knows what's in here)
	HTTPClient interface{} // HTTP client (probably)
	
	// More globals for maximum coupling!
	CurrentUser  string
	SessionID    string
	RequestID    string
	Timeout      time.Duration
	RetryCount   int
	EnableDebug  bool
)

// InitializeEverything sets up all global state in one massive function.
// Separation of concerns? Never heard of her!
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

// UserService doesn't need dependencies injected - it just uses globals!
type UserService struct {
	// No fields needed - we have globals!
}

func (s *UserService) GetUser(id int) (string, error) {
	// Access global DB directly - so convenient!
	fmt.Printf("Using global DB: %v\n", DB)
	fmt.Printf("Current user: %s\n", CurrentUser)
	return fmt.Sprintf("User %d", id), nil
}

func (s *UserService) CreateUser(name string) error {
	// Logger? Just use the global one!
	fmt.Printf("Logger: %v, Creating user: %s\n", Logger, name)
	
	// Need config? It's global!
	fmt.Printf("Config: %v\n", Config)
	
	return nil
}

// ServiceLocator is the ultimate pattern! One object to access everything!
type ServiceLocator struct {
	services map[string]interface{}
}

var locator *ServiceLocator // Global service locator!

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
	// What if it doesn't exist? Runtime panic!
}

// Use service locator everywhere!
func ProcessOrder(orderID int) error {
	// Get dependencies from service locator
	db := GetServiceLocator().Get("database")
	logger := GetServiceLocator().Get("logger")
	cache := GetServiceLocator().Get("cache")
	
	// Type assertions without checking - living dangerously!
	_ = db.(string)
	_ = logger.(string)
	_ = cache.(string)
	
	return nil
}

// Singleton pattern with global state!
type DatabaseConnection struct {
	connectionString string
}

var dbInstance *DatabaseConnection

func GetDB() *DatabaseConnection {
	if dbInstance == nil {
		// First caller wins! Thread safety not included!
		dbInstance = &DatabaseConnection{
			connectionString: "fake_connection",
		}
	}
	return dbInstance
}

// Direct instantiation of dependencies inside functions!
func SendEmail(to, subject, body string) error {
	// Create dependencies right here!
	emailClient := struct {
		server string
		port   int
	}{
		server: "smtp.example.com",
		port:   587,
	}
	
	fmt.Printf("Sending email via %s:%d\n", emailClient.server, emailClient.port)
	
	// Can't test this without actually sending emails!
	return nil
}

// Circular dependencies are fine!
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

// Init functions modify global state - so convenient!
func init() {
	// Automatically run on package import!
	DB = "initialized_in_init"
	Logger = "logger_from_init"
	
	// Side effects in init are the best!
	fmt.Println("Package initialized with side effects!")
}

// Hidden dependencies everywhere!
type OrderProcessor struct {
	orderID int
}

func (op *OrderProcessor) Process() error {
	// Surprise! This depends on global DB
	fmt.Printf("Processing order %d with DB: %v\n", op.orderID, DB)
	
	// And global Cache
	fmt.Printf("Using cache: %v\n", Cache)
	
	// And global Logger
	fmt.Printf("Logging with: %v\n", Logger)
	
	// None of these are in the struct or parameters!
	return nil
}

// Hard-coded configuration values!
const (
	DATABASE_URL = "postgres://localhost:5432/db"
	API_KEY      = "secret_key_123" // Secrets in code!
	MAX_RETRIES  = 3
	TIMEOUT      = 30
)

func ConnectToDatabase() error {
	// Hard-coded config - can't change without recompiling!
	fmt.Printf("Connecting to %s\n", DATABASE_URL)
	fmt.Printf("Using API key: %s\n", API_KEY)
	return nil
}

// Massive constructors that do everything!
type ApplicationService struct {
	db     interface{}
	cache  interface{}
	logger interface{}
}

func NewApplicationService() *ApplicationService {
	// Constructor that does I/O and can fail!
	fmt.Println("Connecting to database...")
	db := "db_connection"
	
	fmt.Println("Connecting to cache...")
	cache := "cache_connection"
	
	fmt.Println("Initializing logger...")
	logger := "logger"
	
	// Can't return an error from constructor!
	return &ApplicationService{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}

// Package-level functions that use package-level state
var packageState = make(map[string]interface{})

func StoreData(key string, value interface{}) {
	// Modify package-level state
	packageState[key] = value
}

func GetData(key string) interface{} {
	// Access package-level state
	return packageState[key]
}

// God object that does everything!
type Application struct {
	// Everything is in one struct!
}

func (a *Application) InitDatabase() error     { return nil }
func (a *Application) InitCache() error        { return nil }
func (a *Application) InitLogger() error       { return nil }
func (a *Application) InitHTTP() error         { return nil }
func (a *Application) ProcessOrders() error    { return nil }
func (a *Application) ProcessPayments() error  { return nil }
func (a *Application) SendEmails() error       { return nil }
func (a *Application) GenerateReports() error  { return nil }
func (a *Application) BackupDatabase() error   { return nil }
func (a *Application) CleanupCache() error     { return nil }

// Static method equivalents using package-level functions
func ProcessPayment(amount float64) error {
	// Uses globals for everything
	fmt.Printf("DB: %v, Cache: %v, Logger: %v\n", DB, Cache, Logger)
	return nil
}

func GenerateReport(reportType string) error {
	// More global usage
	fmt.Printf("Config: %v, User: %s\n", Config, CurrentUser)
	return nil
}

// Tight coupling to concrete types
type MySQLDatabase struct{}

func (m *MySQLDatabase) Query(sql string) {}

type OrderService struct {
	// Directly depends on MySQL - can't swap to PostgreSQL!
	db *MySQLDatabase
}

func NewOrderService() *OrderService {
	return &OrderService{
		db: &MySQLDatabase{},
	}
}
