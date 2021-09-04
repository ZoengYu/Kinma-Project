class FlyingXMiddleware:
	def __init__(self,get_response):
		self.get_response = get_response
	
	def __call__(self, request):
		print("hello world")
		print(request)
		response = self.get_response(request)
		print(response)
		return response

	def process_view(self, request, view_func, view_args, view_kwargs):
		print(f'current path: /{view_func.__name__}')
		pass