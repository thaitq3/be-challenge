package distributedlock

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

//go:generate mockery -name Mutex -inpkg -case=underscore
type Mutex interface {
	Lock() error

	Unlock() error
}

//go:generate mockery -name DistributedLock -inpkg -case=underscore
type DistributedLock interface {
	NewMutex(name string, expiry, wait int) Mutex
}

const (
	setCmd      = "SET"
)

var (
	// lua script to release a lock with correct token
	luaRelease = redis.NewScript( `if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("del", KEYS[1]) else return 0 end`)
)

var (
	// ErrLockNotObtained may be returned by Lock when a lock could not be obtained.
	ErrLockNotObtained = errors.New("lock not obtained")

	errLockUnlockFailed = errors.New("lock unlock failed")
)

type cRedisMutex struct {
	client *redis.Client
	key    string
	token  string
	mutex  sync.Mutex

	retries    int
	retryDelay time.Duration
	expiry     time.Duration
}

func (m *cRedisMutex) Lock() error {

	m.mutex.Lock()
	defer m.mutex.Unlock()

	// Create a random token
	token, err := randomToken()
	if err != nil {
		return err
	}

	// Calculate the timestamp we are willing to wait for and number of attempts
	attempts := m.retries + 1
	for {
		// Try to obtain a lock
		// cmd: SET lock_key token EX expiry NX
		res, err := m.client.Do(context.Background(), setCmd, m.key, token, "EX", m.expiry.Seconds(), "NX").Result()
		if err == nil && res.(string) == "OK" {
			m.token = token
			return nil
		}

		if attempts--; attempts <= 0 {
			return ErrLockNotObtained
		}
		time.Sleep(m.retryDelay)
	}
}

// Unlock releases the lock
func (m *cRedisMutex) Unlock() error {
	m.mutex.Lock()
	err := m.release()
	m.mutex.Unlock()
	return err
}

func (m *cRedisMutex) release() error {
	res, err := luaRelease.Run(context.Background(), m.client, []string{m.key}, m.token).Int64()

	if res != 1 {
		return errLockUnlockFailed
	}

	return err
}

func randomToken() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buf), nil
}

type shardedRedisLock struct {
	client *redis.Client
}

func (l *shardedRedisLock) NewMutex(name string, expiry, wait int) Mutex {
	return &cRedisMutex{
		client:     l.client,
		key:        "lock_" + name,
		expiry:     time.Duration(expiry) * time.Second,
		retryDelay: 50 * time.Millisecond,
		retries:    wait * 2,
	}
}

func New(client *redis.Client) DistributedLock {

	return &shardedRedisLock{
		client: client,
	}
}
