from rest_framework.permissions import SAFE_METHODS, BasePermission

from grocery_helper_api.models import AllowList


class IsAdminOrReadOnly(BasePermission):
    def has_permission(self, request, view):
        return bool(
            request.method in SAFE_METHODS or request.user and request.user.is_authenticated and request.user.is_staff
        )

class IsOwnerOrReadOnly(BasePermission):

    def has_object_permission(self, request, view, obj):
        # Read permissions are allowed to any request,
        # so we'll always allow GET, HEAD or OPTIONS requests.
        if request.method in SAFE_METHODS:
            return True

        # Write permissions are only allowed to the owner of the snippet.
        return obj.added_by == request.user

class AllowListPermission(BasePermission):
    def has_permission(self, request, view):
        ip_addr = request.META["REMOTE_ADDR"]
        allowed = AllowList.objects.filter(ip_address=ip_addr).exists()
        return allowed