#!/bin/bash
diff $1 $2
rv=$?  
if [[ $rv == 1 ]]  
then    
    echo "Error file contents are different"
    exit 1  
fi
