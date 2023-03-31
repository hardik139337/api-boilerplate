package controller

import (
	"bytes"
	"encoding/json"
	"lcode/mocks"
	"lcode/models"
	"lcode/repository"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var rPath = "/test"

func TestMoviecontroller_AddMovie(t *testing.T) {
	tests := []struct {
		name string
		body models.Movie
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Movie{
			Name:     "Titanic",
			Language: "Hindi",
			Length:   2 * time.Hour,
		},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("Create", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusCreated,
	},
		{
			name: "failes for empty movie name",
			want: http.StatusBadRequest,
			mock: func() repository.RepositoryI {
				ri := mocks.RepositoryI{}
				ri.On("Create", mock.Anything).Return(nil)
				return &ri
			},
			body: models.Movie{
				Language: "Hindi",
				Length:   2 * time.Hour,
			},
		},
	}
	router := gin.Default()
	m := &Moviecontroller{}
	router.GET(rPath, m.AddMovie)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.Repository = tt.mock()
			b, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			t.Logf("status: %d", w.Code)
			t.Logf("response: %s", w.Body.String())
			assert.Equal(t, tt.want, w.Code)
		})
	}
}

func TestMoviecontroller_GetMovies(t *testing.T) {
	tests := []struct {
		name string
		body models.Movie
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Movie{Id: "test"},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("QueryAll", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusOK,
	}}

	router := gin.Default()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Moviecontroller{Repository: tt.mock()}
			router.GET(rPath, m.GetMovies)
			b, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			t.Logf("status: %d", w.Code)
			t.Logf("response: %s", w.Body.String())
			assert.Equal(t, tt.want, w.Code)
		})
	}
}

func TestMoviecontroller_UpdateMovie(t *testing.T) {
	tests := []struct {
		name string
		body models.Movie
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Movie{
			Id:       "id",
			Name:     "Titanic",
			Language: "Hindi",
			Length:   2 * time.Hour,
		},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("Update", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusOK,
	}}

	router := gin.Default()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Moviecontroller{Repository: tt.mock()}
			router.GET(rPath, m.UpdateMovie)
			b, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			t.Logf("status: %d", w.Code)
			t.Logf("response: %s", w.Body.String())
			assert.Equal(t, tt.want, w.Code)
		})
	}
}

func TestMoviecontroller_DeleteMovie(t *testing.T) {

	tests := []struct {
		name string
		body models.Movie
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Movie{
			Id:       "id",
			Name:     "Titanic",
			Language: "Hindi",
			Length:   2 * time.Hour,
		},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("Delete", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusOK,
	}}

	router := gin.Default()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Moviecontroller{Repository: tt.mock()}
			router.GET(rPath, m.DeleteMovie)
			b, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest("GET", rPath, bytes.NewReader(b))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			t.Logf("status: %d", w.Code)
			t.Logf("response: %s", w.Body.String())
			assert.Equal(t, tt.want, w.Code)
		})
	}
}
