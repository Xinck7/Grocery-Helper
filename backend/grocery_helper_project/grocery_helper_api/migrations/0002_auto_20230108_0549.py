# Generated by Django 3.2.16 on 2023-01-08 05:49

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('grocery_helper_api', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='ingredient',
            name='ailse',
            field=models.DecimalField(blank=True, decimal_places=0, help_text='where it is in the store approximately', max_digits=10, null=True),
        ),
        migrations.AlterField(
            model_name='ingredient',
            name='calore_per_serving',
            field=models.DecimalField(blank=True, decimal_places=2, help_text='calorie number eg 500', max_digits=10, null=True),
        ),
        migrations.AlterField(
            model_name='ingredient',
            name='price',
            field=models.DecimalField(blank=True, decimal_places=2, help_text='how much in USD? eg 3', max_digits=10, null=True),
        ),
        migrations.AlterField(
            model_name='item',
            name='ailse',
            field=models.DecimalField(blank=True, decimal_places=0, help_text='where it is in the store approximately', max_digits=10, null=True),
        ),
        migrations.AlterField(
            model_name='item',
            name='price',
            field=models.DecimalField(blank=True, decimal_places=2, help_text='how much in USD? eg 3', max_digits=10, null=True),
        ),
        migrations.AlterField(
            model_name='list',
            name='price',
            field=models.DecimalField(blank=True, decimal_places=2, default=0, help_text='total price of items and ingredients', max_digits=10, null=True),
        ),
        migrations.AlterField(
            model_name='recipe',
            name='price',
            field=models.DecimalField(blank=True, decimal_places=2, help_text='how much in USD? eg 3', max_digits=10, null=True),
        ),
        migrations.AlterField(
            model_name='recipe',
            name='time_to_cook',
            field=models.DecimalField(blank=True, decimal_places=1, default=0, help_text='how long to make it start to finish?', max_digits=3, null=True),
        ),
    ]