GOCMD=go

.PHONY: gen-docs
gen-docs:
	$(GOCMD) run tools/gendocs/main.go
