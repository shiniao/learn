

# PEP 487
class AnimalBase:
    subclasses = []

    def __init_subclass__(cls, speak, **kwargs):
        super().__init_subclass__(**kwargs)
        cls.speak = speak

class Dog(AnimalBase, speak="Wang"):
    pass

class Cat(AnimalBase, speak="Miao"):
    pass



