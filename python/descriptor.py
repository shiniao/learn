class Quantity:
    
    def __set_name__(self, owner, name):
        """ 设置被描述实例名称 """
        self.name = name

    def __get__(self, instance, owner):
        return instance.__dict__.get(self.name)

    def __set__(self, instance, value):
        """ 确保属性值大于 1 """
        if value > 0:
            instance.__dict__[self.name] = value
        else:
            raise ValueError('value must > 0')


class Student:
    # * 任何被描述的属性值都拥有了描述符的特性
    age = Quantity()
        
    

    


s = Student()
s.age = 21
s2 = Student()
print(s2.age)


