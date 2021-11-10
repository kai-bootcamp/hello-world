from backend.settings.common import *

# SECURITY WARNING: keep the secret key used in production secret!
SECRET_KEY = 'django-insecure-o67b93tybw8xxu)y)=@k28iw$v3057ltodazi!e$q2c_753_4q'

# SECURITY WARNING: don't run with debug turned on in production!
DEBUG = True

ALLOWED_HOSTS = ['*']

CORS_ALLOWED_ORIGINS = [
  'http://localhost:3000',
]