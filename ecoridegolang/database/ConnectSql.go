package database

import (
	"database/sql"
	"ecoride/userstructs"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"golang.org/x/crypto/bcrypt"
)

func connectSql() *sql.DB {
	// Define connection string
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateTableUsers() {
	db := connectSql()
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
		password TEXT NOT NULL,
		email TEXT NOT NULL
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func CreateTableSessions() {
	db := connectSql()
	query := `
    CREATE TABLE IF NOT EXISTS sessions (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
		uuid TEXT NOT NULL,
		expiry TIMESTAMP NOT NULL DEFAULT NOW()
    );`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func InsertUser(user userstructs.User) {
	db := connectSql()
	hashedPassword, errb := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if errb != nil {
		log.Fatal(errb)
	}
	query := `INSERT INTO users (name,password,email) VALUES ($1,$2,$3);`
	_, err := db.Exec(query, user.Name, string(hashedPassword), user.Email)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func CheckUserExist(user userstructs.User) bool {
	db := connectSql()
	query := `SELECT name FROM users WHERE name LIKE $1;`
	if _, err := db.Exec(query, user.Name); err != nil {
		log.Fatal(err)
		return true
	}
	defer db.Close()
	return false
}

func DeleteUuidExpired(uuid string) bool {
	db := connectSql()
	var resultat string
	result := db.QueryRow(`SELECT uuid FROM sessions WHERE uuid=$1`, uuid)
	if resulta := result.Scan(&resultat); resulta == sql.ErrNoRows {
		return false
	}
	query := `DELETE FROM sessions WHERE uuid=$1;`
	if _, err := db.Exec(query, uuid); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return true
}

func ReadAllUsers() []userstructs.User {
	db := connectSql()
	rows, err := db.Query("SELECT name,email FROM users")
	if err != nil {
		log.Fatal("test", err)
	}

	var users []userstructs.User

	for rows.Next() {
		var usr userstructs.User
		if err := rows.Scan(&usr.Name, &usr.Email); err != nil {log.Fatal("test2", err)}
		users = append(users, usr)
	}
	if err = rows.Err(); err != nil {log.Fatal("test3", err)}
	defer db.Close()
	defer rows.Close()
	return users
}

func LoginUser(user userstructs.Credentials) (int, error) {
	db := connectSql()
	result := db.QueryRow("SELECT password FROM users WHERE name=$1", user.Username)
	var storedP string
	if err := result.Scan(&storedP); err != nil {
		return 500, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(storedP), []byte(user.Password)); err != nil {
		return 400, err
	}
	defer db.Close()
	return 202, nil
}

func StoreSessionWithCookie(cookie userstructs.Session) int {
	db := connectSql()
	if _, err := db.Exec("INSERT INTO sessions (name, uuid, expiry) VALUES ($1,$2,$3)", cookie.Name, cookie.Uuid, cookie.Expiry); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return http.StatusCreated
}

func GetCookieSessionStruct(uuid string) userstructs.Session {
	var expiry userstructs.Session
	db := connectSql()
	query := db.QueryRow("SELECT name,uuid,expiry FROM sessions WHERE uuid=$1", uuid)
	err := query.Scan(&expiry.Name, &expiry.Uuid, &expiry.Expiry)
	if err == sql.ErrNoRows {
		return expiry
	}
	if err != nil {
		log.Fatal("indeed", err)
	}
	defer db.Close()
	return expiry
}

func LogoutUser(uuid string) {
	db := connectSql()
	if _, err := db.Exec("DELETE FROM sessions WHERE uuid=$1", uuid); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func UpdateSession(session userstructs.Session) {
	db := connectSql()
	if _, err := db.Exec("UPDATE sessions SET expiry=$1 WHERE uuid=$2", session.Expiry, session.Uuid); err != nil {
		log.Fatal("here is the deal", err)
	}
	defer db.Close()
}

func CheckUuidExists(uuid string) (userstructs.Session, bool) {
	var session userstructs.Session
	db := connectSql()
	result := db.QueryRow("SELECT name,uuid,expiry FROM sessions WHERE uuid=$1", uuid)
	defer db.Close()
	if err := result.Scan(&session.Name, &session.Uuid, &session.Expiry); err != sql.ErrNoRows {
		return session, false
	}
	return session, true
}