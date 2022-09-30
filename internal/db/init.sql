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
    CONSTRAINT fk_user FOREIGN KEY (userid) REFERENCES users(userid) ON DELETE CASCADE,
    CONSTRAINT fk_coin FOREIGN KEY (coinid) REFERENCES coins(coinid) ON DELETE CASCADE,
    CONSTRAINT unique_user_coin_pairs unique (userid, coinid),
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
INSERT INTO balances(userid, amount, coinid)
VALUES
    (1,1000, 1);
INSERT INTO balances(userid, amount, coinid)
VALUES
    (1,100000, 3);
-- INSERT INTO balances(userid, amount, coinid)
-- VALUES
--     (1,-20, 1);