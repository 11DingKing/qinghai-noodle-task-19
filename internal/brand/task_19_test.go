package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask19(t *testing.T) {
	s := NewService(NewRegistry(), time.Now)
	r := Recall{AffectedStores: []string{"s1", "s2"}}
	first, err := s.AcknowledgeRecall(context.Background(), r, "s1")
	require.NoError(t, err)
	second, err := s.AcknowledgeRecall(context.Background(), first, "s2")
	require.NoError(t, err)
	require.ElementsMatch(t, []string{"s1", "s2"}, second.Acknowledged)
}
