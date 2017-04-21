#!/bin/bash
arrName=("a" "b" "c" "d")
for i in ${arrName[@]};do
echo $i
done

echo ${arrName[0]}

arrName[4]="e"

echo ${arrName[4]}
