FROM python:3.10-slim

#code
COPY . /opt
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

# workdir + migrations
WORKDIR /opt/grocery_helper_project
CMD [ "python3" "manage.py" "makemigrations"]