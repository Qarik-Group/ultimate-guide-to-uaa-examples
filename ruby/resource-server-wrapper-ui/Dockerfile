FROM ruby:2.5

EXPOSE 9393
ENV PORT=9393

ADD . /app

RUN cd /app \
  && bundle install

CMD [ "/app/bin/run.sh" ]