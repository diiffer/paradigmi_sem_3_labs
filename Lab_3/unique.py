#итератор, который пропускает повторяющиеся элементы
class Unique(object):
    def __init__(self, items, **kwargs):
        self.items = iter(items)
        self.ignore_case = kwargs.get('ignore_case', False)
        self.seen = set()

    def __iter__(self):
        return self

    def __next__(self):
        while True:
            try:
                item = next(self.items)
            except StopIteration:
                raise

            if isinstance(item, str) and self.ignore_case:
                key = item.lower()
            else:
                key = item

            if key in self.seen:
                continue

            self.seen.add(key)
            return item


if __name__ == '__main__':
    from gen_random import gen_random

    data = [1, 1, 1, 1, 1, 2, 2, 2, 2, 2]
    print(list(Unique(data)))

    data = gen_random(10, 1, 3)
    print(list(Unique(data)))

    data = ['a', 'A', 'b', 'B', 'a', 'A', 'b', 'B']
    print(list(Unique(data)))
    print(list(Unique(data, ignore_case=True)))
