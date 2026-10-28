#генератор, который принимает список словарей и один или несколько ключей. Он возвращает генератор, который будет выдавать значения этих ключей из каждого словаря в списке. Если ключ отсутствует в словаре или его значение равно None, то этот словарь пропускается.
def field(items, *args):
    assert len(args) > 0
    if len(args) == 1:
        key = args[0]
        for item in items:
            if key in item and item[key] is not None:
                yield item[key]
    else:
        for item in items:
            result_dict = {}
            for key in args:
                if key in item and item[key] is not None:
                    result_dict[key] = item[key]
            if result_dict:
                yield result_dict


if __name__ == '__main__':
    goods = [
        {'title': 'Ковер', 'price': 2000, 'color': 'green'},
        {'title': 'Диван для отдыха', 'color': 'black'}
    ]

    print('field(goods, "title"):')
    for g in field(goods, 'title'):
        print(g)

    print('\nfield(goods, "title", "price"):')
    for g in field(goods, 'title', 'price'):
        print(g)
