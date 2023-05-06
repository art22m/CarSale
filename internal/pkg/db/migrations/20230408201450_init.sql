-- +goose Up
-- +goose StatementBegin
CREATE TABLE car_sale
(
    id         BIGSERIAL PRIMARY KEY                  NOT NULL,
    brand      TEXT                                   NOT NULL,
    model      TEXT                                   NOT NULL,

    seller_id  BIGINT                                 NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE               NULL
);

CREATE TABLE seller
(
    id           BIGSERIAL PRIMARY KEY                  NOT NULL,
    name         TEXT                                   NOT NULL,
    phone_number TEXT                                   NOT NULL,

    created_at   TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
    updated_at   TIMESTAMP WITH TIME ZONE               NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE seller;
DROP TABLE car_sale;

-- +goose StatementEnd
