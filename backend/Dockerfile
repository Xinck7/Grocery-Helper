FROM python:3.10-slim
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

# workdir + runserver
COPY . . 
WORKDIR ./grocery_helper_project
RUN python manage.py
RUN python manage.py makemigrations
RUN python manage.py migrate
# RUN python manage.py runserver 0.0.0.0:8000