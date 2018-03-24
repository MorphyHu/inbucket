package mem

import (
	"testing"

	"github.com/jhillyerd/inbucket/pkg/config"
	"github.com/jhillyerd/inbucket/pkg/storage"
	"github.com/jhillyerd/inbucket/pkg/test"
)

// TestSuite runs storage package test suite on file store.
func TestSuite(t *testing.T) {
	test.StoreSuite(t, func(conf config.Storage) (storage.Store, func(), error) {
		s, _ := New(conf)
		destroy := func() {}
		return s, destroy, nil
	})
}
