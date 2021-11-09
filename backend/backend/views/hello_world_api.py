from django.http import HttpResponse
from random import randint

def hello_world(request):
  return HttpResponse("Hello world", content_type='text/plain')


def random_string(request):
  random_str = "This is a random number between 1 and 100: %d" % randint(1, 100)
  return HttpResponse(random_str, content_type='text/plain')