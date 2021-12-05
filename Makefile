day:
	@mkdir -p day`date +%d`

	@if ! [ -f day`date +%d`/day`date +%d`.go ]; then \
		cat template.go.tmpl | sed "s/DAYNUMBER/`date +%d | sed -E 's/^0//g'`/g" > day`date +%d`/day`date +%d`.go;\
		echo Created: day`date +%d`/day`date +%d`.go ; \
		code day`date +%d`/day`date +%d`.go ; \
	fi

build:
	@for day in $(shell ls | grep day) ; do go build -o bin/$${day} $${day}/$${day}.go ; done

benchmark:
	@time for day in $(shell ls bin/) ; do time bin/$${day} ; done