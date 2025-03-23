CREATE TABLE users (
    user_id BIGSERIAL PRIMARY KEY,
    identity_number VARCHAR,
    full_name VARCHAR,
    legal_name VARCHAR,
    place_of_birth VARCHAR,
    date_of_birth VARCHAR,
    salary BIGINT,
    identity_card_photo_url VARCHAR,
    selfie_photo_url VARCHAR,
    created_at BIGINT,
    updated_at BIGINT,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE transactions (
    transaction_id BIGSERIAL PRIMARY KEY,
    otr_price VARCHAR,
    admin_fee INT,
    installment_amount VARCHAR,
    interest_amount DECIMAL,
    asset_name VARCHAR,
    created_at BIGINT,
    updated_at BIGINT,
    deleted_at BIGINT
);