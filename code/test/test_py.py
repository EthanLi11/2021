import datetime

dt = datetime.datetime.strptime("20210105", '%Y%m%d')
year = dt.isocalendar()[0]
week = dt.isocalendar()[1]



print(b)
print(year)