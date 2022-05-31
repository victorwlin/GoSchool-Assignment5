This assignment is for the Go Microservices 1 module in the GoSchool curriculum. It is a continuation of earlier assignments and is focused on building a microservice, utilizing containers, and accessing databases.

### Database
The database is a MySQL database running in a Docker container. I took the following steps to create it.

Create MySQL container
`$ docker run --name kit_db -p 32769:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql:latest`

Connecting MySQL container with local MySQL client
`$ mysql -P 32769 --protocol=tcp -u root -p`

Create and use database
`mysql> create database kit_db;`
`mysql> use kit_db;`

Create table

	mysql> CREATE TABLE victor (`friend_name` VARCHAR(30) NOT NULL PRIMARY KEY, `group` VARCHAR(30), `desired_freq` INT, `last_contact` DATE);

Insert row
`mysql> INSERT INTO victor VALUES ('Jimmy', 'US', 30, '2022-05-23');`
`mysql> INSERT INTO victor VALUES ('Ben', 'US', 60, '2022-04-09');`
`mysql> INSERT INTO victor VALUES ('Iggy', 'SG', 14, '2022-05-20');`
`mysql> INSERT INTO victor VALUES ('Arjun', 'SG', 60, '2022-05-22');`

### REST API
The REST API allows the client to pull data on friends from the following URLs.
- http://localhost:5000/api/v1/friends - Gets data on all friends
- http://localhost:5000/api/v1/friends/{friend}
	- GET - Gets data for one friend
	- POST - Creates a new friend
	- PUT - Updates details for one friends
	- DELETE - Deletes a friend

Various response codes are used to communicate the status of the request.

### Client
The client is a simplified version of previous assignments. There are now five paths to interact with the API.
- "/" - Gets data on all friends and displays a table
- "/friend/" - Gets data on a single friend and displays the information
- "/addfriend/" - Gathers necessary input from the user to create a new friend and makes a POST request to the API
- "/editfriend/" - Gathers necessary input from the user to edit details for a friend and makes a PUT request to the API
- "deletefriend/" - Deletes friend by making a DELETE request to the API

### Setup Guide
1. Follow the steps in the Database section to set up your own MySQL container with a table. (Pay particular attention to the external port as the app requires the use of 32769 to run.)
2. Go into the api folder and use `go run .` to start the REST API
3. Go into the client folder and use `go run .` to star the client
4. Open your browser and go to [localhost:5221/](localhost:5221/) to use the client

### Testing Guide
1. When you navigate to [localhost:5221/](localhost:5221/), a GET request is made to the API to fetch data on all friends in the database. Check to make sure all the friends and information is identical to what you entered in the database during setup.
2. Each friend in the table is clickable and will lead to another page that displays just information for that one friend. Each click sends a GET request to the API for one friend.
3. From the main page, click on Add Friend to navigate to a new page. Try submitting the form without all fields filled out, and you will receive an error. Try submitting a friend name with spaces, and you will get an error. Try submitting a friend name that already exists, and you will get an error. After successfully adding a friend, you will be directed back to the main page. Check to make sure all of the details you entered for that friend are correct.
4. From the Friend Details page, click on Edit Friend to navigate to that page. Only enter data on the fields that you want to edit. A successful edit will redirect you to the Friend Details page.
5. From the Friend Details page, click on Delete Friend to delete that friend. A successful delete will redirect you to the main page.
