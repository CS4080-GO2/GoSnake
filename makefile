# Run the game
run:
	@clear
	@go get -u github.com/CS4080-GO2/GoSnake/game
	@go run main.go


# Build the file
build:
	@clear
	@go build


# Open Project files in atom
atom:
	@clear
	@cd .. && atom Project


# Open Sample Project file in atom
sample:
	@clear
	@cd ../../ && atom SampleSnakeGame


# Open Project file and Sample Project file
all-atom: atom sample


# Reset to all file in main branch
reset:
	@clear
	@git checkout main .


run-sample:
	@clear
	@cd ../../ && cd SampleSnakeGame && go run main.go
