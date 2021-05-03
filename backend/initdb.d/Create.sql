DROP USER IF EXISTS shop;
CREATE USER shop IDENTIFIED WITH mysql_native_password BY 'shop';
DROP SCHEMA IF EXISTS Shop;
CREATE schema Shop;
/* shopユーザーにShopDBの権限を割り当てる */
GRANT all ON Shop.* TO shop;
