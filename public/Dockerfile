FROM node:carbon
RUN mkdir /usr/src/app
WORKDIR /usr/src/app
ENV PATH /app/node_modules/.bin:$PATH
COPY . .
RUN npm install --silent
RUN npm install react-scripts -g --silent
CMD ["npm", "start"]