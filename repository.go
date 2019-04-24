package main

type Repository interface {
	GetCustomerProfileByID(customerID string) (CustomerModel, error)
}

type CustomerRepository struct {
	// todo keep db instance for separate code query logic
}

type MockMongoDBSession struct {
	// todo keep db instance for separate code query logic
}

func NewCustomerSQLRepository(session string) *CustomerRepository {
	return &CustomerRepository{}
}

func NewCustomerMongoRepository(session string) *MockMongoDBSession {
	return &MockMongoDBSession{}
}

func (c *CustomerRepository) GetCustomerProfileByID(customerID string) (*CustomerModel, error) {
	customerProfile := new(CustomerModel)
	customerProfile.ID = "sql-id"
	customerProfile.Name = "sql-name"
	return customerProfile, nil
}

func (m *MockMongoDBSession) GetCustomerProfileByID(customerID string) (*CustomerModel, error) {
	customerProfile := new(CustomerModel)
	customerProfile.ID = "mongo-id"
	customerProfile.Name = "mongo-name"
	return customerProfile, nil
}
