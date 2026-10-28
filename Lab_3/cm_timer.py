# Контекстный менеджер для измерения времени выполнения блока кода
import time
from contextlib import contextmanager


class cm_timer_1:
    def __enter__(self):
        self.start_time = time.time()
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        self.end_time = time.time()
        print(f'time: {self.end_time - self.start_time}')


@contextmanager
def cm_timer_2():
    start_time = time.time()
    try:
        yield
    finally:
        end_time = time.time()
        print(f'time: {end_time - start_time}')


if __name__ == '__main__':
    from time import sleep

    with cm_timer_1():
        sleep(0.5)

    with cm_timer_2():
        sleep(0.5)
