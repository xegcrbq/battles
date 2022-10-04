CREATE TABLE users(
    userid bigserial primary key,
    public_address varchar(60) not null UNIQUE
);

CREATE TABLE coins(
    coinid serial2 primary key,
    ticker varchar(10) not null UNIQUE
);

CREATE TABLE balances(
    balanceid bigserial primary key,
    userid bigserial not null,
    amount bigint not null,
    coinid int2 not null,
    spent bigint not null,
    CONSTRAINT fk_user FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE,
    CONSTRAINT fk_coin FOREIGN KEY (coinid) REFERENCES coins(coinid) ON DELETE CASCADE,
    CONSTRAINT unique_user_coin_pairs unique (userid, coinid),
    CONSTRAINT positive_amount CHECK ( amount > 0 ),
    CONSTRAINT positive_spent CHECK ( spent > 0 )
);

CREATE TABLE base_balances(
    base_balance_id bigserial primary key,
    userid bigserial not null unique ,
    amount bigint not null,
    CONSTRAINT fk_user FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE,
    CONSTRAINT positive_amount CHECK ( amount > 0 )
);

INSERT INTO coins(ticker)
VALUES
    ('BTC'),
    ('ETH'),
    ('USDC'),
    ('BNB'),
    ('XRP'),
    ('BUSD'),
    ('ADA'),
    ('SOL'),
    ('DOGE'),
    ('DOT');
INSERT INTO users(public_address)
VALUES ('0x8a8cB99FBE417c2fBED13B4982e4fE1BE364d58C');
INSERT INTO users(public_address)
VALUES ('0x8a8cB99FBE417c2fBED13B4982e4fE1BE364d59C');
INSERT INTO balances(userid, amount, coinid, spent)
VALUES
    (1,1000, 1, 1000),
    (1,100000, 3, 100000),
    (1,1000 *pow(10, 8), 8, 2000);

INSERT INTO base_balances(userid, amount)
values
    (1, 100000000000),
    (2, 100000000000);