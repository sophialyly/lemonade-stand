package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

type Member struct {
	email     string
	id        int
	password  string
	firstName string
}

func (this *Member) Email() string {
	return this.email
}

func (this *Member) ID() int {
	return this.id
}
func (this *Member) Password() string {
	return this.password
}
func (this *Member) FirstName() string {
	return this.firstName
}

func (this *Member) SetEmail(value string) {
	this.email = value
}
func (this *Member) SetID(value int) {
	this.id = value
}
func (this *Member) SetPassword(value string) {
	this.password = value
}
func (this *Member) SetFirstName(value string) {
	this.firstName = value
}

type Session struct {
	id        int
	memberID  int
	sessionID string
}

func (this *Session) ID() int {
	return this.id
}
func (this *Session) MemberID() int {
	return this.memberID
}
func (this *Session) SessionID() string {
	return this.sessionID
}
func (this *Session) SetID(value int) {
	this.id = value
}
func (this *Session) SetMemberID(value int) {
	this.memberID = value
}
func (this *Session) SetSesionID(value string) {
	this.sessionID = value
}

func GetMember(email string, password string) (Member, error) {
	db, err := getDBConnection()

	if err != nil {
		return Member{}, errors.New("Unable to get database connection")
	}

	defer db.Close()

	pwd := sha256.Sum256([]byte(password))
	row := db.QueryRow(`SELECT id, email, first_name
    FROM Member
    WHERE email = $1 AND password = $2`, email, hex.EncodeToString(pwd[:]))

	result := Member{}

	err = row.Scan(&result.id, &result.email, &result.firstName)

	if err != nil {
		return result, errors.New("Unable to find Member with email: " + email)
	}

	return result, nil
}

func CreateSession(member Member) (Session, error) {
	result := Session{}
	result.memberID = member.ID()
	sessionID := sha256.Sum256([]byte(member.Email() + time.Now().Format("12:00:00")))
	result.sessionID = hex.EncodeToString(sessionID[:])

	db, err := getDBConnection()

	if err != nil {
		return Session{}, errors.New("Unable to connect to the datbase")
	}

	defer db.Close()

	err = db.QueryRow(`INSERT INTO Session 
        (member_id, session_id) 
        VALUES ($1, $2) 
        RETURNING id`, member.ID(), result.sessionID).Scan(&result.id)

	if err != nil {
		return Session{}, errors.New("Unable to save session in database")
	}

	return result, nil
}
