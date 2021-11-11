from backend.settings.common import *

# SECURITY WARNING: keep the secret key used in production secret!
SECRET_KEY = 'FMJbGYoIZIrtuJmwoLTjToydqAhuvOcmFhEOYVnJKXZedTRsCm'

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = False

ALLOWED_HOSTS = ['hello-world-backend-vokmb666lq-as.a.run.app', '34.107.141.204', '34.149.180.144']

CORS_ALLOWED_ORIGINS = [
  'http://34.117.74.251',
]