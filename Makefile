gen:
	buf generate contracts

lint:
	buf lint

breaking:
	buf breaking --against '.git#branch=main'