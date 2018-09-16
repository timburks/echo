all:	
	mkdir -p generated
	protoc echo.proto --go_out=generated
	go run main.go

clean:
	rm -rf echo generated
