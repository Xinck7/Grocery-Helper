from django.urls import include, path

from rest_framework import routers

from grocery_helper_api.views import IngredientsViewSet, ItemsViewSet, ListViewSet, RecipesViewSet

router = routers.DefaultRouter()
router.register(r'ingredients', IngredientsViewSet)
router.register(r'items', ItemsViewSet)
router.register(r'list', ListViewSet)
router.register(r'recipes', RecipesViewSet)

urlpatterns = [
    path('', include(router.urls)),  
]