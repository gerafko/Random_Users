package postgres

import (
	"test/internal/model"
	"test/internal/storage"
)

type UserRepository struct {
	db *storage.Postgres
}

func NewUserRepository(db *storage.Postgres) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) InsertList(users []model.User) error {
	tx, err := u.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	stmt, err := tx.PrepareNamed(`
        INSERT INTO users (
            gender, email, phone, cell, nat,
            name_title, name_first, name_last,
            location_street_number, location_street_name,
            location_city, location_state, location_country, location_postcode,
            location_latitude, location_longitude, location_timezone_offset, location_timezone_description,
            login_uuid, login_username, login_password, login_salt, login_md5, login_sha1, login_sha256,
            dob_date, dob_age, registered_date, registered_age,
            id_name, id_value,
            picture_large, picture_medium, picture_thumbnail
        ) VALUES (
            :gender, :email, :phone, :cell, :nat,
            :name_title, :name_first, :name_last,
            :location_street_number, :location_street_name,
            :location_city, :location_state, :location_country, :location_postcode,
            :location_latitude, :location_longitude, :location_timezone_offset, :location_timezone_description,
            :login_uuid, :login_username, :login_password, :login_salt, :login_md5, :login_sha1, :login_sha256,
            :dob_date, :dob_age, :registered_date, :registered_age,
            :id_name, :id_value,
            :picture_large, :picture_medium, :picture_thumbnail
        )
    `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, user := range users {
		_, err = stmt.Exec(user)
		if err != nil {
			return err
		}
	}

	return nil
}
