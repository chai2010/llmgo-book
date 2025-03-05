API_URL:=http://localhost:11434/v1/chat/completions
JSON_FILE:=request.json

chat:
	curl -X POST ${API_URL} -H "Content-Type: application/json" -d @request.json

clean:
