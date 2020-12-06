# Run the game
run: build
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