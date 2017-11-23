#!/bin/sh
set -e
set -x

make_image(){
	echo "Building... $1 image"
	cd $2
	docker build -t $1:latest .
	echo "Successfully build $1 image"
}

BUILD_PATH=$(cd $(dirname $0);pwd)
booking=$BUILD_PATH/booking
movies=$BUILD_PATH/movies
users=$BUILD_PATH/users
showtimes=$BUILD_PATH/showtimes
cd $BUILD_PATH
make_image "booking" $booking
make_image "movies" $movies
make_image "users" $users
make_image "showtimes" $showtimes
echo '=================================================================='
echo '                     BUILD SUCCESSFULL                            '
echo '=================================================================='

