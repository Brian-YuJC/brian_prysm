package random

import (
	"testing"

	"github.com/prysmaticlabs/prysm/v5/testing/spectest/shared/electra/sanity"
)

func TestMinimal_Electra_Random(t *testing.T) {
	t.Skip("TODO: Electra")
	sanity.RunBlockProcessingTest(t, "minimal", "random/random/pyspec_tests")
}