# goForum
A monolithic web application written in golang using only standard library plugins for learning purposes. This application uses the net/http multiplexer to serve html templates and sql package to connect to a PostgreSQL database. This app was written using concepts described in the book ["Go Web Programming"](https://www.amazon.com/dp/1617292567/ref=cm_sw_em_r_mt_dp_U_Sm9kCbDEZ9BH2) by Sau Sheong Chang.

## Getting Started
1. Clone repo anywhere you like.
2. Make sure Docker is installed and properly configured
3. Build and Run the app `docker-compose build && docker-compose up`
4. See it in action in your local browser at [127.0.0.1:8080](http://127.0.0.1:8080)