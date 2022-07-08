from random import choices
from urllib import request
from rest_framework import serializers
from django.db.models import Sum
from grocery_helper_api.models import Ingredients, Items, List, Recipes

class IngredientsSerializer(serializers.ModelSerializer):
    class Meta:
        model = Ingredients
        fields = '__all__'


class ItemsSerializer(serializers.ModelSerializer):
    class Meta:
        model = Items
        fields = '__all__'

class ListSerializer(serializers.ModelSerializer):
    ingredients = IngredientsSerializer(many=True)
    items = ItemsSerializer(many=True)
    class Meta:
        model = List
        fields = '__all__'


class RecipesSerializer(serializers.ModelSerializer):
    ingredients = IngredientsSerializer(many=True)
    class Meta:
        model = Recipes
        fields = '__all__'
