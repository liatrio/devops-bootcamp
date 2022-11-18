#!/bin/zsh

if [[ $1 == "post" ]];then
	curl http://localhost:8080/dev \
    		--include \
    		--header "Content-Type: application/json" \
    		--request "POST" \
    		--data '{"name": "dev_ferrets"}'

	curl http://localhost:8080/dev \
      		--include \
      		--header "Content-Type: application/json" \
      		--request "POST" \
      		--data '{"name": "dev_bengal"}'
elif [[ $1 == "put" ]]; then
	curl http://localhost:8080/dev/${2} \
        	--include \
        	--header "Content-Type: application/json" \
        	--request "PUT" \
        	--data '{"name":"'"$3"'"}'
elif [[ $1 == "delete" ]]; then
	curl http://localhost:8080/dev/${2} \
        	 --include \
         	--header "Content-Type: application/json" \
         	--request "DELETE"
elif [[ $1 == "devops" ]]; then
	curl http://localhost:8080/devops/dev/${3} \
          	--include \
          	--header "Content-Type: application/json" \
          	--request "POST" \
          	--data '{"id":"'"$2"'"}'
else
	curl http://localhost:8080/dev
fi
