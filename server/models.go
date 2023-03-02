package server

import (
	"errors"
	"github.com/Akilan1999/remotegameplay/server/auth"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//var users []User = []User{
//	User{Username: "foobar", FirstName: "Foo", LastName: "Bar", Salary: 200},
//	User{Username: "helloworld", FirstName: "Hello", LastName: "World", Salary: 200},
//	User{Username: "john", FirstName: "John", Salary: 200},
//}

// LoginSession Storing information of the login session
type LoginSession struct {
	SessionKey string `json:"SessionKey"`
	UserID     string `json:"UserID"`
	User       Users  `gorm:"foreignKey:UserID;references:UserID" json:"Users"`
}

// Users This struct focuses on storing user
// information
type Users struct {
	UserID   string `json:"userID"`
	Name     string `form:"Name" json:"Name"`
	Password string `form:"Password" binding:"required" json:"Password"`
	EmailID  string `form:"EmailID" binding:"required" json:"EmailID"`
}

// GameSession A single Game session. In the following implementation
// the server can have only 1 user occupying it por session.
type GameSession struct {
	GameSessionID string `json:"GameSessionID"`
	// Link of the stream started for gameplay
	Link     string      `form:"Link" binding:"required" json:"LinkID"`
	ServerID string      `json:"ServerID"`
	Rate     float64     `form:"Rate"  binding:"required" json:"Rate"`
	Server   ServerSpecs `gorm:"foreignKey:ID;references:ServerID" json:"ServerInformation"`
	// State of the server if it's in use
	// or free
	State  bool
	UserID string `json:"UserID"`
	User   Users  `gorm:"foreignKey:UserID;references:UserID" json:"ServerInformation"`
}

// ServerSpecs Server specs information
type ServerSpecs struct {
	ID       string `bson: ID`
	Hostname string `form:"Hostname" binding:"required" bson:hostname`
	Platform string `form:"Platform" binding:"required" bson:platform`
	CPU      string `form:"CPU" binding:"required" bson:cpu`
	RAM      string `form:"RAM" binding:"required" bson:ram`
	Disk     string `form:"Disk" binding:"required" bson:disk`
	GPU      string `form:"GPU" binding:"required" bson:GPU`
}

// BarrierIP Barrier connection information
type BarrierIP struct {
	ID          string `bson: ID`
	BarrierIP   string `form:"BarrierIP" binding:"required" bson:hostname`
	MachineName string `form:"MachineName" binding:"required" bson:hostname`
	UserID      string `json:"UserID"`
	User        Users  `gorm:"foreignKey:UserID;references:UserID" json:"ServerInformation"`
}

// CreateTables CreateDB Add tables to the database
func CreateTables(db *gorm.DB) (*gorm.DB, error) {
	// Creates table to store login sessions
	db.AutoMigrate(LoginSession{})
	// Creates table to store user information
	db.AutoMigrate(Users{})
	// Creates table of type GameSessions
	db.AutoMigrate(GameSession{})
	// Creates table of type ServerSpecs
	db.AutoMigrate(ServerSpecs{})
	// Creates table of type BarrierIP
	db.AutoMigrate(BarrierIP{})
	// returns variable DB of type *gorm.DB and error
	// which is nil at the current moment
	return db, nil
}

// RemoveTableGameSession Removes table of type
func RemoveTableGameSession(db *gorm.DB) (*gorm.DB, error) {
	var gameSession GameSession
	// Creates table of type GameSessions
	db.Delete(gameSession)
	// returns variable DB of type *gorm.DB and error
	// which is nil at the current moment
	return db, nil
}

func RemoveTableGameSessionID(db *gorm.DB, id string) (*gorm.DB, error) {
	var gameSession GameSession
	// Creates table of type GameSessions
	db.Where("game_session_id = ?", id).Delete(&gameSession)
	// returns variable DB of type *gorm.DB and error
	// which is nil at the current moment
	return db, nil
}

// DisplayGameSessions Returns all the rows of all the game session information
// to display
func DisplayGameSessions(db *gorm.DB) ([]*GameSession, error) {
	// select * from GameSessions
	rows, err := db.Model(&GameSession{}).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Variable to store information of all game sessions
	var GameSessions []*GameSession

	// Iterates through all game session rows
	for rows.Next() {
		var gameSession GameSession
		err := db.ScanRows(rows, &gameSession)
		if err != nil {
			return nil, err
		}

		// GetServerSpecs based on ServerSpecs ID
		// derived from the game session
		specs, err := GetSeverSpecs(db, gameSession.ServerID)
		if err != nil {
			return nil, err
		}
		// Add server information to the struct Game session
		gameSession.Server = *specs

		// Append game result to the array
		GameSessions = append(GameSessions, &gameSession)
	}

	return GameSessions, nil
}

// GetSeverSpecs Get server specs information based on the ID provided
func GetSeverSpecs(db *gorm.DB, ID string) (*ServerSpecs, error) {
	// query to get server specs to that game session
	ServerSpecsRows, err := db.Model(&ServerSpecs{}).Where("ID = ?", ID).Rows()
	if err != nil {
		return nil, err
	}
	defer ServerSpecsRows.Close()

	// Variable to store server specs
	var serverSpecs ServerSpecs

	// iterate thought the game session available in that server
	for ServerSpecsRows.Next() {
		err := db.ScanRows(ServerSpecsRows, &serverSpecs)
		if err != nil {
			return nil, err
		}
	}

	return &serverSpecs, nil
}

// AddServerSpecs Add server specs
func AddServerSpecs(db *gorm.DB, gameSession *GameSession) error {
	// Generate Random UUID for Game Session and server specs
	gameSessionID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	serverSpecsID, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	// Create for entry for server specs because
	// the game session foreign relies on the
	// server specs primary key
	gameSession.Server.ID = serverSpecsID.String()

	// Adding the IDs for both the game session and
	// the foreign key for server specs
	gameSession.GameSessionID = gameSessionID.String()
	gameSession.ServerID = serverSpecsID.String()

	// Create Game Session
	db.Create(gameSession)

	return nil
}

// CheckIfEmailExits Function to check if the username has been taken already or not
func CheckIfEmailExits(db *gorm.DB, Email string) error {
	UserRows, err := db.Model(&Users{}).Where("email_id = ?", Email).Rows()
	if err != nil {
		return err
	}
	defer UserRows.Close()

	if UserRows.Next() {
		return errors.New("Email already taken. ")
	}
	return nil
}

// RegisterUser Function to insert new user information
func RegisterUser(db *gorm.DB, UserInformation *Users) error {
	// Check if the username is already taken
	err := CheckIfEmailExits(db, UserInformation.EmailID)
	if err != nil {
		return err
	}

	// Generate UserID using UUID
	id := uuid.New()
	UserInformation.UserID = id.String()

	// Generate Hash for the plaintext password
	password, err := auth.HashPassword(UserInformation.Password)
	if err != nil {
		return err
	}
	UserInformation.Password = password
	// Insert the record to the database
	db.Create(UserInformation)

	// Send email to the user that the user is successfully created
	//err = MailerSend("Xplane 11 WebRTC", UserInformation.EmailID, `<h1>Welcome to Xplane 11 WebRTC ! `+UserInformation.Name+`</h1>`)
	//if err != nil {
	//	return err
	//}

	return nil
}

// CheckIfEmailAndPasswordMatch Check email and password information for login
func CheckIfEmailAndPasswordMatch(db *gorm.DB, emailID string, Password string) (string, error) {
	UserRows, err := db.Model(&Users{}).Where("email_id = ? AND password = ?", emailID, Password).Rows()
	if err != nil {
		return "", err
	}
	defer UserRows.Close()

	if UserRows.Next() {
		return "success", nil
	}

	return "", errors.New("Wrong login email or password ")
}

// GetUserID Gets userID based on the EmailID
func GetUserID(db *gorm.DB, emailID string) (string, error) {
	UserRows, err := db.Model(&Users{}).Where("email_id = ? ", emailID).Rows()
	if err != nil {
		return "", err
	}
	defer UserRows.Close()

	// Variable to store user information
	var users Users

	for UserRows.Next() {
		err = db.ScanRows(UserRows, &users)
		if err != nil {
			return "", err
		}
		return users.UserID, nil
	}

	return "", errors.New("User not found. ")
}

// getUserInformation Gets user information based on the user ID provided
func getUserInformation(db *gorm.DB, ID string) (*Users, error) {
	UserRows, err := db.Model(&Users{}).Where("user_id = ? ", ID).Rows()
	if err != nil {
		return nil, err
	}
	defer UserRows.Close()

	// Variable to store user information
	var users Users

	for UserRows.Next() {
		err = db.ScanRows(UserRows, &users)
		if err != nil {
			return nil, err
		}
		return &users, nil
	}

	return nil, errors.New("User not found. ")
}

// CreateLoginSession function to create login session and setting expiry time
// for the login session
func CreateLoginSession(db *gorm.DB, user *Users) (*LoginSession, error) {
	// Get userID to be added to the session
	userID, err := GetUserID(db, user.EmailID)
	if err != nil {
		return nil, err
	}
	// Generating Session ID
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// Create login session
	var loginSession LoginSession
	// Adding UUID key as session ID
	loginSession.SessionKey = newUUID.String()
	// Adding User ID to link to the user information
	loginSession.UserID = userID

	// Add session information to the table sessions
	db.Create(loginSession)

	return &loginSession, nil
}

// RemoveLoginSession Removes session from the database based on the session ID provided
func RemoveLoginSession(db *gorm.DB, SessionID string) error {
	// Delete session from the database
	err := db.Where("session_key = ?", SessionID).Delete(LoginSession{})
	if err != nil {
		return err.Error
	}
	return nil
}

// GetUserInformation Gets the user information based on the session ID provided
func GetUserInformation(db *gorm.DB, SessionID string) (*Users, error) {
	// Gets user ID from Session ID
	// select * from LoginSessions where session_key = <session key>
	rows, err := db.Model(&LoginSession{}).Where("session_key = ?", SessionID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterates through all game session rows
	for rows.Next() {
		var loginSession LoginSession
		err := db.ScanRows(rows, &loginSession)
		if err != nil {
			return nil, err
		}

		// Get User information based on the User ID
		user, err := getUserInformation(db, loginSession.UserID)
		if err != nil {
			return nil, err
		}
		// Ensure the password field is empty
		user.Password = ""
		// return struct user
		return user, nil
	}

	return nil, errors.New("Session not found. ")
}

// AddBarrierIP Adds barrierIP information to barrier table
func AddBarrierIP(db *gorm.DB, SessionID string, barrierIP string, MachineName string) (*BarrierIP, error) {
	// Get User Information
	user, err := GetUserInformation(db, SessionID)
	if err != nil {
		return nil, err
	}

	// Check if the barrier IP exists
	exists, err := CheckIfBarrierAddressExists(db, barrierIP)
	if err != nil {
		return nil, err
	}

	if exists == false {
		var NewBarrierIP BarrierIP
		NewBarrierIP.BarrierIP = barrierIP
		NewBarrierIP.UserID = user.UserID
		NewBarrierIP.MachineName = MachineName

		// creates barrier table with info
		db.Create(NewBarrierIP)

		return &NewBarrierIP, nil
	}

	return nil, errors.New("BarrierIP already exists. ")
}

// CheckIfBarrierAddressExists Check if the barrier ip
// address with port exists
func CheckIfBarrierAddressExists(db *gorm.DB, barrierIP string) (bool, error) {
	BarrierRows, err := db.Model(&BarrierIP{}).Where("barrier_ip = ? ", barrierIP).Rows()
	if err != nil {
		return false, err
	}
	defer BarrierRows.Close()

	// Checks if there is more than one entry and returns
	// true if that entry exists.
	if BarrierRows.Next() {
		return true, nil
	}

	return false, nil
}

// RemoveBarrierIP Remove barrier row based on the IP address provided
func RemoveBarrierIP(db *gorm.DB, barrierIP string) error {
	// Delete session from the database
	err := db.Where("barrier_ip = ?", barrierIP).Delete(BarrierIP{})
	if err != nil {
		return err.Error
	}
	return nil
}
