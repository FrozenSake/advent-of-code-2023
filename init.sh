git init .

for ((i=1; i<=30; i++)); do
	mkdir day$i
	cd day$i
	go mod init frozensake/aoc2023day$i
	cp ../main.go .
	go mod tidy
	cd ..
done