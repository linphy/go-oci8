create table
(
      id            NUMBER(10) not null,
      name          VARCHAR2(20) not null,
      create_date   DATE default sysdate not null,
      birthday      DATE,
      address       VARCHAR2(200),
      email         VARCHAR2(200),
      mobilephone   VARCHAR2(11),
      telephone     VARCHAR2(20),
      identity_card VARCHAR2(18),
      weight        NUMBER,
      height        NUMBER
);
