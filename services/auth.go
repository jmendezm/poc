package services

import (
	"technical_test/domain"
	"technical_test/memory_db"
	"technical_test/postgres_connection"
	"technical_test/utils"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

var auth_service = &AuthService{}

type AuthService struct {
}

func GetAuthServiceInstance() *AuthService {
	if auth_service == nil {
		auth_service = &AuthService{}
	}
	return auth_service
}

func (as *AuthService) CreateUser(user *domain.User) (*domain.User, error) {
	passHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = passHash
	user.RecordStatus = domain.RecordStatusActive
	user.Active = domain.RecordStatusActive
	user.CreateDate = time.Now()
	user.LastUpdated = time.Now()
	res := postgres_connection.Get().Create(&user)
	if res.Error != nil {
		return nil, &domain.ErrInternalServerError
	}
	return user, nil
}

func (as *AuthService) GetUserByID(userID int64) (*domain.User, error) {
	user := domain.User{UserId: userID}
	res := postgres_connection.Get().First(&user)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, &domain.ErrUserNotFound
		}
		return nil, &domain.ErrInternalServerError
	}
	return &user, nil
}

func (as *AuthService) UpdateUser(newUserInfo *domain.User) (*domain.User, error) {
	user := &domain.User{
		UserId: newUserInfo.UserId,
	}
	pgConn := postgres_connection.Get()
	err := pgConn.Transaction(func(tx *gorm.DB) error {
		res := pgConn.Find(&user)
		if res.Error != nil {
			if res.Error == gorm.ErrRecordNotFound {
				return &domain.ErrUserNotFound
			}
			return &domain.ErrInternalServerError
		}
		user.AccountId = newUserInfo.AccountId
		user.Address = newUserInfo.Address
		user.CompanyName = newUserInfo.CompanyName
		user.FirstName = newUserInfo.FirstName
		user.LastName = newUserInfo.LastName
		user.Email = newUserInfo.Email
		user.Phone = newUserInfo.Phone
		user.EmergencyPhone = newUserInfo.EmergencyPhone
		user.I18N = newUserInfo.I18N
		user.Identification = newUserInfo.Identification
		user.LastUpdated = time.Now()
		res = pgConn.Save(&user)
		if res.Error != nil {
			log.Error(res.Error.Error())
			return &domain.ErrInternalServerError
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (as *AuthService) Login(email string, password string) (string, error) {
	pgConn := postgres_connection.Get()
	user := domain.User{}
	res := pgConn.Where("email = ?", email).First(&user)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return "", &domain.ErrWrongCredentials
		}
		return "", &domain.ErrInternalServerError
	}
	if ok := utils.CheckPasswordHash(password, user.Password); !ok {
		return "", &domain.ErrWrongCredentials
	}
	if user.Active != domain.RecordStatusActive {
		return "", &domain.ErrAccountUnAvailable
	}
	conn := domain.Connection{}
	res = pgConn.Where("user_id = ?", user.UserId).First(&conn)
	if res.Error != nil {
		if res.Error != gorm.ErrRecordNotFound {
			log.Error(res.Error.Error())
			return "", &domain.ErrInternalServerError
		}
		conn.UserId = user.UserId
		conn.ConnectionID = utils.GenerateID()
		conn.Active = user.Active
		conn.RecordStatus = user.RecordStatus
		conn.AccountId = user.AccountId
		conn.UserData = user.GenerateUserData()
		conn.CreateDate = time.Now()
	}
	conn.LastUpdated = time.Now()
	conn.Connected = time.Now().Local().String()
	res = pgConn.Save(&conn)
	if res.Error != nil {
		log.Error(res.Error.Error())
		return "", &domain.ErrInternalServerError
	}
	memory_db.SetConnection(conn.ConnectionID, &conn)
	return conn.ConnectionID, nil
}

func (as *AuthService) DeleteUser(userID int64) error {
	pgConn := postgres_connection.Get()
	err := pgConn.Transaction(func(tx *gorm.DB) error {
		user := domain.User{UserId: userID}
		res := tx.First(&user)
		if res.Error != nil {
			if res.Error == gorm.ErrRecordNotFound {
				return &domain.ErrUserNotFound
			}
			return &domain.ErrInternalServerError
		}
		user.RecordStatus = domain.RecordStatusDeleted
		user.Active = domain.RecordStatusDeleted
		res = pgConn.Save(&user)
		if res.Error != nil {
			log.Error(res.Error.Error())
			return &domain.ErrInternalServerError
		}
		conn := domain.Connection{}
		if res = pgConn.Where("user_id = ?", user.UserId).Delete(&conn); res.Error != nil {
			return &domain.ErrInternalServerError
		}
		return nil
	})
	return err
}

func (as *AuthService) CheckConnection(connectionID string) error {
	conn := memory_db.GetConnectionByID(connectionID)
	pgConn := postgres_connection.Get()
	if conn == nil {
		conn = &domain.Connection{}
		res := pgConn.Where("connection_id = ?", connectionID).First(&conn)
		if res.Error != nil {
			if res.Error == gorm.ErrRecordNotFound {
				return &domain.ErrNotLoggedIn
			}
		}
		log.Error(res.Error.Error())
		return &domain.ErrInternalServerError
	}
	if conn.Connected == "" {
		return &domain.ErrNotLoggedIn
	}
	return nil
}

func (as *AuthService) GetConnectionByID(connectionID string) (*domain.Connection, error) {
	conn := memory_db.GetConnectionByID(connectionID)
	pgConn := postgres_connection.Get()
	if conn == nil {
		conn = &domain.Connection{}
		res := pgConn.Where("connection_id = ?", connectionID).First(&conn)
		if res.Error != nil {
			if res.Error == gorm.ErrRecordNotFound {
				return nil, nil
			}
			log.Error(res.Error.Error())
			return nil, &domain.ErrInternalServerError
		}
	}
	memory_db.SetConnection(conn.ConnectionID, conn)
	return conn, nil
}
