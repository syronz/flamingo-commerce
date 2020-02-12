package contextstore_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os/exec"
	"testing"

	"flamingo.me/flamingo/v3/framework/flamingo"
	"github.com/go-test/deep"
	"github.com/gomodule/redigo/redis"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stvp/tempredis"

	"flamingo.me/flamingo-commerce/v3/checkout/domain/placeorder/process"
	"flamingo.me/flamingo-commerce/v3/checkout/infrastructure/contextstore"
)

const (
	existingKey  = "existing"
	wrongDataKey = "wrong data"
)

var (
	testContext  = process.Context{UUID: "test"}
	emptyContext = process.Context{}
)

func getRedisStore(network, address string) *contextstore.Redis {
	return new(contextstore.Redis).Inject(
		new(flamingo.NullLogger),
		&struct {
			MaxIdle                 float64 `inject:"config:commerce.checkout.placeorder.contextstore.redis.maxIdle"`
			IdleTimeoutMilliseconds float64 `inject:"config:commerce.checkout.placeorder.contextstore.redis.idleTimeoutMilliseconds"`
			Network                 string  `inject:"config:commerce.checkout.placeorder.contextstore.redis.network"`
			Address                 string  `inject:"config:commerce.checkout.placeorder.contextstore.redis.address"`
			Database                float64 `inject:"config:commerce.checkout.placeorder.contextstore.redis.database"`
		}{MaxIdle: 3, IdleTimeoutMilliseconds: 240000, Network: network, Address: address, Database: 0})
}

func prepareData(t *testing.T, conn redis.Conn) {
	buffer := new(bytes.Buffer)
	require.NoError(t, gob.NewEncoder(buffer).Encode(testContext))
	_, err := conn.Do("SET", existingKey, buffer)
	require.NoError(t, err)
	_, err = conn.Do("SET", wrongDataKey, "wrong data")
	require.NoError(t, err)
}

func startUpLocalRedis(t *testing.T) (*tempredis.Server, redis.Conn) {
	t.Helper()
	server, err := tempredis.Start(tempredis.Config{})
	if err != nil {
		t.Fatal(err)
	}
	conn, err := redis.Dial("unix", server.Socket())
	if err != nil {
		t.Fatal(err)
	}
	prepareData(t, conn)

	return server, conn
}

func startUpDockerRedis(t *testing.T) (func(), string, redis.Conn) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Skip("docker not installed")
	}

	res, err := pool.Run("redis", "5.0", nil)
	require.NoError(t, err, "failed to run redis docker image")
	address := fmt.Sprintf("%s:%s", "127.0.0.1", res.GetPort("6379/tcp"))

	var conn redis.Conn
	require.NoError(t, pool.Retry(func() error { conn, err = redis.Dial("tcp", address); return err }))

	prepareData(t, conn)

	return func() { pool.Purge(res) }, address, conn
}

func TestRedis_Get(t *testing.T) {
	runTestCases := func(t *testing.T, store *contextstore.Redis) {
		tests := []struct {
			name          string
			key           string
			expectedFound bool
			expected      process.Context
		}{
			{
				name:          "load existing",
				key:           existingKey,
				expectedFound: true,
				expected:      testContext,
			},
			{
				name:          "load existing with wrong data",
				key:           wrongDataKey,
				expectedFound: false,
				expected:      emptyContext,
			},
			{
				name:          "load non existing",
				key:           "non",
				expectedFound: false,
				expected:      emptyContext,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, ok := store.Get(tt.key)
				assert.Equal(t, tt.expectedFound, ok)
				if diff := deep.Equal(got, tt.expected); diff != nil {
					t.Error("expected response is wrong: ", diff)
				}
			})
		}
	}

	t.Run("local-redis", func(t *testing.T) {
		if _, err := exec.LookPath("redis-server"); err != nil {
			t.Skip("redis-server not installed")
		}
		server, _ := startUpLocalRedis(t)
		store := getRedisStore("unix", server.Socket())
		runTestCases(t, store)
	})
	t.Run("docker-redis", func(t *testing.T) {
		if _, err := exec.LookPath("docker"); err != nil {
			t.Skip("docker not installed")
		}
		shutdown, address, _ := startUpDockerRedis(t)
		defer shutdown()
		store := getRedisStore("tcp", address)
		runTestCases(t, store)
	})

}

func TestRedis_Store(t *testing.T) {
	runTestCases := func(t *testing.T, store *contextstore.Redis, conn redis.Conn) {
		tests := []struct {
			name  string
			key   string
			value process.Context
		}{
			{
				name:  "store new value",
				key:   "test_key",
				value: process.Context{UUID: "test-uuid"},
			},
			{
				name:  "overwrite existing",
				key:   existingKey,
				value: process.Context{UUID: "test-uuid"},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				require.NoError(t, store.Store(tt.key, tt.value))

				result, err := redis.Bytes(conn.Do("GET", tt.key))
				require.NoError(t, err)

				buffer := new(bytes.Buffer)
				require.NoError(t, gob.NewEncoder(buffer).Encode(tt.value))

				assert.Equal(t, buffer.Bytes(), result)
			})
		}
	}

	t.Run("local-redis", func(t *testing.T) {
		if _, err := exec.LookPath("redis-server"); err != nil {
			t.Skip("redis-server not installed")
		}
		server, conn := startUpLocalRedis(t)
		store := getRedisStore("unix", server.Socket())
		runTestCases(t, store, conn)
	})
	t.Run("docker-redis", func(t *testing.T) {
		if _, err := exec.LookPath("docker"); err != nil {
			t.Skip("docker not installed")
		}
		shutdown, address, conn := startUpDockerRedis(t)
		defer shutdown()
		store := getRedisStore("tcp", address)
		runTestCases(t, store, conn)
	})
}

func TestRedis_Delete(t *testing.T) {
	runTestCases := func(t *testing.T, store *contextstore.Redis, conn redis.Conn) {
		tests := []struct {
			name string
			key  string
		}{
			{
				name: "delete existing",
				key:  existingKey,
			},
			{
				name: "delete non existing",
				key:  "test_key",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				require.NoError(t, store.Delete(tt.key))

				_, err := redis.Bytes(conn.Do("GET", tt.key))
				require.Error(t, err, "entry not deleted")
			})
		}
	}

	t.Run("local-redis", func(t *testing.T) {
		if _, err := exec.LookPath("redis-server"); err != nil {
			t.Skip("redis-server not installed")
		}
		server, conn := startUpLocalRedis(t)
		store := getRedisStore("unix", server.Socket())
		runTestCases(t, store, conn)
	})
	t.Run("docker-redis", func(t *testing.T) {
		if _, err := exec.LookPath("docker"); err != nil {
			t.Skip("docker not installed")
		}
		shutdown, address, conn := startUpDockerRedis(t)
		defer shutdown()
		store := getRedisStore("tcp", address)
		runTestCases(t, store, conn)
	})
}