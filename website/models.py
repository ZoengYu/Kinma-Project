from django.db import models
from django.contrib.auth.models import User
from django.core.validators import RegexValidator

class flyingX_User(models.Model):
	name = models.CharField(max_length=50,verbose_name='使用者')
	email = models.CharField(max_length=50, verbose_name='信箱')
	phoneNumberRegex = RegexValidator(regex = r"^\+?1?\d{8,15}$",message="Phone must be entered in the format: '+886930322580/0930322580")
	phoneNumber = models.CharField(validators = [phoneNumberRegex], max_length = 16, unique = True,default="",verbose_name="手機號碼")
	class Meta:
		verbose_name = "flyingX 使用者個人資料"