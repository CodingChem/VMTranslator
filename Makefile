assignment:
	mv ./main VMTranslator
	chmod +x ./VMTranslator

build:
	go build -o main src/app/main.go

clean: 
	rm -f ./tests/*/*.asm
	rm ./VMtranslator

#test: build
