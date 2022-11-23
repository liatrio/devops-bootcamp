#!/bin/zsh

if [[ $1 == "post" ]];then
	curl http://localhost:8080/engineers \
    		--include \
    		--header "Content-Type: application/json" \
    		--request "POST" \
    		--data '{"name": "Ryan","email": "ryan@ferrets.com"}'

	curl http://localhost:8080/engineers \
      		--include \
      		--header "Content-Type: application/json" \
      		--request "POST" \
      		--data '{"name": "zach", "email": "zach@bengal.com"}'
elif [[ $1 == "put" ]]; then
	curl http://localhost:8080/engineers/${2} \
        	--include \
        	--header "Content-Type: application/json" \
        	--request "PUT" \
        	--data '{"name":"'"$3"'", "email": "evan@ferrets.com"}'
elif [[ $1 == "delete" ]]; then
	curl http://localhost:8080/engineers/${2} \
        	 --include \
         	--header "Content-Type: application/json" \
         	--request "DELETE"
elif [[ $1 == "dev" ]]; then
	curl http://localhost:8080/dev/${3} \
          	--include \
          	--header "Content-Type: application/json" \
          	--request "POST" \
          	--data '{"id":"'"$2"'"}'
elif [[ $1 == "op" ]]; then
	curl http://localhost:8080/op/${3} \
          	--include \
          	--header "Content-Type: application/json" \
          	--request "POST" \
          	--data '{"id":"'"$2"'"}'
else
	curl http://localhost:8080/engineers
fi
