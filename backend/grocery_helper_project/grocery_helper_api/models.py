from operator import mod
from django.db import models

class Ingredients(models.Model):
    name = models.CharField(max_length=100, help_text="name of ingredient")
    location = models.CharField(max_length=100, help_text="krogers/sams/etc")
    ailse = models.DecimalField(decimal_places=2, max_digits=10,help_text="where it is in the store approximately")
    amount = models.DecimalField(decimal_places=2, max_digits=10, help_text="how many do you want? - needs a number")
    calore_per_serving = models.DecimalField(decimal_places=2, max_digits=10, help_text="calorie number eg 500", default=0)
    serving_amount = models.DecimalField(decimal_places=2, max_digits=10, help_text="how many ounces/pounds? eg number only like 5", default=0)
    obtained = models.BooleanField(help_text="is it in the cart?", default=False)
    price = models.DecimalField(decimal_places=2, max_digits=10, help_text="how much in USD? eg 3")
    #todo can price be called from a website? or more?

class Items(models.Model):
    name = models.CharField(max_length=100, help_text="name of ingredient")
    location = models.CharField(max_length=100, help_text="krogers/sams/etc")
    ailse = models.DecimalField(decimal_places=2, max_digits=10, help_text="where it is in the store approximately")
    price = models.DecimalField(decimal_places=2, max_digits=10, help_text="how much in USD? eg 3")
    #todo can price be called from a website? or more?

class List(models.Model):
    name = models.CharField(max_length=100)
    ingredients = models.ForeignKey(Ingredients, on_delete=models.DO_NOTHING)
    items = models.ForeignKey(Items, on_delete=models.DO_NOTHING)
    price = models.DecimalField(decimal_places=2, max_digits=10, help_text="how much in USD? eg 3")
    #todo need to calculate total price of items
    #todo setup return of list in the 'ingredients' as it doesn't right now

class Recipes(models.Model):
    name = models.CharField(max_length=100)
    ingredients = models.ForeignKey(Ingredients, on_delete=models.DO_NOTHING)
    price = models.DecimalField(decimal_places=2, max_digits=10, help_text="how much in USD? eg 3")
    #todo calculate the individual prices
    #todo need to implement selectable options of ingredients
    #todo need to implement frontend how to dynamically add a new ingredient in a modal if you haven't already