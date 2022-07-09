package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddPositive(t *testing.T) {
	sum, err := Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddNegative(t *testing.T) {
	_, err := Add(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}
	_, err = Add(1, -2)
	if err == nil {
		t.Error("second arg negative - expected error not be nil")
	}
	_, err = Add(-1, -2)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestAddZero(t *testing.T) {
	_, err := Add(0, 1)
	if err == nil {
		t.Error("first argument is zeroe - expected error not nil")
	}
	_, err = Add(1, 0)
	if err == nil {
		t.Error("second argument is zeroe - expected error not nil")
	}
	_, err = Add(0, 0)
	if err == nil {
		t.Error("both arguments are zeroes - expected error not nil")
	}
}

// bad, as every test should test one type of behaviour
func TestDivision(t *testing.T) {
	result, err := Divide(0, 1)
	require.NoError(t, err)
	assert.Equal(t, 0, result)

	result, err = Divide(4, 2)
	require.NoError(t, err)
	assert.Equal(t, 2, result)

	_, err = Divide(1, 0)
	require.Error(t, ZeroDivisionError)
}

func TestDivisionBetter(t *testing.T) {
	t.Run("ZeroNumerator", func(t *testing.T) {
		result, err := Divide(0, 1)
		t.Log(result)
		require.NoError(t, err)
		assert.Equal(t, 0, result)
	})
	t.Run("BothNonZero", func(t *testing.T) {
		result, err := Divide(4, 2)
		require.NoError(t, err)
		assert.Equal(t, 2, result)
	})
	t.Run("ZeroDenomenator", func(t *testing.T) {
		_, err := Divide(1, 0)
		require.Error(t, err)
	})
}

func TestEstimateValue(t *testing.T) {
	t.Run(string(EstimatedValueSmall), func(t *testing.T) {
		estimation := EstimateValue(9)
		assert.Equal(t, EstimatedValueSmall, estimation)
	})
	t.Run(string(EstimatedValueMedium), func(t *testing.T) {
		estimation := EstimateValue(22)
		assert.Equal(t, EstimatedValueMedium, estimation)
	})
	t.Run(string(EstimatedValueLarge), func(t *testing.T) {
		estimation := EstimateValue(5000)
		assert.Equal(t, EstimatedValueLarge, estimation)
	})
	t.Run(string("negative"), func(t *testing.T) {
		estimation := EstimateValue(-5000)
		assert.Equal(t, "negative", estimation)
	})
}
