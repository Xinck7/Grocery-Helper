FROM python:3.10-slim

COPY . .
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

# workdir + runserver
WORKDIR ./grocery_helper_project
CMD [ "python3" "manage.py" "makemigrations" "&&" "python3" "manage.py" "migrate"]