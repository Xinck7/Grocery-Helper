# Generated by Django 3.2.16 on 2023-01-24 03:34

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ('grocery_helper_api', '0004_alter_user_profile_pic'),
    ]

    operations = [
        migrations.CreateModel(
            name='AllowList',
            fields=[
                ('id', models.BigAutoField(auto_created=True, primary_key=True, serialize=False, verbose_name='ID')),
                ('ip_address', models.GenericIPAddressField()),
            ],
        ),
        migrations.AlterField(
            model_name='ingredient',
            name='added_by',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='ingredients', to=settings.AUTH_USER_MODEL),
        ),
        migrations.AlterField(
            model_name='item',
            name='added_by',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='items', to=settings.AUTH_USER_MODEL),
        ),
        migrations.AlterField(
            model_name='list',
            name='added_by',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='list', to=settings.AUTH_USER_MODEL),
        ),
        migrations.AlterField(
            model_name='recipe',
            name='added_by',
            field=models.ForeignKey(on_delete=django.db.models.deletion.CASCADE, related_name='recipe', to=settings.AUTH_USER_MODEL),
        ),
    ]