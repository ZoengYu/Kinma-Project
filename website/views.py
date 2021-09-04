from django.shortcuts import render
from django.http import HttpResponse

# Create your views here.
def flyingX(response):
	return render(response,"website/main.html",{})

def third_party(response):
	return render(response,"website/third_party.html",{})