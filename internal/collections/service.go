package collections

// Repository used to interact with the database
type Repository interface {
	GetByID(id string) (*Collection, error)
}

// Service used to interact with the collection
type Service struct {
	repo Repository
}

// NewService returns a new struct used to interact with this module
func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

// IsCollectionOwnedBy checks if a collection is owned by user
func (s Service) IsCollectionOwnedBy(collectionID string, userID string) (bool, error) {
	collection, err := s.repo.GetByID(collectionID)
	if err != nil {
		return false, err
	}

	return collection.User == userID, nil
}
