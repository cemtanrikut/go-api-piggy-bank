package api

type UserModel struct {
	ID       string
	Name     string
	Surname  string
	Password string
	IsActive string
}

// Create an user
func Create(user UserModel) {

}

// Get an user by userID
func Get(userID string) {

}

// Get user list
func GetList() {

}

// Update user
func Update(user UserModel) {

}

// Delete user by userID
func Delete(userID string) {

}
