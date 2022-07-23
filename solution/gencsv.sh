#!/bin/bash
input=$@
if [ $input ]
then
number=$input
else
number=10
fi


i=0
for n in `seq "$number"`
do
randomNumber=$RANDOM
echo $i, $randomNumber
i=$(( i + 1 ))
done
