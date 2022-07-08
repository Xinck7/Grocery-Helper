from random import choices
from urllib import request
from rest_framework import serializers
from django.db.models import Sum
from grocery_helper_api.models import Ingredient, Item, List, Recipe

class IngredientSerializer(serializers.ModelSerializer):
    class Meta:
        model = Ingredient
        fields = '__all__'


class ItemSerializer(serializers.ModelSerializer):
    class Meta:
        model = Item
        fields = '__all__'

class ListSerializer(serializers.ModelSerializer):
    ingredients = IngredientSerializer(many=True)
    items = ItemSerializer(many=True)
    class Meta:
        model = List
        fields = '__all__'


class RecipeSerializer(serializers.ModelSerializer):
    ingredients = IngredientSerializer(many=True)
    class Meta:
        model = Recipe
        fields = '__all__'
