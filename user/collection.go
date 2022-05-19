package user

type Collection map[string]*User

var collection Collection

func InitializeCollection() {
	collection = Collection{}
}

func (coll Collection) get(email string) *User {
	return coll[email]
}

func (coll Collection) save(email string, user *User) {
	coll[email] = user
}
