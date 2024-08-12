build:
	go build
	chmod +x ./VMTranslator

clean: 
	rm -f ./tests/*/*.asm
	rm ./VMtranslator

test: build
	./VMTranslator tests/BasicTest/BasicTest.vm
	./VMTranslator tests/PointerTest/PointerTest.vm
	./VMTranslator tests/SimpleAdd/SimpleAdd.vm
	./VMTranslator tests/StackTest/StackTest.vm
	./VMTranslator tests/StaticTest/StaticTest.vm
	

