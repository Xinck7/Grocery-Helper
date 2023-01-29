from django.db import models
from django.contrib.auth.models import AbstractUser
from django.conf import settings

class Ingredient(models.Model):
    name = models.CharField(max_length=100, help_text="name of ingredient")
    location = models.CharField(max_length=100, help_text="krogers/sams/etc")
    ailse = models.DecimalField(null=True, blank=True, decimal_places=0, max_digits=10,help_text="where it is in the store approximately")
    amount = models.DecimalField(decimal_places=0, max_digits=10, help_text="how many do you want? - needs a number", default=1)
    calore_per_serving = models.DecimalField(null=True, blank=True, decimal_places=2, max_digits=10, help_text="calorie number eg 500")
    serving_amount = models.DecimalField(decimal_places=2, max_digits=10, help_text="how many ounces/pounds? eg number only like 5", default=0)
    obtained = models.BooleanField(help_text="is it in the cart?", default=False)
    price = models.DecimalField(null=True, blank=True, decimal_places=2, max_digits=10, help_text="how much in USD? eg 3")
    picture = models.ImageField(null=True, blank=True, help_text='Only pictures are supported through this tool', upload_to='grocery_helper_api/media')
    added_by = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, related_name="ingredients")
    show_only_your_added_by = models.BooleanField(default=False, help_text="Show only your stuff or look at the catalog")

    def __str__(self):
        return self.name

    #todo get price from kroger api

class Item(models.Model):
    name = models.CharField(max_length=100, help_text="name of ingredient")
    location = models.CharField(max_length=100, help_text="krogers/sams/etc")
    ailse = models.DecimalField(null=True, blank=True, decimal_places=0, max_digits=10, help_text="where it is in the store approximately")
    price = models.DecimalField(null=True, blank=True, decimal_places=2, max_digits=10, help_text="how much in USD? eg 3")
    picture = models.ImageField(null=True, blank=True, help_text='Only pictures are supported through this tool', upload_to='grocery_helper_api/media')
    added_by = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, related_name="items")
    show_only_your_added_by = models.BooleanField(default=False, help_text="Show only your stuff or look at the catalog")

    def __str__(self):
        return self.name

    #todo get price from kroger api

class List(models.Model):
    name = models.CharField(max_length=100)
    ingredients = models.ManyToManyField(Ingredient)
    items = models.ManyToManyField(Item)
    price = models.DecimalField(null=True, blank=True, decimal_places=2, max_digits=10, help_text="total price of items and ingredients", default=0)
    owner = models.CharField(max_length=100, help_text="Whose list is it anyways?", default='root')
    added_by = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, related_name="list")
    show_only_your_added_by = models.BooleanField(default=False, help_text="Show only your stuff or look at the catalog")

    def __str__(self):
        return self.name

    #todo calculate price with celery task

class Recipe(models.Model):
    name = models.CharField(max_length=100)
    ingredients = models.ManyToManyField(Ingredient)
    price = models.DecimalField(null=True, blank=True, decimal_places=2, max_digits=10, help_text="how much in USD? eg 3")
    time_to_cook = models.DecimalField(null=True, blank=True, decimal_places=1, max_digits=3, help_text="how long to make it start to finish?", default=0)
    added_by = models.ForeignKey(settings.AUTH_USER_MODEL, on_delete=models.CASCADE, related_name="recipe")
    show_only_your_added_by = models.BooleanField(default=False, help_text="Show only your stuff or look at the catalog")

    def __str__(self):
        return self.name


#Setup Accounts
class User(AbstractUser):
    bio = models.TextField(blank=True)
    profile_pic = models.ImageField(blank=True, null= True)

    def __str__(self):
        return self.username

#Copying from DRF
class AllowList(models.Model):
    ip_address = models.GenericIPAddressField()

    def __str__(self):
        return self.ip_address