package mock

type MockAuth struct{}

func (m *MockAuth) GenerateToken(ID string) (string, error) {
	return "token", nil
}

func (m *MockAuth) ValidateToken(token string) (string, error) {
	return "id", nil
}
