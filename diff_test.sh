diff $1 $2
rv=$?  
if [[ $rv == 1 ]]  
then    
    echo "Error file contents are different" >&2
    exit 1  
fi
