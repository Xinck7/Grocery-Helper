from operator import mod
from django.db import models

class Ingredients(models.Model):
    name = models.CharField(max_length=100, help_text="name of ingredient")
    location = models.CharField(max_length=100, help_text="krogers/sams/etc")
    ailse = models.DecimalField(decimal_places=0, max_digits=10,help_text="where it is in the store approximately")
    amount = models.DecimalField(decimal_places=0, max_digits=10, help_text="how many do you want? - needs a number", default=1)
    calore_per_serving = models.DecimalField(decimal_places=2, max_digits=10, help_text="calorie number eg 500", default=0)
    serving_amount = models.DecimalField(decimal_places=2, max_digits=10, help_text="how many ounces/pounds? eg number only like 5", default=0)
    obtained = models.BooleanField(help_text="is it in the cart?", default=False)
    price = models.DecimalField(decimal_places=2, max_digits=10, help_text="how much in USD? eg 3")

    def __str__(self):
        return self.name

class Items(models.Model):
    name = models.CharField(max_length=100, help_text="name of ingredient")
    location = models.CharField(max_length=100, help_text="krogers/sams/etc")
    ailse = models.DecimalField(decimal_places=0, max_digits=10, help_text="where it is in the store approximately")
    price = models.DecimalField(decimal_places=2, max_digits=10, help_text="how much in USD? eg 3")
    
    def __str__(self):
        return self.name

class List(models.Model):
    name = models.CharField(max_length=100)
    ingredients = models.ManyToManyField(Ingredients)
    items = models.ManyToManyField(Items)
    price = models.DecimalField(decimal_places=2, max_digits=10, help_text="total price of items and ingredients", default=0)

    def __str__(self):
        return self.name

class Recipes(models.Model):
    name = models.CharField(max_length=100)
    ingredients = models.ManyToManyField(Ingredients)
    price = models.DecimalField(decimal_places=2, max_digits=10, help_text="how much in USD? eg 3")
    time_to_cook = models.DecimalField(decimal_places=1, max_digits=3, help_text="how long to make it start to finish?", default=0)

    def __str__(self):
        return self.name