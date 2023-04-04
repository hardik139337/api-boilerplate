package show

import (
	"bytes"
	"encoding/json"
	"lcode/mocks"
	"lcode/models"
	"lcode/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var rPath = "/test"

func Test_Showcontroller_AddShow(t *testing.T) {
	tests := []struct {
		name string
		body models.Show
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Show{MovieId: "id", TheaterId: "id", StartTime: 2},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("Create", mock.Anything).Return(nil)
			ri.On("Query", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusCreated,
	},
		{
			name: "failes for empty Show name",
			want: http.StatusBadRequest,
			mock: func() repository.RepositoryI {
				ri := mocks.RepositoryI{}
				ri.On("Create", mock.Anything).Return(nil)
				return &ri
			},
			body: models.Show{},
		},
	}
	router := gin.Default()
	m := &Showcontroller{}
	router.GET(rPath, m.AddShow)
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

func TestShowcontroller_GetShows(t *testing.T) {
	tests := []struct {
		name string
		body models.Show
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Show{Id: "test"},
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
			m := &Showcontroller{Repository: tt.mock()}
			router.GET(rPath, m.Getshows)
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

func TestShowcontroller_UpdateShow(t *testing.T) {
	tests := []struct {
		name string
		body models.Show
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Show{
			Id: "id",
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
			m := &Showcontroller{Repository: tt.mock()}
			router.GET(rPath, m.Updateshow)
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

func TestShowcontroller_DeleteShow(t *testing.T) {

	tests := []struct {
		name string
		body models.Show
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Show{
			Id: "id",
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
			m := &Showcontroller{Repository: tt.mock()}
			router.GET(rPath, m.Deleteshow)
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

func TestGetByIdShows(t *testing.T) {
	tests := []struct {
		name string
		id   string
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		id:   "/testId",
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("Query", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusOK,
	}}

	router := gin.Default()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Showcontroller{Repository: tt.mock()}
			router.GET(rPath+"/:id", m.GetByIdshows)

			req, _ := http.NewRequest("GET", rPath+"/abc", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			t.Logf("status: %d", w.Code)
			t.Logf("response: %s", w.Body.String())
			assert.Equal(t, tt.want, w.Code)
		})
	}
}
