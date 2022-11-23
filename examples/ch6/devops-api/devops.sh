#!/bin/zsh

if [[ $1 == "post" ]];then
	curl http://localhost:8080/devops \
    		--include \
    		--header "Content-Type: application/json" \
    		--request "POST"

elif [[ $1 == "delete" ]]; then
	curl http://localhost:8080/devops/${2} \
        	 --include \
         	--header "Content-Type: application/json" \
         	--request "DELETE"
else
	curl http://localhost:8080/devops
fi
