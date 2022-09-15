package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Nitin", City: "Rewari", Zipcode: "123401", DateOfBirth: "1996-01-01", Status: "1"},
		{Id: "1002", Name: "Anuj", City: "Rewari", Zipcode: "123402", DateOfBirth: "1999-01-01", Status: "1"},
		{Id: "1003", Name: "Priyanka", City: "Rewari", Zipcode: "123403", DateOfBirth: "1995-01-01", Status: "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func (d CustomerRepositoryStub) FindCustomerById(string) (*Customer, error) {
	return nil, nil
}
