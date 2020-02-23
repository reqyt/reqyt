FROM busybox
ADD ./echo /
HEALTHCHECK --interval=1s --timeout=3s CMD wget -O - -q localhost:8000
CMD ["/echo"]
EXPOSE 8000
