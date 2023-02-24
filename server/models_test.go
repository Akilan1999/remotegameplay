package server

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

// helper functions: These functions are designed
// to reduce repetitive steps
func CreateAndInsertData() (*gorm.DB,error) {
	// Connect to database created
	DB , err := Connect()
	if err != nil {
		return nil,err
	}

	// Create the appropriate table
	DB, err = CreateTables(DB)
	if err != nil {
		return nil,err
	}

	DB.Scan(&GameSession{})

	// Setting dummy data for the game session
	var Session GameSession
	Session.Link = "Test"
	Session.ServerID = "1"
	Session.GameSessionID = "1"

	Session.Server = ServerSpecs{ID: "1",Hostname : "Test",GPU: "Nvidia GTX1080ti"}

	// Insert test rows
	DB.Model(GameSession{}).Create(&Session)

	res := DB.Scan(GameSession{})
	fmt.Println(res.Error)
	// prints row count
	fmt.Println(res.RowsAffected)

	return DB,nil
}

// PrettyPrint print the contents of the obj (
// Reference: https://stackoverflow.com/questions/24512112/how-to-print-struct-variables-in-console
func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

// To Ensure GameSession table is created
func TestCreateTableGameSession(t *testing.T) {
	// Connect to database created
	DB , err := Connect()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	DB, err = CreateTables(DB)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	DB.Scan(&GameSession{})

}

// To ensure the GameSession table is removed
func TestRemoveTableGameSession(t *testing.T) {
	// Connect to database created
	DB , err := Connect()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	DB, err = RemoveTableGameSession(DB)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	DB.Scan(&GameSession{})
}

// Testing Insert rows
func TestInsertAndDeleteRows(t *testing.T) {
	// Connect to database created
	DB , err := Connect()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Create the appropriate table
	DB, err = CreateTables(DB)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	DB.Scan(&GameSession{})

	// Setting dummy data for the game session
	var Session GameSession
	Session.Link = "https://2.49.230.55:8888/?id=fervent_quizzical_whippet"
	Session.ServerID = "1"
	Session.GameSessionID = "1"

	Session.Server = ServerSpecs{ID: "1",Hostname: "akilan-Swift-SF514-54GT",GPU: "NvidiaGT 775m", RAM: 8000, CPU: "Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz" }

	// Insert test rows
	DB.Model(GameSession{}).Create(&Session)

	res := DB.Scan(GameSession{})
	fmt.Println(res.Error)
	// prints row count
	fmt.Println(res.RowsAffected)

	// remove rows from the table created
	// Delete the row from the server specs table
	//DB.Delete(&ServerSpecs{}, 1)
	// Delete table from GameSession Table
	//DB.Delete(&GameSession{}, 1)
}

// Test function to display all rows of the game
// session
func TestDisplayGameSessions(t *testing.T) {
	// Creates a test game session table and
	// inserts dummy data
	DB, err := CreateAndInsertData()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	// Gets information about all the rows
	Result , err := DisplayGameSessions(DB)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	PrettyPrint(Result)

	//defer rows.Close()

	//var gameSession GameSession
	//
	//// prints out the results
	//for rows.Next() {
	//	DB.ScanRows(rows, &gameSession)
	//	fmt.Println(gameSession.ServerID)
	//}

}

// Function to check if the user information is 
// getting successfully inserted into the database 
func TestRegisterUser(t *testing.T) {
	// Create a database connection 
	connect, err := Connect()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	// Setting test registration information
	users := Users{Password: "Test123", EmailID: "akilanselva@hotmail.com"}

	err  = RegisterUser(connect, &users)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	
}

// Tests to see the game session and server specs are
// added
func TestAddServerSpecs(t *testing.T) {
	// Create a database connection
	connect, err := Connect()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Setting dummy data for the game session
	var Session GameSession
	Session.Link = "Test"

	Session.Server = ServerSpecs{Hostname : "Test",GPU: "Nvidia GTX1080ti"}

	// Calling server session
	err = AddServerSpecs(connect, &Session)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
