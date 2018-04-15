package manager

import (
	"c2c/db"
	"c2c/manager"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestDBCreate(t *testing.T) {
	migrationError := db.Migration(true, true)
	if migrationError != nil {
		t.Error(migrationError)
	}
	manager.EmptyDBManager()
}

func TestDBInsert(t *testing.T) {
	migrationError := db.Migration(true, true)
	if migrationError != nil {
		t.Error(migrationError)
	}
	userManager := manager.EmptyDBManager()
	userManager.Insert("user001@test.com", "password")
}

func TestDBFetchByID(t *testing.T) {
	migrationError := db.Migration(true, true)
	if migrationError != nil {
		t.Error(migrationError)
	}
	userManager := manager.EmptyDBManager()
	id, error := userManager.Insert("user001@test.com", "password")
	if error != nil {
		t.Fail()
	}
	if user, error := userManager.FetchByID(id); user != nil && error == nil {
		if user.Email != "user001@test.com" || user.Password != "password" {
			t.Fail()
		}
	} else {
		t.Errorf("%s %s", user, error)
	}
}

func TestDBFetchByMissingID(t *testing.T) {
	migrationError := db.Migration(true, true)
	if migrationError != nil {
		t.Error(migrationError)
	}
	userManager := manager.EmptyDBManager()
	id, error := userManager.Insert("user001@test.com", "password")
	if error != nil {
		t.Fail()
	}
	if user, error := userManager.FetchByID(id + 1); user != nil && error != nil {
		t.Errorf("expect nil but got %s", user)
	}
}

func TestDBFetchByEmail(t *testing.T) {
	migrationError := db.Migration(true, true)
	if migrationError != nil {
		t.Error(migrationError)
	}
	userManager := manager.EmptyDBManager()
	userManager.Insert("user001@test.com", "password")
	if user, error := userManager.FetchByEmail("user001@test.com"); user != nil && error == nil {
		if user.Email != "user001@test.com" || user.Password != "password" {
			t.Errorf("expect inserted user but got %s", user)
		}
	} else {
		t.Errorf("expect inserted user but got nil")
	}
}

func TestDBFetchByMissingEmail(t *testing.T) {
	migrationError := db.Migration(true, true)
	if migrationError != nil {
		t.Error(migrationError)
	}
	userManager := manager.EmptyDBManager()
	userManager.Insert("user001@test.com", "password")
	if user, error := userManager.FetchByEmail(""); user != nil && error != nil {
		t.Errorf("expect nil but got %s", user)
	}
}
