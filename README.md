# Microservices for Mesher Demo
Mesher demo is an example project which demonstrates the use of microservices for a fictional movie theater. Its backend is powered by 4 microservices, all of which happen to be written in Go.
* Movie Service: Provides information like movie ratings, title, etc.
* Show Times Service: Provides show times information.
* Booking Service: Provides booking information.
* Users Service: Provides booking information from booking service for users and movie rating etc. from the Movie service by communicating with other services.

#### Requirements:
1.	Golang
2.	Docker
 
### API and Documentation:
**1. Movie Service(Port 5001)**
* GET   /movies  
To look up all movies. Returns a list of all movies.
It will Return a list of all movies.
```
{"data":[{"id":"8y8y888y808","title":"captain america","director":"christopher","rating":9,"createdon":"0001-01-01T00:00:00Z"},{"id":"3y8473743740973","title":"indiana jones","director":"steven","rating":9,"createdon":"0001-01-01T00:00:00Z"}]}
```
* POST /movies
To create a new one 
```
Curl  -X POST localhost:5001/movies –d ‘{ "data": {
		"id" : "3y8473743740973","title": "indiana jones",
		"director": "steven","rating": 9 } }’
```
*	GET  /movies/{id}
*	DELETE /movies/{id}
		

**2. Booking Servcie(Port 5003)**
* GET   /bookings
To look up Booking information of all users.Returns a list of booking information.
```
{"data":[{"userid":"844845885","showtimeid":"4376498348938","movies":["8y8y888y808","3y8473743740973"]}]}
```
*	POST /booking
Create a new booking information
```
Curl  -X POST localhost:5003/booking –d ’{"data" : {"userid" :"844845885","showtimeid":"4376498348938","movies": ["8y8y888y808","3y8473743740973"]}} ‘
```
* 	GET /bookings/{userid}

**3. ShowTimes Service(Port 5002)**
*	GET   /showtimes
To lookup all showtimes.It will return a list of all showtimes.
```
{"data":[{"id":"4376498348938","date":"25-10-2004","createdon":"0001-01-01T00:00:00Z","movies":["223243fe","dkjlwel3"]}]}
```
*	POST  /showtimes
To create a new showtime
```
Curl  -X POST localhost:5002/showtimes –d ‘{"data" : {
		"id": "4376498348938",
		"date": "25-10-2004",
		"movies": ["223243fe","dkjlwel3"]
	}}’
```
*	GET    /showtimes/{id}
*	DELETE /showtimes/{id}
	

**4.	User service(port 5000)**
*	GET     /users
To lookup all users stored in menmory. Returns a list of users.
```
{"data":[{"id":"844845885","name":"oshank","lastname":"kumar"}]}
```
*	POST /users
Create a new user.
```
curl –X POST localhost:5000/users –d ‘{"data" : {
		"id" : "844845885",
		"name": "oshank",
		"lastname": "kumar"}}’
```
*	DELETE   /user/{id}
*	GET /user/{id}/bookings
It will communicates with other 3 microservices to retrieve booking and movie information.
```
{
    "user": {
        "id": "844845885",
        "name": "oshank",
        "lastname": "kumar"
    },
    "booking": {
        "userid": "844845885",
        "showtimeid": "4376498348938",
        "movies": [
            "8y8y888y808",
            "3y8473743740973"
        ]
    },
    "showtime": {
        "id": "4376498348938",
        "date": "25-10-2004",
        "createdon": "0001-01-01T00:00:00Z",
        "movies": [
            "223243fe",
            "dkjlwel3"
        ]
    },
    "movies": [
        {
            "id": "8y8y888y808",
            "title": "captain america",
            "director": "christopher",
            "rating": 9,
            "createdon": "0001-01-01T00:00:00Z"
        },
        {
            "id": "3y8473743740973",
            "title": "indiana jones",
            "director": "steven",
            "rating": 9,
            "createdon": "0001-01-01T00:00:00Z"
        }
    ]
}
```

### 2.	Create Docker Image 
go get github.com/oshankfriends/mesher-demo  
cd  mesher-demo/booking  
docker build –t booking:latest .  
cd mesher-demo/movies  
docker build –t movies:latest .  
cd mesher-demo/showtimes  
docker build –t showtimes:latest .  
cd mesher-demo/users  
docker build –t users:latest .  

### Run
* **In a docker container**  
example :- docker run -d --net=host booking:latest  
* **Using Binaries**  
    In order to build the services locally make sure the repository directory located in correct $GOPATH  
    For Example :-   
    directory : $GOPATH/src/github.com/oshankfriends/mesher-demo/booking  
    cd $GOPATH/src/github.com/oshankfriends/mesher-demo/booking  
    go build -o booking  
    The result is a binary *booking* in current directory  
    ./booking
	

 
