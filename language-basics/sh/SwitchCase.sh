#!/bin/bash

echo "Input an integer between 1 and 5"
read INPUT

case $INPUT in
"1")
	echo "Your number is 1"
	exit 0
;;
"2")
        echo "Your number is 2"
        exit 0
;;
"3")
        echo "Your number is 3"
        exit 0
;;
"4")
        echo "Your number is 4"
        exit 0
;;
"5")
        echo "Your number is 5"
        exit 0
;;
*)
	echo "Invalid input"
	exit 1
esac

exit 1
