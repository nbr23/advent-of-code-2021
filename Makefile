SHELL := /usr/bin/bash
DAY=`date +%d`

build:
	@mkdir -p bin
	@for day in $(shell ls | grep -E "^day") ; do go build -o bin/$${day} $${day}/$${day}.go ; done

day:
	@mkdir -p day${DAY}
	@mkdir -p inputs/test

	@if ! [ -f day${DAY}/day${DAY}.go ]; then \
		cat templates/template.go.tmpl | sed "s/DAYNUMBER/$$(echo ${DAY} | sed -E 's/^0//g')/g" > day${DAY}/day${DAY}.go; \
		cp templates/tests.go.tmpl day${DAY}/day${DAY}_test.go; \
		echo Created: day${DAY}/day${DAY}.go ; \
		mkdir -p inputs/test/day${DAY}/1 ;\
		touch inputs/test/day${DAY}/1/input.txt ; \
		touch inputs/test/day${DAY}/1/result_p1.txt ; \
		code inputs/test/day${DAY}/1/input.txt ; \
		code inputs/test/day${DAY}/1/result_p1.txt ; \
		code day${DAY}/day${DAY}.go ; \
	fi

testday:
	@echo RUNNING TESTS FOR DAY ${DAY}
	@go test day${DAY}/*.go  -v

testall:
	@for day in $(shell ls inputs/test/) ; do echo TESTING $${day}; go test $${day}/*.go ; done

testallv:
	@for day in $(shell ls inputs/test/) ; do echo TESTING $${day}; go test -v $${day}/*.go ; done

benchmark:
	@time for day in $(shell ls bin/) ; do time bin/$${day} ; done

profile:
	@file=`go run day${DAY}/day${DAY}.go 2>&1 > /dev/null | grep "cpu profiling disabled" | grep -Eo "[^ ]+$$"` ; \
		go tool pprof --pdf ./bin/day${DAY} $${file} > profile_day${DAY}.pdf ; \
		evince profile_day${DAY}.pdf&

clean:
	rm -fv bin/*