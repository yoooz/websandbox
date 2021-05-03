DROP USER IF EXISTS dev;
CREATE USER dev IDENTIFIED WITH mysql_native_password BY 'dev';
DROP SCHEMA IF EXISTS Dev;
CREATE schema Dev;
/* devユーザーにDevDBの権限を割り当てる */
GRANT all ON Dev.* TO dev;
