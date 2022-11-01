#!/bin/bash

set -e

day=$1

if [[ $day = "" ]]; then
  echo "Specify day with argument"
  exit 1
fi

cp -R dayX day$1

cd day${day}

sed -e "s/dayX/day${day}/g" -e "s/DayX/Day${day}/g" <dayX.go >day${day}.go
sed -e "s/dayX/day${day}/g" -e "s/DayX/Day${day}/g" <test.go >day${day}_test.go

rm dayX.go
rm test.go

cd ..





