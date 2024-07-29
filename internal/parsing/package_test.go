package parsing_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/thatmattlove/go-pkg-domain/internal/parsing"
)

func Test_GetPackagePath(t *testing.T) {
	cases := [][]string{{"go.example.com", "/"}, {"example.com/go", "/go", "example.com/go/pkg", "/go/pkg"}}
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			in := c[0]
			exp := c[1]
			req := httptest.NewRequest(http.MethodGet, "https://"+in, nil)
			res, err := parsing.GetPackagePath(req, in)
			testza.AssertNoError(t, err)
			testza.AssertEqual(t, exp, res)
		})
	}
}

func Test_GetModulePath(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		t.Parallel()
		req := httptest.NewRequest(http.MethodGet, "https://example.com/repo", nil)
		result, err := parsing.GetModulePath(req)
		testza.AssertNoError(t, err)
		testza.AssertEqual(t, "example.com/repo", result)
	})
	t.Run("2", func(t *testing.T) {
		t.Parallel()
		req := httptest.NewRequest(http.MethodGet, "https://go.example.com/repo", nil)
		result, err := parsing.GetModulePath(req)
		testza.AssertNoError(t, err)
		testza.AssertEqual(t, "go.example.com/repo", result)
	})
	t.Run("2", func(t *testing.T) {
		t.Parallel()
		req := httptest.NewRequest(http.MethodGet, "https://example.com/go/repo", nil)
		result, err := parsing.GetModulePath(req)
		testza.AssertNoError(t, err)
		testza.AssertEqual(t, "example.com/go/repo", result)
	})
	t.Run("3", func(t *testing.T) {
		t.Parallel()
		req := httptest.NewRequest(http.MethodGet, "https://example.com/go/pkg/repo", nil)
		result, err := parsing.GetModulePath(req)
		testza.AssertNoError(t, err)
		testza.AssertEqual(t, "example.com/go/pkg/repo", result)
	})
}

func Test_IsGoGet(t *testing.T) {
	t.Run("is", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "https://example.com/repo?go-get=1", nil)
		testza.AssertTrue(t, parsing.IsGoGet(req))
	})
	t.Run("is not 1", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "https://example.com/repo?go-get=2", nil)
		testza.AssertFalse(t, parsing.IsGoGet(req))
	})
	t.Run("is not 2", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "https://example.com/repo", nil)
		testza.AssertFalse(t, parsing.IsGoGet(req))
	})
}

func Test_MakeRepoPath(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		result := parsing.MakeRepoPath("https://github.com/user", "/repo")
		testza.AssertEqual(t, "https://github.com/user/repo", result)
	})
	t.Run("2", func(t *testing.T) {
		result := parsing.MakeRepoPath("http://github.com/user", "/repo")
		testza.AssertEqual(t, "https://github.com/user/repo", result)
	})
	t.Run("3", func(t *testing.T) {
		result := parsing.MakeRepoPath("http://github.com/user", "/repo/pkg")
		testza.AssertEqual(t, "https://github.com/user/repo/pkg", result)
	})
}

func Test_PartialMatch(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		t.Parallel()
		req := httptest.NewRequest(http.MethodGet, "https://example.com/repo", nil)
		result := parsing.PartialMatch(req, "example.com")
		testza.AssertTrue(t, result)
	})
	t.Run("2", func(t *testing.T) {
		t.Parallel()
		req := httptest.NewRequest(http.MethodGet, "https://example.com/repo", nil)
		result := parsing.PartialMatch(req, "go.example.com")
		testza.AssertFalse(t, result)
	})
	t.Run("3", func(t *testing.T) {
		t.Parallel()
		req := httptest.NewRequest(http.MethodGet, "https://example.com/repo", nil)
		result := parsing.PartialMatch(req, "example.com/go")
		testza.AssertFalse(t, result)
	})
	t.Run("4", func(t *testing.T) {
		t.Parallel()
		req := httptest.NewRequest(http.MethodGet, "https://example.com/go/repo", nil)
		result := parsing.PartialMatch(req, "example.com/go")
		testza.AssertTrue(t, result)
	})
	t.Run("5", func(t *testing.T) {
		t.Parallel()
		req := httptest.NewRequest(http.MethodGet, "http://localhost:8787/repo", nil)
		result := parsing.PartialMatch(req, "localhost:8787")
		testza.AssertTrue(t, result)
	})
	t.Run("6", func(t *testing.T) {
		t.Parallel()
		req := httptest.NewRequest(http.MethodGet, "http://localhost:8787/repo?go-get=1", nil)
		result := parsing.PartialMatch(req, "localhost:8787")
		testza.AssertTrue(t, result)
	})
}
