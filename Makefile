C2C_WORK_SPACE=./src/c2c

default:
	make build
clean:
	make -C $(C2C_WORK_SPACE) clean
build:
	make build-c2c
check:
	make check-c2c

clean-c2c:
	make -C $(C2C_WORK_SPACE) clean
check-c2c:
	make -C $(C2C_WORK_SPACE) check
build-c2c:
	make -C $(C2C_WORK_SPACE) build
run-c2c:
	make -C $(C2C_WORK_SPACE) run