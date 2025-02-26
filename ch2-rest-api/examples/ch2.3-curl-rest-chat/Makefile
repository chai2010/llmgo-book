API_URL:=http://localhost:11434/v1/chat/completions
JSON_FILE:=request.json

chat1:
	curl -X POST ${API_URL} -H "Content-Type: application/json" -d @request.json

chat2:
	curl -X POST ${API_URL} -H "Content-Type: application/json" -d @request2.json

chat3:
	curl -X POST ${API_URL} -H "Content-Type: application/json" -d @request3.json

clean:
