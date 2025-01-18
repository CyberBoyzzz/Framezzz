FROM node:22 AS build

WORKDIR /app

COPY ./framezzz-frontend/ ./

RUN npm install -g @angular/cli

RUN npm install

EXPOSE 4200

CMD ["ng", "serve", "--host", "0.0.0.0"]