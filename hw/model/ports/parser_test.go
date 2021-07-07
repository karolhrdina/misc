package ports_test

import (
	"strings"
	"testing"

	"github.com/karolhrdina/misc/hw/model/ports"

	"github.com/stretchr/testify/require"
)

func TestPortParser(t *testing.T) {
	require := require.New(t)

	t.Run("success valid json", func(t *testing.T) {
		input := strings.NewReader(`
        {
            "A" : {"name" : "a"},
            "B" : {"city" : "b"},
            "C" : {"country" : "c"}
        }`)

		p := ports.NewPortParser(input)
		res := []ports.Port{}
		var port ports.Port

		for p.Next() {
			err := p.Scan(&port)
			require.NoError(err)
			res = append(res, port)
		}
		// can't user gubrak here, since my Port has pointers to strings
		require.Len(res, 3)
		// but here the order should be that of how it's being parsed
		require.Equal("A", res[0].Key)
		require.Equal("a", *res[0].Name)

		require.Equal("B", res[1].Key)
		require.Equal("b", *res[1].City)

		require.Equal("C", res[2].Key)
		require.Equal("c", *res[2].Country)
	})

	t.Run("success invalid json", func(t *testing.T) {
		// Here  we test that invalid json closing brackets don't matter
		input := strings.NewReader(`
        {
            "A" : {"name" : "a"},
            "B" : {"city" : "b"}
        `) // <-- missing closing '}' bracket

		p := ports.NewPortParser(input)
		res := []ports.Port{}
		var port ports.Port

		for p.Next() {
			err := p.Scan(&port)
			require.NoError(err)
			res = append(res, port)
		}
		require.Len(res, 2)
		require.Equal("A", res[0].Key)
		require.Equal("a", *res[0].Name)

		require.Equal("B", res[1].Key)
		require.Equal("b", *res[1].City)
	})

	t.Run("failure array instead object", func(t *testing.T) {
		// Here we test that we return correct error when array is encountered
		// instead of object
		input := strings.NewReader(`
        [
            {"name" : "a"},
            {"city" : "b"}
        ]
        `)

		p := ports.NewPortParser(input)
		res := []ports.Port{}
		var port ports.Port

		var err error
		for p.Next() {
			err = p.Scan(&port)
			if err == nil {
				res = append(res, port)
			}
		}
		require.Error(err)
		require.ErrorIs(err, ports.ErrJsonObjectExpected)

		require.Len(res, 0)
	})

	t.Run("failure bad json", func(t *testing.T) {
		input := strings.NewReader(`
        {
            "A : {"name" : "a"},
            "B" : {"city" : "b"}
        }
        `)

		p := ports.NewPortParser(input)
		res := []ports.Port{}
		var port ports.Port

		var err error
		for p.Next() {
			err = p.Scan(&port)
			if err == nil {
				res = append(res, port)
			}
		}
		require.Error(err)

		require.Len(res, 0)
	})
}
