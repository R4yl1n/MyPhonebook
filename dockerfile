FROM Python:3

WORKDIR /usr/src/app

COPY ./Flask .
RUN pip install Flask



CMD ["Python3", "-m", "flask", "run","--host=0.0.0.0"]