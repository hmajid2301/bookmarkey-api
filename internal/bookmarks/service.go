package bookmarks

import (
	"errors"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const unsortedCollection = "-1"

// CollectionService is used to interact with the collection service
type CollectionService interface {
	IsCollectionOwnedBy(collectionID string, userID string) (bool, error)
}

// Repository is an interface for interacting with the database
type Repository interface {
	Create(metadata *BookmarkMetaData, collectionID string, userID string) error
	GetByURL(url string) (*BookmarkMetaData, error)
}

// Service used to interact with the bookmark
type Service struct {
	repo          Repository
	collectionSrv CollectionService
}

// NewService returns a new struct used to interact with this module
func NewService(repo Repository, srv CollectionService) Service {
	return Service{
		repo:          repo,
		collectionSrv: srv,
	}
}

// ErrNotAuthorized an error when user is not authorized to perform an action
var ErrNotAuthorized = errors.New("user does not have permission to create bookmark on collection")

// Create used to create a new bookmark
func (s Service) Create(url, collectionID, userID string) error {
	if collectionID != unsortedCollection {
		ownedBy, err := s.collectionSrv.IsCollectionOwnedBy(collectionID, userID)
		if err != nil {
			return err
		}
		if !ownedBy {
			return ErrNotAuthorized
		}
	}

	metadata, err := s.GetMetadata(url)
	if err != nil {
		return err
	}

	err = s.repo.Create(metadata, collectionID, userID)
	if err != nil {
		return err
	}

	return nil
}

// GetMetadata returns the metadata for a URL.
func (s Service) GetMetadata(url string) (*BookmarkMetaData, error) {
	bookmarkMetadata, err := s.repo.GetByURL(url)
	if err == nil {
		return bookmarkMetadata, nil
	}

	metadata := &BookmarkMetaData{URL: url}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		return metadata, err
	}
	defer func() {
		err = res.Body.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return metadata, err
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if property, _ := s.Attr("property"); property == "og:description" {
			metadata.Description, _ = s.Attr("content")
		}
		if property, _ := s.Attr("property"); property == "og:image" {
			metadata.Image, _ = s.Attr("content")
		}
		if property, _ := s.Attr("property"); property == "og:title" {
			metadata.Title, _ = s.Attr("content")
		}
	})
	return metadata, nil
}
