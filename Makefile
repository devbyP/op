full-test:
	nx run-many -t test -p tempUnitConverter greeting

dev-up:
	@nx run-many -r serve -p tempUnitConverter greeting

build-test:
	@nx run-many -r build -p tempUnitConverter greeting
