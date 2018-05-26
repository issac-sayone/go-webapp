"""oscar_web URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/1.11/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  url(r'^$', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  url(r'^$', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.conf.urls import url, include
    2. Add a URL to urlpatterns:  url(r'^blog/', include('blog.urls'))
"""

from django.conf.urls import include, url
from django.contrib import admin
from oscar.app import application

from django.conf import settings
from django.views import static
from django.contrib.staticfiles.urls import staticfiles_urlpatterns

urlpatterns = [

     url(r'^i18n/', include('django.conf.urls.i18n')),

    # The Django admin is not officially supported; expect breakage.
    # Nonetheless, it's often useful for debugging.
    url(r'^admin/', include(admin.site.urls)),

    url(r'', include(application.urls)),
    url(r'^pages/', include('django.contrib.flatpages.urls')),

    url(r'^static/(.*)$', static.serve, {'document_root': settings.STATIC_ROOT}),

]

urlpatterns = [
   url(r'^media/(?P<path>.*)$', static.serve,
   {'document_root': settings.MEDIA_ROOT,}),
   url(r'', include('django.contrib.staticfiles.urls')),
] + urlpatterns

urlpatterns += staticfiles_urlpatterns()


