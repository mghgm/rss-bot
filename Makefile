default: test build


test:
	@echo "Testing..."
	
clean:
	@echo "Cleaning..."
	rm -rf ./out

build:
	@mkdir -p ./out
	go build -o ./out/rss-bot rss_bot.go 
