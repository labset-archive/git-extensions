package pet_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viqueen/go-template/pkg/pet/petmocks"
)

func TestProject(t *testing.T) {
	t.Parallel()
	project := petmocks.NewProject(t)
	project.EXPECT().Get().Return("project")
	assert.Equal(t, "project", project.Get())
}
