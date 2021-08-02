#!/bin/bash

echo "Do you:"
echo "1: Like bash"
echo "2: Not like bash"

read SELECTION

if [ $SELECTION == "1" ]
then

	echo "Awesome!"
	exit 0

elif [ $SELECTION == "2" ]
then
	echo "Thaat's lame!"
	exit 0

else
	echo "That is not an option!"
	exit 1
fi


exit 1
