#from django.db.models import F
#from .models import newstats
class flyingXMiddleware:
	def __init__(self,get_response):
		self.get_response = get_response
	
	def __call__(self, request):  
		print("request:",request)
		print("request HOST:",request.headers['HOST'])
		print("request lan:",request.headers['Accept-Language'])
		print("METHOD:",request.META['REQUEST_METHOD'])
		print("Agent:",request.META['HTTP_USER_AGENT'])
		print("HOST:",request.META['HTTP_HOST'])
		# self.stats(request.META['HTTP_USER_AGENT'])
		response = self.get_response(request)
		print("response:",response)
		print("*************")
		return response
	def stats(self, os_info):
		# if "Windows" in os_info:
		# 	newstats.objects.all().update(win=F('win')+1)
		# elif "mac" in os_info:
		# 	newstats.objects.all().update(win=F('mac')+1)
		# elif "iPhone" in os_info:
		# 	newstats.objects.all().update(win=F('iph')+1)
		# elif "Android" in os_info:
		# 	newstats.objects.all().update(win=F('android')+1)
		# else:
		# 	newstats.objects.all().update(win=F('oth')+1)	
		pass
	def process_view(self, request, view_func, view_args, view_kwargs):
		print(f'current path: /{view_func.__name__}')
		pass 

	def process_exception(self, request, execption):
		pass
		#from django.core import exceptions
		#ls = list(dir(exceptions))