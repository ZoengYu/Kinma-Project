from django.urls import path

from . import views

urlpatterns = [
	path('',views.flyingX,name='index'),
    path('flyingX/', views.flyingX, name='flyingX'),
    path('flyingX/third_party',views.third_party,name="third_party")
]