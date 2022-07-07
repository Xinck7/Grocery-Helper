from rest_framework import serializers

from grocery_helper_api.models import Ingredients, Items, List, Recipes

class IngredientsSerializer(serializers.ModelSerializer):
    class Meta:
        model = Ingredients
        fields = ('name',
                  'location',
                  'ailse',
                  'amount',
                  'calore_per_serving',
                  'serving_amount',
                  'obtained',
                  'price')


class ItemsSerializer(serializers.ModelSerializer):
    class Meta:
        model = Items
        fields = ('name',
                  'location',
                  'ailse',
                  'price')

class ListSerializer(serializers.ModelSerializer):
    class Meta:
        model = List
        fields = ('name',
                  'ingredients',
                  'items',
                  'price')

class RecipesSerializer(serializers.ModelSerializer):
    class Meta:
        model = Recipes
        fields = ('name',
                  'ingredients',
                  'price')