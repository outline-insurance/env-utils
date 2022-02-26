#!/bin/bash

diff $1 $2
rv=$?  
if [[ $rv != 1 ]]  
then    
    echo "Error versions are the same"
    exit 1  
fi
