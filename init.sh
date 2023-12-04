git init .

for ((i=1; i<=30; i++)); do
	if [$i -lt 10]; then
		mkdir day0$i
		cd day0$i
	else 
		mkdir day$i
		cd day$i
	fi
	go mod init frozensake/aoc2023day$i
	cp ../main.go .
	go mod tidy
	cd ..
done