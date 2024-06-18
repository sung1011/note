all: gogo pb

.PHONY: gogo
gogo:
	@cd ./protocols/gogo && protoc -I=.  --gogoslick_out=./gen ./*.proto

.PHONY: pb
pb:
	@protoc -I=. --go_out=./protocols/pb ./protocols/pb/*.proto
