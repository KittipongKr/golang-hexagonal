package mongo_model

type Users struct {
	BaseModel `bson:",inline"`
	Username  string `bson:"username"`
	Password  string `bson:"password"`
	Email     string `bson:"email"`
	Phone     string `bson:"phone"`
	Status    string `bson:"status"`
}
