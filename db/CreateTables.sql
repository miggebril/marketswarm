use Markets

CREATE TABLE Traders (
	TraderID int NOT NULL IDENTITY(1,1) PRIMARY KEY,
	UserName varchar(50) NOT NULL, 
	Email varchar(50) NOT NULL,
	Password varbinary(128),
	IsVerified bit DEFAULT 0
)

INSERT INTO Traders (UserName, Email, IsVerified) values ('miggebril', 'migdam.gebril@gmail.com', 1), ('test', 'test@example.com', 0)

CREATE TABLE Exchanges (
	ExchangeID int NOT NULL IDENTITY(1, 1) PRIMARY KEY,
	Name varchar(50) NOT NULL
)

INSERT INTO Exchanges (Name) values ('ICE'),( 'CME')

CREATE TABLE Products (
	ProductID int NOT NULL IDENTITY(1, 1) PRIMARY KEY,
	Symbol varchar(50) NOT NULL,
	ExchangeCode varchar(50) NOT NULL,
	ExchangeID int NOT NULL
)

 INSERT INTO Products (Symbol, ExchangeCode, ExchangeID) values ('CL', 'Crude Oil', 2), ('BRN', 'IPE e-Brent', 1)

 CREATE TABLE MonthYears (
	Month int NOT NULL,
	Year int NOT NULL,
	MMMYY varchar(5) NOT NULL
 )