CREATE TABLE IF NOT EXISTS users (
                                     id SERIAL PRIMARY KEY,
                                     gender TEXT,
                                     email TEXT,
                                     phone TEXT,
                                     cell TEXT,
                                     nat TEXT,
                                     name_title TEXT,
                                     name_first TEXT,
                                     name_last TEXT,
                                     location_street_number INTEGER,
                                     location_street_name TEXT,
                                     location_city TEXT,
                                     location_state TEXT,
                                     location_country TEXT,
                                     location_postcode INTEGER,
                                     location_latitude TEXT,
                                     location_longitude TEXT,
                                     location_timezone_offset TEXT,
                                     location_timezone_description TEXT,
                                     login_uuid TEXT,
                                     login_username TEXT,
                                     login_password TEXT,
                                     login_salt TEXT,
                                     login_md5 TEXT,
                                     login_sha1 TEXT,
                                     login_sha256 TEXT,
                                     dob_date TEXT,
                                     dob_age INTEGER,
                                     registered_date TEXT,
                                     registered_age INTEGER,
                                     id_name TEXT,
                                     id_value TEXT,
                                     picture_large TEXT,
                                     picture_medium TEXT,
                                     picture_thumbnail TEXT
);