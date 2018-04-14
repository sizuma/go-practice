package manager

import (
	"c2c/manager"
	"testing"
)

func TestCreate(t *testing.T) {
	manager.EmptyInMemoryManager()
}

func TestInsert(t *testing.T) {
	userManager := manager.EmptyInMemoryManager()
	userManager.Insert("user001@test.com", "password")
}

func TestFetchByID(t *testing.T) {
	userManager := manager.EmptyInMemoryManager()
	id := userManager.Insert("user001@test.com", "password")
	if user := userManager.FetchByID(id); user != nil {
		if user.Email != "user001@test.com" || user.Password != "password" {
			t.Fail()
		}
	} else {
		t.Fail()
	}
}

func TestFetchByMissingID(t *testing.T) {
	userManager := manager.EmptyInMemoryManager()
	id := userManager.Insert("user001@test.com", "password")
	if user := userManager.FetchByID(id + 1); user != nil {
		t.Errorf("expect nil but got %s", user)
	}
}

func TestFetchByEmail(t *testing.T) {
	userManager := manager.EmptyInMemoryManager()
	userManager.Insert("user001@test.com", "password")
	if user := userManager.FetchByEmail("user001@test.com"); user != nil {
		if user.Email != "user001@test.com" || user.Password != "password" {
			t.Errorf("expect inserted user but got %s", user)
		}
	} else {
		t.Errorf("expect inserted user but got nil")
	}
}

func TestFetchByMissingEmail(t *testing.T) {
	userManager := manager.EmptyInMemoryManager()
	userManager.Insert("user001@test.com", "password")
	if user := userManager.FetchByEmail(""); user != nil {
		t.Errorf("expect nil but got %s", user)
	}
}
