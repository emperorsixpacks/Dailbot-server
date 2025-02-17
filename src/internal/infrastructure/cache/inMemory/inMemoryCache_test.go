package inmemoryCache

import (
	"testing"
	"time"
)

// Test Set and Get functionality
func TestCache_SetAndGet(t *testing.T) {
	c := GetCache()
	key := "testKey"
	value := "testValue"

	_, err := c.Set(key, value, 2*time.Second)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	got, found := c.Get(key)
	if !found {
		t.Fatalf("Expected to find key %s", key)
	}
	if got != value {
		t.Errorf("Expected value %v, got %v", value, got)
	}
}

// Test expired items
func TestCache_ExpiredItem(t *testing.T) {
	c := GetCache()

	key := "tempKey"
	value := "tempValue"

	_, err := c.Set(key, value, 1*time.Millisecond) // Tiny TTL
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	time.Sleep(2 * time.Millisecond) // Wait for expiration

	_, found := c.Get(key)
	if found {
		t.Errorf("Expected key %s to expire, but it didn't", key)
	}
}

// Test deleting an item
func TestCache_Delete(t *testing.T) {
	c := GetCache()
	key := "delKey"
	value := "delValue"

	c.Set(key, value, 5*time.Second)

	success := c.Delete(key)
	if !success {
		t.Fatalf("Expected Delete to succeed")
	}

	_, found := c.Get(key)
	if found {
		t.Errorf("Expected key %s to be deleted, but it's still present", key)
	}
}

// Test cache length
func TestCache_Len(t *testing.T) {
	c := GetCache()
	c.Set("key1", "val1", 5*time.Second)
	c.Set("key2", "val2", 5*time.Second)

	if c.Len() != 2 {
		t.Errorf("Expected cache length to be 2, got %d", c.Len())
	}
}

// Test flushing the cache
func TestCache_Flush(t *testing.T) {
	c := GetCache()
	c.Set("key1", "val1", 5*time.Second)
	c.Set("key2", "val2", 5*time.Second)

	c.Flush()

	if c.Len() != 0 {
		t.Errorf("Expected cache to be empty after Flush, but got %d items", c.Len())
	}
}

// Test deleting expired items
func TestCache_DeleteExpired(t *testing.T) {
	c := GetCache()
	c.Set("key1", "val1", 2*time.Second)  // Will expire quickly
	c.Set("key2", "val2", 10*time.Second) // Won't expire
	time.Sleep(3 * time.Second)           // Wait for key1 to expire
	if c.Len() != 1 {
		t.Errorf("Expected only non-expired items to remain, but got %d", c.Len())
	}
}
