FROM python:3.10-slim

COPY . .
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

# workdir + runserver
WORKDIR ./grocery_helper_project
CMD [ "python" "manage.py" "makemigrations" ]


# RUN python manage.py runserver 0.0.0.0:8000