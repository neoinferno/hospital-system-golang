# hospital-system-golang

Run docker compose -f docker-compose.yml up to start project

Run sql-script in configs/init.sql

Copy curl --location 'localhost:8888/api/staffs/login' \
--header 'Content-Type: application/json' \
--data '{
"username":"admin",
"password":"1234",
"hospitalId":"eb83e5fb-2662-4edf-b7ec-4f7b8cd6feb5"
}' to postman for login
