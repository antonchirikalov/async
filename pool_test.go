package async

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPool_Async(t *testing.T) {

	tests := []struct {
		name, groupName, want string
		f                     TaskFunc[string]
	}{
		{
			name:      "return test",
			groupName: "gr",
			want:      "test",
			f: func() string {
				<-time.After(1 * time.Microsecond)
				return "test"
			},
		},
		{
			name:      "return test1",
			groupName: "gr",
			want:      "test1",
			f: func() string {
				<-time.After(1 * time.Microsecond)
				return "test1"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewPool[string](10).Async(tt.groupName, tt.f)
			res := f.Wait()
			assert.Equal(t, tt.want, res)
		})
	}

}

func TestPool_Async_DiffGroup(t *testing.T) {

	//id := getGoID()
	pool := NewPool[int](10)

	f := pool.Async("group1", func() int {
		return getGoID()
	})

	group1ID := f.Wait()

	f = pool.Async("group1", func() int {
		return getGoID()
	})

	group1ID2 := f.Wait()
	assert.Equal(t, group1ID, group1ID2, "must be the same goroutine")

	f = pool.Async("group2", func() int {
		return getGoID()
	})

	group2ID := f.Wait()
	assert.NotEqual(t, group1ID, group2ID, "must be different goroutines")

}

func getGoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
