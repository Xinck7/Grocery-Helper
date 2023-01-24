from django.urls import include, path
from rest_framework.routers import DefaultRouter
from grocery_helper_api.views import IngredientsViewSet, ItemsViewSet, ListViewSet, RecipesViewSet, UserViewSet
from django.conf import settings
from django.conf.urls.static import static

router = DefaultRouter()
router.register(r'ingredients', IngredientsViewSet)
router.register(r'items', ItemsViewSet)
router.register(r'list', ListViewSet)
router.register(r'recipes', RecipesViewSet)
router.register(r'users', UserViewSet)

urlpatterns = [
    path('', include(router.urls)), 
]