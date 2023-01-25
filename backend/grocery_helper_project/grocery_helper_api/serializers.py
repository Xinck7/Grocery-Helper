from rest_framework import serializers
# from django.db.models import Sum
from grocery_helper_api.models import Ingredient, Item, List, Recipe, User
from django.contrib.auth import get_user_model

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

# class UserSerializer(serializers.ModelSerializer):

#     class Meta:
#         model = User
#         fields = ('id', 'username', 'bio', 'profile_pic')
#         profile_pic = serializers.ImageField(required=False, max_length=None, 
#                                      allow_empty_file=True, use_url=True)

class UserSerializer(serializers.ModelSerializer):

    class Meta:
        model = get_user_model()
        fields = (
            'id',
            'username',
            'password',
            'email',
            'first_name',
            'last_name',
            'bio',
            'profile_pic'
        )
        extra_kwargs = {
            'password': {'write_only': True},
        }

    def create(self, validated_data):
        user = get_user_model().objects.create_user(**validated_data)
        return user

    def update(self, instance, validated_data):
        if 'password' in validated_data:
            password = validated_data.pop('password')
            instance.set_password(password)
        return super(UserSerializer, self).update(instance, validated_data)
