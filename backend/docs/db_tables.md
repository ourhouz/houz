# Postgres Table Designs

## Houses/Users

### Houses
A house contains users/housemates and pay periods. 

```postgresql
CREATE TABLE Houses (
    HouseId         SERIAL,
    HouseName       VARCHAR(100),
    PRIMARY KEY (HouseId),
);
```

### Users
A user is a single member of a house(s). A user cannot be in multiple houses (a user is created after joining a house). 

```postgresql
CREATE TABLE Users (
    HouseId         INTEGER,
    Username        VARCHAR(100),
    -- and other authentication related columns
    PRIMARY KEY (HouseId, Username),
    FOREIGN KEY (HouseId) REFERENCES Houses (HouseId),
);
```

## Payments

### PayPeriods
A pay period contains a list of pay entries within a specific timeframe. 
A pay period starts and ends when balances are zeroed.
In practice, it is possible for users to not use pay periods (ie. never zero balances). 

```postgresql
CREATE TABLE PayPeriods (
    PayPeriodId     SERIAL,
    HouseId         INTEGER,
    StartTime       TIMESTAMPTZ,
    EndTime         TIMESTAMPTZ,
    Completed       BOOLEAN,
    PRIMARY KEY (PayPeriodId),
    FOREIGN KEY (HouseId) REFERENCES Houses (HouseId),
);
```

### PayEntries
A pay entry contains metadata on a single payment. 

```postgresql
CREATE TABLE PayEntries (
    PayEntryId      SERIAL,
    PayPeriodId     INTEGER,
    StartTime       TIMESTAMPTZ,
    Location        VARCHAR(100),
    Description     VARCHAR(500),
    TotalCost       REAL,
    PRIMARY KEY (PayEntryId),
    FOREIGN KEY (PayPeriodId) REFERENCES PayPeriods (PayPeriodId),
);
```

### PayItems
A pay item is a single item within a pay entry payment. 
It contains individualized cost data. 

```postgresql
CREATE TABLE PayItems (
    PayItemId       SERIAL,
    PayEntryId      INTEGER,
    Name            VARCHAR(100),
    Description     VARCHAR(500),
    CostPerUnit     REAL,
    Quantity        REAL,
    PRIMARY KEY (PayItemId),
    FOREIGN KEY (PayEntryId) REFERENCES PayEntries (PayEntryId),
);
```

### PayDues
A pay due relates a pay item with a user.

```postgresql
CREATE TABLE PayDues (
    PayItemId       INTEGER,
    UserId          INTEGER,
    TotalCost       REAL,   
    FOREIGN KEY (PayItemId) REFERENCES PayItems (PayItemId),
    FOREIGN KEY (UserId) REFERENCES Users (UserId),
);
```
