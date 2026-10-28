#генератор случайных чисел, который принимает два аргумента: количество чисел и диапазон (кортеж из двух чисел). Он возвращает генератор, который будет выдавать случайные числа в указанном диапазоне.
import random


def gen_random(num_count, begin, end):
    for _ in range(num_count):
        yield random.randint(begin, end)


if __name__ == '__main__':
    print(list(gen_random(5, 1, 3)))
