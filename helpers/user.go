package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	i18n "github.com/azizka85/azizka-go-i18n"
	"github.com/azizka85/azizka-go-my-routes/data"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func GetUserIdByEmail(
	email string,
	user *data.User,
	db *gorm.DB,
) (int64, error) {
	result := db.
		Select("id").
		Where(&data.User{Email: email}).
		Take(user)

	return result.RowsAffected, result.Error
}

func GetUserInfoById(
	userId int,
	user *data.User,
	db *gorm.DB,
) (int64, error) {
	result := db.
		Select("full_name", "photo").
		Take(user, userId)

	return result.RowsAffected, result.Error
}

func LogIn(
	user *data.User,
	db *gorm.DB,
) (int64, error) {
	hash := md5.Sum([]byte(user.Password))
	password := hex.EncodeToString(
		hash[:],
	)

	result := db.
		Select("id", "full_name", "photo").
		Where(&data.User{
			Email:    user.Email,
			Password: password,
		}).
		Take(user)

	return result.RowsAffected, result.Error
}

func SignIn(
	user *data.User,
	translator *i18n.Translator,
	session *sessions.Session,
	db *gorm.DB,
) error {
	count, err := LogIn(user, db)

	if err != nil || count == 0 {
		return data.CreateActionError(
			translator.Translate(
				"User with this email and password doesn't exist",
			),
		)
	}

	session.Values["userId"] = user.ID
	delete(session.Values, "service")
	delete(session.Values, "oauth")

	return nil
}

func CreateUser(
	user *data.User,
	db *gorm.DB,
) (int64, error) {
	password := user.Password
	hash := md5.Sum([]byte(password))

	user.Password = hex.EncodeToString(
		hash[:],
	)

	result := db.Create(user)

	user.Password = password

	return result.RowsAffected, result.Error
}

func SignUp(
	user *data.User,
	translator *i18n.Translator,
	session *sessions.Session,
	db *gorm.DB,
) error {
	user.FullName = strings.TrimSpace(user.FullName)
	user.Email = strings.TrimSpace(user.Email)

	if user.FullName == "" {
		return data.CreateActionError(
			translator.Translate(
				"Name required",
			),
		)
	} else if user.Email == "" {
		return data.CreateActionError(
			translator.Translate(
				"Email required",
			),
		)
	} else if user.Password == "" {
		return data.CreateActionError(
			translator.Translate(
				"Password required",
			),
		)
	} else {
		count, _ := GetUserIdByEmail(user.Email, &data.User{}, db)

		if count > 0 {
			return data.CreateActionError(
				translator.Translate(
					"User with this email already exists",
				),
			)
		} else {
			_, err := CreateUser(user, db)

			if err != nil {
				return err
			}

			err = SignIn(
				user,
				translator,
				session,
				db,
			)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GetUserInfoFromSession(
	user *data.User,
	session *sessions.Session,
	db *gorm.DB,
) (int64, error) {
	service, ok := session.Values["service"].(string)
	data, dataOk := session.Values["oauth"].(map[string]map[string]interface{})

	service = strings.TrimSpace(service)

	if ok && service != "" && dataOk && data != nil {
		switch service {
		case "github":
			GetUserInfoFromGithub(user, data)
			return 1, nil
		}
	} else {
		userId, _ := session.Values["userId"].(int)

		return GetUserInfoById(
			userId,
			user,
			db,
		)
	}

	return 0, nil
}

func GetUserInfoFromGithub(
	user *data.User,
	data map[string]map[string]interface{},
) {
	fullName, _ := data["user"]["name"].(string)
	photo, _ := data["user"]["avatar_url"].(string)

	user.FullName = fullName
	user.Photo = photo
}

func SignOut(session *sessions.Session) {
	delete(session.Values, "userId")
	delete(session.Values, "service")
	delete(session.Values, "oauth")
}

func OAuth(
	session *sessions.Session,
	service string,
	data map[string]map[string]interface{},
) {
	session.Values["service"] = service
	session.Values["oauth"] = data
}
