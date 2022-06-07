package redis_cache

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_pckg(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	key := strconv.Itoa(rand.Intn(1000))
	val1 := strconv.Itoa(rand.Intn(1000))
	val2 := strconv.Itoa(rand.Intn(1000))
	var rc *RedisCache
	var res string

	t.Log("\tCreating client")
	{
		rc = New("localhost", "6379", "", 0)
	}

	t.Log("\tSetting value")
	{
		rc.Create(key, val1)
	}

	t.Log("\tGetting value")
	{
		res = rc.Read(key)
		require.Equal(t, val1, res)
	}

	t.Log("\tUpdating value")
	{
		rc.Update(key, val2)
		res = rc.Read(key)
		require.NotEqual(t, val1, res)
	}

	t.Log("\tDeleting value")
	{
		rc.Delete(key)
		res = rc.Read(key)
		require.Empty(t, res)
	}
}
