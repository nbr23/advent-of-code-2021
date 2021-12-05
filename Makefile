build:
	@mkdir -p bin
	@for day in $(shell ls | grep day) ; do go build -o bin/$${day} $${day}/$${day}.go ; done

day:
	@mkdir -p day`date +%d`
	@mkdir -p inputs/test

	@if ! [ -f day`date +%d`/day`date +%d`.go ]; then \
		cat template.go.tmpl | sed "s/DAYNUMBER/`date +%d | sed -E 's/^0//g'`/g" > day`date +%d`/day`date +%d`.go;\
		echo Created: day`date +%d`/day`date +%d`.go ; \
		touch inputs/test/day`date +%d`.txt ; \
		code inputs/test/day`date +%d`.txt ; \
		code day`date +%d`/day`date +%d`.go ; \
	fi

benchmark:
	@time for day in $(shell ls bin/) ; do time bin/$${day} ; done

clean:
	rm -fv bin/*