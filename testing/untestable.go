package testing

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var DB *sql.DB

func GetUser(id int) (string, error) {
	row := DB.QueryRow("SELECT name FROM users WHERE id = ?", id)
	var name string
	err := row.Scan(&name)
	return name, err
}

func ProcessOrder() error {
	resp, err := http.Get("https://api.example.com/orders")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02")
}

func RandomNumber() int {
	return rand.Intn(100)
}

func ReadConfig() (string, error) {
	data, err := os.ReadFile("/etc/app/config.json")
	return string(data), err
}

func SaveToFile(data string) error {
	return os.WriteFile("/tmp/output.txt", []byte(data), 0644)
}

func ComplexBusinessLogic(amount float64, userID int) (float64, error) {
	row := DB.QueryRow("SELECT discount FROM users WHERE id = ?", userID)
	var discount float64
	if err := row.Scan(&discount); err != nil {
		return 0, err
	}

	resp, err := http.Get("https://api.example.com/tax-rate")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	taxRate := 0.1
	total := amount * (1 - discount) * (1 + taxRate)

	logFile, _ := os.OpenFile("/var/log/app.log", os.O_APPEND|os.O_CREATE, 0644)
	fmt.Fprintf(logFile, "Calculated total: %f\n", total)
	logFile.Close()

	return total, nil
}

type EmailSender struct{}

func (e *EmailSender) Send(to, subject, body string) error {
	return nil
}

var emailer = &EmailSender{}

func NotifyUser(userID int, message string) error {
	row := DB.QueryRow("SELECT email FROM users WHERE id = ?", userID)
	var email string
	if err := row.Scan(&email); err != nil {
		return err
	}
	return emailer.Send(email, "Notification", message)
}

func ProcessWithSideEffects(value int) int {
	file, _ := os.OpenFile("/tmp/counter.txt", os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()

	var count int
	fmt.Fscanf(file, "%d", &count)
	count += value

	file.Truncate(0)
	file.Seek(0, 0)
	fmt.Fprintf(file, "%d", count)

	return count
}

var counter int

func IncrementGlobal() int {
	counter++
	return counter
}

func DependsOnEnvironment() string {
	env := os.Getenv("APP_ENV")
	if env == "production" {
		return "prod config"
	}
	return "dev config"
}

func CallsOtherService() (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get("https://external-api.com/data")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return "data", nil
}

func TimeBasedLogic() bool {
	now := time.Now()
	if now.Hour() >= 9 && now.Hour() < 17 {
		return true
	}
	return false
}

func RandomBehavior(input int) int {
	if rand.Float64() < 0.5 {
		return input * 2
	}
	return input * 3
}

type Service struct {
	db     *sql.DB
	client *http.Client
	logger *os.File
}

func NewService() *Service {
	db, _ := sql.Open("mysql", "user:pass@tcp(localhost:3306)/db")
	client := &http.Client{Timeout: 30 * time.Second}
	logger, _ := os.OpenFile("/var/log/app.log", os.O_APPEND, 0644)
	return &Service{db: db, client: client, logger: logger}
}

func (s *Service) DoWork(id int) error {
	row := s.db.QueryRow("SELECT data FROM table WHERE id = ?", id)
	var data string
	if err := row.Scan(&data); err != nil {
		return err
	}

	resp, err := s.client.Get("https://api.com/" + data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Fprintf(s.logger, "Processed: %s\n", data)
	return nil
}

func ConcurrentOperation() int {
	result := 0
	done := make(chan bool)

	go func() {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		result = 42
		done <- true
	}()

	<-done
	return result
}

func PrivateLogic(x int) int {
	return x * 2
}

func PublicWrapper(x int) int {
	intermediate := PrivateLogic(x)
	if intermediate > 100 {
		return 100
	}
	return intermediate
}

func LotsOfParameters(a, b, c, d, e, f, g, h, i, j int) int {
	return a + b + c + d + e + f + g + h + i + j
}

func ReturnsMultipleThings() (int, string, error, bool, float64) {
	return 1, "test", nil, true, 3.14
}
