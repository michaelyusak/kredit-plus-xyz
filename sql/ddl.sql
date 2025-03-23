CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    identity_number VARCHAR NOT NULL,
    full_name VARCHAR NOT NULL,
    legal_name VARCHAR NOT NULL,
    place_of_birth VARCHAR NOT NULL,
    date_of_birth VARCHAR NOT NULL,
    salary BIGINT NOT NULL,
    identity_card_photo_url VARCHAR NOT NULL,
    selfie_photo_url VARCHAR NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at BIGINT DEFAULT NULL
);

CREATE TABLE transactions (
    transaction_id BIGSERIAL PRIMARY KEY NOT NULL,
    otr_price VARCHAR NOT NULL,
    admin_fee INT NOT NULL,
    installment_amount VARCHAR NOT NULL,
    interest_amount DECIMAL NOT NULL,
    asset_name VARCHAR NOT NULL,
    created_at BIGINT NOT NULL,
    updated_at BIGINT NOT NULL,
    deleted_at BIGINT DEFAULT NULL
);