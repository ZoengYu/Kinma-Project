from django.urls import path,include

from website import views

urlpatterns = [
        path('',views.flyingX,name='index'),
        path('flyingX/', views.flyingX, name='flyingX'),
        path('flyingX/third_party',views.third_party,name="third_party"),
        path('test/',views.test,name='test'),
        path('getuser/',views.getuser,name='getuser'),
        path('test/create/',views.create,name='create'),
        path('logout/',views.logout,name='logout'),
]