CREATE TABLE users(
    userid bigserial primary key ,
    public_address varchar(100) UNIQUE
);

CREATE TABLE balances(
    balanceid bigserial primary key ,
    userid bigserial,
    amount bigint,
    coinid bigint,
    CONSTRAINT fk_user FOREIGN KEY (userid) REFERENCES users(userid)
);
