FROM nginx:latest

COPY ./kubecost /kubecost
WORKDIR /
CMD /kubecost
