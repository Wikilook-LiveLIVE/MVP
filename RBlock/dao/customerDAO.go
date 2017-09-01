package dao

type CustomerDAO struct {
}

func NewCustomerDAO () *CustomerDAO{

	if _, err := pool.Prepare("INSERT_CUSTOMER", "INSERT INTO rblock.customer("+
		"customer_email_c, customer_password_h, customer_reg_key, customer_created)"+
		"VALUES ($1, $2, $3, date_trunc('second', NOW())) returning customer_id"); err != nil {
		panic(err)
	}

	if _, err := pool.Prepare("UPDATE_ACTIVATION", "UPDATE rblock.customer SET "+
		"customer_activated=date_trunc('second', NOW()), customer_reg_key=null WHERE customer_reg_key=$1 returning customer_id"); err != nil {
		panic(err)
	}

	return &CustomerDAO{}
}

func (self *CustomerDAO) SaveSignUp(email, password, activationCode string) (int64, error) {
	return InsertReturning("INSERT_CUSTOMER", email, password, activationCode)
}

func (self *CustomerDAO) ActivateByRegKey(activationCode string) (int64, error) {
	return UpdateReturning("UPDATE_ACTIVATION", activationCode)
}