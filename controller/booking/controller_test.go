package booking

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

func TestBookingcontroller_AddBooking(t *testing.T) {
	tests := []struct {
		name string
		body models.BookingR
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.BookingR{UserId: "", ShowId: "test", SeatIdS: []string{""}},
		mock: func() repository.RepositoryI {
			ri := mocks.RepositoryI{}
			ri.On("Create", mock.Anything).Return(nil)
			return &ri
		},
		want: http.StatusCreated,
	},
		{
			name: "failes for empty Booking name",
			want: http.StatusBadRequest,
			mock: func() repository.RepositoryI {
				ri := mocks.RepositoryI{}
				ri.On("Create", mock.Anything).Return(nil)
				return &ri
			},
		},
	}
	router := gin.Default()
	m := &Bookingcontroller{}
	router.GET(rPath, m.AddBooking)
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

func TestBookingcontroller_GetBookings(t *testing.T) {
	tests := []struct {
		name string
		body models.Booking
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Booking{Id: "test"},
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
			m := &Bookingcontroller{Repository: tt.mock()}
			router.GET(rPath, m.GetBookings)
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

func TestBookingcontroller_UpdateBooking(t *testing.T) {
	tests := []struct {
		name string
		body models.Booking
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Booking{},
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
			m := &Bookingcontroller{Repository: tt.mock()}
			router.GET(rPath, m.UpdateBooking)
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

func TestBookingcontroller_DeleteBooking(t *testing.T) {

	tests := []struct {
		name string
		body models.Booking
		mock func() repository.RepositoryI
		want int
	}{{
		name: "pass",
		body: models.Booking{
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
			m := &Bookingcontroller{Repository: tt.mock()}
			router.GET(rPath, m.DeleteBooking)
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
