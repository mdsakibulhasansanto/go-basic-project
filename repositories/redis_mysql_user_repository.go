package repositories

import (
	"encoding/json"
	"final-project/cache"
	"final-project/models"
	"time"
)

type RedisMySQLUserRepository struct {
	repo  UserRepositoryInterface
	cache cache.Cache
}

func NewRedisMySQLUserRepository(repo UserRepositoryInterface, cache cache.Cache) *RedisMySQLUserRepository {
	return &RedisMySQLUserRepository{repo: repo, cache: cache}
}

func (r *RedisMySQLUserRepository) Create(user models.User) error {
	return r.repo.Create(user)
}

func (r *RedisMySQLUserRepository) GetByEmail(email string) (*models.User, error) {
	cacheKey := "user:" + email
	cached, err := r.cache.Get(cacheKey)
	if err == nil {
		var user models.User
		if err := json.Unmarshal([]byte(cached), &user); err == nil {
			return &user, nil
		}
	}

	user, err := r.repo.GetByEmail(email)
	if err != nil || user == nil {
		return nil, err
	}

	userJson, _ := json.Marshal(user)
	r.cache.Set(cacheKey, string(userJson), 10*time.Minute)

	return user, nil
}

// UpdateUserByID updates user data in the base repository and invalidates the cache
func (r *RedisMySQLUserRepository) UpdateUserByID(user *models.User) error {
	err := r.repo.UpdateUserByID(user)
	if err == nil {
		cacheKey := "user:" + user.Email
		r.cache.Set(cacheKey, "", 0) // Invalidate cache (could also delete)
	}
	return err
}

// DeleteUser deletes user from the base repository and invalidates the cache
func (r *RedisMySQLUserRepository) DeleteUser(email string) error {
	err := r.repo.DeleteUser(email)
	if err == nil {
		cacheKey := "user:" + email
		r.cache.Set(cacheKey, "", 0) // Invalidate cache (could also delete)
	}
	return err
}
