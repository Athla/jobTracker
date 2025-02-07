package server

import (
	"bytes"
	"encoding/json"
	"jobTracker/internal/database"
	"jobTracker/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJobHandlers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("CreateJobHandler", func(t *testing.T) {
		tests := []struct {
			name       string
			input      models.Job
			wantStatus int
			wantErr    bool
		}{
			{
				name: "Valid Job Creation",
				input: models.Job{
					Name:        "Software Engineer",
					Company:     "Tech Corp",
					Source:      "LinkedIn",
					Description: "Building awesome stuff",
					JobType:     models.FullTime,
					Status:      models.Wishlist,
				},
				wantStatus: http.StatusCreated,
				wantErr:    false,
			},
			{
				name: "Invalid Job - Missing Required Fields",
				input: models.Job{
					Description: "No name or company",
				},
				wantStatus: http.StatusBadRequest,
				wantErr:    true,
			},
			{
				name: "Invalid Job Status",
				input: models.Job{
					Name:    "Developer",
					Company: "Tech Corp",
					Status:  "INVALID_STATUS",
				},
				wantStatus: http.StatusBadRequest,
				wantErr:    true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)

				jsonInput, err := json.Marshal(tt.input)
				require.NoError(t, err)

				c.Request = httptest.NewRequest(
					http.MethodPost,
					"/api/jobs",
					bytes.NewBuffer(jsonInput),
				)
				c.Request.Header.Set("Content-Type", "application/json")

				// Add auth token
				c.Request.Header.Set("Authorization", "test-token")

				s := NewTestServer(t)
				s.CreateJobHandler(c)

				assert.Equal(t, tt.wantStatus, w.Code)
				if !tt.wantErr {
					var response models.Job
					err := json.Unmarshal(w.Body.Bytes(), &response)
					require.NoError(t, err)
					assert.NotEmpty(t, response.ID)
					assert.Equal(t, tt.input.Name, response.Name)
				}
			})
		}
	})
	t.Run("UpdateJobStatusHandler", func(t *testing.T) {
		tests := []struct {
			name  string
			jobID string
			input struct {
				Status  models.JobStatus `json:"status"`
				Version int              `json:"version"`
			}
			setupFn    func(*testing.T, *Server) // Setup function to prepare test data
			wantStatus int
			wantErr    bool
		}{
			{
				name:  "Valid Status Update",
				jobID: "1",
				input: struct {
					Status  models.JobStatus `json:"status"`
					Version int              `json:"version"`
				}{
					Status:  models.Applied,
					Version: 1,
				},
				setupFn: func(t *testing.T, s *Server) {
					// Create a test job in Wishlist status
					// Implementation needed
				},
				wantStatus: http.StatusOK,
				wantErr:    false,
			},
			{
				name:  "Invalid Status Transition",
				jobID: "1",
				input: struct {
					Status  models.JobStatus `json:"status"`
					Version int              `json:"version"`
				}{
					Status:  models.Offer,
					Version: 1,
				},
				setupFn: func(t *testing.T, s *Server) {
					// Create a test job in Wishlist status
				},
				wantStatus: http.StatusBadRequest,
				wantErr:    true,
			},
			// Add more test cases
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				s := NewTestServer(t)
				if tt.setupFn != nil {
					tt.setupFn(t, s)
				}

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)

				jsonInput, err := json.Marshal(tt.input)
				require.NoError(t, err)

				c.Request = httptest.NewRequest(
					http.MethodPatch,
					"/api/jobs/"+tt.jobID+"/status",
					bytes.NewBuffer(jsonInput),
				)
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request.Header.Set("Authorization", "test-token")
				c.Params = []gin.Param{{Key: "id", Value: tt.jobID}}

				s.UpdateJobStatusHandler(c)

				assert.Equal(t, tt.wantStatus, w.Code)
			})
		}
	})
}

func NewTestServer(t *testing.T) *Server {
	db := setupTestDB(t)
	return &Server{
		port: 8181,
		db:   db,
	}
}

func setupTestDB(t *testing.T) *sqlx.DB {
	// Create an in-memory SQLite database for testing
	db, err := sqlx.Connect("sqlite3", ":memory:")
	require.NoError(t, err)

	// Run migrations
	err = database.RunMigrations(db)
	require.NoError(t, err)

	return db
}

// Add test utilities for auth
func addTestAuth(req *http.Request) {
	// Add test JWT token
	token := createTestToken()
	req.Header.Set("Authorization", "Bearer "+token)
}

func createTestToken() string {
	// Create a test JWT token
	// Implementation needed
	return "test-token"
}
