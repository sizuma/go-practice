package manager

import (
	"c2c/db"
	"c2c/model"
)

// UserManager is a Manager for User
type UserManager interface {
	FetchByID(id int) (*model.User, error)
	FetchByEmail(email string) (*model.User, error)
	Insert(email, password string) (int, error)
}

type userManager struct {
	users  map[int]*model.User
	lastID int
}

func (manager *userManager) FetchByID(id int) (*model.User, error) {
	if user, ok := manager.users[id]; ok {
		return user, nil
	}
	return nil, nil
}

func (manager *userManager) FetchByEmail(email string) (*model.User, error) {
	for _, value := range manager.users {
		if value.Email == email {
			return value, nil
		}
	}
	return nil, nil
}

func (manager *userManager) Insert(email, password string) (int, error) {
	userModel := model.User{
		ID:       manager.lastID,
		Email:    email,
		Password: password,
	}
	manager.users[userModel.ID] = &userModel
	manager.lastID++
	return userModel.ID, nil
}

// EmptyInMemoryManager returns new/empty user manager
func EmptyInMemoryManager() UserManager {
	return &userManager{
		users:  make(map[int]*model.User),
		lastID: 0,
	}
}

type userManagerDBImplementation struct {
}

func (manager *userManagerDBImplementation) FetchByID(id int) (*model.User, error) {
	db, error := db.Connect()
	if error != nil {
		return nil, error
	}
	user := model.User{}
	row := db.QueryRow("select id, email, password from users where id = ? limit 1", id)

	scanRowError := row.Scan(&user.ID, &user.Email, &user.Password)
	if scanRowError != nil {
		return nil, nil
	}
	return &user, scanRowError
}

func (manager *userManagerDBImplementation) FetchByEmail(email string) (*model.User, error) {
	db, error := db.Connect()
	if error != nil {
		return nil, error
	}
	user := model.User{}
	row := db.QueryRow("select id, email, password from users where email = ? limit 1", email)
	scanRowError := row.Scan(&user.ID, &user.Email, &user.Password)

	if scanRowError != nil {
		return nil, nil
	}
	return &user, scanRowError
}

func (manager *userManagerDBImplementation) Insert(email, password string) (int, error) {
	db, error := db.Connect()
	if error != nil {
		return 0, error
	}

	res, insertError := db.Exec("insert into users(email, password) values(?, ?)", email, password)
	if insertError != nil {
		return 0, insertError
	}

	id, lastIDError := res.LastInsertId()
	if lastIDError != nil {
		return 0, lastIDError
	}
	return int(id), nil
}

// EmptyDBManager return empty user manager implemented by db
func EmptyDBManager() UserManager {
	return &userManagerDBImplementation{}
}
