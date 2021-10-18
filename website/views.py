from django.http.response import HttpResponseRedirect
from django.shortcuts import render
from django.http import HttpResponse, JsonResponse
from .models import flyingX_User
from django.contrib import auth

# Create your views here.
def flyingX(response):
	return render(response,"website/main.html",{})

def third_party(response):
	return render(response,"website/third_party.html",{})
def test(response):
	return render(response,"website/test.html",{})
def getuser(request):
	profiles = flyingX_User.objects.all()
	return JsonResponse({"profiles":list(profiles.values())})

def create(request):
	if request.method == 'POST':
		name = request.POST['name']
		email = request.POST['email']
		phone = request.POST['phone']
		user = flyingX_User(name=name,email=email,phone=phone)
		user.save()
		return HttpResponse('New Profile Created Successfully')

def logout(request):
	auth.logout(request)
	return HttpResponseRedirect('/flyingX')

# def signup(request):
# 	if request.method='POST':
# 		form = flyingX_User(request.POST)
# 		if form.is_valid():
# 			form.save()