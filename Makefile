# go install github.com/wa-lang/wabook@latest

default:
	wabook serve

gopher-base64:
	base64 -i gopher.png -o gopher-base64.txt

build:
	-rm book
	wabook build
	-rm book/.gitignore
	-rm -rf book/.git

deploy:
	-rm book
	wabook build
	-rm book/.gitignore
	-rm -rf book/.git

	cd book && git init
	cd book && git add .
	cd book && git commit -m "first commit"
	cd book && git branch -M gh-pages
	cd book && git remote add origin git@github.com:chai2010/llmgo-book.git
	cd book && git push -f origin gh-pages

clean:
	-rm -rf book
