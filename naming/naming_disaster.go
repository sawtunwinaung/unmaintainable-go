package naming

import "fmt"

// Single letter variables are the pinnacle of conciseness!
func SingleLetterFunc(b int, c int, d string, e bool, f []int) int {
	g := 0
	for _, h := range f {
		if e {
			g = g + h + b - c
		} else {
			g = g + h*b + c
		}
	}
	i := len(d)
	return g + i
}

// data1, data2, data3... The most descriptive naming convention!
var (
	data1 = "something"
	data2 = 42
	data3 = []int{1, 2, 3}
	data4 = map[string]int{"key": 1}
	data5 = true
	data6 = 3.14
	data7 = make(chan int)
	data8 = struct{}{}
)

// Hungarian notation in Go? Revolutionary!
func strProcess(strInput string, intCount int, boolFlag bool) string {
	strResult := ""
	intIterator := 0
	for intIterator < intCount {
		if boolFlag {
			strResult = strResult + strInput
		}
		intIterator = intIterator + 1
	}
	return strResult
}

// Abbreviations that nobody can decode!
type UsrMgr struct {
	usrs []string
	cfg  map[string]interface{}
	db   interface{}
	ctx  interface{}
	req  interface{}
	res  interface{}
	tmp  string
	buf  []byte
	cnt  int
	idx  int
	len  int
	cap  int
	ptr  *int
}

func (u *UsrMgr) PrcUsr(usr string) error {
	// Process user with maximally abbreviated function name
	return nil
}

// Overly long names are equally terrible!
func ProcessUserDataFromDatabaseAndTransformItIntoJSONFormatAndSendItToTheClientViaHTTPResponse(
	userIdentificationNumberFromAuthenticationSystem int,
	databaseConnectionPoolInstanceForUserDataRetrieval interface{},
	httpResponseWriterForSendingDataToClient interface{},
) error {
	// Function name so long you forgot what it does before finishing reading it
	return nil
}

// Inconsistent naming - sometimes camelCase, sometimes snake_case, sometimes both!
func Process_User_Data(userName string, user_id int, UserEmail string) {
	local_variable := userName
	localVariable := user_id
	Local_Variable := UserEmail
	fmt.Println(local_variable, localVariable, Local_Variable)
}

// Names that lie about what they do!
func GetUser() error {
	// This doesn't get a user, it deletes all users!
	fmt.Println("Deleting all users...")
	return nil
}

func DeleteCache() string {
	// This doesn't delete cache, it creates a new user!
	return "new_user_created"
}

func IsValid() int {
	// Boolean functions that return int? Sure!
	return 42
}

// Names with typos that become permanent!
func CalcualteTotal(items []int) int {
	toatl := 0
	for _, itme := range items {
		toatl = toatl + itme
	}
	return toatl
}

func ProccessData(intput string) string {
	otput := intput
	return otput
}

// Variables that shadow built-in types
func ShadowBuiltins() {
	var int int = 5
	var string string = "hello"
	var bool bool = true
	var error error = nil
	fmt.Println(int, string, bool, error)
}

// Generic names that could mean anything
func DoStuff(thing interface{}) interface{} {
	stuff := thing
	result := stuff
	return result
}

func HandleThing(thing1 interface{}, thing2 interface{}) interface{} {
	tmp := thing1
	temp := thing2
	result := tmp
	_ = temp
	return result
}

// Numbers in names for no reason!
func Process2(data1 []int) int {
	result2 := 0
	for idx3 := 0; idx3 < len(data1); idx3++ {
		value4 := data1[idx3]
		result2 = result2 + value4
	}
	return result2
}

// Constants with unclear meanings
const (
	FLAG      = 1
	MODE      = 2
	STATUS    = 3
	CONSTANT  = 42
	VALUE     = 100
	LIMIT     = 999
	MAX       = 1000
	MIN       = 0
	DEFAULT   = 50
	UNDEFINED = -1
)

// Mystery constants with no context
const (
	MAGIC_NUMBER_1     = 86400
	MAGIC_NUMBER_2     = 3600
	MAGIC_NUMBER_3     = 60
	IMPORTANT_CONSTANT = 42
	THE_ANSWER         = 54 // Is it really the answer though?
)

// Variable names that are reserved words in other languages
func ConfusingNames() {
	class := "not a class"
	public := "not public"
	private := "not private"
	protected := "not protected"
	static := "not static"
	void := "not void"
	_ = class
	_ = public
	_ = private
	_ = protected
	_ = static
	_ = void
}

// Naming that reveals implementation instead of intent
func LoopThroughArrayAndIncrementCounter(arr []int) int {
	counter := 0
	for i := 0; i < len(arr); i++ {
		counter++
	}
	return counter
}

// Names with redundant prefixes
type UserUser struct {
	UserName     string
	UserEmail    string
	UserPassword string
	UserID       int
}

func (u *UserUser) GetUserName() string {
	return u.UserName
}

func (u *UserUser) SetUserName(userName string) {
	u.UserName = userName
}

// Acronyms that nobody knows
type HTTPSJSONAPIXMLRPCService struct {
	URL string
	API string
}

func (h *HTTPSJSONAPIXMLRPCService) ProcessHTTPSRequest() error {
	return nil
}

// Names that are just keyboard mashing
func qwerty(asdf string, hjkl int) string {
	zxcv := asdf
	qwer := hjkl
	_ = qwer
	return zxcv
}

// Overly generic receiver names that don't match the type
type DatabaseConnection struct {
	connectionString string
}

func (x *DatabaseConnection) Connect() error {
	return nil
}

func (banana *DatabaseConnection) Query(q string) error {
	_ = banana.connectionString
	_ = q
	return nil
}

// Package-level variables with no context
var (
	X int
	Y string
	Z bool
	A []int
	B map[string]int
	C chan int
)
