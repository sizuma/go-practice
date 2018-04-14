package manager

import "c2c/model"

// UserManager is a Manager for User
type UserManager interface {
	FetchByID(id int) *model.User
	FetchByEmail(email string) *model.User
	Insert(email, password string) int
}

type userManager struct {
	users  map[int]*model.User
	lastID int
}

func (manager *userManager) FetchByID(id int) *model.User {
	if user, ok := manager.users[id]; ok {
		return user
	}
	return nil
}

func (manager *userManager) FetchByEmail(email string) *model.User {
	for _, value := range manager.users {
		if value.Email == email {
			return value
		}
	}
	return nil
}

func (manager *userManager) Insert(email, password string) int {
	userModel := model.User{
		ID:       manager.lastID,
		Email:    email,
		Password: password,
	}
	manager.users[userModel.ID] = &userModel
	manager.lastID++
	return userModel.ID
}

// EmptyInMemoryManager returns new/empty user manager
func EmptyInMemoryManager() UserManager {
	return &userManager{
		users:  make(map[int]*model.User),
		lastID: 0,
	}
}
