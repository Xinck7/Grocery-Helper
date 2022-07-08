from django.contrib import admin
from grocery_helper_api.models import Ingredient, Item, List, Recipe
# Register your models here.
admin.site.register(Ingredient)
admin.site.register(Item)
admin.site.register(List)
admin.site.register(Recipe)