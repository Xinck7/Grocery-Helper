# Generated by Django 3.2.16 on 2023-01-23 06:57

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('grocery_helper_api', '0003_alter_user_profile_pic'),
    ]

    operations = [
        migrations.AlterField(
            model_name='user',
            name='profile_pic',
            field=models.ImageField(blank=True, null=True, upload_to=''),
        ),
    ]
