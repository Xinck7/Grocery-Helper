from django.urls import include, path, re_path, reverse_lazy
from django.conf import settings
from django.conf.urls.static import static
from django.views.generic.base import RedirectView
from rest_framework import routers
from django.contrib import admin

from grocery_helper_api.views import IngredientsViewSet, ItemsViewSet, ListViewSet, RecipesViewSet, UserViewSet, UserLogIn

router = routers.DefaultRouter()
router.register(r'ingredients', IngredientsViewSet)
router.register(r'items', ItemsViewSet)
router.register(r'list', ListViewSet)
router.register(r'recipes', RecipesViewSet)
router.register(r'users', UserViewSet)

urlpatterns = [
    # path('admin/', admin.site.urls),
    path('v1/', include(router.urls)),
    path('api-user-login/', UserLogIn.as_view()),
    path('api-auth/', include('rest_framework.urls', namespace='api-auth')),
    re_path(r'^$', RedirectView.as_view(url=reverse_lazy('api-root'), permanent=False)),  
] + static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)