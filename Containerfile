FROM registry.access.redhat.com/ubi9/ubi-micro 

COPY hello_openshift /hello_openshift

CMD [ "/hello_openshift" ]
