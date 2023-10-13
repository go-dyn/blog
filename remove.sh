#!/bin/bash
regex=test
for b in $(git branch -r | grep $regex)
do
    echo $b
done