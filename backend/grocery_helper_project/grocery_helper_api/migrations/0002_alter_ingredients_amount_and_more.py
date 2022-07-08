# Generated by Django 4.0.6 on 2022-07-07 05:46

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('grocery_helper_api', '0001_initial'),
    ]

    operations = [
        migrations.AlterField(
            model_name='ingredients',
            name='amount',
            field=models.DecimalField(decimal_places=2, help_text='how many do you want? - needs a number', max_digits=10),
        ),
        migrations.AlterField(
            model_name='ingredients',
            name='calore_per_serving',
            field=models.DecimalField(decimal_places=2, help_text='calorie number eg 500', max_digits=10),
        ),
        migrations.AlterField(
            model_name='ingredients',
            name='price',
            field=models.DecimalField(decimal_places=2, help_text='how much in USD? eg 3', max_digits=10),
        ),
        migrations.AlterField(
            model_name='ingredients',
            name='serving_amount',
            field=models.DecimalField(decimal_places=2, help_text='how many ounces/pounds? eg number only like 5', max_digits=10),
        ),
        migrations.AlterField(
            model_name='items',
            name='price',
            field=models.DecimalField(decimal_places=2, help_text='how much in USD? eg 3', max_digits=10),
        ),
        migrations.AlterField(
            model_name='list',
            name='price',
            field=models.DecimalField(decimal_places=2, help_text='how much in USD? eg 3', max_digits=10),
        ),
        migrations.AlterField(
            model_name='recipes',
            name='price',
            field=models.DecimalField(decimal_places=2, help_text='how much in USD? eg 3', max_digits=10),
        ),
    ]
