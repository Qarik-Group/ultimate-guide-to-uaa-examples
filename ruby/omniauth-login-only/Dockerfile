FROM ruby:2.5

EXPOSE 9292

ADD . /app

RUN cd /app \
  && bundle install

CMD [ "/app/bin/run.sh" ]