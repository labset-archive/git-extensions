package registry_test

import (
	"os"
	"testing"

	"github.com/dgraph-io/badger/v4"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/viqueen/git-extensions/pkg/registry"
)

type TestData struct {
	Name string `json:"name"`
}

func TestRegistry(t *testing.T) {
	t.Parallel()

	localPath := ".test-registry"
	instance, err := registry.NewBadgerRegistry[TestData](localPath)
	require.NoError(t, err)

	t.Cleanup(func() {
		err = os.RemoveAll(localPath)
		require.NoError(t, err)
	})

	testData := TestData{Name: "test"}
	item := registry.Item[TestData]{ID: "test-id", Value: testData}

	created, err := instance.CreateOrUpdate(item)
	require.NoError(t, err)
	assert.Equal(t, &item, created)

	found, err := instance.Read("test-id")
	require.NoError(t, err)
	assert.Equal(t, &item, found)

	notFound, err := instance.Read("non-existent-id")
	require.Error(t, err)
	assert.Contains(t, err.Error(), badger.ErrKeyNotFound.Error())
	assert.Nil(t, notFound)

	items, err := instance.ListItems(registry.AllFilter[TestData]())
	require.NoError(t, err)
	assert.Len(t, items, 1)
	assert.Equal(t, item, items[0])

	ids, err := instance.ListIds()
	require.NoError(t, err)
	assert.Len(t, ids, 1)
	assert.Equal(t, "test-id", ids[0])

	deleted, err := instance.Delete("test-id")
	require.NoError(t, err)
	assert.Equal(t, &item, deleted)

	notFoundDeleted, err := instance.Delete("test-id")
	require.Error(t, err)
	assert.Contains(t, err.Error(), badger.ErrKeyNotFound.Error())
	assert.Nil(t, notFoundDeleted)
}
